package service

import (
	"errors"
	"fmt"
	"go-ecommerce-app/internal/domain"
	"go-ecommerce-app/internal/dto"
	"go-ecommerce-app/internal/repository"
	"log"
)

type UserService struct {
	Repo repository.UserRepository
}

/*email username.. ????????????? ---> input any || input interface{}*/
func (s UserService) Signup(input dto.UserSignup) (string, error) {

	log.Println(input)

	user, err := s.Repo.CreateUser(domain.User{
		Email:    input.Email,
		Password: input.Password,
		Phone:    input.Phone,
	})

	// generate token
	log.Println(user)

	userInfo := fmt.Sprintf("%v, %v, %v", user.ID, user.Email, user.UserType)

	// call db to create user
	return userInfo, err
}

func (s UserService) UpdateProfile(id uint, input dto.UserUpdateProfile) (string, error) {

	_, err := s.Repo.FindUserById(id)
	if err != nil {
		return "", errors.New("user not found")
	}

	user, err := s.Repo.UpdateUser(id, domain.User{
		Email:    input.Email,
		Password: input.Password,
	})

	userInfo := fmt.Sprintf("%v, %v", user.Email, user.Password)

	return userInfo, err
}

func (s UserService) findUserByEmail(email string) (*domain.User, error) {

	user, err := s.Repo.FindUser(email)

	return &user, err

}

func (s UserService) Login(email, password string) (string, error) {

	user, err := s.findUserByEmail(email)
	if err != nil {
		return "", errors.New("user doest not exist with the provided email id")
	}

	// compare password and generate token
	return user.Email, nil

}

func (s UserService) GetVerificationCode(e domain.User) (int, error) {
	return 0, nil
}

func (s UserService) VerifyCode(id uint, code int) error {
	return nil
}

func (s UserService) CreateProfile(id uint, input any) error {
	return nil
}

func (s UserService) GetProfile(id uint) (*domain.User, error) {

	user, err := s.Repo.FindUserById(id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (s UserService) BecomeSeller(id uint, input any) (string, error) {
	return "", nil
}

func (s UserService) FindCart(id uint) ([]interface{}, error) {
	return nil, nil
}

func (s UserService) CreateCart(input any, u domain.User) ([]interface{}, error) {
	return nil, nil
}

func (s UserService) CreateOrder(u domain.User) (int, error) {
	return 0, nil
}

func (s UserService) GetOrders(u domain.User) ([]interface{}, error) {
	return nil, nil
}

func (s UserService) GetOrderById(id uint, uId uint) (interface{}, error) {
	return nil, nil
}
