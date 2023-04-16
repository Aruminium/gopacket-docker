package main

import (
	"fmt"
	"log"
	"main/conf"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var (
	pcapFile string = ""
	handle   *pcap.Handle
	err      error
)

func main() {
	// Read Config
	config, err := conf.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
	}
	pcapFile = config.PcapFilePath

	// Open file instead of device
	handle, err = pcap.OpenOffline(pcapFile)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Loop throuph packets in file
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		fmt.Println(packet)
	}
}
