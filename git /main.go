package main

import (
	"fmt"
	"os"

	"github.com/go-rod/rod"
	"github.com/mercadolibre/fury_go-platform/pkg/fury"
	"github.com/xuri/excelize/v2"

	//rodbrowser "github.com/mercadolibre/fury_mprc-automation-test/internal/infrastructure/browsers/rodbrowsers"

	service "github.com/mercadolibre/fury_mprc-automation-test/internal/core/services"
	csvservice "github.com/mercadolibre/fury_mprc-automation-test/internal/infrastructure/files-service/csv"
	xlsservice "github.com/mercadolibre/fury_mprc-automation-test/internal/infrastructure/files-service/xls"
	furyhandler "github.com/mercadolibre/fury_mprc-automation-test/internal/infrastructure/handlers/fury"
	mysql "github.com/mercadolibre/fury_mprc-automation-test/internal/infrastructure/repositories/mysql"
	gorodservice "github.com/mercadolibre/fury_mprc-automation-test/internal/infrastructure/scrappers/go-rod"
)

func main() {
	// RunApp da error, mirar luego, simplemente no arranca, algun parametro no esta siendo enviado correctamente.
	// furyHandler, err := RunApp()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// SetupRoutesFuryHandler(furyApp, furyHandler)
	//}

	err := run()
	if err != nil {
		fmt.Println(err)
	}
}

func run() error {

	db, err := mysql.GetConnectionDB()
	if err != nil {
		fmt.Println(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}

	app, err := fury.NewWebApplication()
	if err != nil {
		return err
	}

	rb := rod.New().MustConnect()
	defer rb.MustClose()
	//rodBrowser := rodbrowser.NewGoRodBrowser(rb)

	var xls *excelize.File
	var csv *os.File

	mysqlRepository := mysql.NewPersonRepository(*db)
	serviceMySQL := service.NewAutoTestService(mysqlRepository)
	serviceGoRod := gorodservice.NewGoRodService(rb)
	serviceCSV := csvservice.NewCsvService(csv)
	serviceXLS := xlsservice.NewXlsService(xls)
	furyHandler := furyhandler.NewHTTPFuryHandler(serviceMySQL, serviceGoRod, serviceCSV, serviceXLS)

	SetupRoutesFuryHandler(app, furyHandler)

	return app.Run()
}
