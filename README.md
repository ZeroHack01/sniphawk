<div align="center">

```ascii
███████╗███╗   ██╗██╗██████╗ ██╗  ██╗ █████╗ ██╗    ██╗██╗  ██╗
██╔════╝████╗  ██║██║██╔══██╗██║  ██║██╔══██╗██║    ██║██║ ██╔╝
███████╗██╔██╗ ██║██║██████╔╝███████║███████║██║ █╗ ██║█████╔╝ 
╚════██║██║╚██╗██║██║██╔═══╝ ██╔══██║██╔══██║██║███╗██║██╔═██╗ 
███████║██║ ╚████║██║██║     ██║  ██║██║  ██║╚███╔███╔╝██║  ██╗
╚══════╝╚═╝  ╚═══╝╚═╝╚═╝     ╚═╝  ╚═╝╚═╝  ╚═╝ ╚══╝╚══╝ ╚═╝  ╚═╝
```

[![License: MIT](https://img.shields.io/badge/License-MIT-00ff00.svg?style=for-the-badge&logo=opensourceinitiative&logoColor=white)](LICENSE)
[![Go](https://img.shields.io/badge/Go-1.18+-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://golang.org/dl/)
[![Status](https://img.shields.io/badge/STATUS-🟢_OPERATIONAL-00ff00?style=for-the-badge&logo=statuspage&logoColor=white)](/)

**Advanced Network Packet Sniffer for Cybersecurity Operations**

</div>

---

## 🎯 OVERVIEW

**SNIPHAWK** is a high-performance network packet analyzer built in Go for real-time traffic monitoring and analysis. Designed for cybersecurity professionals and penetration testers.

```bash
┌─[root@cybersec]─[~/sniphawk]
└──╼ $ sudo ./sniphawk -i wlan0 --protocol=tcp
```

---

## ⚡ FEATURES

| CAPABILITY | STATUS | DESCRIPTION |
|------------|--------|-------------|
| 🔍 **Multi-Protocol Capture** | `🟢 ACTIVE` | TCP, UDP, ICMP packet interception |
| 🎯 **Smart Filtering** | `🟢 ACTIVE` | Port and protocol-based filtering |
| 📊 **Real-time Analytics** | `🟢 ACTIVE` | Live traffic summaries every 30s |
| 🚫 **Encryption Bypass** | `🟢 ACTIVE` | Optional HTTPS traffic analysis |
| 🎨 **Terminal UI** | `🟢 ACTIVE` | Colorful CLI with professional output |

---

## 🚀 QUICK DEPLOYMENT

### Prerequisites
```bash
# Install Go 1.18+ and libpcap
sudo apt install libpcap-dev    # Ubuntu/Debian
brew install libpcap            # macOS
# Windows: Install Npcap
```

### Installation
```bash
git clone https://github.com/ZeroHack01/sniphawk.git
cd sniphawk
go mod tidy
go build -o sniphawk sniphawk.go
```

### Basic Usage
```bash
# Find your network interface
ip a                            # Linux
ifconfig                        # macOS

# Start monitoring
sudo ./sniphawk -i wlan0
sudo ./sniphawk -i eth0 --protocol=tcp --port=80
```

---

## 🔧 COMMAND REFERENCE

| FLAG | ALIAS | DESCRIPTION | EXAMPLE |
|------|-------|-------------|---------|
| `--interface` | `-i` | Network interface | `-i wlan0` |
| `--port` | `-p` | Filter by port | `-p 443` |
| `--protocol` | `-proto` | Filter by protocol | `--protocol=tcp` |
| `--skip-encrypted` | `-s` | Skip HTTPS traffic | `-s=false` |

---

## 📊 SAMPLE OUTPUT

```bash
🕐 [2025-06-05T15:42:13Z] ⚡ PACKET INTERCEPTED
┌─────────────────────────────────────────────────────────┐
│ 🎯 ROUTE: 192.168.1.105:52847 ──→ 142.250.196.35:443  │
│ 🛡️ PROTOCOL: TCP | 📦 SIZE: 1,337 bytes               │
│ 🔐 ENCRYPTION: TLS 1.3 | 🎯 TARGET: google.com        │
│ 🚨 STATUS: 🟢 BENIGN | 🔍 TYPE: WEB_TRAFFIC           │
└─────────────────────────────────────────────────────────┘

📊 TRAFFIC SUMMARY (30s)
Top IPs: 192.168.1.105 (342 packets), 142.250.196.35 (198 packets)
Top Ports: 443 (456 packets), 80 (198 packets), 53 (134 packets)
```

---

## 🛡️ SECURITY & LEGAL

**⚠️ AUTHORIZED USE ONLY**

This tool is for legitimate security testing and educational purposes. By using SNIPHAWK:
- ✅ Only monitor networks you own or have explicit permission to test
- ✅ Follow all applicable laws and privacy regulations
- ✅ Do not use for unauthorized surveillance

---

## 🤝 CONTRIBUTING

```bash
git fork https://github.com/ZeroHack01/sniphawk.git
git checkout -b feature/enhancement
git commit -m "Add new feature"
git push origin feature/enhancement
# Open Pull Request
```

---

## 📡 CONTACT

**Mission Commander:** [ZeroHack01](https://github.com/ZeroHack01)  
**Email:** mongwoiching2080@gmail.com  
**License:** MIT

---

<div align="center">

```bash
┌─[SNIPHAWK]─[OPERATIONAL]─[v2.0]
└──╼ $ echo "Happy packet hunting! 🦅"
```

<sub>Built with ⚡ for cybersecurity professionals</sub>

</div>
