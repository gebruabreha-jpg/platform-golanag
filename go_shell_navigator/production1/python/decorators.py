"""Decorator framework for Shell Navigator (classic Decorator Pattern)."""

from abc import ABC, abstractmethod
import re
import time
import logging

logger = logging.getLogger("shell_navigator")


class CommandDecorator(ABC):
    """Base decorator. Subclass and override pre/post hooks."""

    @abstractmethod
    def before_execute(self, command: str, shell_name: str) -> str:
        """Called before command execution. Return (possibly modified) command."""
        return command

    @abstractmethod
    def after_execute(self, command: str, output: str, shell_name: str) -> str:
        """Called after command execution. Return (possibly modified) output."""
        return output


class LoggingDecorator(CommandDecorator):
    def before_execute(self, command: str, shell_name: str) -> str:
        logger.info(f"[{shell_name}] >>> {command}")
        return command

    def after_execute(self, command: str, output: str, shell_name: str) -> str:
        logger.info(f"[{shell_name}] <<< {len(output)} chars")
        return output


class TimingDecorator(CommandDecorator):
    def __init__(self):
        self._start = 0.0

    def before_execute(self, command: str, shell_name: str) -> str:
        self._start = time.time()
        return command

    def after_execute(self, command: str, output: str, shell_name: str) -> str:
        elapsed = time.time() - self._start
        logger.info(f"[{shell_name}] '{command}' took {elapsed:.3f}s")
        return output


class CommandBlocker(CommandDecorator):
    """Blocks dangerous commands from being executed."""

    def __init__(self, blocked: list[str] = None):
        self.blocked = blocked or ["rm -rf /", "format", "mkfs", "dd if="]

    def before_execute(self, command: str, shell_name: str) -> str:
        for pattern in self.blocked:
            if pattern in command:
                raise PermissionError(f"Blocked dangerous command: '{command}'")
        return command

    def after_execute(self, command: str, output: str, shell_name: str) -> str:
        return output


class SecretMasker(CommandDecorator):
    """Masks secrets in output for safe logging."""

    def __init__(self, patterns: list[str] = None):
        self.patterns = patterns or [
            r"(?i)(password|secret|token|key)\s*[:=]\s*\S+",
        ]

    def before_execute(self, command: str, shell_name: str) -> str:
        return command

    def after_execute(self, command: str, output: str, shell_name: str) -> str:
        masked = output
        for p in self.patterns:
            masked = re.sub(p, r"\1=***MASKED***", masked)
        return masked


class RetryDecorator(CommandDecorator):
    """Marks commands for retry on failure (used by navigator)."""

    def __init__(self, max_retries: int = 3):
        self.max_retries = max_retries

    def before_execute(self, command: str, shell_name: str) -> str:
        return command

    def after_execute(self, command: str, output: str, shell_name: str) -> str:
        return output
