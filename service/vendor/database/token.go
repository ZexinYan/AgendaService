package database

import (
	"fmt"
)

// GetToken ..
func GetToken(username string) (string, error) {
	rows, err := pQuery(
		theDB, "SELECT token FROM Login WHERE username = ?", username)
	if err != nil {
		return "", err
	}
	for row := range rows {
		var token string
		row.Scan(&token)
		return token, nil
	}
	return "", fmt.Errorf("Token of %s not found", username)
}

// PutToken ..
func PutToken(username string, token string) error {
	_, err := pExec(
		theDB,
		"INSERT INTO Login (token, username) VALUES (?, ?)",
		token, username)
	return err
}

// DeleteToken ..
func DeleteToken(token string) error {
	_, err := pExec(theDB, "DELETE FROM Login WHERE token = ?", token)
	return err
}

// DeleteTokenByUsername ..
func DeleteTokenByUsername(username string) error {
	_, err := pExec(theDB, "DELETE FROM Login WHERE username = ?", username)
	return err
}
