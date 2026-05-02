from enum import Enum
from dataclasses import dataclass
from typing import Optional
import re


class ShellId(Enum):
    EXEC = "EXEC"
    SHELL = "SHELL"
    ROOT = "ROOT"


@dataclass
class Transition:
    from_shell: ShellId
    to_shell: ShellId
    command: str
    prompt_pattern: str  # regex to detect target shell
    dialog: Optional[dict] = None  # e.g. {"Password:": "mypass"}


# Default prompt patterns
PROMPTS = {
    ShellId.EXEC: r"[\$]\s*$",
    ShellId.SHELL: r"[\$]\s*$",
    ShellId.ROOT: r"[#]\s*$",
}

# Transition graph
TRANSITIONS = [
    Transition(ShellId.EXEC, ShellId.SHELL, "sh", PROMPTS[ShellId.SHELL]),
    Transition(ShellId.SHELL, ShellId.EXEC, "exit", PROMPTS[ShellId.EXEC]),
    Transition(ShellId.SHELL, ShellId.ROOT, "sudo su", PROMPTS[ShellId.ROOT],
               dialog={"[Pp]assword": ""}),  # password filled at runtime
    Transition(ShellId.ROOT, ShellId.SHELL, "exit", PROMPTS[ShellId.SHELL]),
    Transition(ShellId.EXEC, ShellId.ROOT, "sudo su", PROMPTS[ShellId.ROOT],
               dialog={"[Pp]assword": ""}),
    Transition(ShellId.ROOT, ShellId.EXEC, "exit", PROMPTS[ShellId.EXEC]),
]


def detect_shell(output: str) -> Optional[ShellId]:
    """Detect current shell from output using prompt patterns."""
    lines = output.strip().split("\n")
    if not lines:
        return None
    last_line = lines[-1]
    # ROOT must be checked before SHELL/EXEC since # is more specific
    if re.search(PROMPTS[ShellId.ROOT], last_line):
        return ShellId.ROOT
    if re.search(PROMPTS[ShellId.EXEC], last_line):
        return ShellId.EXEC
    return None


def find_transition(from_shell: ShellId, to_shell: ShellId) -> Optional[Transition]:
    """Find direct transition between two shells."""
    for t in TRANSITIONS:
        if t.from_shell == from_shell and t.to_shell == to_shell:
            return t
    return None
