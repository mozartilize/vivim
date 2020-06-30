package repository

import (
	"vivim/db"
	"vivim/user/models"
)

func GetUserById(id string) (models.User, error) {
	user := models.User{}
	err := db.Database.Get(&user, "select id, username, email from users_user where id=$1", id)
	return user, err
}

func GetUserByUsernameOrEmail(v string) (models.User, error) {
	user := models.User{}
	err := db.Database.Get(&user, "select id, username, email, password from users_user where username=$1 or email=$1", v)
	return user, err
}
