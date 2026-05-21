"""
Shell Navigator MVP Demo

Usage:
    python main.py [host] [username] [password] [port]

Defaults to localhost:22 with current user.
"""
import sys
from navigator import ShellNavigator


def main():
    host = sys.argv[1] if len(sys.argv) > 1 else "localhost"
    user = sys.argv[2] if len(sys.argv) > 2 else "user"
    pwd = sys.argv[3] if len(sys.argv) > 3 else "password"
    port = int(sys.argv[4]) if len(sys.argv) > 4 else 22

    nav = ShellNavigator(host, user, pwd, port)

    try:
        nav.connect()
        print("--- Current shell: EXEC ---")

        # Navigate to ROOT
        nav.navigate("ROOT")
        output = nav.execute("whoami")
        print(f"whoami output:\n{output}")

        # Back to EXEC
        nav.navigate("EXEC")
        output = nav.execute("echo 'back to exec'")
        print(f"echo output:\n{output}")

    except Exception as e:
        print(f"Error: {e}")
    finally:
        nav.disconnect()


if __name__ == "__main__":
    main()
