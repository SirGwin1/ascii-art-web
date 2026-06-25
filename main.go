package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "welcome")
}

func main() {
	http.HandleFunc("/", home)

	fmt.Println("server running")

	http.ListenAndServe(":8080", nil)

}
