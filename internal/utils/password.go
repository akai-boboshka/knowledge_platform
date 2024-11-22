package utils

import (
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
)

// HashPassword хеширует заданный пароль с использованием bcrypt
func HashPassword(password string) (string, error) {
	// Генерация хеша пароля
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	encodeToString := hex.EncodeToString(hash)

	return encodeToString, nil
}

// CheckPasswordHash проверяет, соответствует ли пароль заданному хешу
func CheckPasswordHash(password, hash string) bool {
	decodeString, err := hex.DecodeString(hash)
	if err != nil {
		log.Printf("Error checking password: %v", err)
		return false
	}

	// Сравнение пароля с хешем
	err = bcrypt.CompareHashAndPassword(decodeString, []byte(password))
	if err != nil {
		// TODO
		return false
	}

	return true
}

// GenerateRandomString генерирует случайную строку заданной длиной
func GenerateRandomString(length int) string {
	// Генерация случайного пароля
	p := make([]byte, length)
	for i := range p {
		p[i] = byte(97 + rand.Intn(26)) // a-z
	}
	return string(p)
}
