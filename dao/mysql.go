package dao

import (
	"fmt"
	"ginessential/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var (
	DB *gorm.DB
)

func GetDB() *gorm.DB {
	return DB
}

func InitMySQL() (err error) {
	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset,
	)
	DB, err = gorm.Open(driverName, args)
	//DB, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/ginessential?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		return fmt.Errorf("error in InitMySQL.Open, error: [%s]", err.Error())
	}
	DB.AutoMigrate(&models.User{})

	return nil
}

func Close() {
	DB.Close()
}
