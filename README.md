# ⚡ Sniphawk — Advanced Packet Sniffer in Go 🐍

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/ZeroHack01/sniphawk)](https://goreportcard.com/report/github.com/ZeroHack01/sniphawk)
[![GitHub issues](https://img.shields.io/github/issues/ZeroHack01/sniphawk)](https://github.com/ZeroHack01/sniphawk/issues)
[![Go](https://img.shields.io/badge/Go-1.18+-00ADD8?style=flat&logo=go)](https://golang.org/dl/)
[![Last Commit](https://img.shields.io/github/last-commit/ZeroHack01/sniphawk?color=orange)](https://github.com/ZeroHack01/sniphawk/commits/main)

<p align="center">
  <b>A blazing fast network packet analyzer written in Go</b><br>
  <sub>Built for the CodeAlpha Cybersecurity Internship</sub>
</p
  
---

## 🚀 Overview

**Sniphawk** is a powerful, real-time network packet sniffer built in **Go**, designed to capture, analyze, and display live network traffic on any specified interface. This project showcases practical skills in network analysis and Go programming while providing security professionals with an intuitive tool for traffic inspection.

---

## ✨ Features

- 🕵️‍♂️ Capture **TCP**, **UDP**, and **ICMP** packets in real-time
- 🔍 Intelligent filtering by **port** and **protocol**
- 🚫 Option to skip encrypted traffic (e.g., HTTPS on port 443)
- 📊 Real-time summary of **top IPs** and **most active ports** every 30 seconds
- 🎨 Colorful CLI output with a professional banner
- 🧩 Minimal dependencies for easy setup and portability

---

## 🎯 Installation

### Prerequisites

1. [Go](https://golang.org/dl/) installed (version 1.18+ recommended)
2. **libpcap** development files:
   - Ubuntu/Debian: `sudo apt-get install libpcap-dev`
   - CentOS/RHEL: `sudo yum install libpcap-devel`
   - macOS: `brew install libpcap`
   - Windows: Install [Npcap](https://npcap.com/#download) with development files

### Installation Steps

```bash
# Clone the repository
git clone https://github.com/ZeroHack01/sniphawk.git

# Navigate to project directory
cd sniphawk

# Install required Go packages
go get github.com/google/gopacket
go get github.com/google/gopacket/pcap
go get github.com/fatih/color

# Build the binary
go build -o sniphawk sniphawk.go

# Or alternatively with go mod:
go mod init github.com/ZeroHack01/sniphawk
go mod tidy
go build
```

### Troubleshooting Installation

If you encounter errors like `pcap.h: No such file or directory`:
- Make sure libpcap development files are installed
- On Windows, ensure Npcap is installed with the SDK option

---

## 🛠️ Usage

### Step 1: Find your network interface

First, find your active network interface name:

```bash
# On Linux
ip a

# On macOS
ifconfig

# On Windows
ipconfig
```

Look for interfaces like `eth0`, `wlan0`, `en0`, or similar that show your active connection.

### Step 2: Run Sniphawk

Run with `sudo` (required for packet capture):

```bash
# Basic usage with default interface (eth0)
sudo ./sniphawk

# Specify a different interface (example: wlan0)
sudo ./sniphawk -i wlan0

# Or using the long flag format
sudo ./sniphawk --interface=wlan0
```

### Step 3: View network traffic

Once running, Sniphawk will:
1. Display the colorful banner
2. Start capturing and displaying packets in real-time
3. Show traffic summaries every 30 seconds

Press `Ctrl+C` to stop the capture.

### Complete Flag Reference

| Flag | Short | Description | Example | Default |
|------|-------|-------------|---------|---------|
| `--interface` | `-i` | Network interface to capture on | `--interface=wlan0` | `eth0` |
| `--port` | `-p` | Filter by specific port | `--port=80` | `0` (all ports) |
| `--protocol` | `-proto` | Filter by protocol | `--protocol=tcp` | all protocols |
| `--skip-encrypted` | `-s` | Skip HTTPS traffic | `--skip-encrypted=false` | `true` |
| `--help` | `-h` | Show help message | `--help` | n/a |

### Examples

```bash
# Capture only HTTP traffic (port 80)
sudo ./sniphawk -i eth0 -p 80

# Capture only TCP traffic including encrypted traffic
sudo ./sniphawk -i wlan0 --protocol=tcp --skip-encrypted=false

# Capture only UDP traffic (e.g., DNS)
sudo ./sniphawk -i eth0 --protocol=udp
```

---

## 🖥️ Example Output & Interpretation

When running, Sniphawk captures packets and displays them in a colorful, easy-to-read format:

```
🕒 Timestamp: 2025-05-15T19:47:18+06:00
🔗 IP: 192.168.0.109 → 104.18.32.47 | Protocol: TCP
📦 TCP: 47812 → 443
📝 Payload: [Binary or Encrypted Data]
----------------------------------------------------
```

### Understanding the Output

Each packet is broken down into:

| Symbol | Meaning | Example |
|--------|---------|---------|
| 🕒 | When the packet was captured | `2025-05-15T19:47:18+06:00` |
| 🔗 | Source and destination IP addresses + Protocol | `192.168.0.109 → 104.18.32.47 | Protocol: TCP` |
| 📦 | TCP port information (if TCP packet) | `47812 → 443` (local port → destination port) |
| 📨 | UDP port information (if UDP packet) | `53212 → 53` (typically DNS traffic) |
| 💥 | ICMP information (if ICMP packet) | `ICMP Packet Detected` |
| 📝 | Payload/data content (may be encrypted) | `[Binary or Encrypted Data]` |

### Traffic Summary Analysis

Every 30 seconds, Sniphawk displays a traffic summary:

```
📊 ==== Sniphawk Traffic Summary (last 30s) ==== 📊
Top IPs:
  192.168.0.109 : 120 packets (Your device)
  104.18.32.47  : 80 packets  (Remote server)
  142.250.196.35: 50 packets  (Google server)

Top Ports:
  443: 130 packets (HTTPS)
  80 : 60 packets  (HTTP)
  53 : 40 packets  (DNS)
==============================================
```

This summary helps you quickly identify:
- Which devices are communicating most frequently
- Which services (ports) are most active
- Potential security concerns (unexpected connections)

### Common Network Patterns

| Port | Service | What it typically means |
|------|---------|-------------------------|
| 80   | HTTP    | Web browsing (unencrypted) |
| 443  | HTTPS   | Secure web browsing |
| 53   | DNS     | Domain name lookups |
| 67/68| DHCP    | Obtaining IP addresses |
| 22   | SSH     | Secure terminal connections |
| 25   | SMTP    | Email sending |

---

## ⚙️ Dependencies

* **gopacket** — packet processing
* **fatih/color** — colorful CLI output

Install dependencies:

```bash
go mod tidy
```

---

## 🎓 About This Project

This tool was developed during the **CodeAlpha Cybersecurity Internship**, to demonstrate hands-on experience with:

* Network traffic analysis and packet inspection
* Writing performant CLI tools in Go
* Understanding of TCP/IP and network protocols

---

## ⚠️ Disclaimer

This tool is provided for **educational and legitimate security testing purposes only**. By using Sniphawk, you agree to:

1. **Only monitor networks you own or have explicit permission to test**
2. **Follow all applicable laws and regulations** regarding network monitoring and privacy
3. **Not use this tool for any illegal or unauthorized activities**

The author and contributors are not responsible for any misuse or damage caused by this software. Always obtain proper authorization before monitoring any network.

---

## 📄 License

This project is licensed under the **MIT License** — see the [LICENSE](LICENSE) file for details.

---

## 🤝 Contributing

Contributions are welcome! Feel free to open issues or submit pull requests for enhancements or bug fixes.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

---

## 📫 Contact

For questions or collaboration, reach me at:

* GitHub: [ZeroHack01](https://github.com/ZeroHack01)
* Email: mongwoiching2080@gmail.com

---

<p align="center">
  <i>Happy sniffing!</i> 🦅 🐍
</p>

<div align="center">
  <sub>Made with ❤️ by ZeroHack01</sub>
</div>
