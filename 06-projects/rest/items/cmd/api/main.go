package main

import (
	"log"

	ctrl "github.com/mercadolibre/items/internal/adapters/controller"
	"github.com/mercadolibre/items/internal/adapters/repository"
	"github.com/mercadolibre/items/internal/infra/mysql"
	"github.com/mercadolibre/items/internal/infra/web"
	"github.com/mercadolibre/items/internal/usecase"
)

func main() {
	//MySQL
	conn, err := mysql.GetConnectionDB()
	if err != nil {
		log.Fatalln(err)
	}
	boorRepository := repository.NewMySQLBookRepository(conn)
	//boorRepository := repository.NewBookRepository()
	bookUsecase := usecase.NewBookUsecase(boorRepository)
	bookController := ctrl.NewBookController(bookUsecase)

	if err := web.NewHTTPServer(bookController); err != nil {
		log.Fatalln(err)
	}
}
