package main

import (
	"fmt"
	"log"
	"main/conf"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var (
	device      string = ""
	snapshotLen int32  = 1024
	promiscuous bool   = false
	err         error
	timeout     time.Duration = 30 * time.Second
	handle      *pcap.Handle
)

func main() {
	// Read Config
	config, err := conf.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
	}
	device = config.DeviceName

	// Open device
	handle, err = pcap.OpenLive(device, snapshotLen, promiscuous, timeout)

	if err != nil {
		log.Fatal()
	}
	defer handle.Close()

	// Use the handle as a packet source to process all packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// Process packet here
		fmt.Println(packet)
	}
}
