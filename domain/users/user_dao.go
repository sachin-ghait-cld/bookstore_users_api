package users

import (
	"fmt"
	"strings"

	"github.com/sachin-ghait-cld/bookstore_users_api/datasources/mysql/users_db"
	"github.com/sachin-ghait-cld/bookstore_users_api/utils/mysql_utils"
	"github.com/sachin-ghait-cld/bookstore_utils-go/logger"
	"github.com/sachin-ghait-cld/bookstore_utils-go/rest_errors"
)

const (
	queryInsertUser             = "INSERT INTO users(first_name,last_name, email,date_created,status, password) VALUES(?,?,?,?,?,?);"
	queryGetUser                = "SELECT id,first_name,last_name, email,date_created,status from users where id = ? ;"
	queryUpdateUser             = "UPDATE users set first_name = ?,last_name = ?, email = ? where id = ? ;"
	queryDeleteUser             = "DELETE FROM users where id = ? ;"
	queryFindByStutus           = "SELECT id,first_name,last_name, email,date_created,status from users where status = ?;"
	queryFindByEmailAndPassword = "SELECT id,first_name,last_name, email,date_created,status from users where email = ? and password = ? and status=?;"
)

var userDB = make(map[int64]*User)

// Save saves a user
func (user *User) Save() *rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return rest_errors.NewInternalServerError("Error preparing get user query", err)
	}
	// result,err:= users_db.Client.Exec(queryInsertUser,user.FirstName, user.LastName, user.Email, user.DateCreated)

	defer stmt.Close()
	insertRes, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
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
func (user *User) Get() *rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("Error in Prepare query", err)
		return rest_errors.NewInternalServerError("database error", err)
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		return mysql_utils.ParseError(getErr)
	}

	return nil
}

// Update a user
func (user *User) Update() *rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return rest_errors.NewInternalServerError("error preparing update user query", err)
	}
	defer stmt.Close()
	_, updateErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if updateErr != nil {
		return mysql_utils.ParseError(updateErr)
	}

	return nil
}

// Delete a user
func (user *User) Delete() *rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return rest_errors.NewInternalServerError("Error preparing delete user query", err)
	}
	defer stmt.Close()
	_, deleteErr := stmt.Exec(user.ID)
	if deleteErr != nil {
		return mysql_utils.ParseError(deleteErr)
	}

	return nil
}

// Search Find User By Status
func (user *User) Search(status string) ([]User, *rest_errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindByStutus)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("error in preparing user search query", err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(status)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("Error in finding users from DB", err)
	}
	defer rows.Close()
	result := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		result = append(result, user)
	}
	if len(result) == 0 {
		return nil, rest_errors.NewNotFoundError(fmt.Sprintf("No users matching status %s", status))
	}
	return result, nil
}

// FindByEmailAndPassword get a user by email ans password
func (user *User) FindByEmailAndPassword() *rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		logger.Error("Error in Prepare query", err)
		return rest_errors.NewInternalServerError("database error", err)
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Email, user.Password, StatusActive)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		if strings.Contains(getErr.Error(), mysql_utils.ErrorNoRows) {
			return rest_errors.NewInternalServerError("no user found", getErr)
		}
		return mysql_utils.ParseError(getErr)
	}

	return nil
}
