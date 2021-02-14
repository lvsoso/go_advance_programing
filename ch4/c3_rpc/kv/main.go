package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"time"
)

const  KEY = "abc"

func doClientWork(client *rpc.Client) {
	go func() {
		var keyChanged = KEY
		err := client.Call("KVStoreService.Watch", 30, &keyChanged)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("watch:", keyChanged)
	} ()

	err := client.Call(
		"KVStoreService.Set", [2]string{KEY, "abc-value"},
		new(struct{}),
	)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second*1)

	err = client.Call(
		"KVStoreService.Set", [2]string{KEY, "abc-value1"},
		new(struct{}),
	)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second*3)
}

func client()  {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	doClientWork(client)
}

func server(){
	rpc.Register(NewKVStoreService())

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}

	rpc.ServeConn(conn)

}

func main() {
	go server()
	time.Sleep(time.Second*1)
	client()
}