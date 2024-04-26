package main

import (
	"fmt"
	"log"
	"net/http"
)

func rootRouteHandler(response_w http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response_w, "<h1>Yeah Whatever by http.HandleFuck</h1>")
}

func rootRouteHandlerMux(response_w http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response_w, "<h1>Yeah Whatever By ServeMux</h1>")
}

func byHandleFunc() error {
	http.HandleFunc("/", rootRouteHandler)
	fmt.Println("Listening on :3000 ...")

	return http.ListenAndServe(":5", nil) // nil means use the http.DefaultServeMux as the handler
}

func byServerMux() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootRouteHandlerMux)
	fmt.Println("Listening on :3001 ...")
	return http.ListenAndServe(":3001", mux)
}

func main() {
	defer func() {
		err := recover()
		log.Println(err)
		log.Println("Though the :5 serve had a panic attack, i forced the panic defer into a forever so that the :3001 still keeps running.")
		fmt.Println("I Know, im twisting everything. But its 7:40 in the morning, i havent sl;ept yet and im BORED with these. Until i reach a state in this tuto that i start learning new thioooooo")
		for {
		}
	}()
	go byServerMux()
	err := byHandleFunc()

	fmt.Println("Also ListenAndServe can return an error: ")
	panic(err)
}
