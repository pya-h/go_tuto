package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<p style=\"margin: auto; color: red;font-size:20px;\">Home Sweet Home</p>")
}

func foPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<p style=\"margin: auto; color: red;font-size:20px;\">FuckOff</p>")
}

func notFoundPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<p style=\"margin: auto; color: red;font-size:20px;\">Where the fuck you think you're Going?</p>")
}
func withHandlerFunc() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/fo", foPage)
	http.ListenAndServe(":3000", nil)
}

func withHandle() {
	http.Handle("/", http.HandlerFunc(homePage))
	http.Handle("/fo", http.HandlerFunc(foPage))
	http.ListenAndServe(":3000", nil)
}

func routerFunc(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		homePage(res, req)
	case "/fo":
		foPage(res, req)
	default:
		notFoundPage(res, req)
	}

}

func withRouterFunc() {
	http.ListenAndServe(":3000", http.HandlerFunc(routerFunc))
}
func main() {
	withRouterFunc()
}
