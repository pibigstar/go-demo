package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
var (
	user = "root"
	password = "123456"
	host = "192.168.56.1"
	port = 3306
	db = "test"
)
/**
 * 使用gorm链接mysql
 */
func main() {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local&charset=utf8mb4,utf8", user, password, host, port, db)
	db, err := gorm.Open("mysql", dns)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
}
