package models

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Phone    string    `json:"phone"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
	RoleId   string    `json:"role_id"`
	Role     *Role     `gorm:"foreignKey:RoleId" json:"role,omitempty`
	Name     string    `json:"name"`
	Customer *Customer `gorm:"foreignKey:UserId" json:"customer,omitempty"`
	BaseModel
}

// func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
// 	u.Id = ulid.Make().String()
// 	return
// }

func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	fmt.Print(user.Password)
	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)
	}
	return nil
}
