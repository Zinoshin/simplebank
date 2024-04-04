package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hasg password: %w", err)
	}
	return string(hashedPassword), nil
}

// 비밀번호와 해시 비밀번호 비교
func CheckPassword(password string, hashPasssword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPasssword), []byte(password))
}
