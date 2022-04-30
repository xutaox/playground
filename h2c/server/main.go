package main

import (
	"fmt"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
)

func main() {
	h2s := &http2.Server{}
	h1s := &http.Server{
		Addr: ":8972",
		Handler: h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("this is a http2 test sever")
			w.Write([]byte("this is a http2 test sever"))
		}), h2s),
	}
	log.Fatal(h1s.ListenAndServe())
}


