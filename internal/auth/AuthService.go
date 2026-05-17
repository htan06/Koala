package auth

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"koala.com/internal/auth/dto/request"
	"koala.com/internal/auth/dto/response"
	"koala.com/internal/auth/entity"
)

type AuthService interface {
	Login(ctx context.Context, login request.LoginDto) (response.TokenResponse, error)
	Register(ctx context.Context, registerRider request.RegisterDto, roleName string) error
	ChangePassword(ctx context.Context, userId uuid.UUID, ChangePassword request.ChangePasswordDto) error
}

type AuthServiceImpl struct {
	userRepository UserRepository
	jwtService     JwtService
}

func NewAuthService(userRepository UserRepository, jwtService JwtService) *AuthServiceImpl {
	return &AuthServiceImpl{
		userRepository: userRepository, 
		jwtService: jwtService,
	}
}

func (auth *AuthServiceImpl) Login(ctx context.Context, login request.LoginDto) (response.TokenResponse, error) {
	user, errFindUser := auth.userRepository.FindByUsername(ctx, login.Username)

	if errFindUser != nil {
		return response.TokenResponse{}, fmt.Errorf("Login: %w", errFindUser)
	}

	errComparePassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))

	if errComparePassword != nil {
		return response.TokenResponse{}, fmt.Errorf("Login: %w", errComparePassword)
	} 
	
	accessToken, errAccessToken := auth.jwtService.generateAccessToken(user.Id)
	refreshToken, errRefreshToken := auth.jwtService.generateRefreshToken(user.Id)

	if errAccessToken != nil {
		return response.TokenResponse{}, fmt.Errorf("Login: %w", errAccessToken)
	}

	if errRefreshToken != nil {
		return response.TokenResponse{}, fmt.Errorf("Login: %w", errRefreshToken)
	}

	return response.TokenResponse{
		AccessToken: accessToken, 
		RefreshToken: refreshToken,
	}, nil
}

func (auth *AuthServiceImpl) Register(ctx context.Context, registerRider request.RegisterDto, roleName string) error {
	hashPassword, errHashPassword := bcrypt.GenerateFromPassword([]byte(registerRider.Password), 10)

	if errHashPassword != nil {
		return fmt.Errorf("Register rider: %w", errHashPassword)
	}

	newRider := entity.User{
		Username:    registerRider.Username,
		Password:    string(hashPassword),
		PhoneNumber: registerRider.PhoneNumber,
		Email:       registerRider.Email,
	}

	role := entity.Role{
		RoleName: roleName,
	}

	errSave := auth.userRepository.Save(ctx, newRider, role)

	if errSave != nil {
		return fmt.Errorf("Register rider: %w")
	}
	return nil
}

func (auth *AuthServiceImpl) ChangePassword(ctx context.Context, userId uuid.UUID, changePassword request.ChangePasswordDto) error {
	user, errRepository := auth.userRepository.FindById(ctx, userId)

	if errRepository != nil {
		return errRepository
	}

	errComparePw := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(changePassword.CurrentPassword))

	if errComparePw != nil {
		return fmt.Errorf("Change password: %w", errComparePw)
	}

	newHashPassword, errHashPassword := bcrypt.GenerateFromPassword([]byte(changePassword.NewPassword), 10)

	if errHashPassword != nil {
		return fmt.Errorf("Change password: %w", errHashPassword)
	}

	user.Password = string(newHashPassword)
	return auth.userRepository.UpdatePassword(ctx, user)
}
