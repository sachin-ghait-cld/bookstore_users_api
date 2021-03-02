package services

import (
	"github.com/sachin-ghait-cld/bookstore_users_api/domain/users"
	"github.com/sachin-ghait-cld/bookstore_users_api/utils/date_utils"
	"github.com/sachin-ghait-cld/bookstore_users_api/utils/errors"
)

// CreateUser creates user
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowDBFormat()
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUser get user details
func GetUser(userID int64) (*users.User, *errors.RestErr) {
	user := &users.User{ID: userID}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUser Updates a User
func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.ID)
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
func DeleteUser(userID int64) *errors.RestErr {
	user := users.User{ID: userID}
	return user.Delete()
}

// Search find a user by status
func Search(status string) ([]users.User, *errors.RestErr) {
	dao := &users.User{}
	return dao.Search(status)
}
