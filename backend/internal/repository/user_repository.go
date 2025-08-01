package repository

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/grim13/go-api/internal/model"
	"gorm.io/gorm"
)

// UserRepository mendefinisikan kontrak untuk user repository
type UserRepository interface {
	Save(user *model.User) error
	Update(id uuid.UUID, user *model.User) error
	Delete(id uuid.UUID) error
	FindByEmailOrUsername(identifier string) (*model.User, error)
	FindByID(id uuid.UUID) (*model.User, error)
	FindRoleByName(name string) (*model.Role, error)
	FindAll(page, limit int, sortBy, sortOrder string, searchQuery string) ([]model.User, int64, error)
}

// userRepositoryGORM adalah implementasi konkret dari UserRepository menggunakan GORM
type userRepositoryGORM struct {
	db *gorm.DB
}

// NewUserRepositoryGORM membuat instance baru dari userRepositoryGORM
func NewUserRepositoryGORM(db *gorm.DB) UserRepository {
	return &userRepositoryGORM{db: db}
}

// Implementasi setiap fungsi dari interface

func (r *userRepositoryGORM) Save(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *userRepositoryGORM) Update(id uuid.UUID, user *model.User) error {
	return r.db.Model(&model.User{}).Where("id = ?", id).Updates(user).Error
}

func (r *userRepositoryGORM) Delete(id uuid.UUID) error {
	return r.db.Delete(&model.User{}, id).Error
}

func (r *userRepositoryGORM) FindByEmailOrUsername(identifier string) (*model.User, error) {
	var user model.User
	println("Finding user by identifier:", identifier) // Debugging line
	err := r.db.Preload("Role").Preload("Roles").Where("email = ? OR username = ?", identifier, identifier).First(&user).Error
	return &user, err
}

func (r *userRepositoryGORM) FindByID(id uuid.UUID) (*model.User, error) {
	var user model.User
	err := r.db.Preload("Role").First(&user, id).Error
	return &user, err
}

func (r *userRepositoryGORM) FindAll(page, limit int, sortBy, sortOrder string, searchQuery string) ([]model.User, int64, error) {
	var users []model.User
	var totalRecords int64

	// Hitung total record terlebih dahulu untuk paginasi
	if err := r.db.Model(&model.User{}).Count(&totalRecords).Error; err != nil {
		return nil, 0, err
	}

	// Ambil data yang sudah dipaginasi
	offset := (page - 1) * limit
	orderClause := fmt.Sprintf("%s %s", sortBy, sortOrder)

	err := r.db.Preload("Role").Preload("Roles").Order(orderClause).Limit(limit).Offset(offset).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, totalRecords, nil
}

func (r *userRepositoryGORM) FindRoleByName(name string) (*model.Role, error) {
	var role model.Role
	err := r.db.Where("name = ?", name).First(&role).Error
	return &role, err
}
