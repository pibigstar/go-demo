package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go-demo/utils/env"
	"testing"
)

func TestMysql(t *testing.T) {
	if env.IsCI() {
		return
	}
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/oa?charset=utf8")
	if err != nil {
		t.Error(err)
	}

	rows, err := db.Query("select id,username from user where id = ?", "1")
	if err != nil {
		t.Error(err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var username string
		err := rows.Scan(&id, &username)
		if err != nil {
			t.Error(err)
		}
		t.Log(id, username)
	}
}
