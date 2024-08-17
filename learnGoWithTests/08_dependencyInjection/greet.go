package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Use the writer to send the greeting to the buffer in our test.
// Remember fmt.Fprintf is like fmt.Printf but instead takes a Writer to send the string to,
// whereas fmt.Printf defaults to stdout.

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hellow, %s", name)
}

// The Greet function is called with os.Stdout as the writer and "Ellodie" as the name.
/*func main() {
	Greet(os.Stdout, "Ellodie")
}
*/

// The Greet function is called with http.ResponseWriter as the writer.
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
