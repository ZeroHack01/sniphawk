# ⚡ Sniphawk — Advanced Packet Sniffer in Go 🐍

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/ZeroHack01/sniphawk)](https://goreportcard.com/report/github.com/ZeroHack01/sniphawk)
[![GitHub issues](https://img.shields.io/github/issues/ZeroHack01/sniphawk)](https://github.com/ZeroHack01/sniphawk/issues)
[![Go](https://img.shields.io/badge/Go-1.18+-00ADD8?style=flat&logo=go)](https://golang.org/dl/)
[![Last Commit](https://img.shields.io/github/last-commit/ZeroHack01/sniphawk?color=orange)](https://github.com/ZeroHack01/sniphawk/commits/main)

<p align="center">
  <b>A blazing fast network packet analyzer written in Go</b><br>
  <sub>Built for the CodeAlpha Cybersecurity Internship</sub>
</p>

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

Make sure you have [Go](https://golang.org/dl/) installed (version 1.18+ recommended).

```bash
# Clone the repository
git clone https://github.com/ZeroHack01/sniphawk.git

# Navigate to project directory
cd sniphawk

# Install dependencies
go mod tidy

# Build the binary
go build -o sniphawk sniphawk.go
```

---

## 🛠️ Usage

Run with `sudo` to capture network traffic:

```bash
sudo ./sniphawk --interface=wlp2s0 --skip-encrypted=true
```

### Available flags

| Flag | Description | Default |
|------|-------------|---------|
| `--interface` | Network interface to capture on | `wlp2s0` |
| `--port` | Filter by specific port (e.g., 80) | `0` (all) |
| `--protocol` | Filter by protocol (`tcp`, `udp`, `icmp`) | none |
| `--skip-encrypted` | Skip encrypted traffic (port 443) | `true` |

---

## 🖥️ Example Output

```
🕒 Timestamp: 2025-05-15T19:47:18+06:00
🔗 IP: 192.168.0.109 → 104.18.32.47 | Protocol: TCP
📦 TCP: 47812 → 443
📝 Payload: [Binary or Encrypted Data]
----------------------------------------------------
```

### 📊 Live Traffic Summary

Every 30 seconds, Sniphawk prints:
* Top 5 IP addresses by packet count
* Top 5 ports by packet count

Example:

```
📊 ==== Sniphawk Traffic Summary (last 30s) ==== 📊
Top IPs:
  192.168.0.109 : 120 packets
  104.18.32.47  : 80 packets
  142.250.196.35: 50 packets

Top Ports:
  443: 130 packets
  80 : 60 packets
  53 : 40 packets
==============================================
```

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
