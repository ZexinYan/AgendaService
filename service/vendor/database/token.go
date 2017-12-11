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

// HasToken ..
func HasToken(token string) (bool, error) {
	var result bool
	var err error
	<-pQuery(func(r *sql.Rows) {
		result = true
	}, func(e error) {
		err = e
	}, "SELECT token FROM Login WHERE token = ?", token)
	if err != nil {
		return false, err
	}
	return result, err
}

// GetUsername ..
func GetUsername(token string) (string, error) {
	var username string
	var err error
	<-pQuery(func(r *sql.Rows) {
		r.Scan(&username)
	}, func(e error) {
		err = e
	}, "SELECT username FROM Login WHERE token = ?", token)
	if err != nil {
		return "", err
	}
	return username, nil
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
