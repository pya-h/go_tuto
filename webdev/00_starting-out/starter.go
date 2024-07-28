package main

import (
	"fmt"
	"log"
	"net/http"
)

func rootRouteHandler(response_w http.ResponseWriter, request *http.Request) {
	// response_w.Header().Set("Content-Type", "text/plain")
	response_w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	fmt.Fprint(response_w, "<h1>Yeah Whatever by http.HandleFuck</h1>")
}

func contactRouteHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=UTF-8")
	fmt.Fprint(res, "<h1>Contact Me</h1><br><b>You can check me out from <a href=\"https://github.com/pya-h\">Here</a></b>")
}

func serve() error {
	http.HandleFunc("/", rootRouteHandler)
	http.HandleFunc("/contact", contactRouteHandler)
	fmt.Println("Listening on :3000 ...")

	return http.ListenAndServe(":3000", nil) // nil means use the http.DefaultServeMux as the handler
}

func main() {
	defer func() {
		err := recover()
		log.Println(err)
	}()
	// go serveServerMux()
	if err := serve(); err != nil {
		panic(err)
	}

	fmt.Println("Also ListenAndServe can return an error: ")

}
