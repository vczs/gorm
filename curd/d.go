package curd

import (
	"gorm/model"

	"gorm.io/gorm"
)

func D(db *gorm.DB) {
	db.Delete(&model.UserInfo{}, 100)
	db.Delete(&model.UserInfo{ID: 200})
}
