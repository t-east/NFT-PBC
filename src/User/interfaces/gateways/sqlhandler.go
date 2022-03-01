package gateways

import "gorm.io/gorm"

type SQLHandler interface {
	Find(interface{}, ...interface{}) *gorm.DB
	First(interface{}, ...interface{}) *gorm.DB
	Create(interface{}) error
	Save(interface{}) *gorm.DB
	Delete(interface{}) *gorm.DB
	Where(interface{}, ...interface{}) *gorm.DB
}
