package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	log.Println("Running")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		proxyHandler(r)(w, r)
	})

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}

func proxyHandler(r *http.Request) func(http.ResponseWriter, *http.Request) {
	log.Printf("Connection from %s for %s and %s\n", r.RemoteAddr, r.Host, r.URL)
	host := routeForHost(r.Host)
	if isHostAlive(host) {
		urlVal, err := url.Parse("http://" + host)
		if err != nil {
			return IncorrectHostHandler
		}

		return func(w http.ResponseWriter, r *http.Request) {
			p := httputil.NewSingleHostReverseProxy(urlVal)
			r.Header.Add("X-Forwarded-Host", r.Host)
			r.Header.Add("X-Origin-Host", r.Host)
			r.Header.Add("X-Real-IP", r.RemoteAddr)

			// For some services - I use cloudflare on the frontend
			r.Header.Add("X-Forwarded-Proto", "http")
			p.ServeHTTP(w, r)
		}
	} else {
		return CannotContactHostHander
	}
}
