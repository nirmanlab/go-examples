package user

import "database/sql"

// Info of user
type Info struct{
ID int64 `db:"id" json:"id"`
Name	string `db:"name" json:"name"`
Role_id	int `db:"role_id" json:"role_id"`
Email	string `db:"email" json:"email"`
Password string `db:"password" json:"password"`
RememberToken sql.NullString `db:"remember_token" json:"remember_token"`
CreatedAt	sql.NullTime `db:"created_at" json:"created_at"`
UpdatedAt	sql.NullTime `db:"updated_at" json:"updated_at"`
}
//192.168.29.65