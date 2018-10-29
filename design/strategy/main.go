package strategy

// 策略模式


// 实现此接口，则为一个策略
type IStragegy interface {
	do(int, int) int
}

// 加
type add struct {

}

func (*add) do(a, b int) int  {
	return a + b
}

// 减

type reduce struct {

}

func (*reduce) do(a, b int) int {
	return a - b
}

type Operater struct {
	strategy IStragegy
}

func (operater *Operater) setStrategy(strategy IStragegy)  {
	operater.strategy = strategy
}












