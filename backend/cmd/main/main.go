package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/grim13/go-api/database"
	"github.com/grim13/go-api/internal/auth"
	"github.com/grim13/go-api/internal/config" // Impor repository
	"github.com/grim13/go-api/routes"
)

func main() {
	// ... (inisialisasi config, db, auth)
	config.LoadConfig()
	config.ConnectDB()
	auth.LoadKeys()

	// ... (seeder)
	database.RunSeeders()

	// === Dependency Injection ===
	// 1. Buat instance dari repository

	// 2. Setup router dan suntikkan repository ke dalamnya
	r := gin.Default()
	routes.SetupRouter(r, config.DB) // Berikan repo ke router

	// ... (jalankan server)
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("🚀 Server starting on port %s", port)
	r.Run(":" + port)
}
