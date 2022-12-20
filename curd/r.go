package curd

import (
	"fmt"
	"gorm/model"

	"gorm.io/gorm"
)

func R(db *gorm.DB) {
	u3 := &model.UserInfo{}
	db.First(&u3, 100)
	u4 := &model.UserInfo{ID: 200}
	db.First(&u4)
	fmt.Printf("数据查询成功：\n%v\n%v", u3, u4)
}
