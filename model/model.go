package model

// User ...
type User struct {
	ID        int64
	Username  string
	Password  string
	Nickname  string
	IsTwoStep bool
	Secret    string
}
