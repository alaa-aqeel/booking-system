package shared

import "golang.org/x/crypto/bcrypt"

func MakeHash(value string) (string, error) {
	// bcrypt.DefaultCost is 10, safe default
	hashed, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func CheckHash(hashed, value string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(value))
	return err == nil
}
