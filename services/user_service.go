package services

import (
	"github.com/sachin-ghait-cld/bookstore_users_api/domain/users"
	"github.com/sachin-ghait-cld/bookstore_users_api/utils/crypto_utils"
	"github.com/sachin-ghait-cld/bookstore_users_api/utils/date_utils"
	"github.com/sachin-ghait-cld/bookstore_users_api/utils/errors"
)

var (
	// UserService to access methods on this service
	UserService userServiceInterface = &userService{}
)

type userService struct{}

type userServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestErr)
	GetUser(int64) (*users.User, *errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	SearchUser(string) (users.Users, *errors.RestErr)
	LoginUser(users.LoginRequest) (*users.User, *errors.RestErr)
}

// CreateUser creates user
func (s *userService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowDBFormat()
	user.Password = crypto_utils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUser get user details
func (s *userService) GetUser(userID int64) (*users.User, *errors.RestErr) {
	user := &users.User{ID: userID}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUser Updates a User
func (s *userService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := s.GetUser(user.ID)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email

	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

// DeleteUser Delete a User
func (s *userService) DeleteUser(userID int64) *errors.RestErr {
	user := users.User{ID: userID}
	return user.Delete()
}

// Search find a user by status
func (s *userService) SearchUser(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.Search(status)
}

// GetUser get user details
func (s *userService) LoginUser(request users.LoginRequest) (*users.User, *errors.RestErr) {
	user := &users.User{
		Email:    request.Email,
		Password: crypto_utils.GetMd5(request.Password),
	}
	if err := user.FindByEmailAndPassword(); err != nil {
		return nil, err
	}
	return user, nil
}
