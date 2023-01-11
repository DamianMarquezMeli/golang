package api

import (
	"log"
	"net/http"
	"time"
)

func RunHttpServer(port string, routes http.Handler) {
	log.Println("starting chi server")

	sv := &http.Server{
		Addr:         ":" + port,
		Handler:      routes,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	go func() {
		err := sv.ListenAndServe()
		if err != nil {
			if err != http.ErrServerClosed {
				log.Fatalf("could not listen on %s due to %s", sv.Addr, err.Error())
			}
		}
	}()
	log.Printf("the chi server is ready to handle requests %s", sv.Addr)
}
