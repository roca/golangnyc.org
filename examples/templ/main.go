package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Hello("Golang NYC").Render(r.Context(), w)
	})

	fmt.Println("Listening on port 3000")
	http.ListenAndServe(":3000", nil)
}
