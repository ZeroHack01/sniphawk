// sniphawk.go
package main

import (
    "flag"
    "fmt"
    "log"
    "sort"
    "strings"
    "sync"
    "time"

    "github.com/fatih/color"
    "github.com/google/gopacket"
    "github.com/google/gopacket/layers"
    "github.com/google/gopacket/pcap"
)

var (
    device         string
    filterPort     int
    filterProtocol string
    skipEncrypted  bool
    snapshotLen    int32 = 1024
    promiscuous           = true
    timeout               = pcap.BlockForever
    handle        *pcap.Handle
    seenFlows     = make(map[string]bool)
    mu            sync.Mutex
    ipCounts      = make(map[string]int)
    portCounts    = make(map[int]int)
)

func init() {
    flag.StringVar(&device, "interface", "wlp2s0", "Network interface to capture packets from")
    flag.IntVar(&filterPort, "port", 0, "Filter packets by port number")
    flag.StringVar(&filterProtocol, "protocol", "", "Filter packets by protocol (tcp/udp/icmp)")
    flag.BoolVar(&skipEncrypted, "skip-encrypted", true, "Skip common encrypted traffic ports like 443")
    flag.Parse()
}

func main() {
    printBanner()

    var err error
    handle, err = pcap.OpenLive(device, snapshotLen, promiscuous, timeout)
    if err != nil {
        log.Fatalf("❌ Failed to open device %s: %v", device, err)
    }
    defer handle.Close()

    ticker := time.NewTicker(30 * time.Second)
    defer ticker.Stop()

    packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

    go func() {
        for range ticker.C {
            printStats()
        }
    }()

    for packet := range packetSource.Packets() {
        output := processPacket(packet)
        if output != "" {
            fmt.Print(output)
        }
    }
}

func printBanner() {
    banner := `
 ███████╗███╗   ██╗██╗██████╗ ██╗  ██╗ █████╗ ██╗    ██╗██╗  ██╗
 ██╔════╝████╗  ██║██║██╔══██╗██║  ██║██╔══██╗██║    ██║██║ ██╔╝
 ███████╗██╔██╗ ██║██║██████╔╝███████║███████║██║ █╗ ██║█████╔╝ 
 ╚════██║██║╚██╗██║██║██╔═══╝ ██╔══██║██╔══██║██║███╗██║██╔═██╗ 
 ███████║██║ ╚████║██║██║     ██║  ██║██║  ██║╚███╔███╔╝██║  ██╗
 ╚══════╝╚═╝  ╚═══╝╚═╝╚═╝     ╚═╝  ╚═╝╚═╝  ╚═╝ ╚══╝╚══╝ ╚═╝  ╚═╝
     ⚡ Sniphawk - Advanced Packet Sniffer
     🛡️ CodeAlpha Cybersecurity Internship Project
     📡 Real-Time Network Traffic Monitor
`
    color.Cyan(banner)
}

func processPacket(packet gopacket.Packet) string {
    ipLayer := packet.Layer(layers.LayerTypeIPv4)
    if ipLayer == nil {
        return ""
    }
    ip := ipLayer.(*layers.IPv4)

    proto := strings.ToLower(ip.Protocol.String())
    srcIP := ip.SrcIP.String()
    dstIP := ip.DstIP.String()

    key := fmt.Sprintf("%s-%s-%s", srcIP, dstIP, proto)

    mu.Lock()
    if seenFlows[key] {
        mu.Unlock()
        return ""
    }
    seenFlows[key] = true
    ipCounts[srcIP]++
    ipCounts[dstIP]++
    mu.Unlock()

    if skipEncrypted && isEncryptedPort(packet) {
        return ""
    }

    if filterPort != 0 && !matchesPort(packet, filterPort) {
        return ""
    }

    if filterProtocol != "" && filterProtocol != proto {
        return ""
    }

    var sb strings.Builder
    sb.WriteString(fmt.Sprintf("\n🕒 Timestamp: %s\n", packet.Metadata().Timestamp.Format(time.RFC3339)))
    sb.WriteString(fmt.Sprintf("🔗 IP: %s → %s | Protocol: %s\n", srcIP, dstIP, proto))

    if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
        tcp := tcpLayer.(*layers.TCP)
        mu.Lock()
        portCounts[int(tcp.SrcPort)]++
        portCounts[int(tcp.DstPort)]++
        mu.Unlock()
        sb.WriteString(fmt.Sprintf("📦 TCP: %d → %d\n", tcp.SrcPort, tcp.DstPort))
    }
    if udpLayer := packet.Layer(layers.LayerTypeUDP); udpLayer != nil {
        udp := udpLayer.(*layers.UDP)
        mu.Lock()
        portCounts[int(udp.SrcPort)]++
        portCounts[int(udp.DstPort)]++
        mu.Unlock()
        sb.WriteString(fmt.Sprintf("📨 UDP: %d → %d\n", udp.SrcPort, udp.DstPort))
    }
    if icmpLayer := packet.Layer(layers.LayerTypeICMPv4); icmpLayer != nil {
        sb.WriteString("💥 ICMPv4 Packet Detected\n")
    }
    if app := packet.ApplicationLayer(); app != nil {
        if isPrintable(app.Payload()) {
            sb.WriteString(fmt.Sprintf("📝 Payload: %s\n", string(app.Payload())))
        } else {
            sb.WriteString("📝 Payload: [Binary or Encrypted Data]\n")
        }
    }
    sb.WriteString("----------------------------------------------------\n")
    return sb.String()
}

func isEncryptedPort(packet gopacket.Packet) bool {
    if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
        tcp := tcpLayer.(*layers.TCP)
        if tcp.SrcPort == 443 || tcp.DstPort == 443 {
            return true
        }
    }
    if udpLayer := packet.Layer(layers.LayerTypeUDP); udpLayer != nil {
        udp := udpLayer.(*layers.UDP)
        if udp.SrcPort == 443 || udp.DstPort == 443 {
            return true
        }
    }
    return false
}

func matchesPort(packet gopacket.Packet, port int) bool {
    if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
        tcp := tcpLayer.(*layers.TCP)
        if int(tcp.SrcPort) == port || int(tcp.DstPort) == port {
            return true
        }
    }
    if udpLayer := packet.Layer(layers.LayerTypeUDP); udpLayer != nil {
        udp := udpLayer.(*layers.UDP)
        if int(udp.SrcPort) == port || int(udp.DstPort) == port {
            return true
        }
    }
    return false
}

func isPrintable(data []byte) bool {
    for _, b := range data {
        if b < 32 || b > 126 {
            return false
        }
    }
    return true
}

func printStats() {
    mu.Lock()
    defer mu.Unlock()

    fmt.Println("\n📊 ==== Sniphawk Traffic Summary (last 30s) ==== 📊")

    type kv struct {
        Key   string
        Value int
    }

    var ipList []kv
    for k, v := range ipCounts {
        ipList = append(ipList, kv{k, v})
    }
    sort.Slice(ipList, func(i, j int) bool {
        return ipList[i].Value > ipList[j].Value
    })
    fmt.Println("Top IPs:")
    for i := 0; i < len(ipList) && i < 5; i++ {
        fmt.Printf("  %s : %d packets\n", ipList[i].Key, ipList[i].Value)
    }

    var portList []kv
    for k, v := range portCounts {
        portList = append(portList, kv{fmt.Sprintf("%d", k), v})
    }
    sort.Slice(portList, func(i, j int) bool {
        return portList[i].Value > portList[j].Value
    })
    fmt.Println("Top Ports:")
    for i := 0; i < len(portList) && i < 5; i++ {
        fmt.Printf("  %s : %d packets\n", portList[i].Key, portList[i].Value)
    }
    fmt.Println("==============================================\n")

    ipCounts = make(map[string]int)
    portCounts = make(map[int]int)
}
