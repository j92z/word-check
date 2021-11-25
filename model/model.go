package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"sensitive_words_check/config"
)

var Db *gorm.DB

type Model struct {
	ID        int `gorm:"primary_key" json:"id"`
	CreatedOn int `json:"created_on"`
	UpdatedOn int `json:"updated_on"`
	DeletedOn int `json:"deleted_on"`
}

func Setup() {
	var err error
	Db, err = gorm.Open(config.DbConfig.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=True&loc=Local",
		config.DbConfig.User,
		config.DbConfig.Password,
		config.DbConfig.Host,
		config.DbConfig.DbName))
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	Db.LogMode(config.DbConfig.ShowDbLog) //调试开发模式
	Db.SingularTable(true)                            //操作单表
	Db.DB().SetMaxIdleConns(config.DbConfig.MaxIdleConns)
	Db.DB().SetMaxOpenConns(config.DbConfig.MaxOpenConns)
}

func CloseDB() {
	defer Db.Close()
}

func NotDeleted(db *gorm.DB) *gorm.DB {
	return Db.Where("deleted_on = ?", 0)
}
