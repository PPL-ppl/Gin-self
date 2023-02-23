package model

import (
	"gorm.io/gorm"
	"time"
)

// Model gorm.Model 的定义
type Model struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"index:,sort:desc,collate:utf8,type:btree,length:10,where:name3 != 'jinzhu'"`
	UpdatedAt time.Time      `gorm:"index:aaa"`
	DeletedAt gorm.DeletedAt `gorm:"index:aaa"`
}
