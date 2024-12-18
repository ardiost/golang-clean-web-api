package migrations

import (
	"log"

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
	u := models.User{UserName: constants.DefaultUserName, FirstName: "مصطفی", LastName: "Test", MobileNumber: "09130954259", Email: "test@gmail.com"}
	pass := "123456"
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	u.Password = string(hashPassword)
	createAdminUserIfNotExist(database, &u, adminRole.Id)
}

func createRoleIfNotExist(database *gorm.DB, r *models.Role) {
	var count int64
	// بررسی وجود رکورد در جدول roles
	database.Model(&models.Role{}).
		Where("name = ?", r.Name).
		Count(&count)

	if count == 0 {
		// ایجاد رکورد اگر وجود ندارد
		if err := database.Create(r).Error; err != nil {
			log.Println("Error creating role:", err)
		}
	}
}

func createAdminUserIfNotExist(database *gorm.DB, u *models.User, roleId int) {
	var count int64
	// بررسی وجود کاربر در جدول users
	database.Model(&models.User{}).
		Where("user_name = ?", u.UserName).
		Count(&count)

	if count == 0 {
		// ایجاد کاربر
		if err := database.Create(u).Error; err != nil {
			log.Println("Error creating user:", err)
			return
		}

		// ایجاد ارتباط UserRole با استفاده از u.Id به‌روز شده
		ur := &models.UserRole{UserId: u.Id, RoleId: roleId}
		if err := database.Create(ur).Error; err != nil {
			log.Println("Error creating user role:", err)
		}
	}
}

func Down_1() {

}
