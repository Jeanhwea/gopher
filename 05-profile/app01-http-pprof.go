package main

import (
	"net/http"
	_ "net/http/pprof"
)

// http://localhost:8888/debug/pprof

func main() {
	http.ListenAndServe(":8888", nil)
}
