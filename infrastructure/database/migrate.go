package database

import (
	"github.com/ParthPoudyal/E-commerce-project/infrastructure/models"
	"gorm.io/gorm"
	"log"
)

func RunMigrations(db *gorm.DB) error{
	
	log.Print("Running migrations...")
	err := db.AutoMigrate(
		&models.BidModel{}, 
		&models.UserModel{}, 
		&models.CartItemModel{},
		&models.ItemModel{},
		&models.PurchaseModel{},
		&models.UserModel{}, 
		models.WishlistItemModel{}, 
	) 
	if err != nil{ 
		return err
	}  
	log.Println("migrations complete")
	return  nil 
}