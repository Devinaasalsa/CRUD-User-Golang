package controllerAuth

import (
	"fmt"

	"myapp-me/config"
	"myapp-me/dao"
	"myapp-me/dto"
	"myapp-me/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Login(echo echo.Context) error {
	db := config.DB()

	loginRequest := new(dto.RequestLogin)
	user := new(models.User)

	//Binding data
	if err := echo.Bind(loginRequest); err != nil {
		message := dao.Response{
			Status:  http.StatusForbidden,
			Message: err.Error(),
			Payload: nil,
		}

		return echo.JSON(message.Status, message)
	}

	user, err := loginRequest.GetUserByUsername(db, loginRequest.Username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			message := dao.Response{
				Status:  http.StatusForbidden,
				Message: fmt.Sprintf("Username %s does not exist", loginRequest.Username),
				Payload: nil,
			}

			return echo.JSON(message.Status, message)
		}

		message := dao.Response{
			Status:  http.StatusInternalServerError,
			Message: "Database error",
			Payload: nil,
		}

		return echo.JSON(message.Status, message)
	}

	if !loginRequest.IsPasswordMatch(loginRequest.Password, user.Password) {
		message := dao.Response{
			Status:  http.StatusForbidden,
			Message: "Invalid password credential!",
			Payload: nil,
		}

		return echo.JSON(message.Status, message)
	}

	//Generate token
	token, err := loginRequest.GenerateToken(user.Username)
	if err != nil {
		message := dao.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Payload: nil,
		}

		return echo.JSON(message.Status, message)
	}

	message := dao.Response{
		Status:  http.StatusOK,
		Message: fmt.Sprintf("Welcome back! %s", user.Username),
		Payload: token,
	}

	return echo.JSON(message.Status, message)
}

/*
When this method is hit, this will be delete the jwt
from localstorage in front-end.
*/
func Logout(echo echo.Context) error {
	//Restart cookie
	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}
	echo.SetCookie(cookie)

	return echo.NoContent(http.StatusOK)
}
