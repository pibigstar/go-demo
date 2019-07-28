package mysql

import (
	"testing"
)

func TestConnectMysql(t *testing.T) {
	db, err := GetGormDB()
	defer db.Close()
	if err != nil {
		t.Error(err)
	}
	err = db.Exec("SELECT CURRENT_TIMESTAMP();").Error
	if err != nil {
		t.Error(err)
	}
}

func TestMySQLDB(t *testing.T) {
	db, err := GetMySQLDB()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	_, err = db.Exec("SELECT CURRENT_TIMESTAMP();")
	if err != nil {
		t.Error(err)
	}
}
