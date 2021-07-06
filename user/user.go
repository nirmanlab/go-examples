package user

import (
	"OauthPoc/database"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func GetUserInfo(username, password string) (userID string, err error) {
	

	info := Info{}
	mdb, err := database.GetSQLXMysqlInstance("adlevo", "adlevo", "192.168.29.65", "adlevo")
	if err != nil {
		return "", err
	}
	q := "SELECT id, name, role_id, email, password FROM users where id = 1"
	err = mdb.Get(&info, q)
	if err != nil {
		return "info", err
	}

	return fmt.Sprintf("%d",info.ID),err
}
