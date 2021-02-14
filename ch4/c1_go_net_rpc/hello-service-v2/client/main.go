package main

import (
	"fmt"
	"log"
)

func main() {
	client, err := DailHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
	}

	var reply string
	err = client.Hello("hello", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
