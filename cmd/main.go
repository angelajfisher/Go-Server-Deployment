package main

import "fmt"
import "net/http"

func main() {
	fmt.Println("Hello world!")

	mux := http.NewServeMux()

	mux.Handle("/", &homeHandler{})

	http.ListenAndServe(":8080", mux)
}

type homeHandler struct{}

	func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("You hit root!"))
	}
