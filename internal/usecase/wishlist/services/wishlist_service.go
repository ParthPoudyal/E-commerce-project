package services

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/ParthPoudyal/E-commerce-project/internal/domain"
	"github.com/ParthPoudyal/E-commerce-project/internal/usecase/item"
	"github.com/ParthPoudyal/E-commerce-project/internal/usecase/wishlist"
)

type WishlistService struct {
	wishlistRepo wishlist.WishlistRepo
	itemRepo     item.ItemRepo
}

func NewWishlistService(wR wishlist.WishlistRepo, iR item.ItemRepo) *WishlistService {
	return &WishlistService{
		wishlistRepo: wR,
		itemRepo:     iR,
	}
}

func (s *WishlistService) AddToWishlist(userID, itemID uuid.UUID) (*domain.WishlistItem, error) {
	_, err := s.itemRepo.FindItemByID(itemID)
	if err != nil {
		return nil, domain.ErrItemNotFound
	}

	isWishlisted, err := s.wishlistRepo.IsWishlisted(userID, itemID)
	if err == nil && isWishlisted {
		return nil, domain.ErrAlreadyWishlisted
	}

	item := &domain.WishlistItem{
		ID:      uuid.New(),
		UserId:  userID,
		ItemId:  itemID,
		AddedAt: time.Now(),
	}

	if err := s.wishlistRepo.AddItemToWishlist(item); err != nil {
		return nil, fmt.Errorf("add to wishlist: %w", err)
	}

	return item, nil
}

func (s *WishlistService) GetUserWishlist(userID uuid.UUID) ([]domain.WishlistItem, error) {
	return s.wishlistRepo.UserWishList(userID)
}

func (s *WishlistService) RemoveFromWishlist(userID, itemID uuid.UUID) error {
	return s.wishlistRepo.RemoveItemForUser(userID, itemID)
}

func (s *WishlistService) IsWishlisted(userID, itemID uuid.UUID) (bool, error) {
	return s.wishlistRepo.IsWishlisted(userID, itemID)
}

