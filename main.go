package main

import (
	"fmt"
	"gorm/conn"
	"gorm/model"
	"gorm/utils"
)

func main() {
	// 连接数据库
	db, err := conn.ConnDB()
	if err != nil {
		utils.MyLog("数据库连接失败：", err)
		return
	}
	fmt.Println("数据库连接成功！")

	// 创建数据库表
	if err = db.AutoMigrate(&model.UserInfo{}); err != nil {
		utils.MyLog("数据库表创建失败：", err)
		return
	}
	fmt.Println("数据库表创建成功！")

	// 增
	// curd.C(db)

	// 查
	// curd.R(db)

	// 更
	// curd.U(db)

	// 删
	// curd.D(db)
}
