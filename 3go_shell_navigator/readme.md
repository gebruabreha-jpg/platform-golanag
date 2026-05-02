Data first, transport second, logic last.
Always Start From the Bottom Up:-
1. Define your DATA          → What are the things? (shells, states, transitions)
2. Build the TRANSPORT       → How do I talk to the system? (SSH, HTTP, gRPC)
3. Build the DETECTOR        → How do I know what happened? (prompt matching, response parsing)
4. Build the ENGINE          → How do I orchestrate? (state machine, navigator)
5. Add the EXTRAS            → Decorators, dialogs, retry, logging
6. Write the DEMO            → Wire it all up

Bottom = data model because it has zero dependencies. Everything else depends on it.

Step 1: DATA — "What are the things?"
Before writing any logic, answer these questions on paper:
What shells exist? → EXEC, SHELL, ROOT, CONF
How do I recognize each? → $ prompt, # prompt, (config)# prompt
How do I move between them? → sudo su, exit, configure, sh
That becomes shells.py. No logic, just definitions.

Step 2: TRANSPORT — "How do I talk to the system?"
Now you need a way to send commands and receive output. You don't care about shells yet — just:
Can I connect?
Can I send a string?
Can I read the response?
That becomes transport.py. Just a pipe — data in, data out.

Step 3: DETECTOR — "How do I know what happened?"
You can send commands now, but how do you know when the command finished? The SSH channel gives you a continuous stream of bytes. You need something that says "I see root# at the end — the command is done."
That becomes prompt_detector.py. Just reading + pattern matching.


Step 4: ENGINE — "How do I orchestrate?"
NOW you combine everything:
I know what shells exist (step 1)
I can send commands (step 2)
I can detect when they finish (step 3)
So I can navigate: send the right command, wait for the right prompt, update my state
That becomes navigator.py. The brain that uses all the parts.

Step 5: EXTRAS — "Make it production-ready"
The core works. Now add safety and observability:
Log every command → LoggingDecorator
Block dangerous commands → CommandBlocker
Retry on failure → RetryDecorator
Handle password prompts → DialogHandler
That becomes decorators.py + dialogs.py. Plugins that wrap the core.

Step 6: DEMO — "Prove it works"
Wire everything together and run it. That becomes main.py.
This is the same pattern everywhere:-
REST API: define models → build database layer → build routes → add middleware → write tests
Terraform: define variables → build resources → add outputs → add modules
Kubernetes: define pods → build services → add ingress → add monitoring



# Shell Navigator
A multi-shell automation library that navigates layered CLI environments (EXEC → SHELL → ROOT → CONF) over SSH with reliable prompt detection, state-machine transitions, dialog handling, and a decorator framework.

## Problem
Complex systems (telecom, network devices) expose multiple nested shell layers with different prompts (`$`, `#`, `config#`), interactive dialogs (passwords, confirmations), and no built-in way to programmatically navigate between them.

## Architecture
┌─────────────────────────────────────────────┐
│              Shell Navigator                │
├──────────┬──────────┬───────────┬───────────┤
│  Shell   │ Prompt   │Navigation │ Decorator │
│  Model   │ Detector │  Engine   │ Framework │
│ (enums,  │ (regex   │ (BFS/DFS  │ (logging, │
│  graph)  │  match)  │  state    │  security,│
│          │          │  machine) │  retry)   │
├──────────┴──────────┴───────────┴───────────┤
│           Transport Layer (SSH)             │
└─────────────────────────────────────────────┘

### Components
| Component                | Description |
| **Shell Model**          | Shell IDs (enum) + transition graph (enter/exit commands, expected prompts) |
| **Prompt Detector**      | Regex-based detection of current shell from output stream |
| **Navigation Engine**    | State machine that finds shortest path between shells via BFS |
| **Decorator Framework**  | Pluggable filters: logging, command blocking, auto-reconnect, retry |
| **Dialog Handler**       | Responds to interactive prompts (passwords, y/n confirmations) |

## Tech Stack
| Layer                                      | Technology .........................| Why                                |
| Core logic, SSH, prompt matching           | **Python** (paramiko, asyncio)      | Fast iteration, rich SSH ecosystem |
| Mock SSH server, concurrent decorators     | **Go** (x/crypto/ssh, goroutines)   | Speed, concurrency, low memory     |

## Project Structure
shell_navigator/
├── mvp/                    # MVP: Python-only, minimal
│   ├── shells.py           # Shell enums + transitions
│   ├── prompt_detector.py  # Regex prompt matcher
│   ├── navigator.py        # State machine navigation
│   ├── transport.py        # SSH connection wrapper
│   └── main.py             # Demo entry point
├── production/             # Production: Python + Go
│   ├── python/
│   │   ├── shells.py       # Extended shell model
│   │   ├── prompt_detector.py
│   │   ├── navigator.py    # BFS-based navigation engine
│   │   ├── transport.py    # SSH with reconnect
│   │   ├── decorators.py   # Decorator framework
│   │   ├── dialogs.py      # Dialog handlers
│   │   └── main.py
│   └── go/
│       └── mock_ssh/       # Mock SSH server for testing
│           └── main.go
├── requirements.txt
├── go.mod
└── readme.md
```

## Quick Start
### MVP

Terminal 1 — Start the mock server
python mock_server.py

Terminal 2 — Run the MVP client
python main.py localhost user password 2222

```bash
pip install paramiko
cd mvp
python main.py
```

### Production
```bash
pip install -r requirements.txt
cd production/python
python main.py

```

### Mock SSH Server (Go)
```bash
cd production/go/mock_ssh
go run main.go
# Starts a fake multi-shell SSH server on :2222
```

## Usage

### MVP
```python
from navigator import ShellNavigator

nav = ShellNavigator(host="localhost", username="user", password="pass")
nav.connect()
nav.navigate("ROOT")
output = nav.execute("show version")
nav.navigate("EXEC")
nav.disconnect()
```

### Production
```python
from navigator import ShellNavigator
from decorators import LoggingDecorator, CommandBlocker, RetryDecorator

nav = ShellNavigator(host="localhost", username="user", password="pass")
nav.add_decorator(LoggingDecorator())
nav.add_decorator(CommandBlocker(blocked=["rm -rf", "format"]))
nav.add_decorator(RetryDecorator(max_retries=3))

nav.connect()
nav.navigate("CONF")
output = nav.execute("set interface eth0 ip 10.0.0.1/24")
nav.navigate("EXEC")
nav.disconnect()
```

## Testing
Use any of these environments:
1. **Local machine** — `ssh localhost`, navigate bash → sh → sudo → root
2. **Docker** — `docker run -it ubuntu bash`, customize prompts
3. **Go mock server** — `cd production/go/mock_ssh && go run main.go` (simulates multi-shell SSH)
4. **GNS3 / EVE-NG** — simulate routers/switches with layered shells

## Roadmap
- [x] MVP: SSH + 2-3 shells + prompt matching + transitions
- [x] Production: decorators, dialog handlers, BFS navigation
- [x] Go mock SSH server for testing
- [ ] Plugin system for custom decorators
- [ ] Metrics/tracing (OpenTelemetry)
- [ ] Async navigation (asyncio)



MVP (Python only) — 4 files:
shells.py — ShellId enum (EXEC, SHELL, ROOT), transition definitions, prompt patterns, direct transition lookup
prompt_detector.py — reads SSH channel output, matches regex prompts with timeout
transport.py — thin paramiko wrapper (connect, send, disconnect)
navigator.py — state machine that ties it all together: connect → navigate(shell) → execute(cmd) → disconnect


Production (Python + Go) — adds:
shells.py — added CONF shell + BFS path finding so navigate("CONF") from EXEC auto-routes through ROOT
decorators.py — 5 decorators: LoggingDecorator, TimingDecorator, CommandBlocker, SecretMasker, RetryDecorator
dialogs.py — DialogHandler with rules for passwords, y/n confirmations, "Continue?" prompts
navigator.py — full decorator pipeline, auto-reconnect, retry logic, dialog handling mid-command
transport.py — added reconnect() and is_connected property
Go mock SSH server — simulates EXEC→ROOT→CONF shells on port 2222, handles sudo su with password dialog, configure, exit, whoami
How to test:
Start the Go mock server: cd production/go/mock_ssh && go run main.go
Run production demo: cd production/python && python main.py localhost user password 2222












The Problem
Paramiko 4.0 by default tries SSH key authentication first (looks for keys in ~/.ssh/, asks the SSH agent). The Go mock server didn't understand these key auth requests and dropped the connection before password auth was even attempted.

The Timeline
Step	What Happened
1. Python connects to Go server	✅ TCP + SSH handshake works
2. Paramiko tries key auth (automatic)	❌ Go server doesn't handle it, closes transport
3. Paramiko never gets to try password	❌ transport shut down or saw EOF
The Fix
# Before (broken)
c.connect(host, username=user, password=pwd)

# After (working)
c.connect(host, username=user, password=pwd,
          look_for_keys=False, allow_agent=False)

Copy
Two flags:

look_for_keys=False — don't scan ~/.ssh/ for key files

allow_agent=False — don't ask the SSH agent for keys

Lessons
Read the error carefully — transport shut down or saw EOF means the server killed the connection, not that the password was wrong. The auth never reached password stage.

Debug with verbose logging — adding logging.basicConfig(level=logging.DEBUG) showed us Authentication (password) successful! when we disabled key auth. That pinpointed the exact cause.

Client defaults matter — paramiko's default behavior (try keys first) is fine for real SSH servers like OpenSSH, but breaks against minimal/mock servers that only support password auth.

Test one layer at a time — we isolated the problem by testing raw paramiko connection separately (python -c "import paramiko...") instead of debugging through the full navigator stack.

Mixed-stack integration — when Python talks to Go (or any cross-language setup), don't assume defaults match. Each side has its own assumptions about protocol negotiation order.
















1. Architecture Patterns
State Machine — the core of the navigator. Every shell is a state, every command is a transition. This pattern appears everywhere:

Network protocol handlers

Game engines

Workflow engines

Kubernetes pod lifecycle (Pending → Running → Succeeded/Failed)

Decorator Pattern — you saw how LoggingDecorator, TimingDecorator, CommandBlocker all wrap the same execute() without modifying it. This is how real middleware works:

Express.js middleware

Python WSGI middleware

Java servlet filters

Kubernetes admission controllers

Graph + BFS — the production navigator finds the shortest path between any two shells. Same algorithm used in:

Network routing (OSPF)

Kubernetes service mesh routing

Dependency resolution (terraform, npm, pip)

2. Coding Patterns
Separation of concerns — each file does one thing:

transport.py  → only SSH connection
shells.py     → only shell definitions + graph
prompt_detector.py → only output reading
navigator.py  → only orchestration
decorators.py → only filters

Copy
This is why you could swap the Go server without touching any Python navigator code.

Pipeline pattern — decorators run in order:

command → LoggingDecorator → TimingDecorator → CommandBlocker → execute → reverse order

Copy
Same pattern as: CI/CD pipelines, HTTP middleware chains, Unix pipes.

Fail early — CommandBlocker rejects rm -rf / before it reaches the SSH channel. Always validate input before expensive operations.

3. Debugging Methodology
What you actually practiced:

1. Read the error message    → "transport shut down or saw EOF"
2. Isolate the layer         → test raw paramiko, not full navigator
3. Add observability         → DEBUG logging showed auth sequence
4. Fix the root cause        → look_for_keys=False
5. Verify the fix            → run again, works

Copy
This is the exact same layered debugging approach from your today1.md Kubernetes section (LB → Ingress → Service → Pod → App).

4. Cross-Language Integration
Protocol is the contract — Python and Go don't share code, they share SSH protocol. This is how real systems work:

Microservices communicate via HTTP/gRPC, not shared libraries

Your Python navigator doesn't care if the server is Go, Java, or a real Cisco router

The Go server doesn't care if the client is Python or PuTTY

Never assume defaults match — paramiko tries key auth first, Go server only supports password. Same class of bug happens with:

TLS version mismatches

JSON date format differences

Character encoding (UTF-8 vs Latin-1)

5. Concepts to Study Deeper
Concept	Where You Saw It	Where It's Used in Production
State machine	Shell navigation	K8s controllers, workflow engines, protocol handlers
Decorator pattern	Logging/blocking/retry	Middleware in every web framework
BFS graph traversal	find_path()	Network routing, dependency resolution
Prompt detection (regex)	wait_for_prompt()	Log parsing, monitoring, alerting
Transport abstraction	SSHTransport class	Database drivers, message queues, cloud SDKs
Idempotent operations	navigate() — calling twice is safe	Terraform apply, K8s reconciliation loops
Retry with backoff	RetryDecorator	Every cloud SDK, circuit breakers



What Would Make This Production-Ready:-
Things still missing that real infra tools have:
Tests — unit tests for each component, integration test with mock server
Config file — shell definitions in YAML instead of hardcoded Python
Async — asyncio for handling multiple devices concurrently
Metrics — Prometheus counters for transitions, failures, latency
Plugin system — let users add custom shells/decorators without modifying core code
The biggest takeaway: you built a real infrastructure automation tool using the same patterns (state machines, decorators, graph traversal, protocol abstraction) that Ansible, Terraform, and Kubernetes use internally.

server side log:-
2026/04/10 13:43:00 Host key type: ecdsa-sha2-nistp256
2026/04/10 13:43:00 Mock SSH server listening on :2222
2026/04/10 13:43:00 Credentials: user / password
2026/04/10 13:43:19 Handshake/auth failed: ssh: unexpected message type 5 (expected one of [50])
2026/04/10 13:44:34 Auth attempt: user=user
2026/04/10 13:44:34 Auth SUCCESS
2026/04/10 13:44:34 Connected: [::1]:64851 (user: user)
2026/04/10 13:45:02 Auth attempt: user=user
2026/04/10 13:45:02 Auth SUCCESS
2026/04/10 13:45:02 Connected: [::1]:53838 (user: user)
2026/04/10 13:45:02 [EXEC] cmd: sudo su
2026/04/10 13:45:02 [ROOT] cmd: whoami
2026/04/10 13:45:02 [ROOT] cmd: configure
2026/04/10 13:45:02 [CONF] cmd: show running-config
2026/04/10 13:45:02 [CONF] cmd: exit
2026/04/10 13:45:02 [ROOT] cmd: exit
2026/04/10 13:45:02 [EXEC] cmd: echo done