package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Role mendefinisikan peran pengguna (misal: admin, user)
type Role struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
	Users []User `gorm:"many2many:user_roles;"`
}

// User mendefinisikan struktur data pengguna
type User struct {
	BaseModel        // Menggunakan BaseModel kustom
	Name      string `gorm:"not null"`
	Username  string `gorm:"unique;not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	RoleID    uint   // Foreign key untuk Role
	Role      Role   `gorm:"foreignKey:RoleID"` // Relasi ke Role
	Roles     []Role `gorm:"many2many:user_roles;"`
}

type Permission struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"unique;not null"`
	Description string `gorm:"not null"`
}

type RolePermission struct {
	gorm.Model
	RoleID       uint       `gorm:"not null"`
	PermissionID uint       `gorm:"not null"`
	Role         Role       `gorm:"foreignKey:RoleID"`
	Permission   Permission `gorm:"foreignKey:PermissionID"`
}

type UserRole struct {
	gorm.Model
	UserID uuid.UUID `gorm:"not null"`
	RoleID uint      `gorm:"not null"`
	User   User      `gorm:"foreignKey:UserID"`
	Role   Role      `gorm:"foreignKey:RoleID"`
}

// RegisterInput mendefinisikan data untuk registrasi
type RegisterInput struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// LoginInput mendefinisikan data untuk login
type LoginInput struct {
	// Bisa diisi dengan username atau email
	Identifier string `json:"identifier" binding:"required"`
	Password   string `json:"password" binding:"required"`
}
