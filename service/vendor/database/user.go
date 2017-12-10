package database

import (
	"database/sql"
	"entity"
	"log"
)

// GetAllUsers ..
func GetAllUsers() ([]*entity.User, error) {
	result := make([]*entity.User, 0)
	var err error
	<-pQuery(func(r *sql.Rows) {
		var u entity.User
		r.Scan(&u.Username, &u.Password, &u.Email)
		result = append(result, &u)
	}, func(e error) {
		err = e
	}, "SELECT * FROM User")
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreUser ..
func StoreUser(user *entity.User) error {
	log.Printf(
		"Creating User (%s, %s, %s)", user.Username, user.Password, user.Email)
	result, err := pExec(
		"INSERT INTO User (username, password, email) VALUES (?, ?, ?)",
		user.Username, user.Password, user.Email)
	if err != nil {
		log.Printf("Error: %s", err.Error())
	} else {
		r, err := result.RowsAffected()
		if err != nil {
			log.Printf("Error: %s", err.Error())
		} else {
			log.Printf("Success: %d rows affected", r)
		}
	}
	return err
}

// GetUser by username
func GetUser(username string) (*entity.User, error) {
	var user *entity.User
	var err error
	<-pQuery(func(r *sql.Rows) {
		var u entity.User
		r.Scan(&u.Username, &u.Password, &u.Email)
		user = &u
	}, func(e error) {
		err = e
	}, "SELECT * FROM User WHERE username = ?", username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// RemoveUser ..
func RemoveUser(username string) error {
	_, err := pExec("DELETE FROM User WHERE username = ?", username)
	return err
}
