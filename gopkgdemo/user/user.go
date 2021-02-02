package user

// User : Model for user
type User struct {
	Name    string `db:"name" json:"name"`
	Address string `db:"address" json:"address"`
}

// CreateUser : Creating user
func CreateUser(name, address string) *User {
	return &User{
		Name:    name,
		Address: address,
	}
}
