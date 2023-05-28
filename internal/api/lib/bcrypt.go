package lib

import "golang.org/x/crypto/bcrypt"

func HashPassword(password *string) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	*password = string(hashedPass)
	return nil
}

func ComparePassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
