"""
Shell Navigator Production Demo

Usage:
    python main.py [host] [username] [password] [port]
"""
import sys
import logging
from navigator import ShellNavigator
from decorators import LoggingDecorator, TimingDecorator, CommandBlocker, SecretMasker, RetryDecorator

logging.basicConfig(level=logging.INFO, format="%(asctime)s %(name)s %(message)s")


def main():
    host = sys.argv[1] if len(sys.argv) > 1 else "localhost"
    user = sys.argv[2] if len(sys.argv) > 2 else "user"
    pwd = sys.argv[3] if len(sys.argv) > 3 else "password"
    port = int(sys.argv[4]) if len(sys.argv) > 4 else 22

    nav = ShellNavigator(host, user, pwd, port)
    nav.add_decorator(LoggingDecorator())
    nav.add_decorator(TimingDecorator())
    nav.add_decorator(CommandBlocker(blocked=["rm -rf /", "format", "mkfs"]))
    nav.add_decorator(SecretMasker())
    nav.add_decorator(RetryDecorator(max_retries=3))

    try:
        nav.connect()

        # Navigate EXEC → ROOT (auto BFS path)
        nav.navigate("ROOT")
        print(nav.execute("whoami"))

        # Navigate ROOT → CONF (auto BFS path)
        nav.navigate("CONF")
        print(nav.execute("show running-config"))

        # Navigate CONF → EXEC (auto BFS: CONF→ROOT→EXEC)
        nav.navigate("EXEC")
        print(nav.execute("echo done"))

        # This should be blocked:
        try:
            nav.execute("rm -rf /")
        except PermissionError as e:
            print(f"Blocked: {e}")

    except Exception as e:
        print(f"Error: {e}")
    finally:
        nav.disconnect()


if __name__ == "__main__":
    main()
