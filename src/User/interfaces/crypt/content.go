package crypt

import (
	entities "pairing_test/src/user/domains/entities"
)

type UserCrypt interface {
	MakeMetaData(content entities.ContentInput) (entities.Content, error)
}