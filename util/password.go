package util

import (
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	cost, _ := strconv.Atoi(os.Getenv("PW_SALT_ROUND"))
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	PanicIfError(err)
	return string(bytes)
}

func ComparePasswordHash(clearText string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(clearText))
	return err == nil
}
