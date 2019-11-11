package goroutinue

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/sync/errgroup"
)

// 错误组(Errgroup)
// 有时候你想创建多个goroutine，
// 让它们并行地工作，当遇到某种错误或者你不想再输出了，
// 你可能想取消整个goroutine。

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
