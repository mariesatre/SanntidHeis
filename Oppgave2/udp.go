package main

import (
	"net"
	"fmt"
)

func updRecive() {
	// Address we are listening for messages on
	addr := net.UDPAddr{ 
		Port: 30000, // Port we are listening to
		IP: net.ParseIP("0.0.0.0"),
	}

	// Create a UDP connection
	conn, err := net.ListenUDP("udp", &addr) //conn = net.UDPConn. err = error
	if err != nil {	//error handler. non-nil means there is an error
		fmt.Println(err)
		return
	}
	defer conn.Close() // Waiting for the function to return before closing the connection (error)


	buffer := make([]byte, 1024) // Creating a buffer to store the message
	for {
		n, fromAddr, err := conn.ReadFromUDP(buffer) //n = number of bytes read, fromAddr = address of the sender, err = error
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Received ", string(buffer[0:n]), " from ", fromAddr)
	}
}

func udpSender() {
	// Address we are sending messages to
	addr := net.UDPAddr{
		Port: 20007,
		IP: net.ParseIP("255.255.255.255"),
	}

	conn, err := net.DialUDP("udp", nil, &addr) 
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	message := []byte("Hei hei")
	_, err = conn.Write(message)
	if err != nil {
		fmt.Println("Error:", err)
	}
}