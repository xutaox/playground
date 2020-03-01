package main

import (
	"crypto/tls"
	"fmt"
	"golang.org/x/net/http2"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

func main() {
	client := http.Client{
		//// Skip TLS dial, otherwise client doesn't support http
		Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
				return net.Dial(network, addr)
			},
		},
	}
	resp, err := client.Get("http://localhost:8972/")
	if err != nil {
		log.Fatal(fmt.Errorf("error making request: %v", err))
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Proto, resp.StatusCode)
	fmt.Println(string(bytes))
	resp.Body.Close()
}

