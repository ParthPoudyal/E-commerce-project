package services

import (
	"github.com/ParthPoudyal/E-commerce-project/internal/domain"
	"github.com/ParthPoudyal/E-commerce-project/internal/usecase/item"
	"github.com/ParthPoudyal/E-commerce-project/internal/usecase/purchase"
	"github.com/google/uuid"
)

type BuyNowService struct {
	itemRepo     item.ItemRepo
	purchaseRepo purchase.PurchaseRepo
}

func NewBuyNowService(iR item.ItemRepo, pR purchase.PurchaseRepo) (*BuyNowService, error) {
	return &BuyNowService{
		itemRepo:     iR,
		purchaseRepo: pR,
	}, nil
}

func (bn *BuyNowService) Execute(itemID, buyerID uuid.UUID, price float64) (*domain.Purchase, error) {

	it, err := bn.itemRepo.FindItemByID(itemID)
	if err != nil {
		return nil, domain.ErrItemNotFound
	}
	if it.Status != domain.AVAILABE {
		return nil, domain.ErrItemNotAvailable
	}
	if it.PostedBy.ID == buyerID {
		return nil, domain.ErrCannotBuyOwnItem
	}

	newPurchase := &domain.Purchase{
		ID:        uuid.New(),
		UserID:    buyerID,
		ItemID:    itemID,
		PricePaid: price,
	}

	if err := bn.purchaseRepo.AddPurchase(newPurchase); err != nil {
		return nil, err
	}

	it.Status = domain.SOLD
	if err := bn.itemRepo.UpdateItem(it); err != nil {
		return nil, err
	}
	return newPurchase, nil
}
