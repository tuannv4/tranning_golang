package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	connection, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open("../data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		input := scanner.Text() + "\n"
		time.Sleep(1 * time.Second)
		connection.Write([]byte(input))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	connection.Close()
	// nhận message từ bàn phím
	// for {
	// 	msgReder := bufio.NewReader(os.Stdin)
	// 	msg, err := msgReder.ReadString('\n')
	// 	if err != nil {
	// 		break
	// 	}
	// 	connection.Write([]byte(msg))
	// }

}
