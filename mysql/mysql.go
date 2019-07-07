package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	user     = "root"
	password = "123456"
	host     = "192.168.56.1"
	port     = 3306
	db       = "test"
)

/**
 * 使用gorm链接mysql
 */
func GormDB() {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local&charset=utf8mb4,utf8", user, password, host, port, db)
	db, err := gorm.Open("mysql", dns)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
}

/**
 * 原生mysql链接
 */
func MySQLDB() {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local&charset=utf8mb4,utf8", user, password, host, port, db)
	conn, err := sql.Open("mysql", dns)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	fmt.Println("connect mysql success! ")
	test := conn.Ping()
	fmt.Println(test)
}
