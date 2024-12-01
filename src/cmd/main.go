package main

import (
	"log"

	"github.com/ardiost/golang-clean-web-api/api"
	"github.com/ardiost/golang-clean-web-api/config"
	"github.com/ardiost/golang-clean-web-api/data/cache"
	"github.com/ardiost/golang-clean-web-api/data/db"
)

func main() {
	// @securityDefinitions.apikey AuthBearer
	// @in header
	// @name Authorization
	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)
	err := cache.InitRedis(cfg)
	if err != nil {
		logger.Fatal(logging.Redis,logging.StartUp,err.Error(),nil)
	}
	defer cache.CloseRedis()

	err = db.InitDb(cfg)

	if err != nil {
		logger.Fatal(logging.Postgres,logging.StartUp,err.Error(),nil)
	}
	defer db.CloseDb()

	api.InitServer(cfg)
}
