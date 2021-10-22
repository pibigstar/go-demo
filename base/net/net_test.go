package netx_test

import (
	netx "go-demo/base/net"
	"testing"
)

func TestGetIp(t *testing.T) {
	ip, err := netx.GetIp("http://www.baidu.com")
	if err != nil {
		t.Error(err)
	}
	t.Log(ip)
}
