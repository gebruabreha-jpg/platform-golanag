import re
import time
from typing import Optional


class PromptDetector:
    """Reads output from a channel and detects when a prompt appears."""

    def __init__(self, timeout: float = 10.0, read_chunk: int = 4096):
        self.timeout = timeout
        self.read_chunk = read_chunk

    def wait_for_prompt(self, channel, pattern: str) -> str:
        """Read channel output until prompt pattern matches or timeout."""
        buffer = ""
        start = time.time()
        while time.time() - start < self.timeout:
            if channel.recv_ready():
                data = channel.recv(self.read_chunk).decode("utf-8", errors="replace")
                buffer += data
                if re.search(pattern, buffer.split("\n")[-1]):
                    return buffer
            time.sleep(0.1)
        raise TimeoutError(f"Prompt '{pattern}' not detected within {self.timeout}s. Buffer:\n{buffer}")

    def wait_for_any(self, channel, patterns: dict[str, str]) -> tuple[str, str]:
        """Wait for any of the given patterns. Returns (matched_key, buffer)."""
        buffer = ""
        start = time.time()
        while time.time() - start < self.timeout:
            if channel.recv_ready():
                data = channel.recv(self.read_chunk).decode("utf-8", errors="replace")
                buffer += data
                last_line = buffer.split("\n")[-1]
                for key, pattern in patterns.items():
                    if re.search(pattern, last_line):
                        return key, buffer
            time.sleep(0.1)
        raise TimeoutError(f"None of {list(patterns.keys())} detected within {self.timeout}s")
