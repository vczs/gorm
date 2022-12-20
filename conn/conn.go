package conn

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	Host     = "127.0.0.1"
	Port     = 3306
	User     = "root"
	PassWord = "123456"
	DbName   = "vczsdb"
)

func ConnDB() (*gorm.DB, error) {
	// 配置参数
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", User, PassWord, Host, Port, DbName)
	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
