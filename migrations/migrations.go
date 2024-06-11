package migrations

import (
	categoryModel "3-design-core/api/categories/models"
	productModel "3-design-core/api/products/models"
	db "3-design-core/database"
	"3-design-core/utils/logger"
)

//Migrate create all tables
func Migrate() {
	const section = "migrations.Migrate"
	logger.Log.Infoln(section, "starting")

	db.Gorm.AutoMigrate(categoryModel.Category{},productModel.Product{})

	logger.Log.Infoln(section, "finished")
}

//DropTables drop all tables
func DropTables() {
	const section = "migrations.DropTables"
	logger.Log.Infoln(section, "starting")

	db.Gorm.Migrator().DropTable(categoryModel.Category{},productModel.Product{})

	logger.Log.Infoln(section, "finished")
}
