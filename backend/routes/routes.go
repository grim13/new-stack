package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/grim13/go-api/internal/handler"
	"github.com/grim13/go-api/internal/middleware"
	"github.com/grim13/go-api/internal/repository" // Impor repository
	"gorm.io/gorm"
)

// SetupRouter sekarang menerima dependensi UserRepository
func SetupRouter(r *gin.Engine, db *gorm.DB) {
	// Buat instance handler dengan menyuntikkan repository

	userRepo := repository.NewUserRepositoryGORM(db)
	authHandler := handler.NewAuthHandler(userRepo)
	userHandler := handler.NewUserHandler(userRepo)

	// Gunakan method dari instance authHandler
	public := r.Group("/api/auth")
	{
		public.POST("/register", authHandler.Register)
		public.POST("/login", authHandler.Login)
	}

	protected := r.Group("/api/users")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/", userHandler.GetAllUsers)
		protected.GET("/profile", authHandler.Profile)
	}
	// ...
}
