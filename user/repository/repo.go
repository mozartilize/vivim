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
