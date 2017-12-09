package cmd

import(
  "fmt"
)

// Register a user
func Register(user, pass, mail, phone string) int {
  fmt.Println("register")
	return 0
}

// Login Command
func Login(user, pass string) int {
  fmt.Println("Login")
	return 0
}

// Logout Command
func Logout() int {
  fmt.Println("Logout")
	return 0
}

// ShowUsers print all users when logged in
func ShowUsers() int {
  fmt.Println("ShowUsers")
	return 0
}

// DeleteUser delete current login user, and removed from its meeting
func DeleteUser() int {
  fmt.Println("delete User")
	return 0
}
