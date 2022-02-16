package model

import "time"

type User struct {
	BaseColumn
	InstanceID   string    `gorm:"column:instanceID" json:"instanceID" validate:"omitempty"`
	Name         string    `gorm:"column:name;NOT NULL" json:"name" validate:"omitempty"`
	Status       int       `gorm:"column:status;default:1" json:"status" validate:"omitempty"` // 1:可用，0:不可用
	Nickname     string    `gorm:"column:nickname;NOT NULL" json:"nickname" validate:"required,min=1,max=30"`
	Password     string    `gorm:"column:password;NOT NULL" json:"password" validate:"required"`
	Email        string    `gorm:"column:email;NOT NULL" json:"email" validate:"required,email,min=1,max=100"`
	Phone        string    `gorm:"column:phone" json:"phone" validate:"omitempty"`
	IsAdmin      int       `gorm:"column:isAdmin;default:0;NOT NULL" json:"isAdmin" validate:"omitempty"` // 1: administrator\\n0: non-administrator
	ExtendShadow string    `gorm:"column:extendShadow" json:"extendShadow" validate:"omitempty"`
	LoginedAt    time.Time `gorm:"column:loginedAt" json:"loginedAt"` // last login time

}

// UserList is the whole list of all users which have been stored in stroage.
type UserList struct {

	TotalCount int64 `json:"total_count,omitempty"`

	Items []*User `json:"items"`
}

// TableName maps to mysql table name.
func (u *User) TableName() string {
	return "user"
}
