package services

import (
	"TaskFlowAPI/models"
	"TaskFlowAPI/repository"
	"TaskFlowAPI/utils"
	"context"
	"errors"
)

type AuthService struct {
	userRepo *repository.UserRepository
	jwtUtil  *utils.JWTUtil
}

func NewAuthService(userRepo *repository.UserRepository, jwtUtil *utils.JWTUtil) *AuthService {
	return &AuthService{userRepo: userRepo, jwtUtil: jwtUtil}
}

func (s *AuthService) Register(ctx context.Context, email, username, password string) (*models.User, error) {
	hashed, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}
	user := &models.User{
		Email:        email,
		Username:     username,
		PasswordHash: hashed,
	}
	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}
	if err := utils.CheckPassword(password, user.PasswordHash); err != nil {
		return "", errors.New("invalid credentials")
	}
	token, err := s.jwtUtil.GenerateToken(user.ID, 0)
	if err != nil {
		return "", err
	}
	return token, nil
}
