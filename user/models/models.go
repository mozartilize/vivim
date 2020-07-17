package models

import (
	"crypto/sha256"
	"strconv"
	"strings"

	p "github.com/wuriyanto48/go-pbkdf2"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

func (u User) CheckPassword(pwd string) bool {
	passParts := strings.Split(u.Password, "$")
	iter, _ := strconv.Atoi(passParts[1])
	salt := passParts[2]
	cipherText := passParts[3]
	pass := p.NewPassword(sha256.New, len(salt), len(cipherText)-len(salt), iter)
	return pass.VerifyPassword(pwd, cipherText, salt)
}
