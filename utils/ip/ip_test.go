package ip

import (
	"go-demo/utils/ip/address"
	"testing"
)

func Test_Address(t *testing.T) {

	ip := GetInternetIP()
	address, _ := address.GetAddressByIP(ip)
	t.Logf("%+v", address)
}
