package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/grim13/go-api/internal/auth"
	"github.com/grim13/go-api/internal/model"
	"github.com/grim13/go-api/internal/repository" // Impor repository
	"github.com/grim13/go-api/internal/util"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// AuthHandler menampung dependensi untuk auth handler, yaitu user repository
type AuthHandler struct {
	userRepo repository.UserRepository
}

// NewAuthHandler membuat instance baru dari AuthHandler
func NewAuthHandler(userRepo repository.UserRepository) *AuthHandler {
	return &AuthHandler{userRepo: userRepo}
}

// --- Ubah semua method menjadi milik struct AuthHandler ---

func (h *AuthHandler) Register(c *gin.Context) {
	var input model.RegisterInput

	// Tangkap error dari ShouldBindJSON
	if err := c.ShouldBindJSON(&input); err != nil {
		// Panggil fungsi helper kita untuk memformat error
		errors := util.FormatValidationErrors(err)

		// Kembalikan objek JSON dengan key "errors"
		c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	defaultRole, err := h.userRepo.FindRoleByName("user") // Gunakan repo
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Default role not found"})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	user := model.User{
		Name:     input.Name,
		Username: input.Username,
		Email:    input.Email,
		Password: string(hashedPassword),
		RoleID:   defaultRole.ID,
		Roles:    []model.Role{*defaultRole},
	}

	if err := h.userRepo.Save(&user); err != nil { // Gunakan repo
		if strings.Contains(err.Error(), "duplicate key") {
			c.JSON(http.StatusConflict, gin.H{"error": "Username or Email already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Registration successful"})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var input model.LoginInput
	// Tangkap error dari ShouldBindJSON
	if err := c.ShouldBindJSON(&input); err != nil {
		// Panggil fungsi helper kita untuk memformat error
		errors := util.FormatValidationErrors(err)

		// Kembalikan objek JSON dengan key "errors"
		println("Login input error:", errors)
		// c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	// debug input
	// println("Login input:", input.Identifier, input.Password)

	user, err := h.userRepo.FindByEmailOrUsername(input.Identifier) // Gunakan repo
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid identifier or password"})
		return
	}

	println("User found:", user.Roles[0].ID)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid identifier or password"})
		return
	}

	// ... (logika bcrypt dan generate token sama)
	// ...
	token, _ := auth.GenerateToken(user.ID, user.Username, user.Name, user.Email, user.RoleID, user.Role.Name)
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *AuthHandler) Profile(c *gin.Context) {
	userIDContext, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: User ID not found in token"})
		return
	}
	//konversi userIDContext ke uuid.UUID
	if userIDContext == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: User ID is nil"})
		return
	}
	userIDUUID, err := uuid.Parse(userIDContext.(string))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: User ID is not a valid UUID", "userID": userIDContext})
		return
	}
	user, err := h.userRepo.FindByID(userIDUUID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch user profile"})
		return
	}

	// Buat respons yang aman (tanpa password)
	userProfile := gin.H{
		"id":       user.ID,
		"name":     user.Name,
		"username": user.Username,
		"email":    user.Email,
		"role":     user.Role.Name,
	}

	c.JSON(http.StatusOK, userProfile)
}
