import logging
from shells import ShellId, find_path, PROMPTS
from prompt_detector import PromptDetector
from transport import SSHTransport
from decorators import CommandDecorator, RetryDecorator
from dialogs import DialogHandler

logger = logging.getLogger("shell_navigator")


class ShellNavigator:
    """Production-grade navigator with BFS routing, decorators, and dialog handling."""

    def __init__(self, host: str, username: str, password: str, port: int = 22,
                 timeout: float = 10.0):
        self.transport = SSHTransport(host, username, password, port)
        self.detector = PromptDetector(timeout=timeout)
        self.dialog_handler = DialogHandler(default_password=password)
        self.password = password
        self.current_shell: ShellId | None = None
        self.channel = None
        self._decorators: list[CommandDecorator] = []

    def add_decorator(self, decorator: CommandDecorator):
        self._decorators.append(decorator)

    def connect(self) -> str:
        self.channel = self.transport.connect()
        output = self.detector.wait_for_prompt(self.channel, PROMPTS[ShellId.EXEC])
        self.current_shell = ShellId.EXEC
        logger.info(f"Connected, shell={self.current_shell.value}")
        return output

    def navigate(self, target: str | ShellId) -> str:
        """Navigate to target shell via BFS shortest path."""
        if isinstance(target, str):
            target = ShellId(target)
        if self.current_shell == target:
            return ""

        path = find_path(self.current_shell, target)
        output = ""
        for transition in path:
            output = self._execute_transition(transition)
        return output

    def _execute_transition(self, t) -> str:
        self.transport.send(t.command)

        if t.dialog:
            patterns = dict(t.dialog)
            patterns["__target__"] = t.prompt_pattern
            matched, buf = self.detector.wait_for_any(self.channel, patterns)
            if matched != "__target__":
                response = t.dialog[matched] or self.password
                self.transport.send(response)
                buf = self.detector.wait_for_prompt(self.channel, t.prompt_pattern)
        else:
            buf = self.detector.wait_for_prompt(self.channel, t.prompt_pattern)

        self.current_shell = t.to_shell
        logger.info(f"Transitioned to {self.current_shell.value}")
        return buf

    def execute(self, command: str) -> str:
        """Execute command with decorator pipeline and retry support."""
        if not self.current_shell:
            raise RuntimeError("Not connected")

        shell_name = self.current_shell.value

        # Pre-execute decorators
        cmd = command
        for d in self._decorators:
            cmd = d.before_execute(cmd, shell_name)

        # Find retry config
        max_retries = 1
        for d in self._decorators:
            if isinstance(d, RetryDecorator):
                max_retries = d.max_retries
                break

        last_error = None
        for attempt in range(max_retries):
            try:
                # Reconnect if needed
                if not self.transport.is_connected:
                    logger.warning("Connection lost, reconnecting...")
                    self.channel = self.transport.reconnect()
                    self.detector.wait_for_prompt(self.channel, PROMPTS[ShellId.EXEC])
                    self.current_shell = ShellId.EXEC

                prompt = PROMPTS[self.current_shell]
                self.transport.send(cmd)
                output = self.detector.wait_for_prompt(self.channel, prompt)

                # Check for dialogs in output
                matched, response = self.dialog_handler.check(output)
                if matched:
                    self.transport.send(response)
                    output = self.detector.wait_for_prompt(self.channel, prompt)

                # Post-execute decorators
                for d in self._decorators:
                    output = d.after_execute(cmd, output, shell_name)

                return output

            except Exception as e:
                last_error = e
                if attempt < max_retries - 1:
                    logger.warning(f"Attempt {attempt+1} failed: {e}, retrying...")

        raise last_error

    def disconnect(self):
        self.transport.disconnect()
        self.current_shell = None
        logger.info("Disconnected")
