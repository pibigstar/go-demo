package govaluate

import (
	"gopkg.in/Knetic/govaluate.v3"
	"testing"
)

// 值比较
func TestComparison(t *testing.T) {
	expression, err := govaluate.NewEvaluableExpression("foo > 0")
	if err != nil {
		t.Error(err)
	}
	params := make(map[string]interface{})
	params["foo"] = 1

	result, err := expression.Evaluate(params)
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

// 值计算
func TestCalculate(t *testing.T) {
	expression, err := govaluate.NewEvaluableExpression("salary / days")
	if err != nil {
		t.Error(err)
	}
	params := make(map[string]interface{})
	params["salary"] = 12000
	params["days"] = 30

	result, err := expression.Evaluate(params)
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}
