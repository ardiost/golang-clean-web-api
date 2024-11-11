package main

import (
	"log"

	"github.com/ardiost/golang-clean-web-api/api"
	"github.com/ardiost/golang-clean-web-api/config"
	"github.com/ardiost/golang-clean-web-api/data/cache"
	"github.com/ardiost/golang-clean-web-api/data/db"
)

func main() {
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
	defer db.CloseDb()

	api.InitServer(cfg)
}
