package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey"`
	Name      string         `gorm:"size:100;not null"`
	Email     string         `gorm:"size:100;uniqueIndex;not null"`
	Password  string         `gorm:"size:255;not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// BeforeCreate is a GORM hook that runs before a new record is created
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// You can add code here to hash the password, generate a UUID, etc.
	return
}

// TableName overrides the table name used by GORM for this model
func (User) TableName() string {
	return "users"
}

// CreateUser creates a new user in the database
func CreateUser(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}

// GetUserByID retrieves a user by their ID
func GetUserByID(db *gorm.DB, id uint) (*User, error) {
	var user User
	if err := db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByEmail retrieves a user by their email
func GetUserByEmail(db *gorm.DB, email string) (*User, error) {
	var user User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser updates a user's information in the database
func UpdateUser(db *gorm.DB, user *User) error {
	return db.Save(user).Error
}

// DeleteUser soft deletes a user from the database
func DeleteUser(db *gorm.DB, id uint) error {
	return db.Delete(&User{}, id).Error
}
