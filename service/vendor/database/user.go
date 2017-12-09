package database

import "entity"

// GetAllUsers ..
func GetAllUsers() ([]*entity.User, error) {
	rows, err := pQuery(theDB, "SELECT * FROM User")
	if err != nil {
		return nil, err
	}
	result := make([]*entity.User, 0)
	for row := range rows {
		var u entity.User
		row.Scan(&u.Username, &u.Password, &u.Phone, &u.Email)
		result = append(result, &u)
	}
	return result, nil
}

// StoreUser ..
func StoreUser(user *entity.User) error {
	_, err := pExec(
		theDB,
		"INSERT INTO User (username, password, phone, email) VALUES (?, ?, ?, ?)",
		user.Username, user.Password, user.Phone, user.Email)
	return err
}

// GetUser by username
func GetUser(username string) (*entity.User, error) {
	rows, err := pQuery(theDB, "SELECT * FROM User WHERE username = ?", username)
	if err != nil {
		return nil, err
	}
	for row := range rows {
		var u entity.User
		row.Scan(&u.Username, &u.Password, &u.Phone, &u.Email)
		return &u, nil
	}
	return nil, nil
}

// RemoveUser ..
func RemoveUser(username string) error {
	_, err := pExec(theDB, "DELETE FROM User WHERE username = ?", username)
	return err
}
