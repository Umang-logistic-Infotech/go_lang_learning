package models

import (
	"github.com/goravel/framework/database/orm"
)

type Users struct {
	orm.Model
	Name     string `gorm:"not null"`
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	Age      int    `gorm:"default:0"`
	City     string `gorm:"type:varchar(100)"`
	IsActive bool   `gorm:"default:true"`
	// orm.SoftDeletes
}
