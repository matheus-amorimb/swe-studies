package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

const port = ":42069"
const network = "tcp"

func main() {
	listener, err := net.Listen(network, port)
	if err != nil {
		log.Fatalf("Error listening to port %v: %v", port, err.Error())
	}
	defer listener.Close()

	println(fmt.Sprintf("Started listening for TCP traffic on port %v", port))
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Error accepting connection to listener...")
		}

		fmt.Println("Accepted connection from ", conn.RemoteAddr())

		for line := range getLinesChannel(conn) {
			fmt.Println(line)
		}
	}

}

func getLinesChannel(f io.ReadCloser) <-chan string {
	ch := make(chan string)
	go func() {
		defer f.Close()
		defer close(ch)
		line := ""
		for {
			b := make([]byte, 8)
			n, err := f.Read(b)
			if err != nil {
				if line != "" {
					ch <- line
				}
				if errors.Is(err, io.EOF) {
					return
				}
				ch <- fmt.Sprintf("Error reading file: %s\n", err.Error())
				return
			}

			str := string(b[:n])
			lineParts := strings.Split(str, "\n")
			for i := 0; i < len(lineParts)-1; i++ {
				ch <- line + lineParts[i]
				line = ""
			}
			line += lineParts[len(lineParts)-1]
		}
	}()

	return ch
}
