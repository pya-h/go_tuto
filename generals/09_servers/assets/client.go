package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
)

func main() {
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
	// read full line
	in := bufio.NewReader(os.Stdin)
	msg, _ := in.ReadString('\n')
	err = gob.NewEncoder(connection).Encode(msg)
	if err != nil {
		panic(err)
	}
	connection.Close()
}
