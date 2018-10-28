package main

import (
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

func main() {
	listenPort := os.Getenv("LISTEN")
	targetHost := os.Getenv("PROXY")

	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			u := req.URL
			u.Scheme = "http"
			u.Host = targetHost
			u.RawQuery = strings.Replace(u.Query().Encode(), "+", "%20", -1)
		},
	}
	http.HandleFunc("/", proxy.ServeHTTP)
	http.ListenAndServe(":"+listenPort, nil)
}
