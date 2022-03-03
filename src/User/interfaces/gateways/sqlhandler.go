package gateways

import "gorm.io/gorm"

type SQLHandler interface {
	Find(interface{}, ...interface{}) (*gorm.DB, error)
	First(interface{}, ...interface{}) *gorm.DB
	Create(interface{}) (*gorm.DB, error)
	Save(interface{}) *gorm.DB
	Delete(interface{}) *gorm.DB
	Where(interface{}, ...interface{}) *gorm.DB
}
