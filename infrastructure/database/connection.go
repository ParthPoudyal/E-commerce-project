package database

import (
	"os"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func Connectdb() (*gorm.DB, error) {
	
	if err := godotenv.Load() ; err != nil{
		return nil , &DBerror{message : "Could not load a .env file"}
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password='%s' dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

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
