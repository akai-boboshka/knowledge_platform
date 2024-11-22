package models

import (
	"errors"
	"fmt"
	"regexp"
	"unicode"
)

type User struct {
	ID       int     `json:"id" gorm:"id,primaryKey"`
	Username string  `json:"username" binding:"required" gorm:"username"`
	Password *string `json:"password,omitempty"`
	Email    string  `json:"email"`
}

func (User) TableName() string {
	return "public.users"
}

// ValidateUser функция, которая проверяет данные пользователя на соответствие по заданным критериям.
func (u *User) ValidateUser() error {
	if len(u.Username) < 4 {
		return errors.New("The username must be longer than 4 symbols")
	}
	if err := u.ValidatePassword(); err != nil {
		return err
	}

	return nil
}

func (u *User) ValidatePassword() error {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	if u.Password == nil {
		return fmt.Errorf("The password field is empty")
	}

	if len(*u.Password) >= 7 {
		hasMinLen = true
	}

	for _, char := range *u.Password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasMinLen {
		return fmt.Errorf("password must be at least 7 characters long")
	}
	if !hasUpper {
		return fmt.Errorf("password must contain at least one uppercase letter")
	}
	if !hasLower {
		return fmt.Errorf("password must contain at least one lowercase letter")
	}
	if !hasNumber {
		return fmt.Errorf("password must contain at least one number")
	}
	if !hasSpecial {
		return fmt.Errorf("password must contain at least one special character")
	}

	return nil
}

type Profile struct {
	ID      string `json:"id" gorm:"id,primaryKey"`
	UserID  int    `json:"user_id" gorm:"user_id,foreignKey"`
	Email   string `json:"email"`
	Age     int    `json:"age"`
	AboutMe string `json:"about_me"`
}

func (p *Profile) ValidateProfile() error {
	if !p.isValidEmail() {
		return fmt.Errorf("Invalid email")
	}
	return nil
}

func (p *Profile) isValidEmail() bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(p.Email)
}
