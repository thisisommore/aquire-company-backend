// Package user provides model and methods for storing and retriving user identified by Email id
package user

// User user model with Email id and one to many relation with Role
type User struct {
	EmailId        string            `json:"-" gorm:"primaryKey;not null"`
}
