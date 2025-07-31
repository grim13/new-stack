package database

import (
	"log"

	"github.com/grim13/go-api/internal/config" // Impor config untuk akses variabel DB
	"github.com/grim13/go-api/internal/model"  // Impor model yang akan dimigrasi
)

// RunMigrations menjalankan migrasi database otomatis GORM.
func RunMigrations() {
	err := config.DB.AutoMigrate(&model.Role{}, &model.User{}, &model.Permission{}, &model.RolePermission{}, &model.UserRole{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("✅ Database migrated successfully")
}
