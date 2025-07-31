package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Role mendefinisikan peran pengguna (misal: admin, user)
type Role struct {
	ID          uint         `gorm:"primaryKey" json:"id"`
	Name        string       `gorm:"unique;not null" json:"name"`
	Permissions []Permission `gorm:"many2many:role_permissions;" json:"permissions"`
	Users       []User       `gorm:"many2many:user_roles;" json:"users"`
}

// User mendefinisikan struktur data pengguna
type User struct {
	BaseModel        // Menggunakan BaseModel kustom
	Name      string `gorm:"not null" json:"name"`
	Username  string `gorm:"unique;not null" json:"username"`
	Email     string `gorm:"unique;not null" json:"email"`
	Password  string `gorm:"not null" json:"password"`
	RoleID    uint   `json:"role_id"`
	Role      Role   `gorm:"foreignKey:RoleID" json:"role"`
	Roles     []Role `gorm:"many2many:user_roles;" json:"roles"`
}

type Permission struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"unique;not null" json:"name"`
	Description string `gorm:"not null" json:"description"`
}

type RolePermission struct {
	gorm.Model
	RoleID       uint       `gorm:"not null" json:"role_id"`
	PermissionID uint       `gorm:"not null" json:"permission_id"`
	Role         Role       `gorm:"foreignKey:RoleID" json:"role"`
	Permission   Permission `gorm:"foreignKey:PermissionID" json:"permission"`
}

type UserRole struct {
	gorm.Model
	UserID uuid.UUID `gorm:"not null" json:"user_id"`
	RoleID uint      `gorm:"not null" json:"role_id"`
	User   User      `gorm:"foreignKey:UserID" json:"user"`
	Role   Role      `gorm:"foreignKey:RoleID" json:"role"`
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
