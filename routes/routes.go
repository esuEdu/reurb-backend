package routes

import (
	"net/http"

	"github.com/esuEdu/reurb-backend/internal/handlers"
	"github.com/esuEdu/reurb-backend/internal/middleware"
	"github.com/esuEdu/reurb-backend/internal/repositories"
	"github.com/esuEdu/reurb-backend/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	router.POST("/register", userHandler.RegisterUser)
	router.POST("/login", userHandler.AuthenticateUser)

	authorized := router.Group("/")

	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "pong"})
		})

		authorized.GET("/user/:id", userHandler.GetUserByID)
	}

}
