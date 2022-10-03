package app

import (
	"github.com/hadiyanfathur/golang-movies/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=hadiyan password=hadiyan dbname=go_movies port=5432 sslmode=disable TimeZone=Asia/Jakarta",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	helper.PanicIfError(err)

	return db
}

// migrate up migrate -database 'postgres://postgres:password@localhost:5432/example?sslmode=disable' -path db/migrations up
