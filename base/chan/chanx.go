package chanx

import (
	"math/rand"
	"reflect"
	"sync"
)

// 当等待多个信号的时候，如果收到任意一个信号， 就执行业务逻辑，忽略其它的还未收到的信号
func Or(chs ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		var once sync.Once
		for _, c := range chs {
			go func(c <-chan interface{}) {
				select {
				case <-c:
					once.Do(func() {
						close(out) // 关闭out,提醒外部可以继续执行了
					})
				case <-out:
				}
			}(c)
		}
	}()
	return out
}

// 通过select方式，实现起来更简洁
func OrBySelect(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}
	orDone := make(chan interface{})
	go func() {
		defer close(orDone)
		var cases []reflect.SelectCase
		for _, c := range channels {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			})
		}
		reflect.Select(cases)
	}()
	return orDone
}

// 将多个同样类型的输入channel合并成一个同样类型的输出channel
func FanIn(chs ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		var wg sync.WaitGroup
		wg.Add(len(chs))
		for _, c := range chs {
			go func(c <-chan interface{}) {
				for v := range c {
					out <- v
				}
				wg.Done()
			}(c)
		}
		wg.Wait()
		close(out)
	}()
	return out
}

// 将多个同样类型的输入channel合并成一个同样类型的输出channel
func FanInByReflect(chs ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)

		var cases []reflect.SelectCase
		for _, c := range chs {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			})
		}
		if len(cases) == 0 {
			return
		}
		// 执行 select，也就是从 chan 中接受值
		i, v, ok := reflect.Select(cases)
		if !ok {
			cases = append(cases[:i], cases[i+1:]...)
		}
		out <- int(v.Int())
	}()

	return out
}

// 将一个输入channel扇出为多个channel
func FanOut(ch <-chan interface{}, out []chan interface{}) {
	go func() {
		defer func() {
			for i := 0; i < len(out); i++ {
				close(out[i])
			}
		}()

		for v := range ch {
			v := v
			for i := 0; i < len(out); i++ {
				out[i] <- v
			}
		}
	}()
}

// 将一个输入channel扇出为多个channel
func FanOutByReflect(ch <-chan interface{}, out []chan interface{}) {
	go func() {
		defer func() {
			for i := 0; i < len(out); i++ {
				close(out[i])
			}
		}()

		cases := make([]reflect.SelectCase, len(out))

		for i := range cases {
			cases[i].Dir = reflect.SelectSend
			cases[i].Chan = reflect.ValueOf(out[i])
		}

		// 阻塞一直到ch被close
		for v := range ch {
			v := v
			for i := range cases {
				cases[i].Send = reflect.ValueOf(v)
			}
			for range cases {
				// 执行 select，也就是将 v 发送到 所有的cases中
				reflect.Select(cases)
			}
		}
	}()
}

// 将一个输入channel扇出到out chan中任意一个
func FanOutRandom(ch <-chan interface{}, out []chan interface{}) {
	go func() {
		defer func() {
			for i := 0; i < len(out); i++ {
				close(out[i])
			}
		}()

		var n = len(out)
		for v := range ch {
			v := v
			i := rand.Intn(n)
			out[i] <- v
		}
	}()
}

// 将一个输入channel扇出到out chan中任意一个
func FanOutRandomByReflect(ch <-chan interface{}, out []chan interface{}) {
	go func() {
		defer func() {
			for i := 0; i < len(out); i++ {
				close(out[i])
			}
		}()

		cases := make([]reflect.SelectCase, len(out))

		for i := range cases {
			cases[i].Dir = reflect.SelectSend
			cases[i].Chan = reflect.ValueOf(out[i])
		}

		// 阻塞一直到ch被close
		for v := range ch {
			v := v
			for i := range cases {
				cases[i].Send = reflect.ValueOf(v)
			}
			// 执行 select，也就是将 v 发送到 所有的cases中
			reflect.Select(cases)
		}
	}()
}
