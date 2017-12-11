package model

import (
	"encoding/json"
	"os"
	"fmt"
)

type token struct {
	Token string
}

func openFileRewrite(path string) (*os.File, error) {
	return os.OpenFile(path, rewritePerm, 0644)
}

const rewritePerm = os.O_WRONLY | os.O_CREATE | os.O_TRUNC

func LoadLoginFile() *token {
	file, e := os.Open(LoginFile())
	if e != nil {
		return nil
	}
	t := new(token)
	e = json.NewDecoder(file).Decode(t)
	if e != nil {
		Logout()
		return nil
	}
	return t
}

func WriteLoginFile(t string) {
	file, e := openFileRewrite(LoginFile())
	if e != nil {
		fmt.Print(e)
	}
	json.NewEncoder(file).Encode(token{t})
}

// Logout try to delete the login file returns true if success
// returns false if there's no login file
func Logout() bool {
	return os.Remove(LoginFile()) == nil
}

