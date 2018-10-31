package observer

import "testing"

func TestObserver(t *testing.T)  {

	customerA := &CustomerA{}
	customerB := &CustomerB{}

	office := &NewsOffice{}
	// 模拟客户订阅
	office.addCustomer(customerA)
	office.addCustomer(customerB)
	// 新的报纸
	office.newspaperCome()

}
