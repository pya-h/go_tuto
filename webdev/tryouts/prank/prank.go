package main

import (
	"fmt"
	"log"
	"net/http"
)

func rootRouteHandlerMux(response_w http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response_w, "<h1>Yeah Whatever By ServeMux</h1>")
}

func serveServerMux() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootRouteHandlerMux)
	fmt.Println("Listening on :3001 hmm ...")
	return http.ListenAndServe(":3001", mux)
}

func rootRouteHandler(response_w http.ResponseWriter, request *http.Request) {
	// response_w.Header().Set("Content-Type", "text/plain")
	response_w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	fmt.Fprint(response_w, "<h1>Yeah Whatever by http.HandleFuck</h1>")
}
func serve() error {
	http.HandleFunc("/", rootRouteHandler)
	fmt.Println("Listening on :3000 ...")

	return http.ListenAndServe(":3000", nil) // nil means use the http.DefaultServeMux as the handler
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
	// go serveServerMux()
	err := serve()
	fmt.Println("Also ListenAndServe can return an error: ")
	panic(err)
}
