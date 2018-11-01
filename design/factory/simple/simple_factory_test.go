package simple

import "testing"

func TestSimple(t *testing.T)  {

	// 创建工厂
	girlFactory := new(GirlFactory)

	// 传递你喜欢类型的姑娘即可
	girl := girlFactory.CreateGirl("fat")
	girl.weight()

	girl = girlFactory.CreateGirl("thin")
	girl.weight()
}
