package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes a plain text password
func HashPassword(pw string) (string, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

// CheckPassword compares plain with hashed password
func CheckPassword(pw, hashed string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pw))
}
