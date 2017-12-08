package entity

// User struct
type User struct {
	Username string
	Password string
	Phone    string
	Email    string
}

// UserDAO ..
type UserDAO interface {
	GetAllUsers() []*User
	StoreUsers(user *User)
	GetUser(username string) *User
}
