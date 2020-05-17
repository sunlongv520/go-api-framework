package AppInit

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB
func init() {
	var err error
	db, err = gorm.Open("mysql", "root:root@tcp(192.168.1.101:3306)/test?charset=utf8&parseTime=True&loc=Local")
	db.SingularTable(true)
	if err != nil {
		//fmt.Println(err)
		log.Fatal(err)
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(50)
}
func  GetDB() *gorm.DB {
	return db
}
