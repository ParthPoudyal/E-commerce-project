package services

import (
	"time"

	"github.com/ParthPoudyal/E-commerce-project/internal/domain"
	"github.com/ParthPoudyal/E-commerce-project/internal/usecase/item"
	"github.com/ParthPoudyal/E-commerce-project/internal/usecase/purchase"
	"github.com/ParthPoudyal/E-commerce-project/internal/usecase/user"
	"github.com/google/uuid"
)

type CreatePurchaseService struct {
	purchaseRepo purchase.PurchaseRepo
	userRepo     user.UserRepo
	itemRepo     item.ItemRepo
}

func NewCreatePurchaseService(pR purchase.PurchaseRepo, uR user.UserRepo, iR item.ItemRepo) (*CreatePurchaseService, error) {
	return &CreatePurchaseService{
		purchaseRepo: pR,
		itemRepo:     iR,
		userRepo:     uR,
	}, nil
}

func (cp *CreatePurchaseService) Execute(userID uuid.UUID, itemID uuid.UUID, price float64) (*domain.Purchase, error) {
	if _, err := cp.userRepo.FindUserByID(userID); err != nil {
		return nil, err
	}

	it, err := cp.itemRepo.FindItemByID(itemID)
	if err != nil {
		return nil, domain.ErrItemNotFound
	}

	if it.Status != domain.AVAILABLE && it.Status != domain.BIDDING {
		return nil, domain.ErrItemNotAvailable
	}

	if it.PostedBy != nil && it.PostedBy.ID == userID {
		return nil, domain.ErrCannotBuyOwnItem
	}

	newPurchase := &domain.Purchase{
		ID:          uuid.New(),
		UserID:      userID,
		ItemID:      itemID,
		PricePaid:   price,
		PurchasedAt: time.Now(),
	}

	err = cp.purchaseRepo.AddPurchase(newPurchase)
	if err != nil {
		return nil, err
	}

	it.Status = domain.SOLD
	now := time.Now()
	it.SoldAt = &now
	_ = cp.itemRepo.UpdateItem(it)

	return newPurchase, nil
}
