package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm/utils"
)

const (
	Host     = "127.0.0.1"
	Port     = 3306
	User     = "root"
	PassWord = "123456"
	DbName   = "vczsdb"
)

type UserInfo struct {
	ID      uint   `gorm:"primary_key;auto_increment"`          // 主键 自增
	Account string `gorm:"unique;not null"`                     // 唯一 不为空
	Name    string `gorm:"type:varchar(200)"`                   // 指定类型
	Gender  string `gorm:"type:varchar(100);size:10;default:男"` // 指定类型 设置字段大小 默认值
	Age     int    `gorm:"column:user_age"`                     // 零值类型 不为空
	Address string `gorm:"index:addr"`                          // 创建名为addr的索引
	Other   string `gorm:"-"`                                   // 忽略该字段(不在数据库表中关联)
	gorm.Model
}

// 为UserInfo结构体对应数据库表设置表名
func (UserInfo) TableName() string {
	return "db_user_info"
}

func main() {
	// 配置参数
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", User, PassWord, Host, Port, DbName)
	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		utils.MyLog("数据库连接失败：", err)
		return
	}
	fmt.Println("数据库连接成功！")

	// 创建表(自动迁移)：把数据库表和结构体进行对应
	err = db.AutoMigrate(&UserInfo{})
	if err != nil {
		utils.MyLog("数据库表创建失败：", err)
		return
	}
	fmt.Println("数据库表创建成功！")

	// 增
	u1 := UserInfo{ID: 100, Name: "vcz1", Gender: "男", Account: "v01", Age: 15, Address: "US", Other: "000"}
	u2 := UserInfo{ID: 200, Name: "vcz2", Gender: "女", Account: "v02", Age: 18, Address: "CN", Other: "111"}
	db.Create(&u1)
	db.Create(&u2)

	// 查
	u3 := UserInfo{}
	db.First(&u3, 100)
	u4 := UserInfo{ID: 200}
	db.First(&u4)
	fmt.Printf("数据查询成功：\n%v\n%v", u3, u4)

	// 更
	db.Model(&UserInfo{ID: 100}).Update("name", "vcz10")
	db.Updates(&UserInfo{ID: 200, Name: "vcz20", Gender: "男"})
	db.Save(&UserInfo{ID: 300, Name: "vcz30", Gender: "男", Account: "v03", Age: 20, Address: "US", Other: "222"}) // 若数据不包含匹配的主键 则插入数据

	// 删
	db.Delete(&UserInfo{}, 100)
	db.Delete(&UserInfo{ID: 200})
}
