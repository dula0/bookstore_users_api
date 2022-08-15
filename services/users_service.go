package services

import (
	"fmt"

	"github.com/dula0/bookstore_users_api/domain/users"
	"github.com/dula0/bookstore_users_api/utils/crypto_utils"
	"github.com/dula0/bookstore_users_api/utils/date_utils"
	"github.com/dula0/bookstore_users_api/utils/errors"
)

var (
	UserService userServiceInterface = &userService{}
)

type userService struct{}

type userServiceInterface interface {
	GetUser(int64) (*users.User, *errors.RestErr)
	CreateUser(users.User) (*users.User, *errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	SearchUser(string) (users.Users, *errors.RestErr)
	LoginUser(users.LoginRequest) (*users.User, *errors.RestErr)
}

func (s *userService) GetUser(userID int64) (*users.User, *errors.RestErr) {
	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *userService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowDBFormat()
	user.Password = crypto_utils.HashPassword(user.Password)

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *userService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := s.GetUser(user.ID)
	if err != nil {
		return nil, err
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if !isPartial { // Executes during POST request
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	} else { // Executes during PATCH request
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	}

	err = current.Update()
	if err != nil {
		return nil, err
	}
	return current, nil
}

func (s *userService) DeleteUser(userID int64) *errors.RestErr {
	user := &users.User{ID: userID}
	return user.Delete()
}

func (s *userService) SearchUser(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}

func (s *userService) LoginUser(req users.LoginRequest) (*users.User, *errors.RestErr) {
	dao := &users.User{
		Email:    req.Email,
		Password: crypto_utils.HashPassword(req.Password),
	}
	if err := dao.FindByEmailAndPassword(); err != nil {
		fmt.Println(crypto_utils.HashPassword(req.Password))
		return nil, err
	}
	return dao, nil
}
