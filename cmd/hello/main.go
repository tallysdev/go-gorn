// Click here and start typing.
package main

import (
	"fmt"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(hello)
	http.ListenAndServe(":8080", handler)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
