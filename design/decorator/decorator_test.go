package decorator

import (
	"testing"
	"fmt"
)


func TestDecorator(t *testing.T)  {

	laowang := &laowang{}

	jacket := &Jacket{}
	jacket.person = laowang
	jacket.show()

	hat := &Hat{}
	hat.person = jacket
	hat.show()

	fmt.Println("cost:", hat.cost())

}