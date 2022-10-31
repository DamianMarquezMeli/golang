package main

import (
	ginhandler "github.com/devpablocristo/golang/hex-arch-5/internal/infrastructure/handlers/gin"
	"github.com/gin-gonic/gin"
)

func runPatientService(ph *ginhandler.GinHandler) {
	router := gin.New()
	setupPatientsRoutes(router, ph)
	router.Run(webServerPort)
}
