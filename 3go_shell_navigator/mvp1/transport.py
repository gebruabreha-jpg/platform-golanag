import paramiko


class SSHTransport:
    """Thin wrapper around paramiko SSH + interactive shell channel."""

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
