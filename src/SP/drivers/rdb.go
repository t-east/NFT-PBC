package drivers

import (
	"fmt"

	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

type SQLHandler struct {
	Conn *gorm.DB
}

type Config struct {
	Protocol string
	User     string
	Password string
	Name     string
}

// MySQL設定の読み込み関数 -> drivers.Config
func loadConfig() *Config {
	return &Config{"tcp(mysql)", "root", "password", "sample_db"}
}

// MySQLハンドラーのinit関数 -> (*SQLHandler, error)
func NewSQLHandler() (*SQLHandler, error) {
	c := loadConfig()
	connectTemplate := "%s:%s@%s/%s?parseTime=true"
	dsn := fmt.Sprintf(connectTemplate, c.User, c.Password, c.Protocol, c.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return &SQLHandler{}, err
	}
	sqlHandler := new(SQLHandler)
	sqlHandler.Conn = db
	return sqlHandler, nil
}

func (handler *SQLHandler) Find(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.Find(out, where...)
}

func (handler *SQLHandler) First(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.First(out, where...)
}

func (handler *SQLHandler) Create(value interface{}) (err error) {
	err = handler.Conn.Create(value).Error
	return
}

func (handler *SQLHandler) Save(value interface{}) *gorm.DB {
	return handler.Conn.Save(value)
}

func (handler *SQLHandler) Delete(value interface{}) *gorm.DB {
	return handler.Conn.Delete(value)
}

func (handler *SQLHandler) Where(query interface{}, args ...interface{}) *gorm.DB {
	return handler.Conn.Where(query, args...)
}
