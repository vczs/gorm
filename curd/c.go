package curd

import (
	"gorm/model"

	"gorm.io/gorm"
)

func C(db *gorm.DB) {
	u1 := model.UserInfo{ID: 100, Name: "vcz1", Gender: "男", Account: "v01", Age: 15, Address: "US", Other: "000"}
	u2 := model.UserInfo{ID: 200, Name: "vcz2", Gender: "女", Account: "v02", Age: 18, Address: "CN", Other: "111"}
	db.Create(&u1)
	db.Create(&u2)
}
