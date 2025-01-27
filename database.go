package config

import (
	"fmt"

	models "github.com/RaihanMalay21/models_TB_Berkah_Jaya"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func DB_Connection(AKID, SECRETKEY string) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(AKID, SECRETKEY, ""),
	})
	if err != nil {
		panic(err)
	}

	ssmSvc := ssm.New(sess)

	dbUser := getParameter("DB_USER", ssmSvc)
	if dbUser == "" {
		dbUser = "root"
	}
	dbPwdd := getParameter("DB_PASSWORD", ssmSvc)
	if dbPwdd == "" {
		dbPwdd = "90909090"
	}
	dbHost := getParameter("DB_HOST", ssmSvc)
	if dbHost == "" {
		dbHost = "/cloudsql/api-tb-berkah-jaya:us-central1:db-tb-berkah-jaya21"
	}
	dbName := getParameter("DB_NAME", ssmSvc)
	if dbName == "" {
		dbName = "db_tb_berkah_jaya"
	}

	// @unix development
	dbURI := fmt.Sprintf("%s:%s@unix(%s)/%s?parseTime=true", dbUser, dbPwdd, dbHost, dbName)

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
