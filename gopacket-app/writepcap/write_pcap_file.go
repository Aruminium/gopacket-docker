package main

import (
	"fmt"
	"main/conf"
	"os"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
)

const PACKET_COUNT_MAX = 100

var (
	deviceName  string = ""
	snapshotLen int32  = 1024
	promiscuous bool   = false
	err         error
	timeout     time.Duration = -1 * time.Second
	handle      *pcap.Handle
	packetCount int = 0
)

func main() {
	// Read Config
	config, err := conf.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
	}
	deviceName = config.DeviceName

	// Open output pcap file and write header
	fmt.Println("Open output pcap file and write header")
	f, _ := os.Create(config.PcapFilePath)
	w := pcapgo.NewWriter(f)
	w.WriteFileHeader(uint32(snapshotLen), layers.LinkTypeEthernet)
	defer f.Close()

	// Open the device for capturing
	fmt.Println("Open the device for capturing")
	handle, err = pcap.OpenLive(deviceName, snapshotLen, promiscuous, timeout)
	if err != nil {
		fmt.Printf("Error opening device %s: %v", deviceName, err)
		os.Exit(1)
	}
	defer handle.Close()

	// Start processing packets
	fmt.Println("Start processing packets")
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// Process packet here
		fmt.Println(packet)
		w.WritePacket(packet.Metadata().CaptureInfo, packet.Data())
		packetCount++

		// Only capture PACKET_COUNT_MAX and then stop
		if packetCount > PACKET_COUNT_MAX {
			break
		}
	}
}
