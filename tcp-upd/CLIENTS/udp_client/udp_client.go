package udp_client

import (
	"fmt"
	"log"
	"net"
	"time"
)

func UDP_CLIENT() {

	con, err := net.Dial("udp", ":4001")

	if err != nil {
		log.Fatalln(err)
	}

	defer con.Close()

	fmt.Println("UDP client on :4001")

	msg := "Hello UDP"

	_, wErr := con.Write([]byte(msg))
	if wErr != nil {
		log.Fatalln("Failed message", wErr)
	}

	fmt.Println("Sent:", msg)

	con.SetReadDeadline(time.Now().Add(5 * time.Second))

	buf := make([]byte, 1024)
	n, err := con.Read(buf)
	if err != nil {
		log.Fatal("Reading failed")
		return
	}

	fmt.Println("Received:", string(buf[:n]))

}
