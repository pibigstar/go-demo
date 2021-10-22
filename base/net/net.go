package netx

import (
	"net"
	"net/url"
)

// 取url的ip地址
func GetIp(_url string) (string, error) {
	pu, err := url.Parse(_url)
	if err != nil {
		return "", err
	}
	host := pu.Hostname()
	port := pu.Port()
	if port == "" {
		port = "80"
		if pu.Scheme == "https" {
			port = "443"
		}
	}

	addr, err := net.ResolveTCPAddr("tcp", net.JoinHostPort(host, port))
	if err != nil {
		return "", err
	}
	return addr.IP.String(), nil
}
