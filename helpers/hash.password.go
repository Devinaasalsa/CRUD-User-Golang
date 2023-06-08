package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(password *string) error {
	resultHash, err := bcrypt.GenerateFromPassword(
		[]byte(*password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	*password = string(resultHash)

	return nil
}

func Compare(password, hashed string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	if err != nil {
		return false, err
	}

	return true, nil
}
