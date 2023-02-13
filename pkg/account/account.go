package account

import "math/rand"

type Account struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	AvatarUrl string `json:"avatarUrl"`
	DOB       string `json:"dob"`
	Password  string `json:"password"`
}

func GenerateRandomID(n int) string {
	var letters = []int32(
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
