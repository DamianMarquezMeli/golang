package main

import (
	ginhandler "github.com/devpablocristo/golang/hex-arch/internal/infrastructure/handlers/gin"
	"github.com/gin-gonic/gin"
)

const webServerPort string = ":8088"

func runPatientService(ph *ginhandler.GinHandler) {
	router := gin.New()
	setupPatientsRoutes(router, ph)
	router.Run(webServerPort)
}
