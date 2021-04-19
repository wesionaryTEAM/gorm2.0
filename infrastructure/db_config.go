package infrastructure

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm2.0/utils"
)

//Setup Models: initializaing the mysql database
func GetDatabaseInstance() *gorm.DB {
	get := utils.GetEnvWithKey
	USER := get("DB_USER")
	PASS := get("DB_PASS")
	HOST := get("DB_HOST")
	PORT := get("DB_PORT")
	DBNAME := get("DB_NAME")
	ENVIRONMENT := get("ENVIRONMENT")

	createDBDsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", USER, PASS, HOST, PORT)
	database, err := gorm.Open(mysql.Open(createDBDsn), &gorm.Config{})

	_ = database.Exec("CREATE DATABASE IF NOT EXISTS " + DBNAME + ";")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)

	if ENVIRONMENT != "local" {
		dsn = fmt.Sprintf(
			"%s:%s@unix(/cloudsql/%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			USER,
			PASS,
			HOST,
			DBNAME,
		)
	}

	fmt.Println("The DSN is:::", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	return db
}
