package config

import (
	"fmt"
	"os"

	models "github.com/RaihanMalay21/models_TB_Berkah_Jaya"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func DB_Connection() {
	dbUser := os.Getenv("DB_USER")
	dbPwdd := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	fmt.Println(dbUser, dbPwdd, dbHost, dbName)
	// @unix development
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", dbUser, dbPwdd, dbHost, 3306, dbName)

	db, err := gorm.Open(mysql.Open(dbURI))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Barang{})
	db.AutoMigrate(&models.Hadiah{})
	db.AutoMigrate(&models.Pembelian{})
	db.AutoMigrate(&models.Pembelian_Per_Item{})
	db.AutoMigrate(&models.HadiahUser{})

	DB = db
}
