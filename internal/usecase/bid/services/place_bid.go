package services

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"github.com/ParthPoudyal/E-commerce-project/internal/domain"
	"github.com/ParthPoudyal/E-commerce-project/internal/usecase/bid"
	"github.com/ParthPoudyal/E-commerce-project/internal/usecase/item"
)

type PlaceBidService struct {
	bidRepo  bid.BidRepository
	itemRepo item.ItemRepo
}

func NewPlaceBidService(bR bid.BidRepository, iR item.ItemRepo) *PlaceBidService {
	return &PlaceBidService{
		bidRepo:  bR,
		itemRepo: iR,
	}
}

type PlaceBidResult struct {
	Bid                 *domain.Bid
	PreviousHighestUser *uuid.UUID
	IsNewHighest        bool
}

func (s *PlaceBidService) Execute(itemID, bidderID uuid.UUID, amount decimal.Decimal) (*PlaceBidResult, error) {
	itm, err := s.itemRepo.FindItemByID(itemID)
	if err != nil {
		return nil, domain.ErrItemNotFound
	}

	if !itm.IsAuction {
		return nil, fmt.Errorf("item is not listed for auction")
	}

	if itm.Status != domain.BIDDING && itm.Status != domain.AVAILABLE {
		return nil, domain.ErrAuctionNotActive
	}

	now := time.Now()
	if itm.ValidTo != nil && now.After(*itm.ValidTo) {
		return nil, domain.ErrAuctionClosed
	}

	if itm.PostedBy != nil && itm.PostedBy.ID == bidderID {
		return nil, domain.ErrCannotBidOnOwnItem
	}

	if itm.StartingBid != nil && amount.LessThan(*itm.StartingBid) {
		return nil, domain.ErrBidBelowStartingPrice
	}

	// Check current highest bid
	highestBid, err := s.bidRepo.GetHighestBidForItem(itemID)
	var prevHighestUser *uuid.UUID

	if err == nil && highestBid != nil {
		prevHighestUser = &highestBid.BidderID

		// Minimum increment calculation: 1% of highest bid or 1.0, whichever is greater
		minInc := highestBid.Amount.Mul(decimal.NewFromFloat(0.01))
		if minInc.LessThan(decimal.NewFromInt(1)) {
			minInc = decimal.NewFromInt(1)
		}

		minRequired := highestBid.Amount.Add(minInc)
		if amount.LessThan(minRequired) {
			return nil, fmt.Errorf("%w: minimum bid must be at least %s", domain.ErrBidTooLow, minRequired.StringFixed(2))
		}
	}

	newBid := &domain.Bid{
		ID:        uuid.New(),
		ItemID:    itemID,
		BidderID:  bidderID,
		Amount:    amount,
		CreatedAt: now,
	}

	if err := s.bidRepo.Create(newBid); err != nil {
		return nil, fmt.Errorf("create bid: %w", err)
	}

	// Update item current bid and price
	itm.CurrentBidID = &newBid.ID
	itm.Price = amount
	itm.Status = domain.BIDDING
	_ = s.itemRepo.UpdateItem(itm)

	return &PlaceBidResult{
		Bid:                 newBid,
		PreviousHighestUser: prevHighestUser,
		IsNewHighest:        true,
	}, nil
}

