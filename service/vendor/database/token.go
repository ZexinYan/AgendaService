package database

import (
	"database/sql"
	"fmt"
)

// GetToken ..
func GetToken(username string) (string, error) {
	var token string
	var err error
	<-pQuery(func(r *sql.Rows) {
		r.Scan(&token)
	}, func(e error) {
		err = e
	}, "SELECT token FROM Login WHERE username = ?", username)
	if err != nil {
		return "", err
	}
	if token == "" {
		return "", fmt.Errorf("Token of %s not found", username)
	}
	return token, err
}

// PutToken ..
func PutToken(username string, token string) error {
	_, err := pExec(
		"INSERT INTO Login (token, username) VALUES (?, ?)",
		token, username)
	return err
}

// DeleteToken ..
func DeleteToken(token string) error {
	_, err := pExec("DELETE FROM Login WHERE token = ?", token)
	return err
}

// DeleteTokenByUsername ..
func DeleteTokenByUsername(username string) error {
	_, err := pExec("DELETE FROM Login WHERE username = ?", username)
	return err
}
