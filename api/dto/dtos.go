package dto

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"github.com/ParthPoudyal/E-commerce-project/internal/domain"
)

// Auth DTOs
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Address  string `json:"address"`
	PhoneNo  string `json:"phone_no"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Token string       `json:"token"`
	User  *domain.User `json:"user"`
}

// Item DTOs
type CreateItemRequest struct {
	Name        string           `json:"name" binding:"required"`
	Description string           `json:"description"`
	CategoryID  *uuid.UUID       `json:"category_id"`
	ImageURL    string           `json:"image_url"`
	Price       decimal.Decimal  `json:"price" binding:"required"`
	IsAuction   bool             `json:"is_auction"`
	StartingBid *decimal.Decimal `json:"starting_bid"`
	DurationDays int             `json:"duration_days"` // e.g., 3, 5, 7 days for auction
}

type UpdateItemRequest struct {
	Name        string           `json:"name"`
	Description string           `json:"description"`
	CategoryID  *uuid.UUID       `json:"category_id"`
	ImageURL    string           `json:"image_url"`
	Price       *decimal.Decimal `json:"price"`
	Status      *domain.ItemStatus `json:"status"`
}

type ItemResponse struct {
	Item        domain.Item `json:"item"`
	HighestBid  *domain.Bid `json:"highest_bid,omitempty"`
	TotalBids   int         `json:"total_bids"`
	TimeLeftSec int64       `json:"time_left_seconds,omitempty"`
}

// Bid DTOs
type PlaceBidRequest struct {
	ItemID uuid.UUID       `json:"item_id" binding:"required"`
	Amount decimal.Decimal `json:"amount" binding:"required"`
}

type BidResponse struct {
	Bid                 *domain.Bid `json:"bid"`
	PreviousHighestUser *uuid.UUID  `json:"previous_highest_user,omitempty"`
	IsNewHighest        bool        `json:"is_new_highest"`
}

// Cart DTOs
type AddToCartRequest struct {
	ItemID uuid.UUID `json:"item_id" binding:"required"`
}

type CartResponseItem struct {
	CartItem domain.CartItem `json:"cart_item"`
	Item     *domain.Item    `json:"item"`
}

// Wishlist DTOs
type AddToWishlistRequest struct {
	ItemID uuid.UUID `json:"item_id" binding:"required"`
}

type WishlistResponseItem struct {
	WishlistItem domain.WishlistItem `json:"wishlist_item"`
	Item         *domain.Item        `json:"item"`
}

// Payment DTOs
type InitiatePaymentRequest struct {
	ItemID uuid.UUID       `json:"item_id" binding:"required"`
	Amount decimal.Decimal `json:"amount"`
}

type VerifyPaymentRequest struct {
	PIDX string `json:"pidx" binding:"required"`
}

// Category DTOs
type CreateCategoryRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
}

// Generic API response
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

