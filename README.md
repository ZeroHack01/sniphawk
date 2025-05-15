# <img src="https://raw.githubusercontent.com/Tarikul-Islam-Anik/Animated-Fluent-Emojis/master/Emojis/Animals/Eagle.png" width="35" /> Sniphawk — Advanced Packet Sniffer in Go <img src="https://raw.githubusercontent.com/Tarikul-Islam-Anik/Animated-Fluent-Emojis/master/Emojis/Animals/Snake.png" width="25" />

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/ZeroHack01/sniphawk)](https://goreportcard.com/report/github.com/ZeroHack01/sniphawk)
[![GitHub issues](https://img.shields.io/github/issues/ZeroHack01/sniphawk)](https://github.com/ZeroHack01/sniphawk/issues)
[![Go Version](https://img.shields.io/github/go-mod/go-version/ZeroHack01/sniphawk?color=00ADD8)](go.mod)
[![Last Commit](https://img.shields.io/github/last-commit/ZeroHack01/sniphawk?color=orange)](https://github.com/ZeroHack01/sniphawk/commits/main)

<div align="center">
  <img src="https://raw.githubusercontent.com/ZeroHack01/sniphawk/main/assets/banner.png" alt="Sniphawk Banner" width="70%">
</div>

<p align="center">
  <b>A blazing fast network packet analyzer written in Go</b><br>
  <sub>Built for the CodeAlpha Cybersecurity Internship</sub>
</p>

---

## <img src="https://raw.githubusercontent.com/Tarikul-Islam-Anik/Animated-Fluent-Emojis/master/Emojis/Objects/Rocket.png" width="25" /> Overview

**Sniphawk** is a powerful, real-time network packet sniffer built in **Go**, designed to capture, analyze, and display live network traffic on any specified interface. This project showcases practical skills in network analysis and Go programming while providing security professionals with an intuitive tool for traffic inspection.

<details>
  <summary><b>🎬 See Sniphawk in action!</b></summary>
  <img src="https://raw.githubusercontent.com/ZeroHack01/sniphawk/main/assets/demo.gif" alt="Sniphawk Demo" width="85%">
</details>

---

## <img src="https://raw.githubusercontent.com/Tarikul-Islam-Anik/Animated-Fluent-Emojis/master/Emojis/Objects/Sparkles.png" width="25" /> Features

<table>
  <tr>
    <td><img src="https://raw.githubusercontent.com/Tarikul-Islam-Anik/Animated-Fluent-Emojis/master/Emojis/People/Detective.png" width="17" /></td>
    <td>Capture <b>TCP</b>, <b>UDP</b>, and <b>ICMP</b> packets in real-time</td>
  </tr>
  <tr>
    <td><img src="https://raw.githubusercontent.com/Tarikul-Islam-Anik/Animated-Fluent-Emojis/master/Emojis/Objects/Magnifying%20Glass%20Tilted%20Right.png" width="17" /></td>
    <td>Intelligent filtering by <b>port</b> and <b>protocol</b></td>
  </tr>
  <tr>
    <td><img src="https://raw.githubusercontent.com/Tarikul-Islam-Anik/Animated-Fluent-Emojis/master/Emojis/Objects/Prohibited.png" width="17" /></td>
    <td>Option to skip encrypted traffic (e.g., HTTPS on port 443)</td>
  </tr>
  <tr>
    <td><img src="https://raw.githubusercontent.com/Tarikul-Islam-Anik/Animated-Fluent-Emojis/master/Emojis/Objects/Bar%20Chart.png" width="17" /></td>
    <td>Real-time summary of <b>top IPs</b> and <b>most active ports</b> every 30 seconds</td>
  </tr>
  <tr>
    <td><img src="https://raw.githubusercontent.com/Tarikul-Islam-Anik/Animated-Fluent-Emojis/master/Emojis/Activities/Artist%20Palette.png" width="17" /></td>
    <td>Colorful CLI output with a professional banner</td>
  </tr>
  <tr>
    <td><img src="https://raw.githubusercontent.com/Tarikul-Islam-Anik/Animated-Fluent-Emojis/master/Emojis/Objects/Puzzle%20Piece.png" width="17" /></td>
    <td>Minimal dependencies for easy setup and portability</td>
  </tr>
</table>

---

## <img src="https://raw.githubusercontent.com/Tarikul-Islam-Anik/Animated-Fluent-Emojis/master/Emojis/Objects/Bullseye.png" width="25" /> Installation

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

## <img src="https://raw.githubusercontent.com/Tarikul-Islam-Anik/Animated-Fluent-Emojis/master/Emojis/Objects/Hammer%20and%20Wrench.png" width="25" /> Usage

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

## <img src="https://raw.githubusercontent.com/Tarikul-Islam-Anik/Animated-Fluent-Emojis/master/Emojis/Objects/Desktop%20Computer.png" width="25" /> Example Output

```
🕒 Timestamp: 2025-05-15T19:47:18+06:00
🔗 IP: 192.168.0.109 → 104.18.32.47 | Protocol: TCP
📦 TCP: 47812 → 443
📝 Payload: [Binary or Encrypted Data]
----------------------------------------------------
```

### <img src="https://raw.githubusercontent.com/Tarikul-Islam-Anik/Animated-Fluent-Emojis/master/Emojis/Objects/Bar%20Chart.png" width="20" /> Live Traffic Summary

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

## <img src="https://raw.githubusercontent.com/Tarikul-Islam-Anik/Animated-Fluent-Emojis/master/Emojis/Objects/Gear.png" width="25" /> Dependencies

* **gopacket** — packet processing
* **fatih/color** — colorful CLI output

Install dependencies:

```bash
go mod tidy
```

---

## <img src="https://raw.githubusercontent.com/Tarikul-Islam-Anik/Animated-Fluent-Emojis/master/Emojis/Objects/Graduation%20Cap.png" width="25" /> About This Project

This tool was developed during the **CodeAlpha Cybersecurity Internship**, to demonstrate hands-on experience with:

* Network traffic analysis and packet inspection
* Writing performant CLI tools in Go
* Understanding of TCP/IP and network protocols

<div align="center">
  <img src="https://raw.githubusercontent.com/ZeroHack01/sniphawk/main/assets/codealpha.png" alt="CodeAlpha Logo" width="200px">
</div>

---

## <img src="https://raw.githubusercontent.com/Tarikul-Islam-Anik/Animated-Fluent-Emojis/master/Emojis/Objects/Page%20with%20Curl.png" width="25" /> License

This project is licensed under the **MIT License** — see the [LICENSE](LICENSE) file for details.

---

## <img src="https://raw.githubusercontent.com/Tarikul-Islam-Anik/Animated-Fluent-Emojis/master/Emojis/Hand%20gestures/Handshake.png" width="25" /> Contributing

Contributions are welcome! Feel free to open issues or submit pull requests for enhancements or bug fixes.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

---

## <img src="https://raw.githubusercontent.com/Tarikul-Islam-Anik/Animated-Fluent-Emojis/master/Emojis/Objects/Incoming%20Envelope.png" width="25" /> Contact

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
