package migrations

import (
	"github.com/ardiost/golang-clean-web-api/constants"
	"github.com/ardiost/golang-clean-web-api/data/db"
	"github.com/ardiost/golang-clean-web-api/data/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Up_1() {
	database := db.GetDb()

	createTables(database)
	createDefaultInformation(database)

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

	database.Migrator().CreateTable(tables...)

}

func addNewTable(database *gorm.DB, models interface{}, tables []interface{}) []interface{} {
	if !database.Migrator().HasTable(models) {
		tables = append(tables, models)
	}

	return tables
}
func createDefaultInformation(database *gorm.DB) {
	adminRole := models.Role{Name: constants.AdminRoleName}
	createRoleIfNotExist(database, &adminRole)
	defaultRole := models.Role{Name: constants.AdminRoleName}
	createRoleIfNotExist(database, &defaultRole)
	u := models.User{UserName: constants.DefaultUserName, FirstName: "Test", LastName: "Test", MobileNumber: "09130954259", Email: "test@gmail.com"}
	pass := "123456"
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	u.Password = string(hashPassword)
	createAdminUserIfNotExist(database, &u, adminRole.Id)
}

func createRoleIfNotExist(database *gorm.DB, r *models.Role) {
	exist := 0
	database.Model(&models.Role{}).
		Select("1").
		Where("name = ?", r.Name).
		First(&exist)
	if exist == 0 {
		database.Create(r)
	}
}

func createAdminUserIfNotExist(database *gorm.DB, u *models.User, roleId int) {
	exist := 0
	database.Model(&models.User{}).
		Select("1").
		Where("user_name = ?", u.UserName).
		First(&exist)
	if exist == 0 {
		database.Create(u)
		ur := models.UserRole{UserId: u.Id, RoleId: roleId}
		database.Create(&ur)
	}
}
func Down_1() {

}
