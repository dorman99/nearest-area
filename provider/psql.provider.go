package provider

import (
	"fmt"

	"github.com/dorman99/area-scan/config"
	"github.com/dorman99/area-scan/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPsql() *gorm.DB {
	psqlConfig := config.GetDatabaseConfig()
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", psqlConfig.Host, psqlConfig.User, psqlConfig.Database, psqlConfig.Password)

	db, err := gorm.Open(postgres.Open(dbUri))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.AddressEntity{})

	return db
}

func ClosePsqlDBConnection(db *gorm.DB) {
	dbSql, err := db.DB()
	if err != nil {
		panic(err)
	}
	dbSql.Close()
}
