!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
Linux Boot Process (BIOS/UEFI → GRUB → Kernel → initramfs → systemd)
Process lifecycle and scheduling
Virtual memory and paging
Page cache and buffer cache
CPU scheduler (CFS)
Filesystem internals (ext4/XFS)
LVM and RAID
Networking stack (TCP/IP, sockets, routing)
System calls and strace
/proc and /sys virtual filesystems
cgroups and namespaces (containers)
Docker and OverlayFS
SELinux
Kernel modules
Signals and process management
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
Week 1: CPU (top, vmstat, /proc/stat, load average)
Week 2: Memory (free, /proc/meminfo, swap, OOM)
Week 3: Storage (df, du, lsblk, iostat, /proc/diskstats)
Week 4: Networking (ip, ss, /proc/net, tcpdump)
Week 5: Processes (ps, top, strace, /proc/<pid>)
Week 6: Boot process, systemd, logging, and performance tuning
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
You/user
  │
  ▼
┌──────────────────┐
│ Terminal (app)   │   ← A window/application you interact with
└──────────────────┘
          │
          ▼
┌──────────────────┐
│ Shell (program)  │   ← Reads and interprets your commands
└──────────────────┘
          │
          ▼
┌──────────────────┐
│ Operating System │   ← Executes the requested actions
└──────────────────┘
          │
          ▼
      Output
          │
          ▼
┌──────────────────┐
│ Terminal (app)   │   ← Displays the output
└──────────────────┘


Learn Linux by layers:-

Instead of memorizing commands, understand this stack:

Applications
     │
Commands (top, ps, df, ip)
     │
System calls
     │
Kernel
     │
/proc      /sys
     │
Hardware



# Linux System Administration Learning Roadmap
---
# 1. CPU

## Commands

```bash
lscpu
cat /proc/cpuinfo
top
vmstat 1
sar -u
mpstat -P ALL
pidstat -u
```

## Understand

* CPU Architecture
* Physical CPU vs Logical CPU
* Core vs Thread
* Hyper-Threading
* CPU Scheduling
* Process Scheduler (CFS)
* Context Switch
* Interrupts (Hardware & Software)
* User Mode
* Kernel Mode
* System Calls
* CPU Utilization
* Load Average
* CPU Affinity
* CPU Steal Time
* Nice Value

## Kernel Files

```text
/proc/stat
/proc/cpuinfo
/proc/interrupts
/proc/loadavg
/proc/schedstat
/sys/devices/system/cpu/
```

---

# 2. Memory

## Commands

```bash
free -h
vmstat
cat /proc/meminfo
cat /proc/vmstat
cat /proc/swaps
swapon --show
```

## Understand

* Physical Memory (RAM)
* Virtual Memory
* Memory Allocation
* Memory Paging
* Page Cache
* Buffer Cache
* Anonymous Memory
* Shared Memory
* Slab Cache
* Huge Pages
* Swap
* OOM Killer
* Dirty Pages
* Memory Fragmentation

## Kernel Files

```text
/proc/meminfo
/proc/vmstat
/proc/swaps
/proc/zoneinfo
/proc/buddyinfo
/sys/kernel/mm/
```

---

# 3. Storage

## Commands

```bash
df -h
du -sh /*
lsblk
blkid
findmnt
mount
iostat -x
fdisk -l
```

## Understand

* Block Device
* Filesystem
* Inode
* Superblock
* Directory Entry
* ext4
* XFS
* LVM
* RAID
* Device Mapper
* Mount Point
* Page Cache
* Read/Write I/O
* Disk Queue
* Disk Latency
* Filesystem Journaling

## Kernel Files

```text
/proc/diskstats
/proc/mounts
/proc/filesystems
/sys/block/
```

---

# 4. Networking

## Commands

```bash
ip addr
ip route
ss -tulnp
ping
traceroute
tcpdump
ethtool
arp -a
```

## Understand

* OSI Model
* TCP
* UDP
* ICMP
* IP Address
* CIDR
* Subnet Mask
* Gateway
* Routing
* Routing Table
* ARP
* DNS
* MTU
* Socket
* Port
* Connection States
* Network Interface
* Bonding / Teaming
* VLAN

## Kernel Files

```text
/proc/net/
/proc/net/dev
/proc/net/tcp
/proc/net/udp
/proc/net/route
/sys/class/net/
```

---

# 5. Process Management

## Commands

```bash
ps -ef
top
htop
kill
killall
pkill
nice
renice
pstree
```

## Understand

* Process
* PID
* PPID
* Parent Process
* Child Process
* Zombie Process
* Orphan Process
* Daemon
* Thread
* Process States
* Scheduler
* Signals
* Foreground
* Background

## Kernel Files

```text
/proc/<PID>/
/proc/<PID>/status
/proc/<PID>/fd
/proc/<PID>/maps
/proc/<PID>/stat
```

---

# 6. Services

## Commands

```bash
systemctl
systemctl status
systemctl list-units
systemctl list-unit-files
journalctl
```

## Understand

* systemd
* Units
* Services
* Targets
* Timers
* Dependencies
* Boot Targets
* Enable vs Start
* Restart Policies

---

# 7. Logs

## Commands

```bash
journalctl
journalctl -xe
journalctl -k
dmesg -T
tail -f /var/log/messages
tail -f /var/log/secure
```

## Understand

* Kernel Logs
* System Logs
* Application Logs
* Authentication Logs
* Audit Logs
* Boot Logs
* Journal
* Log Rotation

---

# 8. Filesystem

## Commands

```bash
mount
findmnt
lsblk
blkid
df -Th
du -sh
```

## Understand

* Filesystem
* Partition
* Block Device
* Inode
* Superblock
* Mount Point
* UUID
* Label
* Journaling
* File Permissions

---

# 9. Security

## Commands

```bash
id
groups
whoami
sudo -l
getenforce
sestatus
getfacl
setfacl
```

## Understand

* Users
* Groups
* UID
* GID
* File Permissions
* chmod
* chown
* ACL
* SELinux
* Capabilities
* sudo
* PAM

---

# 10. Performance Monitoring

## Commands

```bash
uptime
top
vmstat 1
iostat -x 1
sar
pidstat
mpstat
iotop
```

## Understand

* CPU Bottleneck
* Memory Bottleneck
* Disk Bottleneck
* Network Bottleneck
* Load Average
* CPU Utilization
* I/O Wait
* Context Switches
* Paging
* Cache Hit Ratio
* Throughput
* Latency

---

# Performance Troubleshooting Checklist

## If the server is slow, check in this order:

```text
1. uptime                 # Check load average
2. top                    # Check CPU and memory
3. vmstat 1 5             # Check CPU, memory, swap
4. free -h                # Check memory
5. iostat -x 1 5          # Check disk I/O
6. sar                    # Check historical performance
7. ps -eo pid,%cpu,%mem,cmd --sort=-%cpu | head
8. pidstat -u 1           # Per-process CPU usage
9. ss -tulnp              # Check network sockets
10. journalctl -xe        # Review logs
11. dmesg -T              # Check kernel messages
12. df -h                 # Check disk space
13. lsblk                 # Verify storage devices
14. mount                 # Check mounted filesystems
15. systemctl --failed    # Check failed services
```

---

# Important Linux Virtual Filesystems

```text
/proc            Runtime kernel and process information
/sys             Kernel device and driver information
/dev             Device files
/run             Runtime state information
/etc             System configuration
/var/log         Persistent log files
```

---

# Master These Linux Topics

* Linux Boot Process
* GRUB
* Kernel Initialization
* initramfs
* systemd
* Process Scheduling
* Virtual Memory
* Paging
* Swap
* OOM Killer
* Filesystem Internals
* ext4
* XFS
* LVM
* RAID
* Networking Stack
* TCP/IP
* DNS
* ARP
* Routing
* iptables/nftables
* SELinux
* Kernel Modules
* cgroups
* Namespaces
* Containers (Docker/Podman)
* Performance Tuning
* System Calls (`strace`)
* `/proc` and `/sys` internals
