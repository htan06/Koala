package auth

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"koala.com/dto/auth/request"
	"koala.com/dto/auth/response"
)

type AuthService interface {
	Login(ctx context.Context, login request.LoginDto) (response.TokenResponse, error)
	RegisterRider(ctx context.Context, registerRider request.RegisterRider) error
	ChangePassword(ctx context.Context, username string, ChangePassword request.ChangePasswordDto) error
}

type AuthServiceImpl struct {
	userRepository *UserRepository
	jwtService     *JwtService
}

func NewAuthService(userRepository *UserRepository, jwtService *JwtService) *AuthServiceImpl {
	return &AuthServiceImpl{userRepository, jwtService}
}

func (auth *AuthServiceImpl) Login(ctx context.Context, login request.LoginDto) (response.TokenResponse, error) {
	user, errRepository := auth.userRepository.FindByUsername(ctx, login.Username)

	if errRepository != nil {
		return response.TokenResponse{}, errRepository
	}

	errComparePassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if errComparePassword != nil {
		return response.TokenResponse{}, fmt.Errorf("Compare password err: %w", errComparePassword)
	} else {
		accessToken := auth.jwtService.generateAccessToken(user.Username)
		refreshToken := auth.jwtService.generateRefreshToken(user.Username)
		
		return response.TokenResponse{
			AccessToken: accessToken, 
			RefreshToken: refreshToken,
		}, nil
	}
}

func (auth *AuthServiceImpl) RegisterRider(ctx context.Context, registerRider request.RegisterRider) error {
	hashPassword, errHashPassword := bcrypt.GenerateFromPassword([]byte(registerRider.Password), 10)

	if errHashPassword != nil {
		return fmt.Errorf("Hash password err: %w", errHashPassword)
	}

	newRider := User{
		Username:    registerRider.Username,
		Password:    string(hashPassword),
		PhoneNumber: registerRider.PhoneNumber,
		Email:       registerRider.Email,
	}

	errSave := auth.userRepository.Save(ctx, newRider)

	if errSave != nil {
		return fmt.Errorf("Save user: %w")
	}
	return nil
}

func (auth *AuthServiceImpl) ChangePassword(ctx context.Context, username string, changePassword request.ChangePasswordDto) error {
	user, errRepository := auth.userRepository.FindByUsername(ctx, username)

	if errRepository != nil {
		return errRepository
	}

	errComparePw := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(changePassword.CurrentPassword))

	if errComparePw != nil {
		return fmt.Errorf("Compare password err: %w", errComparePw)
	}

	newHashPassword, errHashPassword := bcrypt.GenerateFromPassword([]byte(changePassword.NewPassword), 10)

	if errHashPassword != nil {
		return fmt.Errorf("Hash password err: %w", errHashPassword)
	}

	user.Password = string(newHashPassword)
	return auth.userRepository.updatePassword(ctx, user)
}
