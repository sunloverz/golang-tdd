package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
	// Greet(os.Stdout, "Elodie\n")
}

func Greet(writer io.Writer, name string) {
	// fmt.Printf("hello, %s", name)
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}
