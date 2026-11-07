package main

import (
	"fmt"
	"log"

	"github.com/ParthPoudyal/E-commerce-project/infrastructure/database"
)

func main() {
	
	//* : the *gorm.DB struct to be used in repositories 
	db , err := database.Connectdb()
	if err != nil{ 
		log.Panicf("Error connecting to db : %v", err)
	}

	err = database.RunMigrations(db)
	if err != nil{
		log.Panicf("Migration Error : %v" , err)
	}
	// TODO : the use of DB is to be implemented in Repositories 
	fmt.Println(db)
}