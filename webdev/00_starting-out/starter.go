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

type Router struct {
	// doing this will give us previllages like having usefule fields and data on our router, such as database connection,
	// and we can pass it to our route handlers.
}

func (Router) ServeHTTP(res http.ResponseWriter, req *http.Request) { // implements the Handler interface
	path := req.URL.Path

	switch path {
	case "/":
		rootRouteHandler(res, req)
	case "/contact":
		contactRouteHandler(res, req)
	default:
		// res.WriteHeader(http.StatusNotFound) // Note: When you directly Write() the response, Go assumes the status is 200
		// fmt.Fprint(res, "Page Not Found.")
		http.Error(res, http.StatusText(http.StatusNotFound), http.StatusNotFound) // alternative
	}
}

func serveWithCustomRouter() error {
	var router Router
	fmt.Println("Listening on :3000 ...")

	return http.ListenAndServe(":3000", router) // nil means use the http.DefaultServeMux as the handler
}

func main() {
	defer func() {
		err := recover()
		log.Println(err)
	}()
	// go serveServerMux()
	if err := serveWithCustomRouter(); err != nil {
		panic(err)
	}

	fmt.Println("Also ListenAndServe can return an error: ")

}
