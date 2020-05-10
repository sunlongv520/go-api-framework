package AppInit

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *gorm.DB
func init() {
	var err error
	db, err = gorm.Open("mysql", "root:123456@tcp(192.168.1.188:3307)/test?charset=utf8mb4&parseTime=True&loc=Local")
	db.SingularTable(true)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(50)
}
func  GetDB() *gorm.DB {
	return db
}
