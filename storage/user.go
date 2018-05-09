package storage

import (
	"errors"
	"time"
)

var (
	ErrUserExist = errors.New("user already exist")
)

// UserInfo represents the model of table user_infos.
type UserInfo struct {
	ID         uint64
	Name       string
	CreateTime time.Time
}

// UserStorage controls the read/write operations of table user_infos.
type UserStorage struct {
}

// CreateUser inserts one user to table user_infos.
func (u *UserStorage) CreateUser(user *UserInfo) error {
	created, err := db.Model(user).
		Column("id").
		Where("name = ?name").
		OnConflict("DO NOTHING").
		Returning("id").
		SelectOrInsert()
	if err != nil {
		return err
	}
	if !created {
		return ErrUserExist
	}
	return nil
}

// ListAllUsers retrieves all the users from table user_infos.
func (u *UserStorage) ListAllUsers() ([]*UserInfo, error) {
	var users []*UserInfo
	_, err := db.Query(&users, `SELECT * FROM user_infos`)
	if err != nil {
		return nil, err
	}
	return users, nil
}
