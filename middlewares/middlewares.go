package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/dgrijalva/jwt-go" // Pastikan untuk menambahkan dependensi ini
)

// AuthMiddleware untuk memeriksa JWT
func AuthMiddleware(c *fiber.Ctx) error {
	// Mendapatkan token dari header Authorization
	token := c.Get("Authorization")

	if token == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Authorization header is missing"})
	}

	// Menghapus "Bearer " dari token
	token = strings.TrimPrefix(token, "Bearer ")

	// Memverifikasi token
	claims := jwt.MapClaims{}
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Pastikan metode signing adalah HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil // Ganti "secret" dengan kunci signing Anda
	})

	if err != nil || !jwtToken.Valid {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
	}

	// Menyimpan klaim ke dalam konteks
	c.Locals("claims", claims)

	return c.Next()
}
