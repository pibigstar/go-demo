package mock

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetGoVersion(t *testing.T) {

	// 1. 新建一个mock控制器
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	// 2. 调用mock生成代码里面为我们实现的接口对象
	mockSpider := NewMockSpider(mockCtl)
	// 3. 调用EXPECT()得到实现的对象 并调用对象的方法 指定其返回值
	mockSpider.EXPECT().GetBody().Return("go1.8.3")
	// 4. 再实现Init方法，现在mockSpider 就是一个实现了Spider接口的对象
	mockSpider.EXPECT().Init()

	// 5. 将mock出的Spider接口对象传递过去 (GetGoVersion 必须要使用到Init和GetBody方法)
	goVer := GetGoVersion(mockSpider)

	fmt.Println(goVer)
}
