package services

import (
	"github.com/sachin-ghait-cld/bookstore_users_api/domain/users"
	"github.com/sachin-ghait-cld/bookstore_users_api/utils/errors"
)

// CreateUser creates user
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
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

// FindUser find a user
func FindUser() {

}
