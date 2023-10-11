package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// udp listen
	conn, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4(0, 0, 0, 0), Port: 5000})
	if err != nil {
		panic(err)
	}

	go func() {
		log.Println("SERVER | listening on port 5000")
	}()

	for {
		// read
		buf := make([]byte, 512)
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			panic(err)
		}

		fmt.Println("SERVER | received message from ", addr)

		// write
		_, err = conn.WriteToUDP(buf[:n], addr)

		if err != nil {
			panic(err)
		}
	}
}
