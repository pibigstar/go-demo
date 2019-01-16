package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)
/**
 * 原生mysql链接
 */
func main() {
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