package entity

// User struct
type User struct {
	Username string
	Password string
	Email    string
	Phone    string
}

// UserSlice ..
type UserSlice []*User

func (s UserSlice) Len() int {
	return len(s)
}

func (s UserSlice) Swap(i, j int) {
	t := s[i]
	s[i] = s[j]
	s[j] = t
}

func (s UserSlice) Less(i, j int) bool {
	return s[i].Username < s[j].Username
}
