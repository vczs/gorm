package main

import (
	"fmt"
	"gorm/conn"
	"gorm/model"
	"gorm/utils"

	"gorm.io/gorm"
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
	if err = createTable(db, model.UserInfo{}); err != nil {
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

// 创建数据库表
func createTable(db *gorm.DB, table interface{}) error {
	// 创建表(自动迁移)：把数据库表和结构体进行对应
	err := db.AutoMigrate(table)
	if err != nil {
		return err
	}
	return nil
}
