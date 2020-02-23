package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique" default:"张三"`
	Password string `minLength:"4" maxLength:"16" default:"12345678"`
	Name     string //姓名
	Email    string //邮箱
	Mobile   string //手机
	QQ       string
	Gender   int    //0男 1女
	Age      int    `minimum:"10" maximum:"20" default:"15"` //年龄
	Remark   string //备注
	Token    string `gorm:"-"`
	Session  string `gorm:"-"`
}
