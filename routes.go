package main

import (
	"log"
	"net"
	"strings"
)

func routeForHost(host string) string {
	split := strings.Split(host, ".")
	if len(split) <= 1 {
		return "home"
	}
	log.Println(split[0])
	return split[0]
}

func isHostAlive(host string) bool {
	_, err := net.LookupIP(host)
	return err == nil
}
