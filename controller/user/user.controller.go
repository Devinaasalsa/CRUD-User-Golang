package controllerUser

import (
	"fmt"
	"myapp-me/config"
	"myapp-me/dao"
	"myapp-me/dto"
	"myapp-me/helpers"
	"myapp-me/models"
	"net/http"

	validators "myapp-me/helpers/validator"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)



func CreateUser(echo echo.Context) error {
	user := new(dto.UserCreateInput)
	var db = config.DB()

	if err := echo.Bind(&user); err != nil {
		message := dao.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Payload: nil,
		}

		return echo.JSON(http.StatusInternalServerError, message)
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			return err
		}

		var errors []map[string]string
		for _, err := range validationErrors {
			field := err.Field()
			message := validators.GetValidationErrorMessage(err.Tag(), field)

			errorObj := map[string]string{
				"field":   field,
				"message": message,
			}
			errors = append(errors, errorObj)
		}

		return echo.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors": errors,
		})
	}
	data := &models.User{
		Username: user.Username,
		Password: user.Password,
	}

	//Status OK
	if err := helpers.HashPassword(&data.Password); err != nil {
		return err
	}

	if err := db.Create(&data).Error; err != nil {
		message := dao.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Payload: nil,
		}

		return echo.JSON(http.StatusInternalServerError, message)
	}

	response := dao.Response{
		Status:  200,
		Message: http.StatusText(200),
		Payload: data,
	}

	return echo.JSON(http.StatusOK, response)
}

func GetUser(echo echo.Context) error {
	var db = config.DB()
	var users []*models.User

	db.Find(&users)

	response := dao.Response{
		Status:  200,
		Message: http.StatusText(200),
		Payload: users,
	}

	return echo.JSON(http.StatusOK, response)
}

func GetUserWhereDeletedAtIsNotNull(echo echo.Context) error {
	db := config.DB()
	users := []*models.User{}

	if err := db.Unscoped().Where("deleted_at IS NOT NULL").Find(&users).Error; err != nil {
		return err
	}

	response := dao.Response{
		Status:  200,
		Message: http.StatusText(200),
		Payload: users,
	}

	return echo.JSON(response.Status, response)
}

func DeleteUserById(echo echo.Context) error {
	id := echo.Param("id")
	db := config.DB()

	user := new(models.User)

	err := db.Delete(&user, id).Error

	if err != nil {
		message := dao.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Payload: nil,
		}

		return echo.JSON(message.Status, message)
	}

	response := dao.Response{
		Status:  http.StatusOK,
		Message: fmt.Sprintf("Record id %s has deleted.", id),
		Payload: nil,
	}

	return echo.JSON(response.Status, response)
}
