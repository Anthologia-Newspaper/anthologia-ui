package auth

import (
	src "auth/sources"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

/*
DatabaseInit initializes the database for the microservice
*/
func DatabaseInit() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database/auth.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	db.AutoMigrate(&src.User{})

	return db
}