package main

import (
	"BlogApplication/pkg/Router"
	"fmt"
)

func main() {
	fmt.Println("Hello")
	Router.StartServer()
	// r := chi.NewRouter()
	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	json.NewEncoder(w).Encode("Hey you all i got first result")
	// 	// w.Write([]byte("Welcome"))
	// })
	// http.ListenAndServe(":8080", r)
}
