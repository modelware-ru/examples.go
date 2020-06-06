package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

)

func init() {
	var err error
	if DB, err = gorm.Open("mysql", "root:root@tcp(localhost:3307)/cities?charset=utf8&parseTime=True"); err != nil {
		panic(err)
	}

	DB.DB().SetMaxIdleConns(50)
	DB.DB().SetMaxOpenConns(100)
}

var DB *gorm.DB

