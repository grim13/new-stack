package repository

import (
	"github.com/grim13/go-api/internal/model"
	"gorm.io/gorm"
)

// UserRepository mendefinisikan kontrak untuk user repository
type UserRepository interface {
	Save(user *model.User) error
	Update(id uint, user *model.User) error
	Delete(id uint) error
	FindByEmailOrUsername(identifier string) (*model.User, error)
	FindByID(id uint) (*model.User, error)
	FindRoleByName(name string) (*model.Role, error)
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

func (r *userRepositoryGORM) Update(id uint, user *model.User) error {
	return r.db.Model(&model.User{}).Where("id = ?", id).Updates(user).Error
}

func (r *userRepositoryGORM) Delete(id uint) error {
	return r.db.Delete(&model.User{}, id).Error
}

func (r *userRepositoryGORM) FindByEmailOrUsername(identifier string) (*model.User, error) {
	var user model.User
	println("Finding user by identifier:", identifier) // Debugging line
	err := r.db.Preload("Role").Preload("Roles").Where("email = ? OR username = ?", identifier, identifier).First(&user).Error
	return &user, err
}

func (r *userRepositoryGORM) FindByID(id uint) (*model.User, error) {
	var user model.User
	err := r.db.Preload("Role").First(&user, id).Error
	return &user, err
}

func (r *userRepositoryGORM) FindRoleByName(name string) (*model.Role, error) {
	var role model.Role
	err := r.db.Where("name = ?", name).First(&role).Error
	return &role, err
}
