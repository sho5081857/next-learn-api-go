package usecase

import (
	"context"
	"errors"

	"next-learn-go/model"
	"next-learn-go/repository"
	"next-learn-go/validator"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(user model.User) (model.UserResponse, error)
	Login(user model.User) (model.LoginResponse, error)
	GetUserById(userId uint) (model.UserResponse, error)
	GetUserByEmail(email string) (model.UserResponse, error)
	RefreshToken(refreshTokenString string) (model.TokenResponse, error)
}

type userUsecase struct {
	ur repository.IUserRepository
	uv validator.IUserValidator
}

func NewUserUsecase(ur repository.IUserRepository, uv validator.IUserValidator) IUserUsecase {
	return &userUsecase{ur, uv}
}

func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
	if err := uu.uv.UserValidate(user); err != nil {
		return model.UserResponse{}, err
	}

	if err := uu.ur.GetUserByEmail(context.Background(), &model.User{}, user.Email); err == nil {
		return model.UserResponse{}, errors.New("email already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.UserResponse{}, err
	}

	newUser := model.User{Name: user.Name, Email: user.Email, Password: string(hash)}
	if err := uu.ur.CreateUser(context.Background(), &newUser); err != nil {
		return model.UserResponse{}, err
	}

	resUser := model.UserResponse{
		ID:    newUser.ID,
		Name:  newUser.Name,
		Email: newUser.Email,
	}

	return resUser, nil
}

func (uu *userUsecase) Login(user model.User) (model.LoginResponse, error) {
	if err := uu.uv.UserValidate(user); err != nil {
		return model.LoginResponse{}, err
	}
	storedUser := model.User{}
	ctx := context.Background()
	if err := uu.ur.GetUserByEmail(ctx, &storedUser, user.Email); err != nil {
		return model.LoginResponse{}, err
	}
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return model.LoginResponse{}, err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return model.LoginResponse{}, err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	refreshTokenString, err := refreshToken.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return model.LoginResponse{}, err
	}

	resLogin := model.LoginResponse{
		ID:           storedUser.ID,
		Email:        storedUser.Email,
		AccessToken:  tokenString,
		RefreshToken: refreshTokenString,
	}
	return resLogin, nil
}

func (uu *userUsecase) GetUserById(userId uint) (model.UserResponse, error) {
	user := model.User{}
	ctx := context.Background()
	if err := uu.ur.GetUserById(ctx, &user, userId); err != nil {
		return model.UserResponse{}, err
	}
	resUser := model.UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	return resUser, nil
}

func (uu *userUsecase) GetUserByEmail(email string) (model.UserResponse, error) {
	user := model.User{}
	ctx := context.Background()
	if err := uu.ur.GetUserByEmail(ctx, &user, email); err != nil {
		return model.UserResponse{}, err
	}
	resUser := model.UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	return resUser, nil
}

func (uu *userUsecase) RefreshToken(refreshTokenString string) (model.TokenResponse, error) {

	token, err := jwt.Parse(refreshTokenString, func(token *jwt.Token) (any, error) {
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		return model.TokenResponse{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return model.TokenResponse{}, errors.New("invalid refresh token")
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": claims["user_id"],
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})

	newTokenString, err := newToken.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return model.TokenResponse{}, err
	}

	return model.TokenResponse{
		AccessToken:  newTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}
