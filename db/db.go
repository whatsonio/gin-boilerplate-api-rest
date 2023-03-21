package db

import (
	"app/config"
	"app/infra/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var conn *gorm.DB

func Init() *gorm.DB {

	var err error
	c := config.GetConfig()

	conn, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  c.Db.GetPostgresConnectionInfo(),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		logger.Fatalf("Connexion a la bdd impossible")
		return nil
	}

	return conn
}

func GetDB() *gorm.DB {

	return conn
}
