package foo

import "github.com/pkg/errors"

func Foo() int {
	_ = errors.New("test")
	return 1
}
