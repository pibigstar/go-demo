package ip

import (
	"fmt"
	"go-demo/utils/ip/address"
	"testing"
)

func Test_Address(t *testing.T) {

	ip := GetInternetIP()
	address, _ := address.GetAddressByIp(ip)
	fmt.Printf("%+v", address)
}
