package main

import (
	ginhandler "github.com/devpablocristo/golang/hex-arch-5/internal/infrastructure/handlers/gin"
	"github.com/gin-gonic/gin"
)

func setupPatientsRoutes(r *gin.Engine, ph *ginhandler.GinHandler) {

	r.GET("/patient/:id", ph.GetPatient)
	r.POST("/patient", ph.CreatePatient)
}
