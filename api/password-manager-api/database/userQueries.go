package database

import (
	"context"
	"log"

	"password-manager-service/types"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

const (
	queryAllUsersData  string = "SELECT USER_ID, EMAIL_ID, USERNAME, PASSWORD FROM USER"
	queryUserByEmail   string = "SELECT * FROM USER WHERE EMAIL_ID = ?"
	insertUser         string = "INSERT INTO USER (USER_ID, EMAIL_ID, USERNAME, PASSWORD) VALUES (?, ?, ?, ?)"
	updateUserPassword string = "UPDATE USER SET PASSWORD=? WHERE USER_ID=?"
)

func (conn *DatabaseConnection) GetAllUsers() ([]types.UserData, error) {
	var users []types.UserData

	results, err := conn.db.Query(queryAllUsersData)
	if err != nil {
		return users, err
	}

	for results.Next() {
		var user types.UserData

		err := results.Scan(&user.UserId, &user.Email, &user.Username, &user.Password)
		if err != nil {
			log.Printf("Could not process the row data: %s\n", err)
		}

		users = append(users, user)
	}

	return users, nil
}

func (conn *DatabaseConnection) CheckEmailExists(email string) (bool, error) {

	result, err := conn.db.Query(queryUserByEmail, email)
	if err != nil {
		return false, err
	}

	for result.Next() {
		return true, nil
	}

	return false, nil
}

func (conn *DatabaseConnection) AddUser(user types.UserData) error {

	user.UserId = uuid.New().String()

	// query := fmt.Sprintf("SELECT * FROM USER WHERE USERNAME = '?'", uname)
	_, err := conn.db.Query(insertUser, user.UserId, user.Email, user.Username, user.Password)

	return err
}

func (conn *DatabaseConnection) GetUserByEmail(email string) (types.UserData, error) {

	var user types.UserData

	results, err := conn.db.Query(queryUserByEmail, email)
	if err != nil {
		log.Println(err.Error())
		return user, err
	}

	for results.Next() {
		err = results.Scan(&user.UserId, &user.Email, &user.Username, &user.Password)
		if err != nil {
			log.Printf("Could not process the row data: %s\n", err)
			break
		}
	}

	return user, err
}

func (conn *DatabaseConnection) UpdateUserPassword(userId string, password string) error {

	_, err := conn.db.ExecContext(context.Background(), updateUserPassword, password, userId)

	return err
}
