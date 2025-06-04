⚡ SN1PH4WK — Cyber Packet Sniffer in Go 🐍

  
  
  
  
  



  A blazing-fast network packet analyzer forged in Go
  Crafted for the CodeAlpha Cybersecurity Internship



// Cyber Mission 🚀
SN1PH4WK is your ultimate weapon for real-time network traffic analysis. Built in Go, this elite packet sniffer captures, dissects, and visualizes live data streams with surgical precision. Designed for cybersecurity warriors and network ninjas, it’s your neon-lit portal into the heart of the digital grid. 🌌

// Elite Features ✨

🕵️‍♂️ Live Packet Capture: Snag TCP, UDP, and ICMP packets in real-time with zero lag.
🔍 Precision Filtering: Target specific ports or protocols like a digital sharpshooter.
🚫 Encryption Toggle: Optionally skip encrypted traffic (e.g., HTTPS on port 443).
📊 Neon Analytics: Real-time breakdowns of top IPs and active ports every 30 seconds.
🎨 Cyberpunk CLI: Vibrant, neon-green output with a pro-grade banner for max hacker vibes.
🧩 Lightweight Core: Minimal dependencies for instant deployment across platforms.


// System Boot 🎯
> Prerequisites

Go: Version 1.18+ (Download).
libpcap: Essential for packet capture.
Ubuntu/Debian: sudo apt-get install libpcap-dev
CentOS/RHEL: sudo yum install libpcap-devel
macOS: brew install libpcap
Windows: Install Npcap with dev files.



> Installation Protocol
# Clone the repository
git clone https://github.com/ZeroHack01/sniphawk.git

# Navigate to the project directory
cd sniphawk

# Install Go dependencies
go get github.com/google/gopacket
go get github.com/google/gopacket/pcap
go get github.com/fatih/color

# Build the binary
go build -o sniphawk sniphawk.go

# Or use go mod for a clean build
go mod init github.com/ZeroHack01/sniphawk
go mod tidy
go build

> Debug Matrix

Error: pcap.h: No such file or directory
Verify libpcap development files are installed.
Windows users: Ensure Npcap SDK is enabled during installation.




// Operational Uplink 🛠️
> Step 1: Locate Interface
Identify your active network interface:
# Linux
ip a

# macOS
ifconfig

# Windows
ipconfig

Look for interfaces like eth0, wlan0, or en0.
> Step 2: Engage SN1PH4WK
Run with elevated privileges:
# Default interface (eth0)
sudo ./sniphawk

# Custom interface (e.g., wlan0)
sudo ./sniphawk -i wlan0

# Long flag format
sudo ./sniphawk --interface=wlan0

> Step 3: Scan the Grid
SN1PH4WK will:

Flash a neon-charged banner.
Stream live packet data in vivid colors.
Deliver analytics every 30 seconds.

Stop with Ctrl+C.
> Command Flags



Flag
Short
Description
Example
Default



--interface
-i
Network interface to sniff
--interface=wlan0
eth0


--port
-p
Filter by port
--port=80
0 (all)


--protocol
-proto
Filter by protocol
--protocol=tcp
All protocols


--skip-encrypted
-s
Skip HTTPS traffic
--skip-encrypted=false
true


--help
-h
Show help
--help
n/a


> Usage Examples
# Capture HTTP traffic (port 80)
sudo ./sniphawk -i eth0 -p 80

# Capture TCP traffic, including HTTPS
sudo ./sniphawk -i wlan0 --protocol=tcp --skip-encrypted=false

# Capture UDP traffic (e.g., DNS)
sudo ./sniphawk -i eth0 --protocol=udp


// Data Stream Breakdown 🖥️
SN1PH4WK outputs packets in a neon-lit, hacker-friendly format:
<span style="color: #00FF00; font-family: 'Courier New', monospace;">🕒 Timestamp: 2025-06-05T12:19:00+06:00</span>
<span style="color: #39FF14; font-family: 'Courier New', monospace;">🔗 IP: 192.168.0.109 → 104.18.32.47 | Protocol: TCP</span>
<span style="color: #FF4500; font-family: 'Courier New', monospace;">📦 TCP: 47812 → 443</span>
<span style="color: #FFD700; font-family: 'Courier New', monospace;">📝 Payload: [Binary or Encrypted Data]</span>
<span style="color: #00FFFF; font-family: 'Courier New', monospace;">══════════════════════════════════════</span>

> Output Symbols



Symbol
Meaning
Example



🕒
Timestamp
2025-06-05T12:19:00+06:00


🔗
IPs + Protocol
192.168.0.109 → 104.18.32.47


📦
TCP ports
47812 → 443


📨
UDP ports
53212 → 53


💥
ICMP data
ICMP Packet Detected


📝
Payload
[Binary or Encrypted Data]


> Traffic Analytics
Every 30 seconds, a glowing summary lights up:
<span style="color: #00FFFF; font-family: 'Courier New', monospace;">📊 ════ SN1PH4WK Traffic Matrix (Last 30s) ════ 📊</span>
<span style="color: #39FF14; font-family: 'Courier New', monospace;">Top IPs:</span>
  <span style="color: #FF69B4;">192.168.0.109</span> : 120 packets (Local device)
  <span style="color: #FF69B4;">104.18.32.47</span>  : 80 packets (Remote server)
  <span style="color: #FF69B4;">142.250.196.35</span>: 50 packets (Google server)

<span style="color: #39FF14; font-family: 'Courier New', monospace;">Top Ports:</span>
  <span style="color: #FF4500;">443</span>: 130 packets (HTTPS)
  <span style="color: #FF4500;">80</span> : 60 packets (HTTP)
  <span style="color: #FF4500;">53</span> : 40 packets (DNS)
<span style="color: #00FFFF; font-family: 'Courier New', monospace;">══════════════════════════════════════</span>

This reveals:

Dominant devices in the traffic stream.
Hotspot ports and services.
Potential security anomalies.

> Network Patterns



Port
Service
Typical Use



80
HTTP
Unencrypted web


443
HTTPS
Secure web


53
DNS
Domain lookups


67/68
DHCP
IP assignment


22
SSH
Secure remote access


25
SMTP
Email delivery



// Core Dependencies ⚙️

gopacket: Packet processing engine.
fatih/color: Neon CLI output.

Install with:
go mod tidy


// Project Genesis 🎓
Forged during the CodeAlpha Cybersecurity Internship, SN1PH4WK demonstrates mastery in:

Network traffic dissection and analysis.
High-performance Go CLI development.
TCP/IP protocol expertise.


// Ethical Directive ⚠️
SN1PH4WK is for authorized security testing and education only. Users must:

Monitor only networks with explicit permission.
Adhere to all legal and privacy regulations.
Avoid unauthorized or illegal activities.

Developers are not liable for misuse. Secure authorization before scanning.

// License 📜
MIT License. See LICENSE for details.

// Join the Cyber Collective 🤝
Contribute to the grid! Open issues or PRs to level up SN1PH4WK.

Fork the repo.
Branch out: git checkout -b feature/cyber-upgrade
Commit: git commit -m 'Added cyber upgrade'
Push: git push origin feature/cyber-upgrade
Submit a Pull Request.


// Connect to the Source 📫

GitHub: ZeroHack01
Email: mongwoiching2080@gmail.com



  Crack the network, own the grid! 🦅 🐍



  Forged with ❤️ by ZeroHack01
