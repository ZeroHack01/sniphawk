⚡ SN1PH4WK — Elite Packet Sniffer in Go 🐍

  
  
  
  
  



  A high-octane network packet analyzer forged in Go
  Crafted for the CodeAlpha Cybersecurity Internship



// Mission Overview 🚀
SN1PH4WK is a cutting-edge, real-time packet sniffer built in Go, engineered to capture, dissect, and visualize live network traffic with surgical precision. Designed for cybersecurity pros and network enthusiasts, this tool delivers unparalleled insights into your network's pulse, all wrapped in a sleek, neon-charged CLI interface.

// Core Features ✨

🕵️‍♂️ Real-Time Capture: Snag TCP, UDP, and ICMP packets as they fly across the wire.
🔍 Smart Filtering: Zero in on specific ports or protocols with pinpoint accuracy.
🚫 Encryption Bypass Option: Toggle to skip encrypted traffic (e.g., HTTPS on port 443).
📊 Dynamic Analytics: Get real-time summaries of top IPs and active ports every 30 seconds.
🎨 Vivid CLI: Neon-green output with a professional-grade banner for that hacker aesthetic.
🧩 Lightweight Design: Minimal dependencies for rapid deployment and portability.


// System Setup 🎯
Prerequisites

Go: Version 1.18+ (Download).
libpcap: Required for packet capture.
Ubuntu/Debian: sudo apt-get install libpcap-dev
CentOS/RHEL: sudo yum install libpcap-devel
macOS: brew install libpcap
Windows: Install Npcap with dev files.



Installation Protocol
# Clone the repository
git clone https://github.com/ZeroHack01/sniphawk.git

# Enter the project directory
cd sniphawk

# Fetch Go dependencies
go get github.com/google/gopacket
go get github.com/google/gopacket/pcap
go get github.com/fatih/color

# Compile the binary
go build -o sniphawk sniphawk.go

# Or use go mod for a clean build
go mod init github.com/ZeroHack01/sniphawk
go mod tidy
go build

Troubleshooting

Error: pcap.h: No such file or directory
Ensure libpcap development files are installed.
Windows users: Verify Npcap SDK is included during installation.




// Operational Guide 🛠️
Step 1: Identify Network Interface
Locate your active network interface:
# Linux
ip a

# macOS
ifconfig

# Windows
ipconfig

Look for interfaces like eth0, wlan0, or en0.
Step 2: Launch SN1PH4WK
Run with elevated privileges (required for packet capture):
# Default interface (eth0)
sudo ./sniphawk

# Custom interface (e.g., wlan0)
sudo ./sniphawk -i wlan0

# Long flag format
sudo ./sniphawk --interface=wlan0

Step 3: Monitor the Grid
SN1PH4WK will:

Flash a vibrant neon banner.
Stream live packet data in real-time.
Deliver traffic summaries every 30 seconds.

Stop capture with Ctrl+C.
Command Flags



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
Display help
--help
n/a


Usage Examples
# Capture HTTP traffic (port 80)
sudo ./sniphawk -i eth0 -p 80

# Capture TCP traffic, including HTTPS
sudo ./sniphawk -i wlan0 --protocol=tcp --skip-encrypted=false

# Capture UDP traffic (e.g., DNS)
sudo ./sniphawk -i eth0 --protocol=udp


// Output Breakdown 🖥️
SN1PH4WK delivers a vivid, color-coded stream of packet data:
<span style="color: #00FF00;">🕒 Timestamp: 2025-06-05T12:16:00+06:00</span>
<span style="color: #39FF14;">🔗 IP: 192.168.0.109 → 104.18.32.47 | Protocol: TCP</span>
<span style="color: #FF4500;">📦 TCP: 47812 → 443</span>
<span style="color: #FFD700;">📝 Payload: [Binary or Encrypted Data]</span>
<span style="color: #00FFFF;">----------------------------------------------------</span>

Output Symbols



Symbol
Meaning
Example



🕒
Capture timestamp
2025-06-05T12:16:00+06:00


🔗
Source/Dest IPs + Protocol
`192.168.0.109 → 104.18.32.47


📦
TCP ports
47812 → 443


📨
UDP ports
53212 → 53 (DNS traffic)


💥
ICMP data
ICMP Packet Detected


📝
Payload content
[Binary or Encrypted Data]


Traffic Summary
Every 30 seconds, a neon-lit summary highlights:
<span style="color: #00FFFF;">📊 ==== SN1PH4WK Traffic Report (Last 30s) ==== 📊</span>
<span style="color: #39FF14;">Top IPs:</span>
  <span style="color: #FF69B4;">192.168.0.109</span> : 120 packets (Local device)
  <span style="color: #FF69B4;">104.18.32.47</span>  : 80 packets (Remote server)
  <span style="color: #FF69B4;">142.250.196.35</span>: 50 packets (Google server)

<span style="color: #39FF14;">Top Ports:</span>
  <span style="color: #FF4500;">443</span>: 130 packets (HTTPS)
  <span style="color: #FF4500;">80</span> : 60 packets (HTTP)
  <span style="color: #FF4500;">53</span> : 40 packets (DNS)
<span style="color: #00FFFF;">==============================================</span>

This helps you spot:

High-traffic devices.
Active services/ports.
Potential security red flags.

Common Network Patterns



Port
Service
Typical Use Case



80
HTTP
Unencrypted web traffic


443
HTTPS
Secure web traffic


53
DNS
Domain lookups


67/68
DHCP
IP address assignment


22
SSH
Secure remote access


25
SMTP
Email delivery



// Dependencies ⚙️

gopacket: Packet processing powerhouse.
fatih/color: Neon CLI styling.

Install with:
go mod tidy


// Project Origins 🎓
Developed during the CodeAlpha Cybersecurity Internship, SN1PH4WK showcases expertise in:

Network traffic analysis and packet dissection.
High-performance Go for CLI tools.
TCP/IP mastery and protocol deep-dives.


// Ethical Disclaimer ⚠️
SN1PH4WK is for educational and authorized security testing only. By using this tool, you agree to:

Monitor only networks you own or have explicit permission to analyze.
Comply with all laws and regulations on network monitoring and privacy.
Avoid any illegal or unauthorized activities.

The developers are not liable for misuse or damages. Always secure proper authorization.

// License 📜
Licensed under the MIT License. See LICENSE for details.

// Contribute to the Grid 🤝
Join the syndicate! Open issues or submit PRs to enhance SN1PH4WK.

Fork the repo.
Create a feature branch: git checkout -b feature/epic-upgrade
Commit changes: git commit -m 'Added epic upgrade'
Push: git push origin feature/epic-upgrade
Submit a Pull Request.


// Connect with the Architect 📫

GitHub: ZeroHack01
Email: mongwoiching2080@gmail.com



  Unleash the packets, rule the network! 🦅 🐍



  Crafted with ❤️ by ZeroHack01
