package http

import (
	m "github.com/Mubinabd/project_control/api/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/Mubinabd/project_control/api/docs"
	"github.com/Mubinabd/project_control/api/handlers"
)

// @title Project Control API Documentation
// @version 1.0
// @description API for Instant Delivery resources
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewGin(h *handlers.Handlers) *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// enforcer, err := casbin.NewEnforcer("./internal/http/casbin/model.conf", "./internal/http/casbin/policy.csv")
	// if err != nil {
	// 	log.Println("Error while creating enforcer: ", err)
	// }
	// router.Use(middlerware.NewAuth(enforcer))

	group := router.Group("/v1/group")
	{
		group.POST("/create", h.CreateGroup)
		group.GET("/:id", h.GetGroup)
		group.PUT("/update/:id", h.UpdateGroup)
		group.DELETE("/delete/:id", h.DeleteGroup)
		group.GET("/list", h.ListGroups)
	}
	private := router.Group("/v1/private")
	{
		private.POST("/create", h.CreatePrivate)
		private.GET("/:id", h.GetPrivate)
		private.PUT("/:id", h.UpdatePrivate)
		private.DELETE("/delete/:id", h.DeletePrivate)
		private.GET("/list", h.ListPrivates)
	}
	router.POST("/register", h.RegisterUser).Use(m.Middleware())
	router.POST("/login", h.LoginUser).Use(m.Middleware())
	router.POST("/forgot-password", h.ForgotPassword)
	router.POST("/reset-password", h.ResetPassword)
	router.GET("/developers", h.GetAllUsers).Use(m.JWTMiddleware())

	user := router.Group("/v1/user").Use(m.JWTMiddleware())
	{
		user.GET("/profiles", h.GetProfile)
		user.PUT("/profiles", h.EditProfile)
		user.PUT("/passwords", h.ChangePassword)
		user.GET("/setting", h.GetSetting)
		user.PUT("/setting", h.EditSetting)
		user.DELETE("/", h.DeleteUser)
	}

	return router
}
