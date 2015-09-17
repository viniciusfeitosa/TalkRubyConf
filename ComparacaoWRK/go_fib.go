package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	mux := httprouter.New()
	mux.GET("/:number", indexHandle)

	log.Fatal(http.ListenAndServe(":4567", mux))
}

func indexHandle(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	n, _ := strconv.Atoi(ps.ByName("number"))
	fmt.Fprintf(w, "Go + httprouter fib(%s): %d", ps.ByName("number"), fib(n))
}

func fib(n int) int {
	if n == 0 {
		return n
	}
	if n == 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
