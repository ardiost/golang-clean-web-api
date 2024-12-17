package migrations

import (
	"github.com/ardiost/golang-clean-web-api/data/db"
	"github.com/ardiost/golang-clean-web-api/data/models"
)

func Up_1() {
	database := db.GetDb()

	tables := []interface{}{}

	country := models.Country{}

	city := models.City{}

	if !database.Migrator().HasTable(country) {
		tables = append(tables, country)
	}

	if !database.Migrator().HasTable(city) {
		tables = append(tables, city)
	}

	database.Migrator().CreateTable(tables...)

}

func Down_1() {

}
