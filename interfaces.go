package main

import (
    "fmt"
    "log"

    "github.com/google/gopacket/pcap"
)

func main() {
    devices, err := pcap.FindAllDevs()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("📡 Available Interfaces:")
    for _, dev := range devices {
        fmt.Printf("- %s (%s)\n", dev.Name, dev.Description)
    }
}
