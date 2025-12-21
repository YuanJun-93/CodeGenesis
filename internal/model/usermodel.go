package model

import (
	"context"

	"gorm.io/gorm"
)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	customUserModel struct {
		db *gorm.DB
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(db *gorm.DB) UserModel {
	return &customUserModel{
		db: db,
	}
}

func (m *customUserModel) Insert(ctx context.Context, data *User) error {
	return m.db.WithContext(ctx).Create(data).Error
}

func (m *customUserModel) FindOne(ctx context.Context, id int64) (*User, error) {
	var user User
	err := m.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (m *customUserModel) FindOneByUserAccount(ctx context.Context, userAccount string) (*User, error) {
	var user User
	err := m.db.WithContext(ctx).Where("userAccount = ?", userAccount).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (m *customUserModel) Update(ctx context.Context, data *User) error {
	return m.db.WithContext(ctx).Save(data).Error
}

func (m *customUserModel) Delete(ctx context.Context, id int64) error {
	return m.db.WithContext(ctx).Delete(&User{}, id).Error
}

// TableName matches the sqlx table definition
func (User) TableName() string {
	return "user"
}
