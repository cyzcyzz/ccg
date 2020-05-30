package main

import (
	"ccg/ccg"
	"fmt"
	"net/http"
)

func main() {
	r := ccg.New()
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})
}
