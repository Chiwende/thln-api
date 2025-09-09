package services

import "github.com/google/uuid"

func GenerateUUIDv4() (string, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return u.String(), nil
}
