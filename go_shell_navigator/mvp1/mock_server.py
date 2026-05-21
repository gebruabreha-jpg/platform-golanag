"""
Mock SSH server simulating multi-shell environment (EXEC → ROOT → CONF).
Run in Terminal 1, then run main.py in Terminal 2.

Usage: python mock_server.py
Listens on :2222, credentials: user / password
"""
import socket
import threading
import paramiko
from pathlib import Path

HOST_KEY_FILE = "test_host_key"
if not Path(HOST_KEY_FILE).exists():
    paramiko.RSAKey.generate(2048).write_private_key_file(HOST_KEY_FILE)
HOST_KEY = paramiko.RSAKey(filename=HOST_KEY_FILE)


class MockServer(paramiko.ServerInterface):
    def __init__(self):
        self.shell_event = threading.Event()

    def check_channel_request(self, kind, chanid):
        return paramiko.OPEN_SUCCEEDED if kind == "session" else paramiko.OPEN_FAILED_ADMINISTRATIVELY_PROHIBITED

    def check_auth_password(self, username, password):
        if username == "user" and password == "password":
            return paramiko.AUTH_SUCCESSFUL
        return paramiko.AUTH_FAILED

    def get_allowed_auths(self, username):
        return "password"

    def check_channel_shell_request(self, channel):
        self.shell_event.set()
        return True

    def check_channel_pty_request(self, channel, term, width, height, pixelwidth, pixelheight, modes):
        return True

    def check_channel_env_request(self, channel, name, value):
        return True


def handle_shell(channel):
    shell = "EXEC"

    def prompt():
        return {"EXEC": "user$ ", "SHELL": "user$ ", "ROOT": "root# ", "CONF": "(config)# "}.get(shell, "$ ")

    channel.sendall(prompt().encode())
    line_buf = ""

    while True:
        try:
            data = channel.recv(1024)
            if not data:
                break
            line_buf += data.decode("utf-8", errors="replace")

            while "\n" in line_buf or "\r" in line_buf:
                # Find first line terminator
                idx = -1
                for i, c in enumerate(line_buf):
                    if c in "\r\n":
                        idx = i
                        break
                if idx == -1:
                    break

                cmd = line_buf[:idx].strip()
                # Skip past \r\n
                rest = line_buf[idx:]
                rest = rest.lstrip("\r\n")
                line_buf = rest

                if not cmd:
                    channel.sendall(f"\r\n{prompt()}".encode())
                    continue

                if cmd == "sudo su" and shell in ("EXEC", "SHELL"):
                    channel.sendall(b"\r\nPassword: ")
                    # Read password
                    pw_buf = ""
                    while "\n" not in pw_buf and "\r" not in pw_buf:
                        pw_data = channel.recv(1024)
                        if not pw_data:
                            return
                        pw_buf += pw_data.decode("utf-8", errors="replace")
                    pw = pw_buf.strip("\r\n").strip()
                    if pw == "password":
                        shell = "ROOT"
                        channel.sendall(f"\r\n{prompt()}".encode())
                    else:
                        channel.sendall(f"\r\nAuth failure\r\n{prompt()}".encode())

                elif cmd == "sh" and shell == "EXEC":
                    shell = "SHELL"
                    channel.sendall(f"\r\n{prompt()}".encode())

                elif cmd == "configure" and shell == "ROOT":
                    shell = "CONF"
                    channel.sendall(f"\r\nEntering config mode\r\n{prompt()}".encode())

                elif cmd == "exit":
                    transitions = {"CONF": "ROOT", "ROOT": "EXEC", "SHELL": "EXEC"}
                    if shell in transitions:
                        shell = transitions[shell]
                        channel.sendall(f"\r\n{prompt()}".encode())
                    else:
                        channel.sendall(b"\r\nlogout\r\n")
                        return

                elif cmd == "whoami":
                    user = "root" if shell in ("ROOT", "CONF") else "user"
                    channel.sendall(f"\r\n{user}\r\n{prompt()}".encode())

                else:
                    channel.sendall(f"\r\n{cmd}: executed\r\n{prompt()}".encode())

        except Exception as e:
            print(f"Shell error: {e}")
            break

    channel.close()


def handle_client(client_sock):
    transport = paramiko.Transport(client_sock)
    transport.add_server_key(HOST_KEY)
    server = MockServer()

    try:
        transport.start_server(server=server)
    except Exception as e:
        print(f"SSH negotiation failed: {e}")
        return

    channel = transport.accept(20)
    if channel is None:
        print("No channel opened")
        return

    if not server.shell_event.wait(10):
        print("Shell request timeout")
        return

    print(f"Shell session started")
    handle_shell(channel)
    print(f"Shell session ended")
    transport.close()


def main():
    sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    sock.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
    sock.bind(("0.0.0.0", 2222))
    sock.listen(5)
    print("Mock SSH server listening on :2222")
    print("Credentials: user / password")
    print("Press Ctrl+C to stop\n")

    try:
        while True:
            client, addr = sock.accept()
            print(f"Connection from {addr}")
            threading.Thread(target=handle_client, args=(client,), daemon=True).start()
    except KeyboardInterrupt:
        print("\nShutting down")
    finally:
        sock.close()


if __name__ == "__main__":
    main()
