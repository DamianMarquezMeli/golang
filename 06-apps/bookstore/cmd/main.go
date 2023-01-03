package main

import (
	"os"
)

const defaultPort = ":8080"

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = defaultPort
	}

	Start(port)
}

func Start(port string) {
	r := routes()
	server := newServer(port, r)
	server.Start()
}
