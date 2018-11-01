package simple

import "fmt"

// 简单工厂模式

type Girl interface {
	weight()
}

// 胖女孩
type FatGirl struct {
}
func (FatGirl) weight() {
	fmt.Println("80kg")
}

// 瘦女孩
type ThinGirl struct {
}

func (ThinGirl) weight() {
	fmt.Println("50kg")
}


type GirlFactory struct {

}

func (*GirlFactory) CreateGirl(like string) Girl {
	if like == "fat" {
		return &FatGirl{}
	} else if like == "thin" {
		return &ThinGirl{}
	}
	return nil
}
