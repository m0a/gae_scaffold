package gaescaffold

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello GAE/go")
}

var _ http.HandlerFunc = handler

func init() {
	http.HandleFunc("/", handler)
}
