package net

import (
	"fmt"
	"net"
)

func sendMessage(payload []byte) {
	addrList, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("no interfaces")
		return
	}

	var myAddr net.IP = nil

	for _, addr := range addrList {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				myAddr = ipnet.IP
			}

		}
	}

	srcAddr := net.UDPAddr{
		IP:   myAddr,
		Port: 3000,
	}

	destAddr := net.UDPAddr{
		IP:   net.IPv4(230, 0, 0, 1),
		Port: 3000,
	}

	conn, err := net.DialUDP("udp", &srcAddr, &destAddr)
	defer conn.Close()
	if err != nil {
		fmt.Println("error dial network")
	}

	conn.Write(payload)

}
