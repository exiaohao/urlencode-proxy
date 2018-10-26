package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	listenPort := os.Getenv("LISTEN")
	targetHost := os.Getenv("PROXY")

	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			u := req.URL
			u.Scheme = "http"
			u.Host = targetHost
			u.RawQuery = url.QueryEscape(u.RawQuery)
		},
	}
	http.HandleFunc("/", proxy.ServeHTTP)
	http.ListenAndServe(":"+listenPort, nil)
}
