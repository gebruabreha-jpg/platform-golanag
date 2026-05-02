from shells import ShellId, Transition, find_transition, detect_shell, PROMPTS
from prompt_detector import PromptDetector
from transport import SSHTransport


class ShellNavigator:
    """MVP state-machine navigator for multi-shell SSH sessions."""

    def __init__(self, host: str, username: str, password: str, port: int = 22,
                 timeout: float = 10.0):
        self.transport = SSHTransport(host, username, password, port)
        self.detector = PromptDetector(timeout=timeout)
        self.password = password
        self.current_shell: ShellId | None = None
        self.channel = None

    def connect(self) -> str:
        self.channel = self.transport.connect()
        output = self.detector.wait_for_prompt(self.channel, PROMPTS[ShellId.EXEC])
        self.current_shell = ShellId.EXEC
        print(f"[nav] connected, shell={self.current_shell.value}")
        return output

    def navigate(self, target: str | ShellId) -> str:
        if isinstance(target, str):
            target = ShellId(target)
        if self.current_shell == target:
            return ""

        transition = find_transition(self.current_shell, target)
        if not transition:
            raise ValueError(f"No transition from {self.current_shell.value} to {target.value}")

        return self._execute_transition(transition)

    def _execute_transition(self, t: Transition) -> str:
        self.transport.send(t.command)

        # Handle dialogs (e.g. password prompt)
        if t.dialog:
            patterns = dict(t.dialog)
            patterns["__target__"] = t.prompt_pattern
            matched, buf = self.detector.wait_for_any(self.channel, patterns)
            if matched != "__target__":
                # Respond to dialog
                response = t.dialog[matched]
                if not response:
                    response = self.password  # default to SSH password
                self.transport.send(response)
                buf = self.detector.wait_for_prompt(self.channel, t.prompt_pattern)
        else:
            buf = self.detector.wait_for_prompt(self.channel, t.prompt_pattern)

        self.current_shell = t.to_shell
        print(f"[nav] transitioned to {self.current_shell.value}")
        return buf

    def execute(self, command: str) -> str:
        if not self.current_shell:
            raise RuntimeError("Not connected")
        prompt = PROMPTS[self.current_shell]
        self.transport.send(command)
        return self.detector.wait_for_prompt(self.channel, prompt)

    def disconnect(self):
        self.transport.disconnect()
        self.current_shell = None
        print("[nav] disconnected")
