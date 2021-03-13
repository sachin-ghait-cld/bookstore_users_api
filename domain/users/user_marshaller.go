package users

import "encoding/json"

// PublicUser Public User
type PublicUser struct {
	ID          int64  `json:"id"`
	DateCreated string `json:"dateCreated"`
	Status      string `json:"status"`
}

// PrivateUser Internal User
type PrivateUser struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	DateCreated string `json:"dateCreated"`
	Status      string `json:"status"`
}

// Marshall user
func (user *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			ID:          user.ID,
			Status:      user.Status,
			DateCreated: user.DateCreated,
		}
	}
	userJSON, _ := json.Marshal(user)
	var privateUser PrivateUser
	json.Unmarshal(userJSON, &privateUser)
	return privateUser
}

// Marshall users
func (users Users) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshall(isPublic)
	}
	return result
}
