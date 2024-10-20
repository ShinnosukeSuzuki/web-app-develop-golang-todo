package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	err := http.ListenAndServe(
		":18080",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, r.URL.Path[1:])
		}),
	)
	if err != nil {
		fmt.Println("failed to start server", err)
		os.Exit(1)
	}
}
