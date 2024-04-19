package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Server struct {}

func (this Server) Negate (param int64, response *int64) error {
	*response = -param
	return nil
}

func (this Server) Identify (param string, response *string) error {
	*response = fmt.Sprint("Kose nane ", param)
	return nil
}

func server() {
	err := rpc.Register(new(Server))
	if err != nil {
		fmt.Println("Something went wrong while registering a server: ", err)
	}
	ln, err := net.Listen("tcp", ":4200")
	if err != nil {
		fmt.Println("Something went wrong:", err)
		return
	}
	fmt.Println("Server is up and running on: 127.0.0.1:4200")
	for {
		client, err := ln.Accept()
		if err != nil {
			fmt.Println("Something went wrong while accepting a connection.")
			continue
		}
		go rpc.ServeConn(client)
	}

}

func client() {
	connection, err := rpc.Dial("tcp", "127.0.0.1:4200")
	if err != nil {
		fmt.Println("Something went wrong: ", err)
		return
	}
	var response int64
	err = connection.Call("Server.Negate", int64(999), &response)
	if err != nil {
		fmt.Println("Something went wrong while calling a rpc function: ", err)
	} else {
		fmt.Println("Server.Negate(999) = ", response)
	}
	var res2 string
	err = connection.Call("Server.Identify", "Sadeghi Far", &res2)
	if err != nil {
		fmt.Println("Something went wrong while calling a rpc function: ", err)
	} else {
		fmt.Println("Server.Identify(Sadeghi Far) = ", res2)
	}

}
func main() {
	go server()
	go client()

	var input string
	fmt.Scanln(&input)
}