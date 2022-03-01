package drivers

import (
	"fmt"
	"os"

	handler "pairing_test/src/user/interfaces/gateways"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DATABASE"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return  db, nil
}

func (h *handler.SQLHandler) Find(out interface{}, where ...interface{}) *gorm.DB {
	return h.Conn.Find(out, where...)
}

func (h *handler.SQLHandler) First(out interface{}, where ...interface{}) *gorm.DB {
	return h.Conn.First(out, where...)
}

func (h *handler.SQLHandler) Create(value interface{}) (err error) {
	err = handler.Conn.Create(value).Error
	return
}

func (h *handler.SQLHandler) Save(value interface{}) *gorm.DB {
	return h.Conn.Save(value)
}

func (h *handler.SQLHandler) Delete(value interface{}) *gorm.DB {
	return h.Conn.Delete(value)
}

func (h *handler.SQLHandler) Where(query interface{}, args ...interface{}) *gorm.DB {
	return h.Conn.Where(query, args...)
}
