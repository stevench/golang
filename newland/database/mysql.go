package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var SqlDB *gorm.DB

func init() {
	var err error
	SqlDB, err = gorm.Open("mysql", "root:steven@tcp(127.0.0.1:3306)/test?parseTime=true")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
