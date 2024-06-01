package database

import (
	"fmt"
	"os"

	"3-design-core/utils/logger"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

//Gorm Database orm connection
var Gorm *gorm.DB

// ConnectToDbUsingOrm connects to any mysql db giving a connectionString using gorm
func ConnectToDbUsingOrm(connectionString string) (*gorm.DB, error) {
	const section = "database.ConnectToDbUsingOrm"
	logger.Log.Infoln(section, "starting")

	config := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	database, err := gorm.Open(postgres.Open(connectionString), config)
	if err != nil {
		err = errors.Wrap(err, "Something happened creating orm db")
	}

	logger.Log.Infoln(section, "finished")

	return database, err
}

// Init Initialize the DB connection
func Init() {
	const section = "database.Init"
	var err error

	logger.Log.Infoln(section, "starting")

	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")

	connectionString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?TimeZone=America/Bogota",dbUser,dbPassword,dbHost,dbPort,dbName)

	Gorm, err = ConnectToDbUsingOrm(connectionString)
	if err != nil {
		logger.Log.Errorln(section, err.Error())
	}

	logger.Log.Infoln(section, "finished")
}
