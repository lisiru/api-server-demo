package model

import "time"

type BaseColumn struct {
	ID        uint64    `json:"id,omitempty" gorm:"primary_key;AUTO_INCRMENT;column:id"`
	CreatedAt time.Time `json:"createdAt,omitempty" gorm:"column:createdAt"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" gorm:"column:updatedAt"`
}
