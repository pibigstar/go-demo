package ip

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"
)

// 获取公网地址
func GetInternetIP() string {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	return strings.TrimSpace(string(content))
}

// 获取本地IP地址
func GetLocalIP() string {
	address, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, address := range address {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return strings.TrimSpace(ipnet.IP.String())
			}
		}
	}
	return ""
}
