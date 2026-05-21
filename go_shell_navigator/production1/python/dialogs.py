"""Dialog handlers for interactive prompts (passwords, confirmations)."""

from dataclasses import dataclass
import re


@dataclass
class DialogRule:
    pattern: str       # regex to match in output
    response: str      # what to send back
    is_secret: bool = False  # if True, don't log the response


# Common dialog rules
DEFAULT_DIALOGS = [
    DialogRule(r"[Pp]assword\s*:", "", is_secret=True),  # filled at runtime
    DialogRule(r"Are you sure\??\s*\(y/n\)", "y"),
    DialogRule(r"\[yes/no\]", "yes"),
    DialogRule(r"Continue\?", "y"),
    DialogRule(r"Press any key", ""),
]


class DialogHandler:
    def __init__(self, rules: list[DialogRule] = None, default_password: str = ""):
        self.rules = rules or DEFAULT_DIALOGS
        self.default_password = default_password

    def check(self, output: str) -> tuple[bool, str]:
        """Check if output contains a dialog prompt. Returns (matched, response)."""
        last_line = output.strip().split("\n")[-1] if output.strip() else ""
        for rule in self.rules:
            if re.search(rule.pattern, last_line):
                response = rule.response if rule.response else self.default_password
                return True, response
        return False, ""
