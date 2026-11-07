package services

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/ParthPoudyal/E-commerce-project/internal/domain"
	"github.com/ParthPoudyal/E-commerce-project/internal/usecase/bid"
	"github.com/ParthPoudyal/E-commerce-project/internal/usecase/item"
	"github.com/ParthPoudyal/E-commerce-project/internal/usecase/purchase"
)

type AuctionService struct {
	itemRepo     item.ItemRepo
	bidRepo      bid.BidRepository
	purchaseRepo purchase.PurchaseRepo
}

func NewAuctionService(iR item.ItemRepo, bR bid.BidRepository, pR purchase.PurchaseRepo) *AuctionService {
	return &AuctionService{
		itemRepo:     iR,
		bidRepo:      bR,
		purchaseRepo: pR,
	}
}

func (s *AuctionService) CloseAuction(itemID uuid.UUID) (*domain.Item, error) {
	itm, err := s.itemRepo.FindItemByID(itemID)
	if err != nil {
		return nil, domain.ErrItemNotFound
	}

	if !itm.IsAuction {
		return nil, fmt.Errorf("item is not an auction")
	}

	if itm.Status == domain.SOLD {
		return itm, nil
	}

	highestBid, err := s.bidRepo.GetHighestBidForItem(itemID)
	now := time.Now()

	if err == nil && highestBid != nil {
		itm.Status = domain.SOLD
		itm.SoldPrice = &highestBid.Amount
		itm.SoldAt = &now

		// Create purchase for the winner
		purch := &domain.Purchase{
			ID:          uuid.New(),
			UserID:      highestBid.BidderID,
			ItemID:      itm.ID,
			PricePaid:   highestBid.Amount.InexactFloat64(),
			PurchasedAt: now,
		}
		_ = s.purchaseRepo.AddPurchase(purch)
	} else {
		itm.Status = domain.UNAVAILABLE
	}

	if err := s.itemRepo.UpdateItem(itm); err != nil {
		return nil, fmt.Errorf("update item during close: %w", err)
	}

	return itm, nil
}

