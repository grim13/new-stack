package auth

import (
	"crypto/rsa"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

func LoadKeys() {
	privKeyPath := os.Getenv("JWT_PRIVATE_KEY_FILE")
	pubKeyPath := os.Getenv("JWT_PUBLIC_KEY_FILE")

	privBytes, err := os.ReadFile(privKeyPath)
	if err != nil {
		panic("Failed to read private key")
	}
	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privBytes)
	if err != nil {
		panic("Failed to parse private key")
	}

	pubBytes, err := os.ReadFile(pubKeyPath)
	if err != nil {
		panic("Failed to read public key")
	}
	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(pubBytes)
	if err != nil {
		panic("Failed to parse public key")
	}
}

func GenerateToken(userID uuid.UUID, username, name, email string, roleId uint, roleName string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"userdata": map[string]interface{}{
			"username": username,
			"name":     name,
			"email":    email,
			"roleId":   roleId,
			"roleName": roleName,
		},
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}
	println("Generating token for user:", claims)

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(privateKey)
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})
}
