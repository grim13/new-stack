package util

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// FormatValidationErrors mengubah error validasi menjadi map[string]string
func FormatValidationErrors(err error) map[string]string {
	// Inisialisasi map untuk menyimpan pesan error
	errors := make(map[string]string)

	// Lakukan type assertion ke validator.ValidationErrors
	// Ini akan berisi slice dari semua field yang gagal validasi
	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		// Jika error bukan dari validator, kembalikan pesan umum
		errors["error"] = "An unexpected error occurred"
		return errors
	}

	for _, fieldErr := range validationErrors {
		// Dapatkan nama field JSON (atau gunakan nama struct jika tag json tidak ada)
		// Kita sederhanakan dengan mengubah ke huruf kecil
		fieldName := strings.ToLower(fieldErr.Field())

		// Buat pesan error yang lebih ramah berdasarkan tag validasi
		switch fieldErr.Tag() {
		case "required":
			errors[fieldName] = "This field is required."
		case "email":
			errors[fieldName] = "Must be a valid email address."
		case "min":
			errors[fieldName] = fmt.Sprintf("This field must be at least %s characters long.", fieldErr.Param())
		case "max":
			errors[fieldName] = fmt.Sprintf("This field must be at most %s characters long.", fieldErr.Param())
		default:
			errors[fieldName] = "This field is not valid."
		}
	}

	return errors
}
