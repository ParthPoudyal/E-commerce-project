package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/ParthPoudyal/E-commerce-project/api/handlers"
	"github.com/ParthPoudyal/E-commerce-project/api/ws"
	"github.com/ParthPoudyal/E-commerce-project/infrastructure/database"
	"github.com/ParthPoudyal/E-commerce-project/infrastructure/repository"
	authservices "github.com/ParthPoudyal/E-commerce-project/internal/usecase/auth/services"
	bidservices "github.com/ParthPoudyal/E-commerce-project/internal/usecase/bid/services"
	cartservices "github.com/ParthPoudyal/E-commerce-project/internal/usecase/cart/services"
	categoryservices "github.com/ParthPoudyal/E-commerce-project/internal/usecase/category/services"
	paymentservices "github.com/ParthPoudyal/E-commerce-project/internal/usecase/payment/services"
	purchaseservices "github.com/ParthPoudyal/E-commerce-project/internal/usecase/purchase/services"
	wishlistservices "github.com/ParthPoudyal/E-commerce-project/internal/usecase/wishlist/services"
)

func main() {
	_ = godotenv.Load()

	// 1. Connect to Database
	db, err := database.Connectdb()
	if err != nil {
		log.Printf("Warning: Database connection failed: %v. Running in offline/mock mode.", err)
	} else {
		log.Println("Connected to PostgreSQL successfully.")

		// 2. Run Database AutoMigrations
		if err := database.RunMigrations(db); err != nil {
			log.Printf("Migration error: %v", err)
		} else {
			// 3. Seed Initial Database Content
			_ = database.SeedDatabase(db)
		}
	}

	// 4. Repositories
	userRepo := repository.NewPostgresUserRepository(db)
	itemRepo := repository.NewPostgresItemRepository(db)
	bidRepo := repository.NewPostgresBidRepository(db)
	cartRepo := repository.NewPostgresCartRepository(db)
	purchaseRepo := repository.NewPostgresPurchaseRepository(db)
	wishlistRepo := repository.NewPostgresWishlistRepository(db)
	categoryRepo := repository.NewPostgresCategoryRepository(db)
	paymentRepo := repository.NewPostgresPaymentRepository(db)

	// 5. Use Case Services
	authService := authservices.NewAuthService(userRepo)
	placeBidService := bidservices.NewPlaceBidService(bidRepo, itemRepo)
	addToCartService := cartservices.NewAddtoCartUseCase(cartRepo, itemRepo)
	removeFromCartService := cartservices.NewRemoveFromCartService(cartRepo)
	buyNowService, _ := purchaseservices.NewBuyNowService(itemRepo, purchaseRepo)
	wishlistService := wishlistservices.NewWishlistService(wishlistRepo, itemRepo)
	categoryService := categoryservices.NewCategoryService(categoryRepo)
	khaltiService := paymentservices.NewKhaltiService(paymentRepo, itemRepo, purchaseRepo, userRepo)

	// 6. WebSocket Live Auction Hub
	wsHub := ws.NewHub()
	go wsHub.Run()

	// 7. HTTP Handlers
	authHandler := handlers.NewAuthHandler(authService)
	itemHandler := handlers.NewItemHandler(itemRepo, bidRepo, userRepo)
	bidHandler := handlers.NewBidHandler(placeBidService, bidRepo, wsHub)
	cartHandler := handlers.NewCartHandler(cartRepo, itemRepo, addToCartService, removeFromCartService)
	wishlistHandler := handlers.NewWishlistHandler(wishlistService, itemRepo)
	purchaseHandler := handlers.NewPurchaseHandler(purchaseRepo, itemRepo, buyNowService)
	paymentHandler := handlers.NewPaymentHandler(khaltiService, paymentRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)
	uploadHandler := handlers.NewUploadHandler("./uploads")

	// 8. Router Setup
	router := SetupRouter(&RouterDependencies{
		AuthHandler:     authHandler,
		ItemHandler:     itemHandler,
		BidHandler:      bidHandler,
		CartHandler:     cartHandler,
		WishlistHandler: wishlistHandler,
		PurchaseHandler: purchaseHandler,
		PaymentHandler:  paymentHandler,
		CategoryHandler: categoryHandler,
		UploadHandler:   uploadHandler,
		WSHub:           wsHub,
	})

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	addr := fmt.Sprintf(":%s", port)
	log.Printf("AuctionHub Go API Server running on http://localhost%s", addr)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Server failed to run: %v", err)
	}
}
