package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	serverAddr := "72.194.77.242:5000" // Change this to your server's address and port
	// Create a UDP address structure to resolve the server address
	serverUdpAddr, err := net.ResolveUDPAddr("udp", serverAddr)
	if err != nil {
		fmt.Println("Error resolving address:", err)
		panic(err)
	}

	localAddr, err := net.ResolveUDPAddr("udp", "0.0.0.0:3000")
	if err != nil {
		fmt.Println("Error resolving address:", err)
		panic(err)
	}

	// Create a UDP connection
	conn, err := net.ListenUDP("udp", localAddr)
	if err != nil {
		fmt.Println("Error creating UDP connection:", err)
		panic(err)
	}
	defer conn.Close()

	go func() {
		fmt.Println("CLIENT | listening on", conn.LocalAddr())

		for {
			// read
			buf := make([]byte, 512)
			n, addr, err := conn.ReadFromUDP(buf)
			if err != nil {
				panic(err)
			}

			fmt.Println("CLIENT | received message from ", addr.IP, ":", addr.Port, " with n: ", n)
		}
	}()

	// Message to send
	message := []byte("Hello server, I am the client!")

	go func() {
		for {
			// write
			_, err = conn.WriteToUDP(message, serverUdpAddr)
			if err != nil {
				panic(err)
			}

			// sleep for 5 seconds
			time.Sleep(5 * time.Second)
		}
	}()

	// wait
	select {}
}
