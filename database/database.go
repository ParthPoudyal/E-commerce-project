package database

import (

	"os"
	
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func Connectdb() (*gorm.DB, error) {

	dsn := os.Getenv("DB_DSN")

	if dsn == "" {
		return nil,&DBerror{message: "DB_DSN not set"}
	}

	db , err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{ 
		return nil , &DBerror{message: "Error connecting to the Database" , cause: err}
	} 
	
	sqlDB, err := db.DB()
	if err != nil{ 
		return nil , err
	}

	if err := sqlDB.Ping(); err != nil { 
		return nil , &DBerror{message: "Cannot ping DB server", cause: err}
	}

	return db , nil 
}
