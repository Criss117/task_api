package task_entity

import "gorm.io/gorm"

type Task struct {
	gorm.Model

	Title 			string `gorm:"type:varchar(100);not null" json:"title"`
	Description string `gorm:"type:varchar(255);not null" json:"description"`
	Done 				bool `gorm:"default:false" json:"done"`
	UserID 			uint `gorm:"not null" json:"user_id"`
}