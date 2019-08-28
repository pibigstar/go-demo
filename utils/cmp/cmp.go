package cmp

import "github.com/google/go-cmp/cmp"

// 类似于git的diff
func Diff(x, y interface{}, opts ...cmp.Option) string {
	return cmp.Diff(x, y, opts...)
}
