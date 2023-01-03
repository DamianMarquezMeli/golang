package chiAdapter

import (
	"log"
	"net/http"
)

func (h *ChiHandler) RunChiServer() {
	log.Println("starting chi server")
	go func() {
		err := h.httpServer.ListenAndServe()
		if err != nil {
			if err != http.ErrServerClosed {
				log.Fatalf("could not listen on %s due to %s", h.httpServer.Addr, err.Error())
			}
		}
	}()
	log.Printf("the chi server is ready to handle requests %s", h.httpServer.Addr)
	//gracefulChiServerShutdown()
}
