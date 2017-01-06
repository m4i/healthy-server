package main

import "net/http"

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
