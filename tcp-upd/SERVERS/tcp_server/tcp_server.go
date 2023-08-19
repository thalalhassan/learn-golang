package tcp_server

import (
	"fmt"
	"log"
	"net"
	"time"
)

func handleConnection(con net.Conn) {
	defer con.Close()

	buf := make([]byte, 1024)

	n, err := con.Read(buf)
	if err != nil {
		log.Fatal("Reading failed")
	}

	fmt.Print(string(buf[:n]))

	fmt.Fprintf(con, "Echo %s", string(buf[:n]))
}

func handleConnectionNonBlocking(con net.Conn) {
	defer con.Close()

	for {
		con.SetReadDeadline(time.Now().Add(time.Second))
		buf := make([]byte, 1024)

		n, err := con.Read(buf)
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				continue
			} else {
				log.Fatal("Reading failed")
			}
		}
		fmt.Print("Received:", string(buf[:n]))

		con.SetWriteDeadline(time.Now().Add(time.Second))

		fmt.Fprintf(con, "Echo %s", string(buf[:n]))
	}

}

func TCP_SERVER() {

	fmt.Println()

	listener, err := net.Listen("tcp", ":4000")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(" TCP server listening on 4000 ")

	defer listener.Close()

	for {
		con, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConnectionNonBlocking(con)
		// go handleConnection(con)
	}

}
