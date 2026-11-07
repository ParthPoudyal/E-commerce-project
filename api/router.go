package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ParthPoudyal/E-commerce-project/api/handlers"
	"github.com/ParthPoudyal/E-commerce-project/api/middleware"
	"github.com/ParthPoudyal/E-commerce-project/api/ws"
)

type RouterDependencies struct {
	AuthHandler     *handlers.AuthHandler
	ItemHandler     *handlers.ItemHandler
	BidHandler      *handlers.BidHandler
	CartHandler     *handlers.CartHandler
	WishlistHandler *handlers.WishlistHandler
	PurchaseHandler *handlers.PurchaseHandler
	PaymentHandler  *handlers.PaymentHandler
	CategoryHandler *handlers.CategoryHandler
	UploadHandler   *handlers.UploadHandler
	WSHub           *ws.Hub
}

func SetupRouter(deps *RouterDependencies) *gin.Engine {
	r := gin.Default()

	// Apply CORS
	r.Use(middleware.CORSMiddleware())

	// Static files for uploaded images
	r.Static("/uploads", "./uploads")

	// Health check
	r.GET("/api/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy", "service": "auction-hub-api"})
	})

	// Live Auction WebSocket
	r.GET("/ws/auction/:id", deps.WSHub.HandleWebSocket)

	api := r.Group("/api")
	{
		// Public Auth
		api.POST("/auth/register", deps.AuthHandler.Register)
		api.POST("/auth/login", deps.AuthHandler.Login)

		// Public Categories
		api.GET("/categories", deps.CategoryHandler.ListCategories)
		api.GET("/categories/:slug", deps.CategoryHandler.GetCategory)

		// Public Items & Auctions
		api.GET("/items", deps.ItemHandler.ListItems)
		api.GET("/items/auctions", deps.ItemHandler.ListAuctions)
		api.GET("/items/search", deps.ItemHandler.SearchItems)
		api.GET("/items/user/:id", deps.ItemHandler.GetItemsByUser)
		api.GET("/items/:id", deps.ItemHandler.GetItem)

		// Public Bids history
		api.GET("/bids/item/:id", deps.BidHandler.GetBidHistory)

		// Public Payment verification callback
		api.POST("/payments/khalti/verify", deps.PaymentHandler.VerifyKhaltiPayment)

		// Protected endpoints
		protected := api.Group("")
		protected.Use(middleware.AuthMiddleware())
		{
			// Auth
			protected.GET("/auth/me", deps.AuthHandler.GetMe)

			// Items management
			protected.POST("/items", deps.ItemHandler.CreateItem)
			protected.PUT("/items/:id", deps.ItemHandler.UpdateItem)
			protected.DELETE("/items/:id", deps.ItemHandler.DeleteItem)
			protected.POST("/upload", deps.UploadHandler.UploadImage)

			// Bidding
			protected.POST("/bids", deps.BidHandler.PlaceBid)
			protected.GET("/bids/user", deps.BidHandler.GetUserBids)

			// Cart
			protected.GET("/cart", deps.CartHandler.GetCart)
			protected.POST("/cart", deps.CartHandler.AddToCart)
			protected.DELETE("/cart/:itemId", deps.CartHandler.RemoveFromCart)
			protected.DELETE("/cart", deps.CartHandler.ClearCart)

			// Wishlist
			protected.GET("/wishlist", deps.WishlistHandler.GetWishlist)
			protected.POST("/wishlist", deps.WishlistHandler.AddToWishlist)
			protected.DELETE("/wishlist/:itemId", deps.WishlistHandler.RemoveFromWishlist)
			protected.GET("/wishlist/check/:itemId", deps.WishlistHandler.CheckWishlisted)

			// Purchases
			protected.GET("/purchases", deps.PurchaseHandler.GetPurchases)
			protected.POST("/purchases/buy-now/:itemId", deps.PurchaseHandler.BuyNow)

			// Payments
			protected.POST("/payments/khalti/initiate", deps.PaymentHandler.InitiateKhaltiPayment)
			protected.GET("/payments/user", deps.PaymentHandler.GetUserPayments)
		}
	}

	return r
}

