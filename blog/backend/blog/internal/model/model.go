package model

import "time"

type Model struct {
	Id        int64     `json:"id,omitempty" gorm:"primary_key,column:id"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
	DeletedAt time.Time `json:"deletedAt" gorm:"column:deleted_at"`
}
