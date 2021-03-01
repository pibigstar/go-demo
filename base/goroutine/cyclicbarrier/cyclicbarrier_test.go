package cyclicbarrier

import (
	"context"
	"fmt"
	"github.com/marusama/cyclicbarrier"
	"golang.org/x/sync/semaphore"
	"sort"
	"sync"
	"testing"
)

// 循环栅栏，允许一组 goroutine 彼此等待，到达一个共同的执行点
// 具体的机制是，大家都在栅栏前等待，等全部都到齐了，就抬起栅栏放行。

// 有一个名叫大自然的搬运工的工厂，
// 生产一种叫做一氧化二氢的神秘液体。这种液体的分子是由一个氧原子和两个氢原子组成的，也就是水。
// 这个工厂有多条生产线，每条生产线负责生产氧原子或者是氢原子，每条生产线由一个 goroutine 负责。
// 这些生产线会通过一个栅栏，只有一个氧原子生产线和两个氢原子生产线都准备好，才能生成出一个水分子，
// 否则所有的生产线都会处于等待状态。也就是说，一个水分子必须由三个不同的生产线提供原子，
// 而且水分子是一个一个按照顺序产生的，每生产一个水分子，
// 就会打印出 HHO，其他形式如：HHH、OOH、OHO、HOO、OOO 都是不允许的。
func TestCyclicBarrier(t *testing.T) {
	//用来存放水分子结果的channel
	var ch chan string
	releaseHydrogen := func() {
		ch <- "H"
	}
	releaseOxygen := func() {
		ch <- "O"
	}

	// 30个原子，30个goroutine,每个goroutine并发的产生一个原子
	var N = 10
	ch = make(chan string, N*3)

	h2o := New()

	// 用来等待所有的goroutine完成
	var wg sync.WaitGroup
	wg.Add(N * 3)

	// 20个氢原子goroutine
	for i := 0; i < 2*N; i++ {
		go func() {
			h2o.hydrogen(releaseHydrogen)
			wg.Done()
		}()
	}

	// 10个氧原子goroutine
	for i := 0; i < N; i++ {
		go func() {
			h2o.oxygen(releaseOxygen)
			wg.Done()
		}()
	}

	//等待所有的goroutine执行完
	wg.Wait()

	// 结果中肯定是300个原子
	if len(ch) != N*3 {
		t.Fatalf("expect %d atom but got %d", N*3, len(ch))
	}

	// 每三个原子一组，分别进行检查。要求这一组原子中必须包含两个氢原子和一个氧原子，这样才能正确组成一个水分子。
	var s = make([]string, 3)
	for i := 0; i < N; i++ {
		s[0] = <-ch
		s[1] = <-ch
		s[2] = <-ch
		sort.Strings(s)

		water := s[0] + s[1] + s[2]
		if water != "HHO" {
			t.Fatalf("expect a water molecule but got %s", water)
		}
		fmt.Println(water)
	}
}

// 定义水分子合成的辅助数据结构
type H2O struct {
	semaH *semaphore.Weighted         // 氢原子的信号量
	semaO *semaphore.Weighted         // 氧原子的信号量
	cb    cyclicbarrier.CyclicBarrier // 循环栅栏，用来控制合成
}

func New() *H2O {
	return &H2O{
		semaH: semaphore.NewWeighted(2), // 氢原子需要两个
		semaO: semaphore.NewWeighted(1), // 氧原子需要一个
		cb:    cyclicbarrier.New(3),     // 需要三个原子才能合成
	}
}

// 制造氢原子
func (h2o *H2O) hydrogen(releaseHydrogen func()) {
	h2o.semaH.Acquire(context.Background(), 1)

	releaseHydrogen()                  // 输出H
	h2o.cb.Await(context.Background()) // 等待栅栏放行
	h2o.semaH.Release(1)               // 释放氢原子空槽
}

// 制造氧原子
func (h2o *H2O) oxygen(releaseOxygen func()) {
	h2o.semaO.Acquire(context.Background(), 1)

	releaseOxygen()                    // 输出O
	h2o.cb.Await(context.Background()) // 等待栅栏放行
	h2o.semaO.Release(1)               // 释放氧原子空槽
}
