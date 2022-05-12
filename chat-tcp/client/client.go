package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	connection, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("your name:")
	nameReader := bufio.NewReader(os.Stdin)
	nameInput, _ := nameReader.ReadString('\n')
	nameInput = nameInput[:len(nameInput)-1]

	for {
		msgReader := bufio.NewReader(os.Stdin)
		msg, err := msgReader.ReadSlice('\n')
		if err != nil {
			break
		}
		msg = []byte(fmt.Sprintf("%s: %s\n", nameInput,
			msg[:len(msg)-1]))

		connection.Write(msg)
	}
	connection.Close()
}
