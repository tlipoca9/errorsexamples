package main

import (
	"fmt"

	"github.com/fioepq9/errors"
)

func foo(method string) error {
	switch method {
	case "new":
		return errors.New("foo")
	case "wrap":
		return errors.Wrap(foo("new"), "this is wrapeed foo")
	default:
		panic("invalid method")
	}
}

func bar(method string) error {
	switch method {
	case "new":
		return foo("new")
	case "wrap":
		return errors.Wrap(foo("wrap"), "this is wrapeed bar")
	default:
		panic("invalid method")
	}
}

func baz(method string) error {
	switch method {
	case "new":
		return bar("new")
	case "wrap":
		return errors.Wrap(bar("wrap"), "this is wrapeed baz")
	default:
		panic("invalid method")
	}
}

func main() {
	errors.C.Style = errors.StyleStack

	fmt.Println(baz("new"))

	fmt.Println("=====================================")

	fmt.Println(baz("wrap"))
}

/*
foo
  D:/code/projects/errorsexamples/quickstart/main.go:12 (0x457eda) main.foo()
  D:/code/projects/errorsexamples/quickstart/main.go:23 (0x457f9a) main.bar()
  D:/code/projects/errorsexamples/quickstart/main.go:34 (0x45805a) main.baz()
  D:/code/projects/errorsexamples/quickstart/main.go:45 (0x458133) main.main()
  D:/apps/scoop/apps/go/current/src/runtime/proc.go:267 (0x405231) runtime.main()
=====================================
foo
  D:/code/projects/errorsexamples/quickstart/main.go:12 (0x457eda) main.foo()
this is wrapeed foo
  D:/code/projects/errorsexamples/quickstart/main.go:14 (0x457f15) main.foo()
this is wrapeed bar
  D:/code/projects/errorsexamples/quickstart/main.go:25 (0x457fd5) main.bar()
this is wrapeed baz
  D:/code/projects/errorsexamples/quickstart/main.go:36 (0x458095) main.baz()
  D:/code/projects/errorsexamples/quickstart/main.go:49 (0x4581bb) main.main()
  D:/apps/scoop/apps/go/current/src/runtime/proc.go:267 (0x405231) runtime.main()
*/
