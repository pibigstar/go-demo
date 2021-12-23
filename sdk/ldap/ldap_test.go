package ldap

import "testing"

func TestLogin(t *testing.T) {
	_, err := LoginBind("hello", "123456")
	if err != nil {
		t.Error(err)
	}
}

func TestAddUser(t *testing.T) {
	// 登录管理员账号
	conn, err := LoginBind("admin", "123456")
	if err != nil {
		t.Error(err)
	}

	username := "xiao ming"
	// 创建新用户
	err = AddUser(conn, username, "123456")
	if err != nil {
		t.Error(err)
	}

	// 验证新增用户登录
	_, err = LoginBind(username, "123456")
	if err != nil {
		t.Error(err)
	}

	// 删除新用户
	err = DeleteUser(conn, username)
	if err != nil {
		t.Error(err)
	}
}

func TestListUser(t *testing.T) {
	// 登录管理员账号
	conn, err := LoginBind("admin", "pibigstar")
	if err != nil {
		t.Error(err)
	}

	users, err := ListUser(conn)
	if err != nil {
		t.Error(err)
	}
	t.Log(users)
}
