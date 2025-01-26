package usecase

import (
	"context"
	"errors"

	"next-learn-go/entity"
	"next-learn-go/repository"
	"next-learn-go/validator"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	SignUp(user entity.User) (entity.UserResponse, error)
	Login(user entity.User) (entity.LoginResponse, error)
	GetUserById(userId uint) (entity.UserResponse, error)
	GetUserByEmail(email string) (entity.UserResponse, error)
}

type userUseCase struct {
	ur repository.UserRepository
	uv validator.UserValidator
}

func NewUserUseCase(ur repository.UserRepository, uv validator.UserValidator) UserUseCase {
	return &userUseCase{ur, uv}
}

func (uu *userUseCase) SignUp(user entity.User) (entity.UserResponse, error) {
	if err := uu.uv.UserValidate(user); err != nil {
		return entity.UserResponse{}, err
	}

	if err := uu.ur.GetUserByEmail(context.Background(), &entity.User{}, user.Email); err == nil {
		return entity.UserResponse{}, errors.New("email already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return entity.UserResponse{}, err
	}

	newUser := entity.User{Name: user.Name, Email: user.Email, Password: string(hash)}
	if err := uu.ur.CreateUser(context.Background(), &newUser); err != nil {
		return entity.UserResponse{}, err
	}

	resUser := entity.UserResponse{
		ID:    newUser.ID,
		Name:  newUser.Name,
		Email: newUser.Email,
	}

	return resUser, nil
}

func (uu *userUseCase) Login(user entity.User) (entity.LoginResponse, error) {
	if err := uu.uv.UserValidate(user); err != nil {
		return entity.LoginResponse{}, err
	}
	storedUser := entity.User{}
	ctx := context.Background()
	if err := uu.ur.GetUserByEmail(ctx, &storedUser, user.Email); err != nil {
		return entity.LoginResponse{}, err
	}
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return entity.LoginResponse{}, err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return entity.LoginResponse{}, err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	refreshTokenString, err := refreshToken.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return entity.LoginResponse{}, err
	}

	resLogin := entity.LoginResponse{
		ID:           storedUser.ID,
		Email:        storedUser.Email,
		AccessToken:  tokenString,
		RefreshToken: refreshTokenString,
	}
	return resLogin, nil
}

func (uu *userUseCase) GetUserById(userId uint) (entity.UserResponse, error) {
	user := entity.User{}
	ctx := context.Background()
	if err := uu.ur.GetUserById(ctx, &user, userId); err != nil {
		return entity.UserResponse{}, err
	}
	resUser := entity.UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	return resUser, nil
}

func (uu *userUseCase) GetUserByEmail(email string) (entity.UserResponse, error) {
	user := entity.User{}
	ctx := context.Background()
	if err := uu.ur.GetUserByEmail(ctx, &user, email); err != nil {
		return entity.UserResponse{}, err
	}
	resUser := entity.UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	return resUser, nil
}
