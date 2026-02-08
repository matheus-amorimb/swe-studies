package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const address = "localhost:42069"

func main() {
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		log.Fatalf("Error starting UDP sender %v", err)
	}

	println("Started listening to address: ", address)

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatalf("Error establishing connection %v", err)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("%s", err.Error())
		}

		_, err = conn.Write([]byte(line))
		if err != nil {
			log.Fatalf("%s", err.Error())
		}
	}
}
