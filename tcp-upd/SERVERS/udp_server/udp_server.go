package udp_server

import (
	"fmt"
	"log"
	"net"
)

func UDP_SERVER() {

	addr, err := net.ResolveUDPAddr("udp", ":4001")

	if err != nil {
		log.Fatalln(err)
	}

	con, err := net.ListenUDP("udp", addr)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("UDP server on :4001")

	defer con.Close()

	for {
		buf := make([]byte, 1024)

		n, clientAddr, err := con.ReadFromUDP(buf)

		if err != nil {
			log.Println("Error reading:", err)
			continue
		}

		fmt.Printf("Received %s, from %s", string(buf[:n]), clientAddr)

		_, cErr := con.WriteToUDP([]byte("Hello"), clientAddr)

		if cErr != nil {
			log.Println("Error writing:", cErr)
		}
	}
}
