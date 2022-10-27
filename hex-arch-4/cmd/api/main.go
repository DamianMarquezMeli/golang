package main

import (
	"github.com/gin-gonic/gin"
)

const webPort = 8888

func main() {

	db := storage.ConectDB()
	defer db.Close()

	// gin.SetMode(gin.ReleaseMode)
	// r := gin.Default()

	router := gin.New()

}
