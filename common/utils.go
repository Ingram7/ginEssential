package common

import (
	"ginessential/dao"
	"ginessential/models"
	"math/rand"
	"time"
)

func IsTelephoneExist(telephone string) bool {
	var user models.User
	dao.GetTelephone(&user, telephone)
	if user.ID != 0 {
		return true
	}
	return false
}

func RandomString(n int) string {
	letters := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]byte, n)
	rand.Seed(time.Now().UnixNano())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func ToUserDto(user models.User) models.UserDto {
	return models.UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}
