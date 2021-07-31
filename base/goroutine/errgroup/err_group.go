package errgroup

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/sync/errgroup"
)

// 错误组是用来将一个通用的父任务拆成几个小任务并发执行
// 一旦有一个子任务返回错误，或者是Wait调用返回
// 那么整个context将会被取消，所有任务终止执行
func Run(work []int) error {
	eg, ctx := errgroup.WithContext(context.Background())
	for _, w := range work {
		w := w
		eg.Go(func() error {
			return DoSomething(ctx, w)
		})
	}
	if err := eg.Wait(); err != nil {
		return err
	}
	fmt.Println("All success!")
	return nil
}

func DoSomething(ctx context.Context, w int) error {
	if w == 2 {
		return errors.New("w is two")
	}
	fmt.Printf("w is %d \n", w)
	return nil
}
