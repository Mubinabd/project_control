package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/Mubinabd/project_control/api/docs"
	"github.com/Mubinabd/project_control/api/handlers"
	"github.com/Mubinabd/project_control/api/middleware"
)

// @title Authentication Service API
// @version 1.0
// @description API for Authentication Service
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func Engine(handler *handlers.Handlers) *gin.Engine {
	router := gin.Default()

	router.Use(CORSMiddleware())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/register", handler.RegisterUser).Use(middleware.Middleware())
	router.POST("/login", handler.LoginUser).Use(middleware.Middleware())
	router.POST("/forgot-password", handler.ForgotPassword)
	router.POST("/reset-password", handler.ResetPassword)
	router.GET("/users", handler.GetAllUsers).Use(middleware.JWTMiddleware())

	user := router.Group("/user").Use(middleware.JWTMiddleware())
	{
		user.GET("/profiles", handler.GetProfile)
		user.PUT("/profiles", handler.EditProfile)
		user.PUT("/passwords", handler.ChangePassword)
		user.GET("/setting", handler.GetSetting)
		user.PUT("/setting", handler.EditSetting)
		user.DELETE("/", handler.DeleteUser)
	}

	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}
