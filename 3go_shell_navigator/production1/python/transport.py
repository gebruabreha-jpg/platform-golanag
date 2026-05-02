import paramiko
import time


class SSHTransport:
    def __init__(self, host: str, username: str, password: str, port: int = 22):
        self.host = host
        self.port = port
        self.username = username
        self.password = password
        self.client: paramiko.SSHClient | None = None
        self.channel = None

    def connect(self):
        self.client = paramiko.SSHClient()
        self.client.set_missing_host_key_policy(paramiko.AutoAddPolicy())
        self.client.connect(self.host, port=self.port,
                            username=self.username, password=self.password,
                            look_for_keys=False, allow_agent=False)
        self.channel = self.client.invoke_shell()
        return self.channel

    def reconnect(self, max_retries: int = 3, delay: float = 2.0):
        for attempt in range(1, max_retries + 1):
            try:
                self.disconnect()
                self.connect()
                return self.channel
            except Exception as e:
                if attempt == max_retries:
                    raise
                time.sleep(delay)

    @property
    def is_connected(self) -> bool:
        return (self.client is not None and
                self.client.get_transport() is not None and
                self.client.get_transport().is_active())

    def send(self, command: str):
        if not self.channel:
            raise RuntimeError("Not connected")
        self.channel.send(command + "\n")

    def disconnect(self):
        if self.channel:
            self.channel.close()
        if self.client:
            self.client.close()
        self.channel = None
        self.client = None
