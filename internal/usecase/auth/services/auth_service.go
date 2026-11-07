package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/ParthPoudyal/E-commerce-project/infrastructure/auth"
	"github.com/ParthPoudyal/E-commerce-project/internal/domain"
	"github.com/ParthPoudyal/E-commerce-project/internal/usecase/user"
)

type AuthService struct {
	userRepo user.UserRepo
}

func NewAuthService(uR user.UserRepo) *AuthService {
	return &AuthService{
		userRepo: uR,
	}
}

type RegisterInput struct {
	Username string
	Email    string
	Password string
	Address  string
	PhoneNo  string
}

type AuthOutput struct {
	Token string       `json:"token"`
	User  *domain.User `json:"user"`
}

func (s *AuthService) Register(input RegisterInput) (*AuthOutput, error) {
	if input.Username == "" || input.Email == "" || input.Password == "" {
		return nil, domain.ErrInvalidUserData
	}
	if len(input.Password) < 6 {
		return nil, domain.ErrPasswordTooWeak
	}

	// Check if email already exists
	existingUser, err := s.userRepo.FindUserByEmail(input.Email)
	if err == nil && existingUser != nil {
		return nil, domain.ErrEmailAlreadyExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("hash password: %w", err)
	}

	newUser := &domain.User{
		ID:       uuid.New(),
		Username: input.Username,
		Email:    input.Email,
		Address:  input.Address,
		PhoneNo:  input.PhoneNo,
		Base: domain.Base{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	if err := s.userRepo.AddUser(newUser); err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	pwdStr := string(hashedPassword)
	cred := &domain.AuthCredential{
		UserID:       newUser.ID,
		PasswordHash: &pwdStr,
		Provider:     "local",
	}

	if err := s.userRepo.SaveAuthCredential(cred); err != nil {
		return nil, fmt.Errorf("save credentials: %w", err)
	}

	token, err := auth.GenerateToken(newUser.ID, newUser.Email, newUser.Username)
	if err != nil {
		return nil, err
	}

	return &AuthOutput{
		Token: token,
		User:  newUser,
	}, nil
}

func (s *AuthService) Login(email, password string) (*AuthOutput, error) {
	if email == "" || password == "" {
		return nil, domain.ErrInvalidCredentials
	}

	u, err := s.userRepo.FindUserByEmail(email)
	if err != nil {
		return nil, domain.ErrInvalidCredentials
	}

	cred, err := s.userRepo.FindAuthCredentialByUserID(u.ID)
	if err != nil || cred == nil || cred.PasswordHash == nil {
		return nil, domain.ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(*cred.PasswordHash), []byte(password)); err != nil {
		return nil, domain.ErrInvalidCredentials
	}

	token, err := auth.GenerateToken(u.ID, u.Email, u.Username)
	if err != nil {
		return nil, err
	}

	return &AuthOutput{
		Token: token,
		User:  u,
	}, nil
}

func (s *AuthService) GetUserByID(userID uuid.UUID) (*domain.User, error) {
	u, err := s.userRepo.FindUserByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return u, nil
}

