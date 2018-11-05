package template

import "testing"

func TestTemplate(t *testing.T)  {

	// 做西红柿
	xihongshi := &XiHongShi{}
	doCook(xihongshi)

	// 做炒鸡蛋
	chaojidan := &ChaoJiDan{}
	doCook(chaojidan)

}
