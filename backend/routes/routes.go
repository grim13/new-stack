package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/grim13/go-api/internal/handler"
	"github.com/grim13/go-api/internal/middleware"
	"github.com/grim13/go-api/internal/repository" // Impor repository
)

// SetupRouter sekarang menerima dependensi UserRepository
func SetupRouter(r *gin.Engine, userRepo repository.UserRepository) {
	// Buat instance handler dengan menyuntikkan repository
	authHandler := handler.NewAuthHandler(userRepo)

	// Gunakan method dari instance authHandler
	public := r.Group("/api/auth")
	{
		public.POST("/register", authHandler.Register)
		public.POST("/login", authHandler.Login)
	}

	protected := r.Group("/api/users")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/profile", authHandler.Profile)
	}
	// ...
}
