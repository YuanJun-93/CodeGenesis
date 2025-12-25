package model

import (
	"context"

	"gorm.io/gorm"
)

// ErrNotFound is a standard error for record not found
var ErrNotFound = gorm.ErrRecordNotFound

// UserModel defines the interface for user operations
type UserModel interface {
	Insert(ctx context.Context, data *User) error
	FindOne(ctx context.Context, id int64) (*User, error)
	FindOneByUserAccount(ctx context.Context, userAccount string) (*User, error)
	Update(ctx context.Context, data *User) error
	Delete(ctx context.Context, id int64) error
}

type defaultUserModel struct {
	db *gorm.DB
}

// NewUserModel creates a new UserModel backed by GORM
func NewUserModel(db *gorm.DB) UserModel {
	return &defaultUserModel{
		db: db,
	}
}

func (m *defaultUserModel) Insert(ctx context.Context, data *User) error {
	return m.db.WithContext(ctx).Create(data).Error
}

func (m *defaultUserModel) FindOne(ctx context.Context, id int64) (*User, error) {
	var user User
	err := m.db.WithContext(ctx).Where(&User{ID: id}).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (m *defaultUserModel) FindOneByUserAccount(ctx context.Context, userAccount string) (*User, error) {
	var user User
	err := m.db.WithContext(ctx).Where(&User{UserAccount: userAccount}).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (m *defaultUserModel) Update(ctx context.Context, data *User) error {
	return m.db.WithContext(ctx).Save(data).Error
}

func (m *defaultUserModel) Delete(ctx context.Context, id int64) error {
	return m.db.WithContext(ctx).Delete(&User{}, id).Error
}
