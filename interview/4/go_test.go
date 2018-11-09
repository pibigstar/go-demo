package interview

import (
	"fmt"
	"testing"
)

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func TestGo(t *testing.T)  {
	teacher := Teacher{}
	teacher.ShowA()
}
/**
showA
showB
 */

 /**
 Go中没有继承！ 没有继承！没有继承！是叫组合！组合！组合！
这里People是匿名组合People。被组合的类型People所包含的方法虽然升级成了外部类型Teacher这个组合类型的方法，
 但他们的方法(ShowA())调用时接受者并没有发生变化。这里仍然是People。
 毕竟这个People类型并不知道自己会被什么类型组合，当然也就无法调用方法时去使用未知的组合者Teacher类型的功能。
因此这里执行t.ShowA()时，在执行ShowB()时该函数的接受者是People，而非Teacher。
  */