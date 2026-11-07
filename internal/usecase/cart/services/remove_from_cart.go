package services

import (
	"github.com/ParthPoudyal/E-commerce-project/internal/domain"
	"github.com/ParthPoudyal/E-commerce-project/internal/usecase/cart"
	"github.com/google/uuid"
)

type RemoveFromCartService struct {
	cartRepo cart.CartRepo
}

func (cs *RemoveFromCartService) Execute(userID, itemID uuid.UUID) error {
	itm, err := cs.cartRepo.FindItemInCart(userID, itemID)
	if err != nil {
		return domain.ErrItemNotInCart
	}
	
	return cs.cartRepo.RemoveFromCart(userID, itm.ItemID)
}