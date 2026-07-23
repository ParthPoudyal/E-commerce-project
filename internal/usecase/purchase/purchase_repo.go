package purchase

import (
	"time"

	"github.com/ParthPoudyal/E-commerce-project/internal/domain"
	"github.com/google/uuid"
)

type PurchaseRepo interface { 
	// ADD Purchase
	AddPurchase(purchase *domain.Purchase) error

	// List All Purchases by a User
	ListPurchasesByUser(userID uuid.UUID) ([]domain.Purchase, error)

	// List Purchase by Date 
	ListPurchasesbyDate (user uuid.UUID , date time.Time) ([]domain.Purchase , error) 

	// List All purchase between time A and time B  by a User
	ListPurchasesDuring (user uuid.UUID, timeA , timeB time.Time) ([]domain.Purchase , error)
}