# Linux CPU Interview Summary

## Where does `lscpu` get its information?

`lscpu` does not communicate directly with the CPU. It reads information exposed by the Linux kernel from:

```bash
/proc/cpuinfo
/sys/devices/system/cpu/
```

You can view the raw information yourself:

```bash
cat /proc/cpuinfo
ls /sys/devices/system/cpu/
```

---

## CPU Topology of This Server

```
Socket(s):             2
Core(s) per socket:    24
Thread(s) per core:    2
CPU(s):                96
```

### Physical CPUs (Sockets)

* **2 sockets**
* Each socket is one physical CPU installed on the motherboard.

### Physical Cores

```
24 cores/socket × 2 sockets = 48 physical cores
```

A **core** is a real hardware processing unit capable of executing instructions independently.

### Logical CPUs

```
48 physical cores × 2 threads/core = 96 logical CPUs
```

Linux sees **96 logical CPUs** and schedules processes on these logical CPUs.

---

## Hyper-Threading (SMT)

Hyper-Threading (Intel's Simultaneous Multithreading - SMT) allows **one physical core to appear as two logical CPUs**.

Example:

```
1 Physical Core
      │
 ┌────┴────┐
 │         │
Thread 0  Thread 1
```

Benefits:

* Better CPU utilization
* More efficient execution when one thread is waiting for memory or I/O
* Does **not** double performance

---

## NUMA (Non-Uniform Memory Access)

NUMA divides a large server into multiple CPU and memory nodes.

```
Socket 1  <--> Local Memory
Socket 2  <--> Local Memory
```

A CPU accesses:

* **Local memory** → Faster
* **Remote memory** → Slower

Linux tries to schedule processes on CPUs that are close to the memory they use to improve performance.

---

## CPU Cache

CPU cache is very fast memory located inside or near the CPU.

```
CPU Registers
      ↓
L1 Cache   (Fastest, Smallest)
      ↓
L2 Cache
      ↓
L3 Cache   (Shared by cores)
      ↓
RAM
      ↓
Disk (Slowest)
```

Why cache is important:

* Stores frequently used data and instructions
* Reduces RAM accesses
* Improves application performance
* Lowers CPU wait time

---

## Key Interview Questions

### Where does `lscpu` get its information?

```
/proc/cpuinfo
/sys/devices/system/cpu/
```

---

### How many physical cores does this server have?

```
24 cores/socket × 2 sockets = 48 physical cores
```

---

### How many logical CPUs does Linux see?

```
96 logical CPUs
```

---

### What is Hyper-Threading?

Intel's Simultaneous Multithreading (SMT), where one physical core provides two logical CPUs for better CPU utilization.

---

### What is NUMA?

A memory architecture where each CPU has faster access to its own local memory than to memory attached to another CPU.

---

### Why is CPU cache important?

Cache stores frequently accessed data close to the CPU, reducing memory access time and significantly improving performance.

---

## Useful Commands

```bash
# CPU information
lscpu

# Raw CPU information
cat /proc/cpuinfo

# Logical CPUs
cat /proc/cpuinfo | grep "^processor"

# CPU topology
lscpu -e

# CPU statistics
cat /proc/stat

# NUMA information
numactl --hardware

# CPU online status
ls /sys/devices/system/cpu/
```
A NUMA node is not a virtual machine or a software-only object. It is a logical grouping created by the hardware architecture that contains:
A set of CPU cores (logical CPUs)
A region of local RAM

NUMA (Non-Uniform Memory Access) is a hardware memory architecture where the system is divided into multiple NUMA nodes. Each NUMA node consists of a group of CPUs and its own local memory. Accessing local memory is faster than accessing memory attached to another NUMA node, so Linux tries to schedule processes and allocate memory within the same NUMA node for better performance.

Interview Summary

Q: How many NUMA nodes does this server have?

4 NUMA nodes.

Q: Why are there 4 NUMA nodes if there are only 2 CPU sockets?

Each physical socket is internally divided into 2 NUMA domains, so the operating system sees 4 NUMA nodes.

Q: What does a NUMA node contain?

A NUMA node contains a set of logical CPUs and a region of local memory (RAM).

Q: What do the NUMA distance values mean?

They represent the relative cost of accessing memory. A smaller value means faster access. Local memory has the lowest cost (10), while remote memory has higher costs (20 or 30).

Understanding NUMA is especially important for high-performance workloads such as databases, virtualization, HPC, and large Java applications, where keeping CPU execution and memory allocation on the same NUMA node can significantly improve performance.
# Show NUMA topology
numactl --hardware

# Display NUMA statistics
numastat

# Show NUMA policy for a process
numactl --show

# Display CPU and NUMA mapping
lscpu -e

# See NUMA information in sysfs
ls /sys/devices/system/node/