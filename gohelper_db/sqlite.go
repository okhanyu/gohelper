package gohelper_db

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB //后续用于操作数据库

func InitDB(dsn string) *gorm.DB {
	gdb, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	DB = gdb
	return gdb
}
