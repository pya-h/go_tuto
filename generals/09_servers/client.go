package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
)


func handleConnection(connection net.Conn) {
	var msg string
	err := gob.NewDecoder(connection).Decode(&msg)

	if err != nil {
		fmt.Println("Error while trying to get data: ", err)
	} else {
		fmt.Println("Server Says: ", msg)
	}
	connection.Close()

}


func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	for {
		connection, err := net.Dial("tcp", "127.0.0.1:4200")
		if err != nil {
			fmt.Println("Cannot connect to the server: ", err)
			return
		}
		go handleConnection(connection)

		// read full line
		fmt.Print("Benal: ")
		in := bufio.NewReader(os.Stdin)
		msg, _ := in.ReadString('\n')
		err = gob.NewEncoder(connection).Encode(msg)
		if err != nil {
			panic(err)
		}

	}

}
