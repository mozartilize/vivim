package auth

import (
	"net/http"
	"time"
	userRepo "vivim/user/repository"
	"vivim/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func auth(c echo.Context) error {
	loginInfo := new(LoginSchema)
	if err := c.Bind(loginInfo); err != nil {
		return &echo.HTTPError{
			Code:     http.StatusUnauthorized,
			Message:  "Authentication failed",
			Internal: err,
		}
	}

	u, err := userRepo.GetUserByUsernameOrEmail(loginInfo.Identity)
	if err != nil {
		return &echo.HTTPError{
			Code:     http.StatusUnauthorized,
			Message:  "Authentication failed",
			Internal: err,
		}
	}
	if !u.CheckPassword(loginInfo.Password) {
		return &echo.HTTPError{
			Code:     http.StatusUnauthorized,
			Message:  "Authentication failed",
			Internal: err,
		}
	}
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = u.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(c.Get("config").(*viper.Viper).GetString("secret_key")))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"token": t,
	})
}

var UnauthApiHandlers = []utils.RouteDef{
	{Method: http.MethodPost, Path: "/auth/", Handler: auth},
}
