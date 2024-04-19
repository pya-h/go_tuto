package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"time"
)

func handleConnection(connection net.Conn) {
	var msg string
	err := gob.NewDecoder(connection).Decode(&msg)

	if err != nil {
		fmt.Println("Error while trying to get data: ", err)
	} else {
		fmt.Println("Received: ", msg)
	}
	connection.Close()
}

func server() {
	listener, err := net.Listen("tcp", ":4200")

	if err != nil {
		panic(err)
	}

	for {
		connection, err := listener.Accept()

		if err != nil {
			fmt.Println("Some client tried to connect but it failed: ", err)
			continue
		}

		go handleConnection(connection)
	}
}

func client(msg string) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	connection, err := net.Dial("tcp", "127.0.0.1:4200")
	if err != nil {
		fmt.Println("Cannot connect to the server: ", err)
		return
	}

	err = gob.NewEncoder(connection).Encode(msg)
	if err != nil {
		panic(err)
	}
	connection.Close()
}

func main() {
	go server()
	for i := 0; ; i++{
		go client(fmt.Sprint("Test #", i))
		time.Sleep(time.Second)
	}
	var input string
	fmt.Scanln(&input)

}