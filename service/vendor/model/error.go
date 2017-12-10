package model

// ErrorCode ..
type ErrorCode int

const (
	// OK no error
	OK ErrorCode = iota
	// InvalidToken of login state
	InvalidToken
	// DatabaseFail is a server error
	DatabaseFail
	// DuplicateUser when creating user
	DuplicateUser
	// DuplicateLogin when login multiple times
	DuplicateLogin
	// WrongLoginState when login or logout is needed
	WrongLoginState
	// AuthenticationFail when username don't match password
	AuthenticationFail
)
