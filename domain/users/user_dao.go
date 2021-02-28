package users

import (
	"github.com/sachin-ghait-cld/bookstore_users_api/datasources/mysql/users_db"
	"github.com/sachin-ghait-cld/bookstore_users_api/utils/date_utils"
	"github.com/sachin-ghait-cld/bookstore_users_api/utils/errors"
	"github.com/sachin-ghait-cld/bookstore_users_api/utils/mysql_utils"
)

const (
	queryInsertUser = "INSERT INTO users(first_name,last_name, email,date_created) VALUES(?,?,?,?);"
	queryGetUser    = "SELECT id,first_name,last_name, email,date_created from users where id = ? ;"
	queryUpdateUser = "UPDATE users set first_name = ?,last_name = ?, email = ? where id = ? ;"
	queryDeleteUser = "DELETE FROM users where id = ? ;"
)

var userDB = make(map[int64]*User)

// Save saves a user
func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	// result,err:= users_db.Client.Exec(queryInsertUser,user.FirstName, user.LastName, user.Email, user.DateCreated)

	defer stmt.Close()
	user.DateCreated = date_utils.GetNowString()
	insertRes, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}

	userID, err := insertRes.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	user.ID = userID
	return nil
}

// Get get a user
func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		return mysql_utils.ParseError(getErr)
	}

	return nil
}

// Update a user
func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	_, updateErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if updateErr != nil {
		return mysql_utils.ParseError(updateErr)
	}

	return nil
}

// Delete a user
func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	_, deleteErr := stmt.Exec(user.ID)
	if deleteErr != nil {
		return mysql_utils.ParseError(deleteErr)
	}

	return nil
}
