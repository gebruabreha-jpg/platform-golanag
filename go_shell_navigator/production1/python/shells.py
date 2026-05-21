from enum import Enum
from dataclasses import dataclass, field
from typing import Optional
from collections import deque
import re


class ShellId(Enum):
    EXEC = "EXEC"
    SHELL = "SHELL"
    ROOT = "ROOT"
    CONF = "CONF"


@dataclass
class Transition:
    from_shell: ShellId
    to_shell: ShellId
    command: str
    prompt_pattern: str
    dialog: Optional[dict] = None


PROMPTS = {
    ShellId.EXEC: r"[\$]\s*$",
    ShellId.SHELL: r"[\$]\s*$",
    ShellId.ROOT: r"[^(config)]#\s*$",
    ShellId.CONF: r"\(config\)#\s*$",
}

TRANSITIONS = [
    Transition(ShellId.EXEC, ShellId.SHELL, "sh", PROMPTS[ShellId.SHELL]),
    Transition(ShellId.SHELL, ShellId.EXEC, "exit", PROMPTS[ShellId.EXEC]),
    Transition(ShellId.SHELL, ShellId.ROOT, "sudo su", PROMPTS[ShellId.ROOT],
               dialog={"[Pp]assword": ""}),
    Transition(ShellId.ROOT, ShellId.SHELL, "exit", PROMPTS[ShellId.SHELL]),
    Transition(ShellId.EXEC, ShellId.ROOT, "sudo su", PROMPTS[ShellId.ROOT],
               dialog={"[Pp]assword": ""}),
    Transition(ShellId.ROOT, ShellId.EXEC, "exit", PROMPTS[ShellId.EXEC]),
    Transition(ShellId.ROOT, ShellId.CONF, "configure", PROMPTS[ShellId.CONF]),
    Transition(ShellId.CONF, ShellId.ROOT, "exit", PROMPTS[ShellId.ROOT]),
]


def find_transition(from_shell: ShellId, to_shell: ShellId) -> Optional[Transition]:
    for t in TRANSITIONS:
        if t.from_shell == from_shell and t.to_shell == to_shell:
            return t
    return None


def find_path(from_shell: ShellId, to_shell: ShellId) -> list[Transition]:
    """BFS to find shortest transition path between any two shells."""
    if from_shell == to_shell:
        return []

    # Build adjacency
    adj: dict[ShellId, list[Transition]] = {}
    for t in TRANSITIONS:
        adj.setdefault(t.from_shell, []).append(t)

    visited = {from_shell}
    queue = deque([(from_shell, [])])

    while queue:
        current, path = queue.popleft()
        for t in adj.get(current, []):
            if t.to_shell == to_shell:
                return path + [t]
            if t.to_shell not in visited:
                visited.add(t.to_shell)
                queue.append((t.to_shell, path + [t]))

    raise ValueError(f"No path from {from_shell.value} to {to_shell.value}")


def detect_shell(output: str) -> Optional[ShellId]:
    lines = output.strip().split("\n")
    if not lines:
        return None
    last_line = lines[-1]
    # Check CONF before ROOT (more specific pattern first)
    if re.search(PROMPTS[ShellId.CONF], last_line):
        return ShellId.CONF
    if re.search(PROMPTS[ShellId.ROOT], last_line):
        return ShellId.ROOT
    if re.search(PROMPTS[ShellId.EXEC], last_line):
        return ShellId.EXEC
    return None
