package services

import (
	"github.com/ParthPoudyal/E-commerce-project/internal/domain"
	"github.com/ParthPoudyal/E-commerce-project/internal/usecase/item"
	"github.com/ParthPoudyal/E-commerce-project/internal/usecase/cart"
	"github.com/google/uuid"
)

type AddtoCartService struct {
	cartRepo cart.CartRepo
	itemRepo item.ItemRepo
}

func NewAddtoCartUseCase(cartRepo cart.CartRepo, itemRepo item.ItemRepo) *AddtoCartService {
	return &AddtoCartService{
		cartRepo: cartRepo,
		itemRepo: itemRepo,
	}
}

func (cs *AddtoCartService) Execute(userID, itemID uuid.UUID) (*domain.CartItem, error) {

	it, err := cs.itemRepo.FindItemByID(itemID)
	if err != nil {
		return nil, domain.ErrItemNotFound
	}
	if it.Status != domain.AVAILABE {
		return nil, domain.ErrItemNotAvailable
	}

	if exists, _ := cs.cartRepo.FindItemInCart(userID, it.ID); exists != nil {
		return nil, domain.ErrAlreadyInCart
	}

	newCartItem := &domain.CartItem{
		Id:     uuid.New(),
		UserID: userID,
		ItemID: it.ID,
	}

	if err := cs.cartRepo.AddItemToCart(newCartItem); err != nil {
		return nil, err
	}

	return newCartItem, nil

}
