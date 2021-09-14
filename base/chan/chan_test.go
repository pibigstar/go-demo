package chanx_test

import (
	chanx "go-demo/base/chan"
	"math/rand"
	"testing"
	"time"
)

func TestChanOrOne(t *testing.T) {
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})

	go func() {
		time.AfterFunc(100*time.Millisecond, func() {
			ch1 <- 1
		})
	}()

	go func() {
		time.AfterFunc(200*time.Millisecond, func() {
			ch2 <- 2
		})
	}()

	<-chanx.Or(ch1, ch2)
	t.Log("收到了信号,开始执行业务逻辑")
}

func TestFanIn(t *testing.T) {
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})

	go func() {
		for i := 0; i < 100; i++ {
			if i%2 != 0 {
				ch1 <- i
			} else {
				ch2 <- i
			}
		}
	}()

	out := chanx.FanIn(ch1, ch2)

	go func() {
		for {
			select {
			case i := <-out:
				t.Log("out: ", i)
			}
		}
	}()

	time.Sleep(1 * time.Second)
}

func TestFanInByReflect(t *testing.T) {
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})

	go func() {
		for i := 0; i < 100; i++ {
			if i%2 != 0 {
				ch1 <- i
			} else {
				ch2 <- i
			}
		}
	}()

	go func() {
		for {
			select {
			case i := <-chanx.FanInByReflect(ch1, ch2):
				t.Log("out: ", i)
			}
		}
	}()

	time.Sleep(1 * time.Second)
}

func TestFanOut(t *testing.T) {
	ch1 := make(chan interface{})
	chs := []chan interface{}{
		make(chan interface{}),
		make(chan interface{}),
		make(chan interface{}),
	}
	// 将ch1收到值，扇出到chs每个chan中
	chanx.FanOut(ch1, chs)

	go func() {
		t := time.NewTicker(100 * time.Millisecond)
		for {
			select {
			case <-t.C:
				ch1 <- rand.Intn(100)
			}
		}
	}()

	go func() {
		for {
			for _, c := range chs {
				t.Log(<-c)
			}
		}
	}()

	time.Sleep(1 * time.Second)
}

func TestFanOutByReflect(t *testing.T) {
	ch1 := make(chan interface{})
	chs := []chan interface{}{
		make(chan interface{}),
		make(chan interface{}),
		make(chan interface{}),
	}
	// 将ch1收到值，扇出到chs每个chan中
	chanx.FanOutByReflect(ch1, chs)

	go func() {
		t := time.NewTicker(100 * time.Millisecond)
		for {
			select {
			case <-t.C:
				ch1 <- rand.Intn(100)
			}
		}
	}()

	go func() {
		for {
			for _, c := range chs {
				if v, ok := <-c; ok {
					t.Log(v)
				}
			}
		}
	}()

	time.Sleep(1 * time.Second)
}

func TestFanOutRandom(t *testing.T) {
	ch1 := make(chan interface{})
	chs := []chan interface{}{
		make(chan interface{}, 10),
		make(chan interface{}, 10),
		make(chan interface{}, 10),
	}
	rand.Seed(time.Now().UnixNano())

	// 将ch1收到值，扇出到chan中任意一个
	chanx.FanOutRandom(ch1, chs)

	go func() {
		t := time.NewTicker(100 * time.Millisecond)
		for {
			select {
			case <-t.C:
				ch1 <- rand.Intn(100)
			}
		}
	}()

	go func() {
		for {
			for i, c := range chs {
				if v, ok := <-c; ok {
					t.Log("i:", i, "v:", v)
				}
			}
		}
	}()

	time.Sleep(1 * time.Second)
}

func TestFanOutRandomByReflect(t *testing.T) {
	ch1 := make(chan interface{})
	chs := []chan interface{}{
		make(chan interface{}),
		make(chan interface{}),
		make(chan interface{}),
	}
	// 将ch1收到值，扇出到chan中任意一个
	chanx.FanOutRandomByReflect(ch1, chs)

	go func() {
		t := time.NewTicker(100 * time.Millisecond)
		for {
			select {
			case <-t.C:
				ch1 <- rand.Intn(100)
			}
		}
	}()

	go func() {
		for {
			for i, c := range chs {
				if v, ok := <-c; ok {
					t.Log("i:", i, "v:", v)
				}
			}
		}
	}()

	time.Sleep(1 * time.Second)
}
