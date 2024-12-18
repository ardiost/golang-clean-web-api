package migrations

import (
	"github.com/ardiost/golang-clean-web-api/constants"
	"github.com/ardiost/golang-clean-web-api/data/db"
	"github.com/ardiost/golang-clean-web-api/data/models"
	"gorm.io/gorm"
)

func Up_1() {
	database := db.GetDb()

	createTables(database)

}

func createTables(database *gorm.DB) {
	tables := []interface{}{}

	country := models.Country{}
	city := models.City{}
	user := models.User{}
	role := models.Role{}
	userRole := models.UserRole{}

	tables = addNewTable(database, country, tables)
	tables = addNewTable(database, city, tables)
	tables = addNewTable(database, user, tables)
	tables = addNewTable(database, role, tables)
	tables = addNewTable(database, userRole, tables)
}

func addNewTable(database *gorm.DB, models interface{}, tables []interface{}) []interface{} {
	if !database.Migrator().HasTable(models) {
		tables = append(tables, models)
	}

	database.Migrator().CreateTable(tables...)
	return tables
}
func createDefaultInformation(database *gorm.DB) {
	adminRole := models.Role{Name: constants.AdminRoleName}
	createRoleIfNotExist(database, &adminRole)
	defaultRole := models.Role{Name: constants.AdminRoleName}
	createRoleIfNotExist(database, &defaultRole)
}

func createRoleIfNotExist(database *gorm.DB, r *models.Role) {
	exist := 0
	database.Model(&models.Role{}).Select("1").Where("name = ?", r.Name).First(&exist)
	if exist == 0 {
		r := &models.Role{Name: constants.DefaultRoleName}
		database.Create(r)
	}
}
func Down_1() {

}
