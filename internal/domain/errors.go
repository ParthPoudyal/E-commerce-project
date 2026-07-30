// domain/errors.go
package domain

import "errors"

var (
	// user errors
	ErrUserNotFound      = errors.New("user not found")
	ErrEmailAlreadyExists = errors.New("email is already registered")
	ErrUsernameAlreadyExists = errors.New("username is already taken")
	ErrInvalidUserData   = errors.New("invalid user data")

	// auth errors
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrWrongLoginMethod   = errors.New("this account was registered with a different login method")
	ErrPasswordTooWeak    = errors.New("password does not meet minimum requirements")

	// token errors
	ErrInvalidToken     = errors.New("invalid or malformed token")
	ErrTokenExpired     = errors.New("token has expired")
	ErrTokenRevoked     = errors.New("token has been revoked")

	// oauth errors
	ErrInvalidGoogleToken   = errors.New("google id token could not be verified")
	ErrGoogleEmailNotVerified = errors.New("google account email is not verified")
	
	// item errors
	ErrItemNotFound      = errors.New("item not found")
	ErrItemNotAvailable  = errors.New("item is not available for purchase")
	ErrCannotBuyOwnItem  = errors.New("cannot purchase your own item")

	// cart errors
	ErrAlreadyInCart     = errors.New("item is already in cart")
	ErrItemNotInCart	 = errors.New("item is not present in cart")

	// wishlist errors
	ErrAlreadyWishlisted = errors.New("item is already in wishlist")

	// bidding errors
	ErrAuctionNotActive       = errors.New("auction is not currently active")
	ErrAuctionClosed          = errors.New("auction has already closed")
	ErrCannotBidOnOwnItem     = errors.New("cannot bid on your own item")
	ErrBidTooLow              = errors.New("bid amount must exceed current highest bid")
	ErrBidBelowStartingPrice  = errors.New("bid amount is below starting price")

	
)