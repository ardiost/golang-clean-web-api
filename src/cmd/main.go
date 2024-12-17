package main

import (
	"log"

	"github.com/ardiost/golang-clean-web-api/api"
	"github.com/ardiost/golang-clean-web-api/config"
	"github.com/ardiost/golang-clean-web-api/data/cache"
	"github.com/ardiost/golang-clean-web-api/data/db"
	"github.com/ardiost/golang-clean-web-api/data/db/migrations"
)

func main() {
	// @securityDefinitions.apikey AuthBearer
	// @in header
	// @name Authorization
	cfg := config.GetConfig()
	err := cache.InitRedis(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer cache.CloseRedis()

	err = db.InitDb(cfg)

	if err != nil {
		log.Fatal(err)
	}

	migrations.Up_1()

	defer db.CloseDb()

	api.InitServer(cfg)
}
