package common

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func SaveFile(file *multipart.FileHeader, path string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	uniqueID := GenerateUniqueID()

	photoData := fmt.Sprintf("%d-%s", uniqueID, file.Filename)

	dst, err := os.Create(path + photoData)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = src.Seek(0, 0); err != nil {
		return err
	}

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return nil
}

func GenerateUniqueID() uint {
	now := time.Now().Unix()
	uniqueID := uint(now)

	return uniqueID
}
