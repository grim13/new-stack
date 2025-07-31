// File: backend/internal/handler/user_handler.go

package handler

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/grim13/go-api/internal/repository"
)

type UserHandler struct {
	userRepo repository.UserRepository
}

func NewUserHandler(userRepo repository.UserRepository) *UserHandler {
	return &UserHandler{userRepo: userRepo}
}

// Struct untuk respons paginasi
type PaginatedResponse struct {
	Data interface{}    `json:"data"`
	Meta PaginationMeta `json:"meta"`
}

// Struct untuk metadata paginasi
type PaginationMeta struct {
	Page         int   `json:"page"`
	Limit        int   `json:"limit"`
	TotalRecords int64 `json:"total_records"`
	TotalPages   int   `json:"total_pages"`
}

// GetAllUsers menangani permintaan untuk mendapatkan daftar pengguna dengan paginasi
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	sortBy := c.DefaultQuery("sortBy", "created_at")
	sortOrder := c.DefaultQuery("sortOrder", "desc")

	// Pastikan limit dan page valid
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	// 2. Panggil repository untuk mendapatkan data
	users, totalRecords, err := h.userRepo.FindAll(page, limit, sortBy, sortOrder)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	// 3. Hitung total halaman
	totalPages := 0
	if totalRecords > 0 {
		totalPages = int(math.Ceil(float64(totalRecords) / float64(limit)))
	}

	// 4. Buat respons
	response := PaginatedResponse{
		Data: users,
		Meta: PaginationMeta{
			Page:         page,
			Limit:        limit,
			TotalRecords: totalRecords,
			TotalPages:   totalPages,
		},
	}
	c.JSON(http.StatusOK, response)
}
