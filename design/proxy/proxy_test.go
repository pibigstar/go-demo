package proxy

import "testing"

func TestProxy(t *testing.T) {

	station := &Station{3}
	proxy := &StationProxy{station}
	station.sell("小华")
	proxy.sell("派大星")
	proxy.sell("小明")
	proxy.sell("小兰")
}
