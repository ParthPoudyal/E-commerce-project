package main

import (
	"log"
	"github.com/ParthPoudyal/E-commerce-project/database"
)

func main() {
	
	db, err := database.Connectdb()
	if err != nil{ 
		log.Fatal(err)
	}
	defer db.Close()

	
}