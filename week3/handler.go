package week3

import (
	"fmt"
	"net/http"
)

type MyHandler struct {
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, world")
}
