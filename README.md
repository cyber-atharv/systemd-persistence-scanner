# 🛡️ Linux Persistence Mechanism & Backdoor Scanner (`sentinel`)

> A comprehensive, single-binary Linux persistence scanner written in Go that hunts down stealthy backdoors, malicious cron jobs, hooked systemd services, rogue SSH authorized keys, and poisoned shell profiles.

[![Author](https://img.shields.io/badge/Made%20by-cyber--atharv-00ffcc?style=flat-square&logo=github)](https://github.com/cyber-atharv)
[![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat-square&logo=go&logoColor=white)](https://go.dev)
[![MITRE ATT&CK](https://img.shields.io/badge/MITRE-TA0003_Persistence-orange?style=flat-square)](https://attack.mitre.org/tactics/TA0003/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg?style=flat-square)](LICENSE)

---

## 📌 What is Linux Persistence?

When an attacker breaches a Linux system, their next objective is **Persistence (MITRE ATT&CK TA0003)**: ensuring they maintain access even after a system reboot or password change.

Attackers rarely leave backdoors out in the open. Instead, they hide malicious triggers inside:
- Scheduled Cron jobs & Systemd timers
- Poisoned shell profiles (`~/.bashrc`, `/etc/profile.d/`)
- Rogue SSH authorized keys with forced commands
- Dynamic linker hijacks (`/etc/ld.so.preload`)
- Kernel modules and udev hardware rules

This scanner, developed by **cyber-atharv**, inspects **12+ Linux persistence vectors** and applies heuristic pattern matching (detecting reverse shells, `curl | bash` pipelines, base64 payloads) to flag backdoors immediately.

---

## ✨ Monitored Persistence Vectors

| Vector | What Sentinel Checks |
|---|---|
| **🕒 Cron & At Jobs** | `/etc/crontab`, `/etc/cron.*`, user crontabs for hidden reverse shells or suspicious download chains. |
| **⚙️ Systemd Units** | Unit files and timers in `/etc/systemd/system/` executing binaries from `/tmp`, `/dev/shm`, or hidden folders. |
| **🐚 Shell Profiles** | User `.bashrc`, `.profile`, `.zshrc`, and `/etc/profile.d/` for alias hijacking or command injections. |
| **🔑 SSH Keys & Configs** | `authorized_keys` with `command=` forced execution and non-standard `AuthorizedKeysFile` paths. |
| **🔗 Dynamic Linker** | `/etc/ld.so.preload` and `/etc/ld.so.conf.d/` for stealth user-space rootkits. |
| **🧩 PAM & Udev Rules** | Backdoored authentication modules and hardware trigger rules. |

---

## 🚀 Quick Start & Usage

### 1. Build the Static Binary
```bash
cd systemd-persistence-scanner
go build -o sentinel ./cmd/sentinel
```

### 2. Examples

#### 🔹 Run a full system persistence scan
```bash
sudo ./sentinel scan
```

#### 🔹 Show only High & Critical threats
```bash
sudo ./sentinel scan --min-severity high
```

#### 🔹 Scan a mounted forensic drive / image
```bash
./sentinel scan --root /mnt/target_evidence_drive
```

#### 🔹 Baseline & Drift Detection
```bash
# Save a clean snapshot of the system
sudo ./sentinel baseline save

# Compare current system against clean baseline
sudo ./sentinel baseline diff
```

---

## 🧠 Why I Built This

Incident response and digital forensics require deep knowledge of how the Linux operating system handles startup sequences, user sessions, and service management. Building Sentinel helped me understand real-world APT persistence tradecraft and how to write reliable forensic parsers for system files.

---

## 📜 Author & License

- **Author:** [cyber-atharv](https://github.com/cyber-atharv)
- **License:** Open source under the MIT / AGPL License.
