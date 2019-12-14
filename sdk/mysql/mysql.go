package mysql

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

var (
	user     = "root"
	password = ""
	host     = "127.0.0.1"
	port     = 3306
	db       = "test"
)

func init() {
	// for ci success
	name, _ := os.Hostname()
	if name == "pibigstar" {
		password = "123456"
	}
}

//使用gorm链接mysql
func GetGormDB() (*gorm.DB, error) {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local&charset=utf8mb4,utf8",
		user, password, host, port, db)
	return gorm.Open("mysql", dns)
}

// 原生mysql链接
func GetMySQLDB() (*sql.DB, error) {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local&charset=utf8mb4,utf8",
		user, password, host, port, db)
	return sql.Open("mysql", dns)
}
