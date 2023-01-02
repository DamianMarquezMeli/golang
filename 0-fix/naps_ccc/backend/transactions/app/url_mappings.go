package app

import (
	"github.com/devpablocristo/transactions/controllers/trss"
)

// trss = transactions
// trs = transaction
func urlMap() {

	router.POST("/trs", trss.CreateTrs)
	router.GET("/trs", trss.GetTrss)
	router.GET("/trs/:id", trss.GetTrs)
	router.PUT("/trs/:id", trss.UpdateTrs)
	router.DELETE("/trs/:id", trss.DeleteTrs)
}
