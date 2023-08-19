package tcp_client

import (
	"fmt"
	"log"
	"net"
)

func TCP_CLIENT() {

	fmt.Println(" TCP client hitting on 4000 ")

	con, err := net.Dial("tcp", ":4000")

	if err != nil {
		log.Fatal(err)
	}

	defer con.Close()

	fmt.Fprintf(con, "Hello TCP")

	buf := make([]byte, 1024)

	n, err := con.Read(buf)
	if err != nil {
		log.Fatal("Reading failed")
	}

	fmt.Print(string(buf[:n]))

}
