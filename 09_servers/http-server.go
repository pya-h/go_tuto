package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func testRoute(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")

	html_bytes, err := ioutil.ReadFile("test.html")
	if err != nil {
		io.WriteString(res, "<h1>404: NOT FOUND</h1>")
		return
	}
	html := string(html_bytes)
	fmt.Println(html)

	io.WriteString(res, html)

}

func main() {
	fmt.Println("starting to listen on :4200...")
	http.HandleFunc("/test", testRoute)
	http.Handle("/assets/",
		http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.ListenAndServe(":4200", nil)
}