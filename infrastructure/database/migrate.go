package database

import (
	"github.com/ParthPoudyal/E-commerce-project/infrastructure/models"
	"gorm.io/gorm"
	"log"
)

func RunMigrations(db *gorm.DB) error {
	log.Print("Running migrations...")
	err := db.AutoMigrate(
		&models.UserModel{},
		&models.AuthCredentialModel{},
		&models.CategoryModel{},
		&models.ItemModel{},
		&models.BidModel{},
		&models.CartItemModel{},
		&models.WishlistItemModel{},
		&models.PurchaseModel{},
		&models.PaymentTransactionModel{},
	)
	if err != nil {
		return err
	}
	log.Println("migrations complete")
	return nil
}