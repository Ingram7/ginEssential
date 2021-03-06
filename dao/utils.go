package dao

import "ginessential/models"

func GetTelephone(user *models.User, telephone string) {
	DB.Debug().Where("telephone = ?", telephone).First(user)
}

func CreateUser(user *models.User) {
	DB.Debug().Create(user)
}

func GetUserId(user *models.User, userId uint) {
	DB.Debug().First(&user, userId)
}
