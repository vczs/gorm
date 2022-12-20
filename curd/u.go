package curd

import (
	"gorm/model"

	"gorm.io/gorm"
)

func U(db *gorm.DB) {
	db.Model(&model.UserInfo{ID: 100}).Update("name", "vcz10")
	db.Updates(&model.UserInfo{ID: 200, Name: "vcz20", Gender: "男"})
	db.Save(&model.UserInfo{ID: 300, Name: "vcz30", Gender: "男", Account: "v03", Age: 20, Address: "US", Other: "222"}) // 若数据不包含匹配的主键 则插入数据
}
