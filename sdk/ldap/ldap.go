package ldap

import (
	"fmt"
	"github.com/go-ldap/ldap"
	"strings"
)

const (
	ldapURL      = "ldap://127.0.0.1:389"
	organisation = "dc=pibigstar,dc=com"
)

// 登录
func LoginBind(ldapUser, ldapPassword string) (*ldap.Conn, error) {
	l, err := ldap.DialURL(ldapURL)
	if err != nil {
		return nil, err
	}
	_, err = l.SimpleBind(&ldap.SimpleBindRequest{
		Username: fmt.Sprintf("cn=%s,%s", ldapUser, organisation),
		Password: ldapPassword,
	})

	if err != nil {
		fmt.Println("ldap password is error: ", ldap.LDAPResultInvalidCredentials)
		return nil, err
	}
	fmt.Println(ldapUser, "登录成功")
	return l, nil
}

// 必须管理员才能添加
func AddUser(conn *ldap.Conn, userName string, password string) error {
	addRequest := ldap.NewAddRequest(fmt.Sprintf("cn=%s,%s", userName, organisation), nil)
	addRequest.Attribute("userPassword", []string{password})
	addRequest.Attribute("homeDirectory", []string{fmt.Sprintf("/home/%s", userName)})
	addRequest.Attribute("cn", []string{userName})
	addRequest.Attribute("uid", []string{userName})
	addRequest.Attribute("objectClass", []string{"shadowAccount", "posixAccount", "account"})
	addRequest.Attribute("uidNumber", []string{"110"})
	addRequest.Attribute("gidNumber", []string{"110"})
	addRequest.Attribute("loginShell", []string{"/bin/bash"})
	if err := conn.Add(addRequest); err != nil {
		return err
	}
	return nil
}

func DeleteUser(conn *ldap.Conn, username string) error {
	req := ldap.NewDelRequest(fmt.Sprintf("cn=%s,%s", username, organisation), nil)
	if err := conn.Del(req); err != nil {
		return err
	}
	return nil
}

func ListUser(conn *ldap.Conn) ([]string, error) {
	var usernames []string
	req := ldap.NewSearchRequest(organisation,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		"(objectClass=*)",
		[]string{"dn", "com", "objectClass"},
		nil)

	cur, err := conn.Search(req)
	if err != nil {
		return nil, err
	}

	if len(cur.Entries) > 0 {
		for _, item := range cur.Entries {
			cn := strings.ReplaceAll(item.DN, fmt.Sprintf(",%s", organisation), "")
			if ss := strings.Split(cn, "="); len(ss) == 2 {
				usernames = append(usernames, ss[1])
			}
		}
		return usernames, nil
	}
	return nil, nil
}
