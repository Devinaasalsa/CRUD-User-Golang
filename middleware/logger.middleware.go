package middleware

import (
	_ "fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func MakeLogEntry(echo echo.Context) *log.Entry {
	if echo == nil {
		return log.WithFields(
			log.Fields{
				"at": time.Now().Format("2006-01-02 15:04:05"),
			},
		)
	}

	return log.WithFields(
		log.Fields{
			"at":     time.Now().Format("2006-01-02 15:04:05"),
			"method": echo.Request().Method,
			"uri":    echo.Request().URL.String(),
			"ip":     echo.Request().RemoteAddr,
		},
	)
}

func LoggingRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		MakeLogEntry(c).Info("Incoming request")

		return next(c)
	}
}

// func ErrorHandler(err error, c echo.Context) {
// 	report, ok := err.(*echo.HTTPError)
// 	if ok {
// 		report.Message = map[string]interface{}{
// 			"error":       "HTTP Error",
// 			"message":     report.Message,
// 			"status code": report.Code,
// 		}
// 	} else {
// 		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
// 	}

// 	MakeLogEntry(c).Error(report.Message)

// 	c.JSON(report.Code, report.Message.(map[string]interface{}))
// }
func ErrorHandler(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if !ok || report.Message == nil {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	MakeLogEntry(c).Error(report.Message)

	var message string
	if errMsg, ok := report.Message.(string); ok {
		message = errMsg
	} else {
		message = report.Error()
	}

	c.JSON(report.Code, message)
}