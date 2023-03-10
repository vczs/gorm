package model

import "gorm.io/gorm"

type UserInfo struct {
	ID         uint   `gorm:"primary_key;auto_increment"`          // 主键 自增
	Account    string `gorm:"unique;not null;comment:账号"`          // 唯一 不为空 注释
	Name       string `gorm:"type:varchar(200)"`                   // 指定类型
	Gender     string `gorm:"type:varchar(100);size:10;default:男"` // 指定类型 设置字段大小 默认值
	Age        int    `gorm:"column:user_age"`                     // 指定列名
	Address    string `gorm:"index:addr;comment:地址"`               // 创建名为addr的索引 注释
	Other      string `gorm:"-"`                                   // 忽略该字段(不在数据库表中关联)
	gorm.Model        // gorm内置model 包含ID、创建时间、更新时间、删除时间
}

// 为UserInfo结构体对应数据库表设置表名
func (UserInfo) TableName() string {
	return "db_user_info"
}
