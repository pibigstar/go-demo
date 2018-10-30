package strategy

import (
	"testing"
	"fmt"
)

func TestStrategy(t *testing.T)  {
	operator := Operator{}

	operator.setStrategy(&add{})
	result := operator.calculate(1, 2)
	fmt.Println("add:",result)

	operator.setStrategy(&reduce{})
	result = operator.calculate(2,1)
	fmt.Println("reduce:",result)
}
