package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type FaqQuestion struct {
	Text   string
	Answer string
}
type FaqTemplateContext struct {
	Title     string
	Questions []FaqQuestion
}

func renderTemplate(res http.ResponseWriter, template_name string, context FaqTemplateContext) {
	temp, err := template.ParseFiles(template_name)

	if err != nil {
		http.Error(res, "Failed rendering the page.", http.StatusInternalServerError)
		log.Println("Failed rendering template: ", template_name)
		return
	}
	if err = temp.Execute(res, context); err != nil {
		http.Error(res, "Failed loading page data ...", http.StatusBadRequest)
		log.Println("Failed loading context data into ", template_name, " template.")
		return
	}
}

func faqRouteHandler(res http.ResponseWriter, req *http.Request) {
	questions := make([]FaqQuestion, 0)
	questions = append(questions,
		FaqQuestion{Text: "What is Go programming language?", Answer: "Go, also known as Golang, is a statically typed, compiled programming language designed at Google. It is known for its simplicity, performance, and strong concurrency support."},
		FaqQuestion{Text: "How do I install Go on my computer?", Answer: "You can install Go by downloading it from the official website at golang.org and following the installation instructions for your operating system."},
		FaqQuestion{Text: "What are the main features of Go?", Answer: "Go is known for its simplicity, fast compilation, concurrency support through goroutines, and garbage collection, making it ideal for scalable and high-performance applications."},
		FaqQuestion{Text: "How do I create a new Go module?", Answer: "You can create a new Go module by running `go mod init <module_name>` in your project directory. This initializes the module with a `go.mod` file."},
		FaqQuestion{Text: "What is a goroutine?", Answer: "A goroutine is a lightweight thread managed by the Go runtime. Goroutines allow concurrent execution of functions, making it easier to perform multiple tasks simultaneously."},
		FaqQuestion{Text: "How do I handle errors in Go?", Answer: "Errors in Go are handled by returning an error value from functions. You can check if an error is `nil` to determine if an operation succeeded or failed."},
		FaqQuestion{Text: "What is the purpose of the `defer` statement?", Answer: "The `defer` statement in Go is used to delay the execution of a function until the surrounding function returns. It's commonly used for cleanup tasks like closing files or releasing resources."},
		FaqQuestion{Text: "How do I define a struct in Go?", Answer: "A struct in Go is defined using the `type` keyword followed by the struct name and fields. For example, `type Person struct { Name string; Age int }` defines a `Person` struct with `Name` and `Age` fields."},
		FaqQuestion{Text: "Can Go be used for web development?", Answer: "Yes, Go is widely used for web development. It has a powerful `net/http` package that makes it easy to build HTTP servers, and popular frameworks like Gin and Echo provide more features for web development."},
		FaqQuestion{Text: "What is the `interface{}` type in Go?", Answer: "`interface{}` is the empty interface in Go, which can hold values of any type. It is commonly used for functions that need to accept values of various types."},
	)
	context := FaqTemplateContext{
		Title:     "FAQ Section",
		Questions: questions,
	}

	renderTemplate(res, "faq.html", context)
}

type MainRouter struct {
}

func (MainRouter) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path

	switch path {
	case "/":
		faqRouteHandler(res, req)
	default:
		http.Error(res, http.StatusText(http.StatusNotFound), http.StatusNotFound) // alternative
	}
}

func main() {
	defer func() {
		err := recover()
		log.Println(err)
	}()
	var router MainRouter
	fmt.Println("Listening on :3000 ...")

	http.ListenAndServe(":3000", router)
}
