package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("Hello! This is netbot-hello service."))
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
