package main

import (
	"encoding/json"
	"net"
	"net/http"
	"github.com/labstack/echo"
)

type User struct {
	Name string `json:"name"`
	Id uint8	`json:"id"`
}

func main() {
	http.HandleFunc("/user", handleGetUser)
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(map[string]string {"msg": "Welcome asshole!!"})
	})
	if listener, err := net.Listen("tcp", ":8000"); err == nil {
		http.Serve(listener, nil)
	}
}

func handleGetUser(res http.ResponseWriter, req *http.Request) {
	if user, err := getUser(); err == nil {
		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(user)
	} else {
		res.WriteHeader(http.StatusNotFound)
		json.NewEncoder(res).Encode(map[string]interface{} {"msg": err.Error()})
	}
}

func getUser() (*User, error){
	return &User{"whatever", 1}, nil;
}