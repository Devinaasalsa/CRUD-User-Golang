package dto

import (
	"myapp-me/config"
	"myapp-me/helpers"
	"myapp-me/models"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type RequestLogin struct {
	Username string `json:"username" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
}

func (request *RequestLogin) GetUserByUsername(db *gorm.DB, username string) (*models.User, error) {
	recordUser := new(models.User)

	err := db.First(recordUser, "username = ?", username).Error

	return recordUser, err
}

func (request *RequestLogin) IsPasswordMatch(password, hashedPassword string) bool {
	isMatch, err := helpers.Compare(password, hashedPassword)

	if err != nil || !isMatch {
		return false
	}

	return true
}

func (request *RequestLogin) GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = username
	claims["admin"] = true
	claims["exp"] = config.ExpiresTime

	return token.SignedString([]byte(config.SignatureKey))
}
