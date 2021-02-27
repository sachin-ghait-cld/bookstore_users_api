package users

import (
	"fmt"

	"github.com/sachin-ghait-cld/bookstore_users_api/utils/errors"
)

var userDB = make(map[int64]*User)

// Save saves a user
func (user *User) Save() *errors.RestErr {
	current := userDB[user.ID]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError("user email already registered")
		}
		return errors.NewBadRequestError("user already exist")
	}
	userDB[user.ID] = user
	return nil
}

// Get get a user
func (user *User) Get() *errors.RestErr {
	res := userDB[user.ID]
	if res == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d does not exist", user.ID))
	}
	user.ID = res.ID
	user.Email = res.Email
	user.FirstName = res.FirstName
	user.LastName = res.LastName
	user.DateCreated = res.DateCreated
	return nil
}
