package models

import "time"

type User struct {
	Username  string    `gorm:"type:varchar(50); primaryKey" json:"username"`
	Email     string    `gorm:"type:varchar(100)" json:"email"`
	Name      string    `gorm:"type:varchar(100)" json:"name"`
	Password  string    `gorm:"type:varchar(50)" json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Answer    string    `gorm:"type:varchar(50)" json:"answer"`
}