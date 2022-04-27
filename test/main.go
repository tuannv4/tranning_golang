package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	books := [4]string{"toan", "ly", "hoa", "sinh"}

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Tell me the book you look!!!")
		return
	}

	query := strings.ToLower(args[0])
	fmt.Println("searh value ", query)

	var check bool

	for _, b := range books {
		if strings.Contains(strings.ToLower(b), query) {
			fmt.Printf("find %q in books: ", b)
			check = true
		}
	}
	if !check {
		fmt.Printf("Not fond the %q\n", query)
	}

}
