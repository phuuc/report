package main

import (
	"flag"
	"log"

	"github.com/finnpn/overview/config"
	"github.com/finnpn/overview/infras"
	"github.com/finnpn/overview/infras/repo"
	"github.com/finnpn/overview/infras/usecases"
	"github.com/finnpn/overview/interfaces/restapi/handlers"
	"github.com/finnpn/overview/server"
)

func main() {
	configFile := flag.String("config-file", "", "config file path")
	flag.Parse()

	cfg, err := config.Load(*configFile)
	if err != nil {
		log.Fatalf("failed read config file with err=%v", err)
	}
	db := infras.NewDB(cfg)
	sqlDb := db.SetupDB()
	db.RunMigration()

	// repo
	overviewRepo := repo.NewOverviewRepo(sqlDb)
	flightRepo := repo.NewFlightRepo(sqlDb)
	hotelRepo := repo.NewHotelRepo(sqlDb)

	//usecases

	overviewUc := usecases.NewOverviewUc(overviewRepo)
	fligtUc := usecases.NewFlightUc(flightRepo)
	hotelUc := usecases.NewHotelUc(hotelRepo)

	//handler
	overviewHandler := handlers.NewHandlers(overviewUc, fligtUc, hotelUc)
	server := server.NewRouter(cfg, overviewHandler)
	server.Run()

}
