# Linux Storage — Top-Down Guide & Interview Prep

## 1. The Big Picture (Layers, top to bottom)
start with RAID and LVM concepts and performance monitoring
reat basic mount/fstab stuff as assumed knowledge.

Filesystem basics + mounting (df, du, mount, /etc/fstab) 
Inodes vs blocks, disk-full troubleshooting (lsof +L1) — classic interview trap question
Partitioning + device naming (lsblk, fdisk) — now the earlier commands make more sense
LVM — the #1 topic interviewers probe deeply, so give it the most hands-on time
RAID — conceptual understanding + mdadm basics
Monitoring/performance (iostat, iotop) — usually a "nice to have" unless it's a senior/SRE role
Swap, permissions — quick refreshers

```
Applications
    |
File System (ext4, XFS, Btrfs, ZFS, etc.)
    |
Volume Manager (LVM) / Software RAID (mdadm)  [optional layer]
    |
Block Device Layer (/dev/sda, /dev/nvme0n1, etc.)
    |
Device Drivers (SCSI, NVMe, SATA)
    |
Physical Storage (HDD, SSD, NVMe, network storage)
```

Everything in Linux is a file, including storage devices — they show up under `/dev`.

---

## 2. Device Naming

| Device Type | Naming Pattern | Example |
|---|---|---|
| SATA/SCSI/USB disks | `/dev/sdX` | `/dev/sda`, `/dev/sdb1` (partition 1) |
| NVMe SSDs | `/dev/nvmeXnYpZ` | `/dev/nvme0n1p1` |
| Virtual disks (KVM) | `/dev/vdX` | `/dev/vda1` |
| LVM logical volumes | `/dev/mapper/VG-LV` | `/dev/mapper/vg0-root` |
| Loopback devices | `/dev/loopX` | `/dev/loop0` |

- Letters = physical disk (`sda`, `sdb`)
- Numbers = partition on that disk (`sda1`, `sda2`)

---

## 3. Partitioning

Two schemes:
- **MBR (Master Boot Record)** — legacy, max 2TB disks, 4 primary partitions
- **GPT (GUID Partition Table)** — modern standard, supports huge disks, up to 128 partitions

Tools:
- `fdisk` — MBR/GPT partition editor (interactive)
- `parted` / `gparted` — modern, scriptable, GUI available
- `lsblk` — list block devices and partitions
- `blkid` — show UUIDs and filesystem types

```bash
lsblk -f
sudo fdisk -l
sudo parted /dev/sda print
```

---

## 4. Filesystems

| Filesystem | Notes |
|---|---|
| **ext4** | Default on most distros. Journaling, mature, reliable |
| **XFS** | Great for large files, high performance, used by RHEL default |
| **Btrfs** | Copy-on-write, snapshots, built-in RAID, checksums |
| **ZFS** | Enterprise-grade, snapshots, dedup, checksums (not native kernel due to licensing) |
| **FAT32/exFAT** | Cross-platform compatibility, no journaling |
| **NTFS** | Windows compatibility |
| **tmpfs** | RAM-backed, temporary |
| **swap** | Not a "filesystem" for files — used for virtual memory |

Creating and checking filesystems:
```bash
sudo mkfs.ext4 /dev/sdb1
sudo mkfs.xfs /dev/sdb1
sudo fsck /dev/sdb1          # check/repair
sudo tune2fs -l /dev/sdb1    # ext4 metadata
```

---

## 5. Mounting

```bash
sudo mount /dev/sdb1 /mnt/data
sudo umount /mnt/data
```

Persistent mounts go in `/etc/fstab`:
```
UUID=xxxx-xxxx  /mnt/data  ext4  defaults  0  2
```
Columns: device — mountpoint — fstype — options — dump — fsck order

Use `mount -a` to test fstab without rebooting.

---

## 6. LVM (Logical Volume Manager)

Adds flexibility: resize volumes, span multiple disks, take snapshots.

```
Physical Volume (PV) -> Volume Group (VG) -> Logical Volume (LV)
```

```bash
sudo pvcreate /dev/sdb1
sudo vgcreate vg0 /dev/sdb1
sudo lvcreate -L 20G -n lv_data vg0
sudo mkfs.ext4 /dev/vg0/lv_data

# Resizing (online, ext4/xfs)
sudo lvextend -L +10G /dev/vg0/lv_data
sudo resize2fs /dev/vg0/lv_data      # for ext4
sudo xfs_growfs /mnt/data            # for xfs
```

**Interview point:** LVM snapshots are copy-on-write — good for backups, but can fill up if not sized correctly.

---

## 7. RAID (Software — mdadm)

| Level | Description | Fault Tolerance |
|---|---|---|
| RAID 0 | Striping | None (performance only) |
| RAID 1 | Mirroring | 1 disk failure |
| RAID 5 | Striping + parity | 1 disk failure |
| RAID 6 | Striping + double parity | 2 disk failures |
| RAID 10 | Mirror + stripe | Depends on layout |

```bash
sudo mdadm --create /dev/md0 --level=1 --raid-devices=2 /dev/sdb1 /dev/sdc1
cat /proc/mdstat
sudo mdadm --detail /dev/md0
```

---

## 8. Disk Usage & Monitoring

```bash
df -h              # filesystem-level free space
du -sh /path        # directory size
du -h --max-depth=1 /var | sort -rh   # find big directories
lsblk               # block device tree
iostat -x 1         # I/O performance stats
iotop               # per-process I/O usage (like top for disk)
smartctl -a /dev/sda # disk health (SMART data)
```

**Interview classic:** "`df` says disk is full, but `du` doesn't show why?"
→ Usually a **deleted-but-open file** still held by a process. Find it with:
```bash
sudo lsof +L1
```
Killing/restarting the process (or `> /proc/<pid>/fd/<fd>`) frees the space.

---

## 9. Swap

```bash
sudo fallocate -l 2G /swapfile
sudo chmod 600 /swapfile
sudo mkswap /swapfile
sudo swapon /swapfile
swapon --show
free -h
```
Add to `/etc/fstab` for persistence.

**Interview point:** know `vm.swappiness` — controls kernel's preference to swap (0–100, lower = avoid swap more).

---

## 10. Network / Distributed Storage

- **NFS** — Network File System, mount remote dirs (`mount -t nfs server:/share /mnt`)
- **Samba/CIFS** — Windows-compatible file sharing
- **iSCSI** — block-level storage over IP (looks like a local disk)
- **Ceph / GlusterFS** — distributed storage clusters

---

## 11. Permissions & Ownership (quick refresher)

```bash
chmod 755 file
chown user:group file
ls -l               # rwxr-xr-x etc.
```
- `r=4, w=2, x=1`
- Special bits: SUID (4000), SGID (2000), Sticky bit (1000, common on `/tmp`)

---

## 12. Common Interview Questions & Answers

**Q: What's the difference between a hard link and a symbolic link?**
- Hard link: another name pointing to the same inode; can't cross filesystems; deleting original doesn't remove data.
- Symlink: a separate file containing a path; can cross filesystems; breaks if target is deleted.

**Q: How do you find what's using the most disk space?**
```bash
du -h --max-depth=1 / | sort -rh | head -20
```
Or use `ncdu` for an interactive view.

**Q: Explain inodes.**
Inodes store metadata (permissions, owner, size, timestamps, pointers to data blocks) — not the filename. Running out of inodes (even with free space) causes "no space left on device."
```bash
df -i
```

**Q: How would you extend a root filesystem without downtime?**
If on LVM: `lvextend` + `resize2fs`/`xfs_growfs` while mounted (online resize supported by ext4/xfs).

**Q: Difference between RAID 5 and RAID 6?**
RAID 6 uses double parity, tolerating 2 disk failures vs. RAID 5's 1 — costs more usable capacity but safer, especially important as disks get larger (rebuild time risk).

**Q: What happens during a `fsck`?**
Checks filesystem consistency: verifies inode/block bitmaps, fixes orphaned inodes, corrects superblock issues. Should be run on unmounted filesystems (or read-only mounted) to avoid corruption.

**Q: What's the difference between mount options `noatime` and `relatime`?**
- `atime`: updates access time on every read (I/O overhead).
- `noatime`: never updates it (best performance).
- `relatime`: default on most modern systems — only updates atime if mtime/ctime changed or previous atime is older than a day.

**Q: How do you troubleshoot slow disk I/O?**
```bash
iostat -x 1
iotop
vmstat 1
```
Look at `%util`, `await`, `svctm` in iostat; identify if bottleneck is a specific process or disk saturation.

**Q: What's the difference between a partition, volume group, and logical volume?**
Partition = physical disk slice; Volume Group = pool combining PVs; Logical Volume = flexible "virtual partition" carved from the VG — resizable without repartitioning.

**Q: Journaling filesystems — why do they matter?**
They log changes before committing them, so after a crash the system can replay/roll back the journal instead of running a full fsck — much faster recovery.

**Q: How do you securely wipe a disk?**
```bash
sudo shred -v -n 3 /dev/sdX      # overwrite HDD
sudo blkdiscard /dev/nvme0n1     # SSD (TRIM-based)
```

---

## 13. Quick Command Cheat Sheet

| Task | Command |
|---|---|
| List block devices | `lsblk` |
| Disk free space | `df -h` |
| Directory size | `du -sh` |
| Partition disk | `fdisk` / `parted` |
| Format filesystem | `mkfs.ext4`, `mkfs.xfs` |
| Mount/unmount | `mount`, `umount` |
| Persistent mounts | `/etc/fstab` |
| Check/repair FS | `fsck` |
| LVM info | `pvs`, `vgs`, `lvs` |
| RAID status | `cat /proc/mdstat` |
| I/O stats | `iostat -x`, `iotop` |
| Disk health | `smartctl -a /dev/sdX` |
| Swap status | `free -h`, `swapon --show` |

---

## 14. Suggested Study Path

1. Understand block devices → partitions → filesystems → mount points (the stack).
2. Practice `lsblk`, `df`, `du`, `fdisk` on a test VM.
3. Set up LVM from scratch (PV → VG → LV → resize).
4. Build a software RAID 1 with `mdadm` and simulate a disk failure.
5. Diagnose a full disk caused by a deleted-but-open file (`lsof +L1`).
6. Review inodes vs. blocks (`df -i` vs `df -h`).
7. Time yourself explaining the layered storage stack out loud — this is the #1 "walk me through X" interview question.





"A user says they can't create a new file, but df -h shows plenty of free space. What do you check?"
→ df -i — if IUse% is at 100%, you've run out of inodes, not space. Common cause: some process spinning off tons of tiny files (a caching bug, a runaway logger, a build system leaving temp files).


Data blocks — where the actual file content lives (this is what df -h tracks/Space (bytes))
An inode — a metadata record (permissions, owner, timestamps, pointers to data blocks df -i/Inode(count)).
Every single file, directory, symlink — even a 0-byte file — consumes exactly one inode.

Inodes store metadata (permissions, owner, size, timestamps, pointers to data blocks) — not the filename. Running out of inodes (even with free space) causes "no space left on device."

What's the difference between a partition, volume group, and logical volume?
Partition = physical disk slice; Volume Group = pool combining PVs; Logical Volume = flexible "virtual partition" carved from the VG — resizable without repartitioning.


What happens during a fsck?
Checks filesystem consistency: verifies inode/block bitmaps, fixes orphaned inodes, corrects superblock issues. Should be run on unmounted filesystems (or read-only mounted) to avoid corruption.


What's the difference between a hard link and a symbolic link?
Hard link: another name pointing to the same inode; can't cross filesystems; deleting original doesn't remove data.
Symlink: a separate file containing a path; can cross filesystems; breaks if target is deleted.


Network / Distributed Storage:-
NFS — Network File System, mount remote dirs (mount -t nfs server:/share /mnt)
Samba/CIFS — Windows-compatible file sharing
iSCSI — block-level storage over IP (looks like a local disk)
Ceph / GlusterFS — distributed storage clusters



Here's a tight list of the questions that come up most in real Linux storage interviews — ordered roughly by how often they get asked, using your own server as a mental reference where it helps.
Tier 1 — Almost guaranteed
1. Walk me through how you'd investigate a full disk on a server you've never seen.
df -hT → lsblk -f → du -h --max-depth=1 drilling into the biggest mount → check df -i too → check for deleted-but-open files with lsof +L1.
2. df -h shows plenty of space but I can't create a file — why?
Out of inodes (df -i), not space. Fixed inode count on ext4; dynamic but still limited on xfs.
3. df says disk is full, du doesn't show why — explain.
A process is holding a file open that's already been deleted (unlinked but not released). Space isn't freed until the process closes the file handle or restarts. Find with lsof +L1 or lsof | grep deleted.
4. What's an inode? What does it store, and what does it NOT store?
Stores metadata: permissions, owner, timestamps, size, pointers to data blocks. Does not store the filename — filenames live in the parent directory's entries, which map name → inode number.
5. Hard link vs symbolic link.
Hard link = another directory entry pointing to the same inode, same filesystem only, survives deletion of the "original" name. Symlink = a small file containing a path, can cross filesystems, breaks if target moves/deletes.
6. Explain the LVM stack: PV → VG → LV.
Physical Volume (raw disk/partition) → pooled into a Volume Group → carved into Logical Volumes. Point of it: resize without repartitioning, span multiple disks, take snapshots.
7. How do you extend a filesystem with no downtime?
lvextend -L +XG /dev/vg/lv then resize2fs (ext4) or xfs_growfs (xfs, mounted filesystem, online-only — you can't shrink xfs).
8. Why is /boot often a separate, non-LVM partition?
Bootloader (GRUB) may not reliably read LVM or certain filesystems at boot time — keeping /boot as plain ext4 avoids chicken-and-egg boot failures. (You saw this exact layout on your own server.)
Tier 2 — Very common
9. RAID 5 vs RAID 6 — why would you pick one over the other?
RAID 6 = double parity, survives 2 disk failures, costs more usable capacity. Matters more as disk sizes grow, because RAID 5 rebuild time creates a real window where a second failure loses everything.
10. What happens during fsck? When should you run it?
Checks/repairs filesystem consistency — bitmaps, orphaned inodes, superblock. Should run unmounted (or read-only) to avoid making things worse.
11. Why do journaling filesystems recover faster after a crash?
They log intended changes before committing; after a crash they replay the journal instead of scanning the whole filesystem.
12. noatime vs relatime vs default atime — and why does it matter?
atime updates on every read = extra writes for a read-only op. relatime (modern default) only updates if mtime/ctime changed or the existing atime is >1 day old. noatime disables it entirely — best performance, sometimes breaks tools that depend on access time (rare).
13. Explain the difference between /dev/sda, /dev/sda1, /dev/mapper/vg-lv.
Whole disk → partition on that disk → LVM logical volume (a virtual block device sitting on top of a PV that lives on a partition). Good chance to demo you understand the layering, like we traced on your server (sda3 → rootvg-root → /).
14. How would you identify what's causing high disk I/O right now?
iostat -x 1 — look at %util (near 100% = saturated), await (latency), then iotop to find the specific process.
Tier 3 — Good to have, shows depth
15. tmpfs — what is it, and what's the risk of using it heavily?
RAM-backed filesystem, extremely fast, but contents vanish on reboot and heavy usage eats into actual memory (you saw this directly: /tmp as tmpfs using 127G on your server).
16. NFS vs local disk — tradeoffs?
NFS centralizes storage, easy to share across many hosts, but adds network latency/dependency and a single point of failure if the NFS server goes down. Local disk is fast and independent but doesn't scale across a fleet or share easily.
17. What's a copy-on-write filesystem, and name one.
Writes go to new blocks instead of overwriting in place; old blocks stay until no longer referenced — enables cheap snapshots. Examples: Btrfs, ZFS. (XFS/ext4 are not COW by default.)
18. Software RAID (mdadm) vs hardware RAID — pros/cons?
Software RAID: no extra hardware cost, flexible, but uses CPU cycles and is tied to that OS/kernel. Hardware RAID: dedicated controller handles it, faster, OS-independent, but costs money and is a black box if the controller fails.
19. What's swappiness, and how would you tune it on a memory-heavy server?
vm.swappiness (0–100) controls how aggressively the kernel swaps out memory pages. On a high-RAM server like the one you're looking at (1TB RAM, swap barely touched), you'd typically lower it (e.g., 10) to avoid unnecessary swapping when plenty of RAM is free.
20. GPT vs MBR partitioning.
MBR: legacy, 2TB disk limit, max 4 primary partitions. GPT: modern, huge disk support, up to 128 partitions, required for UEFI boot (which is why your server has that /boot/efi vfat partition).