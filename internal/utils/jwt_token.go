package utils

import (
	"awesomeProject/pkg/models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// Секретный ключ для подписи токенов
var jwtSecret = []byte("SecretKeyForSignature")

// GenerateJWT создает новый JWT токен для пользователя
func GenerateJWT(user models.User) (string, error) {
	// Определяем срок действия токена
	expirationTime := time.Now().Add(15 * time.Minute)

	//Создаем токен с помощью стандарта HMAC и алгоритма подписи
	claims := &jwt.MapClaims{
		"user_id": user.ID,
		"role_id": user.RoleID,
		"exp":     expirationTime.Unix(),
	}

	// Создание токена с подписью
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписание токена с использованием секретного ключа
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateJWT проверяет валидность JWT токена
func ValidateJWT(tokenString string) (int, error) {
	// Парсинг и проверка токена
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Проверка алгоритма подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return 0, err
	}

	// Проверка валидности токена и извлечение данных
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := claims["user_id"].(float64)
		return int(userID), nil
	}

	return 0, fmt.Errorf("invalid token")
}

func ExampleV1() {
	// Пример использования
	username := models.User{
		Username: "john_doe",
	}

	// Создание JWT токена
	token, err := GenerateJWT(username)
	if err != nil {
		fmt.Println("Error generating JWT:", err)
		return
	}

	fmt.Println("Generated JWT Token:", token)

	// Проверка JWT токена
	validUsername, err := ValidateJWT(token)
	if err != nil {
		fmt.Println("Error validating JWT:", err)
		return
	}

	fmt.Println("Validated username from token:", validUsername)
}
