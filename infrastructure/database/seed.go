package database

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/ParthPoudyal/E-commerce-project/infrastructure/models"
)

func SeedDatabase(db *gorm.DB) error {
	var count int64
	db.Model(&models.CategoryModel{}).Count(&count)
	if count > 0 {
		return nil // Already seeded
	}

	log.Println("Seeding initial categories, demo user, and live auction lots...")

	// 1. Categories
	categories := []models.CategoryModel{
		{
			ID:          uuid.New(),
			Name:        "Electronics",
			Slug:        "electronics",
			Description: "High-end smartphones, premium audio gear, laptops, and cameras.",
			ImageURL:    "https://images.unsplash.com/photo-1550009158-9ebf69173e03?auto=format&fit=crop&w=600&q=80",
		},
		{
			ID:          uuid.New(),
			Name:        "Watches & Jewelry",
			Slug:        "watches-jewelry",
			Description: "Luxury Swiss timepieces, fine gold jewelry, diamonds and accessories.",
			ImageURL:    "https://images.unsplash.com/photo-1523275335684-37898b6baf30?auto=format&fit=crop&w=600&q=80",
		},
		{
			ID:          uuid.New(),
			Name:        "Art & Collectibles",
			Slug:        "art-collectibles",
			Description: "Original paintings, rare sculptures, vintage coins, and memorabilia.",
			ImageURL:    "https://images.unsplash.com/photo-1579783902614-a3fb3927b675?auto=format&fit=crop&w=600&q=80",
		},
		{
			ID:          uuid.New(),
			Name:        "Sneakers & Fashion",
			Slug:        "sneakers-fashion",
			Description: "Limited edition sneakers, designer fashion wear, and luxury bags.",
			ImageURL:    "https://images.unsplash.com/photo-1552346154-21d32810aba3?auto=format&fit=crop&w=600&q=80",
		},
		{
			ID:          uuid.New(),
			Name:        "Antiques & Vintage",
			Slug:        "antiques-vintage",
			Description: "Authentic mid-century furniture, vintage vinyl players, and historic items.",
			ImageURL:    "https://images.unsplash.com/photo-1533090161767-e6ffed986c88?auto=format&fit=crop&w=600&q=80",
		},
		{
			ID:          uuid.New(),
			Name:        "Sports Memorabilia",
			Slug:        "sports-memorabilia",
			Description: "Autographed jerseys, game-used equipment, and iconic sports trading cards.",
			ImageURL:    "https://images.unsplash.com/photo-1518611012118-696072aa579a?auto=format&fit=crop&w=600&q=80",
		},
		{
			ID:          uuid.New(),
			Name:        "Automotive & Motors",
			Slug:        "automotive-motors",
			Description: "Classic vehicles, vintage cafe racers, supercars, and premium parts.",
			ImageURL:    "https://images.unsplash.com/photo-1503376780353-7e6692767b70?auto=format&fit=crop&w=600&q=80",
		},
		{
			ID:          uuid.New(),
			Name:        "Books & Manuscripts",
			Slug:        "books-manuscripts",
			Description: "First edition classics, signed literature, historical documents, and rare prints.",
			ImageURL:    "https://images.unsplash.com/photo-1512820790803-83ca734da794?auto=format&fit=crop&w=600&q=80",
		},
	}

	catMap := make(map[string]uuid.UUID)
	for _, cat := range categories {
		_ = db.Create(&cat)
		catMap[cat.Slug] = cat.ID
	}

	// 2. Demo User (Seller)
	sellerID := uuid.New()
	hashedPwd, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	pwdStr := string(hashedPwd)

	demoUser := models.UserModel{
		ID:        sellerID,
		Username:  "parth_consignor",
		Address:   "Kathmandu, Nepal",
		PhoneNo:   "9800000001",
		Email:     "parth@ecom.np",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_ = db.Create(&demoUser)

	authCred := models.AuthCredentialModel{
		UserID:       sellerID,
		PasswordHash: &pwdStr,
		Provider:     "local",
	}
	_ = db.Create(&authCred)

	// Demo Bidder User
	bidderID := uuid.New()
	demoBidder := models.UserModel{
		ID:        bidderID,
		Username:  "viraj_collector",
		Address:   "Pokhara, Nepal",
		PhoneNo:   "9800000002",
		Email:     "viraj@ecom.np",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_ = db.Create(&demoBidder)

	bidderCred := models.AuthCredentialModel{
		UserID:       bidderID,
		PasswordHash: &pwdStr,
		Provider:     "local",
	}
	_ = db.Create(&bidderCred)

	// 3. Sample Live Auctions and Fixed Price Items
	now := time.Now()
	twoDays := now.Add(48 * time.Hour)
	fiveDays := now.Add(120 * time.Hour)
	sevenDays := now.Add(168 * time.Hour)

	watchCat := catMap["watches-jewelry"]
	elecCat := catMap["electronics"]
	artCat := catMap["art-collectibles"]
	snkCat := catMap["sneakers-fashion"]
	autCat := catMap["automotive-motors"]

	items := []models.ItemModel{
		{
			ID:          uuid.New(),
			Name:        "1972 Vintage Rolex Submariner Ref. 5513 'Matte Dial'",
			Description: "Pristine collector-grade Rolex Submariner Ref 5513 with original tritium dial and fat-font bezel. Serviced in Geneva with documentation and period box.",
			CategoryID:  &watchCat,
			ImageURL:    "https://images.unsplash.com/photo-1523275335684-37898b6baf30?auto=format&fit=crop&w=800&q=80",
			PostedByID:  sellerID,
			Price:       decimal.NewFromFloat(185000),
			Status:      models.StatusBidding,
			ValidFrom:   now,
			ValidTo:     &twoDays,
			IsAuction:   true,
			StartingBid: decimalPtr(decimal.NewFromFloat(150000)),
		},
		{
			ID:          uuid.New(),
			Name:        "Leica M3 Rangefinder Single Stroke w/ 50mm Summicron f/2",
			Description: "Iconic 1957 Leica M3 single-stroke rangefinder body with clean dual-range 50mm Summicron lens. Fully functioning shutter speeds with original leather case.",
			CategoryID:  &elecCat,
			ImageURL:    "https://images.unsplash.com/photo-1516035069371-29a1b244cc32?auto=format&fit=crop&w=800&q=80",
			PostedByID:  sellerID,
			Price:       decimal.NewFromFloat(85000),
			Status:      models.StatusBidding,
			ValidFrom:   now,
			ValidTo:     &fiveDays,
			IsAuction:   true,
			StartingBid: decimalPtr(decimal.NewFromFloat(75000)),
		},
		{
			ID:          uuid.New(),
			Name:        "Nike Air Jordan 1 Retro High 'Chicago' 1985 Original",
			Description: "Authentic 1985 Nike Air Jordan 1 High Chicago. Original laces, leather patina, preserved ankle collars with provenance certificate.",
			CategoryID:  &snkCat,
			ImageURL:    "https://images.unsplash.com/photo-1552346154-21d32810aba3?auto=format&fit=crop&w=800&q=80",
			PostedByID:  sellerID,
			Price:       decimal.NewFromFloat(120000),
			Status:      models.StatusBidding,
			ValidFrom:   now,
			ValidTo:     &sevenDays,
			IsAuction:   true,
			StartingBid: decimalPtr(decimal.NewFromFloat(90000)),
		},
		{
			ID:          uuid.New(),
			Name:        "Original Contemporary Acrylic on Canvas by Master Artist",
			Description: "Signed 36x48 inch original abstract expressionist canvas titled 'Solitude and Echoes'. Framed in solid oak floating frame.",
			CategoryID:  &artCat,
			ImageURL:    "https://images.unsplash.com/photo-1579783902614-a3fb3927b675?auto=format&fit=crop&w=800&q=80",
			PostedByID:  sellerID,
			Price:       decimal.NewFromFloat(45000),
			Status:      models.StatusAvailable,
			ValidFrom:   now,
			IsAuction:   false,
		},
		{
			ID:          uuid.New(),
			Name:        "1968 Bell Star Vintage Motorcycle Helmet (Restored)",
			Description: "Authentic restored Bell Star helmet from 1968 in vintage cream with custom hand-painted gold pin-striping. Display ready with wall mount stand.",
			CategoryID:  &autCat,
			ImageURL:    "https://images.unsplash.com/photo-1558981403-c5f9899a28bc?auto=format&fit=crop&w=800&q=80",
			PostedByID:  sellerID,
			Price:       decimal.NewFromFloat(32000),
			Status:      models.StatusAvailable,
			ValidFrom:   now,
			IsAuction:   false,
		},
	}

	for _, itm := range items {
		_ = db.Create(&itm)

		// Seed initial bid on the first item
		if itm.IsAuction {
			bidAmount := itm.Price
			bid := models.BidModel{
				ID:        uuid.New(),
				ItemID:    itm.ID,
				BidderID:  bidderID,
				Amount:    bidAmount,
				CreatedAt: now.Add(-2 * time.Hour),
			}
			_ = db.Create(&bid)
		}
	}

	log.Printf("Database seeded successfully with categories, demo users, and live auction lots.")
	return nil
}

func decimalPtr(d decimal.Decimal) *decimal.Decimal {
	return &d
}
