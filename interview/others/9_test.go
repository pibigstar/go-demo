package others

import (
	"fmt"
	"testing"
)

/**
考题中的 func (stu *Student) Speak(think string) (talk string) 是表示结构类型*Student的指针有提供该方法
但该方法并不属于结构类型Student的方法。因为struct是值类型。

修改方法：
定义为指针 go var peo People = &Student{}
方法定义在值类型上,指针类型本身是包含值类型的方法
go func (stu Student) Speak(think string) (talk string) { //... }

解析: 类型T方法集包含所有 receiver T 方法, 类型*T方法集包含所有 receiver T 和 receiver *T 方法
*/
type People interface {
	Speak(string) string
}

func (stu *Student) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func TestInterface(t *testing.T) {
	// var peo People = Student{} 不可以这样写, 因为是*Student实现了People，不是Student
	var peo People = &Student{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
