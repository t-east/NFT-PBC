package drivers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SQLHandler struct {
	Conn *gorm.DB
}

func GetDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DATABASE"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return  db, nil
}

func NewSQLHandler() (*SQLHandler, error) {
	db, err := GetDB()
	if err != nil {
		log.Fatal(err)
	}
	sqlHandler := new(SQLHandler)
	sqlHandler.Conn = db
	return sqlHandler, nil
}