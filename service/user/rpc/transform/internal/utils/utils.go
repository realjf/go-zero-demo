package utils

import "golang.org/x/crypto/bcrypt"

func EncodePassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	encodePwd := string(hash)
	return encodePwd, nil
}

func CheckPassword(password, cryptedpwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(cryptedpwd), []byte(password))
	return err == nil
}
