package usermethods

import (
	"errors"
	"fmt"
	"template-app/models/user"
	"template-app/pkg/store"

	"gorm.io/gorm"
)

// Add adds user with given Email id to database
func Add(emailId string) error {
	db := store.DB
	newUser := user.User{
		EmailId: emailId,
	}
	err := db.Model(&newUser).Create(&newUser).Error
	if err != nil {
		return nil
	}
	return err
}

func AddIfNotExist(emailId string) error {
	_, err := Get(emailId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = Add(emailId)
	}

	return err
}

// Get returns user with given Email id from database
func Get(emailId string) (*user.User, error) {
	db := store.DB
	var _user user.User
	res := db.First(&_user, user.User{
		EmailId: emailId,
	})
	if err := res.Error; err != nil {
		err = fmt.Errorf("failed to get user from database: %w", err)
		return nil, err
	}

	return &_user, nil
}
