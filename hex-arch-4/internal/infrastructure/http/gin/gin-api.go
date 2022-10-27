package ginhandler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	ports "github.com/devpablocristo/golang/hex-arch-4/internal/core/application/ports"
)

type GinHandler struct {
	patientService ports.PersonService
	//doctorService  ports.DoctorService
	ginRouter *gin.Engine
}

func NewGinHandler(r *gin.Engine, ps ports.PersonService) *GinHandler {
	return &GinHandler{
		patientService: ps,
		ginRouter * gin.Engine,
	}
}

func (h *GinHandler) Run(port string) {
	if port == "default" {
		port = ":8000"
	}

	log.Printf("Gin Server listening on %s\n", port)
	log.Fatal(http.ListenAndServe(port, h.GinRouter))
}

func (h *GinHandler) SetupRoutes() {
	h.GinRouter.GET("/person", func(ctx *gin.Context) {
		h.List(ctx)
	})
	h.GinRouter.POST("/person", func(ctx *gin.Context) {
		h.Register(ctx)
	})
}

// package gamehdl

// type HTTPHandler struct {
// 	gamesService ports.GamesService
// }

// func NewHTTPHandler(gamesService ports.GamesService) *HTTPHandler {
// 	return &HTTPHandler{
// 		gamesService: gamesService,
// 	}
// }

// func (hdl *HTTPHandler) Get(c *gin.Context) {
// 	game, err := hdl.gamesService.Get(c.Param("id"))
// 	if err != nil {
// 		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
// 		return
// 	}

// 	c.JSON(200, game)
// }
