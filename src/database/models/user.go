package models

import "time"

type User struct {
	UserId    uint      `gorm:"primaryKey" json:"id,omitempty"`
	FullName  string    `json:"full_name,omitempty"`
	UserName  string    `json:"user_name,omitempty"`
	Email     string    `json:"email,omitempty" validate:"required"`
	Role      string    `json:"role,omitempty"`
	Password  string    `json:"Password,omitempty" validate:"required"`
	Address   string    `json:"address,omitempty"`
	Phone     string    `json:"phone,omitempty"`
	CreatedAt time.Time `gorm:"default:now(); not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:now(); not null" json:"updated_at"`
}

type Users []User
