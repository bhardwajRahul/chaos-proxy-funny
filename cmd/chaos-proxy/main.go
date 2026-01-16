package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	upstream := "https://jsonplaceholder.typicode.com"
	upstreamURL, err := url.Parse((upstream))
	if err != nil {
		log.Fatal("Invalid upstream URL:", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(upstreamURL)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[PROXY] %s %s -> %s%s\n", r.Method, r.URL.Path, upstream, r.URL.Path)
		proxy.ServeHTTP(w, r)
	})

	addr := ":8080"
	fmt.Printf("Starting chaos-proxy on %s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
