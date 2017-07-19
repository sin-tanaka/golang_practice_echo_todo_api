package models

import (
	"github.com/jinzhu/gorm"
)

type Task struct {
	// id, created_at, updated_at, deleted_atを自動で付与
	// ついでにsafe deleteになる
	gorm.Model
	Task      string    `json:"task"`
}
