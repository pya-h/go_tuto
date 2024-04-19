package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(response_w http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response_w, "<h1>Yeah Whatever</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Listening on :3000 ...")
	http.ListenAndServe(":3000", nil)
}
