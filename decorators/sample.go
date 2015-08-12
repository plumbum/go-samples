// From http://andyrees.github.io/2015/your-code-a-mess-maybe-its-time-to-bring-in-the-decorators/

package main

import (
	"fmt"
	"log"
	"net/http"
)

func simpleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func decorator(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Do extra stuff here, e.g. check API keys in header,
		// restrict hosts etc
		log.Println("Started", r)
		f(w, r) // call function here
		log.Println("Done")
	}
}

func stringDecorator(fn func(s string)) func(s string) {
	return func(s string) {
		log.Println("BEGIN string decorator")
		fn(s)
		log.Println("END string decorator")
	}
}

func myFunction(s string) {
    fmt.Println(s)
}

func main() {
	f := stringDecorator(myFunction)
	f("Hello Decorator")

	stringDecorator(myFunction)("One line decorator")

	http.HandleFunc("/decorated", decorator(simpleHandler))
	http.HandleFunc("/notdecorated", simpleHandler)

	http.ListenAndServe("127.0.0.1:8080", nil)
}
