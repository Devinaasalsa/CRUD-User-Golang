package routes

import (
	controllerAuth "myapp-me/controller/auth"
	controllerUser "myapp-me/controller/user"

	middleware "myapp-me/middleware"

	"github.com/labstack/echo/v4"
)

func RoutesInit(echo *echo.Echo) {

	globalPath := echo.Group("/api/v1")

	userRoutes := globalPath.Group("/user", middleware.IsAuthenticated)
	userRoutes.GET("/get", controllerUser.GetUser)
	userRoutes.POST("/create", controllerUser.CreateUser)
	userRoutes.GET("/get/deleted", controllerUser.GetUserWhereDeletedAtIsNotNull)
	userRoutes.DELETE("/delete/:id", controllerUser.DeleteUserById)


	authRoutes := globalPath.Group("/auth")
	authRoutes.POST("/login", controllerAuth.Login)
	authRoutes.POST("/logout", controllerAuth.Logout)

}
