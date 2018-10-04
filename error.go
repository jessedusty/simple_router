package main

import (
	"fmt"
	"net/http"
)

func IncorrectHostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "No route for host %s. Possibly a mistype?\n", r.Host)
	w.WriteHeader(404)
}

func CannotContactHostHander(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sorry, I could not reach \"%s\" at the moment. Try again in a little while, it should be back then...\n", routeForHost(r.Host))
	w.WriteHeader(500)
}

func WeirdHander(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Woah something weird happened, but I'm fixing it, try it again later\n")
	w.WriteHeader(500)
}
