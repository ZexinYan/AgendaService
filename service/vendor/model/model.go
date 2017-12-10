package model

import (
	"database"
	"entity"

	"github.com/dchest/uniuri"
)

type needLogin struct {
	username string
}

func authenticate(username string, token string) (*needLogin, ErrorCode) {
	tok, err := database.GetToken(username)
	if err != nil {
		return nil, InvalidToken
	}
	if tok == token {
		return &needLogin{username}, OK
	}
	return nil, InvalidToken
}

// CreateUser ..
func CreateUser(u *entity.User) ErrorCode {
	uu, err := database.GetUser(u.Username)
	if err != nil {
		return DatabaseFail
	}
	if uu != nil {
		return DuplicateUser
	}
	err = database.StoreUser(u)
	if err != nil {
		return DatabaseFail
	}
	return OK
}

func getAllUsers(needLogin) ([]*entity.User, ErrorCode) {
	us, e := database.GetAllUsers()
	if e != nil {
		return nil, DatabaseFail
	}
	return us, OK
}

func removeUser(l needLogin) ErrorCode {
	err := database.RemoveUser(l.username)
	if err != nil {
		return DatabaseFail
	}
	return OK
}

// GetAllUsers ..
func GetAllUsers(username, token string) ([]*entity.User, ErrorCode) {
	lp, ec := authenticate(username, token)
	if ec != OK {
		return nil, ec
	}
	return getAllUsers(*lp)
}

// RemoveUser ..
func RemoveUser(username, token string) ErrorCode {
	lp, ec := authenticate(username, token)
	if ec != OK {
		return ec
	}
	return removeUser(*lp)
}

// Login ..
func Login(username, password string) (string, ErrorCode) {
	u, err := database.GetUser(username)
	if err != nil {
		return "", DatabaseFail
	}
	if u == nil || u.Password != password {
		return "", AuthenticationFail
	}
	t, err := database.GetToken(username)
	if err == nil {
		return "", DuplicateLogin
	}
	t = uniuri.New()
	err = database.PutToken(username, t)
	if err != nil {
		return "", DatabaseFail
	}
	return t, OK
}

// Logout ..
func Logout(username, token string) ErrorCode {
	lp, ec := authenticate(username, token)
	if ec != OK {
		return ec
	}
	return logout(*lp)
}

func logout(l needLogin) ErrorCode {
	err := database.DeleteTokenByUsername(l.username)
	if err != nil {
		return DatabaseFail
	}
	return OK
}
