package database

import (
	"log"

	"github.com/grim13/go-api/internal/config"
	"github.com/grim13/go-api/internal/model"
)

// RunSeeders menjalankan semua seeder untuk mengisi data awal.
func RunSeeders() {
	seedRoles()
	// Anda bisa menambahkan fungsi seeder lain di sini nanti
	// contoh: seedAdminUser()
}

// seedRoles mengisi data awal untuk tabel roles.
func seedRoles() {
	roles := []model.Role{
		{Name: "admin"},
		{Name: "user"},
	}

	for _, role := range roles {
		// FirstOrCreate akan membuat record hanya jika belum ada.
		config.DB.FirstOrCreate(&role, model.Role{Name: role.Name})
	}
	log.Println("👍 Roles seeded successfully")
}
