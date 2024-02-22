package main

import (
	"fmt"

	gerrors "errors"

	pkgerrors "github.com/pkg/errors"

	"github.com/fioepq9/errors"
)

func foo(method string) error {
	switch method {
	case "new":
		return errors.New("foo")
	case "wrap":
		return errors.Wrap(foo("new"), "this is wrapeed foo")
	case "g-new":
		return gerrors.New("foo")
	case "g-wrap":
		return fmt.Errorf("%s: %w", "this is wrapeed foo", foo("g-new"))
	case "pkg-new":
		return pkgerrors.New("foo")
	case "pkg-wrap":
		return pkgerrors.Wrap(foo("pkg-new"), "this is wrapeed foo")
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
	case "g-new":
		return foo("g-new")
	case "g-wrap":
		return fmt.Errorf("%s: %w", "this is wrapeed bar", foo("g-wrap"))
	case "pkg-new":
		return foo("pkg-new")
	case "pkg-wrap":
		return pkgerrors.Wrap(foo("pkg-wrap"), "this is wrapeed bar")
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
	case "g-new":
		return bar("g-new")
	case "g-wrap":
		return fmt.Errorf("%s: %w", "this is wrapeed baz", bar("g-wrap"))
	case "pkg-new":
		return bar("pkg-new")
	case "pkg-wrap":
		return pkgerrors.Wrap(bar("pkg-wrap"), "this is wrapeed baz")
	default:
		panic("invalid method")
	}
}

type stackTracer interface {
	StackTrace() pkgerrors.StackTrace
}

func main() {
	fmt.Println("==== golang errors.New ====")
	fmt.Println(baz("g-new"))
	fmt.Printf("===========================\n\n")

	fmt.Println("==== golang fmt.Errorf(%w) ====")
	fmt.Println(baz("g-wrap"))
	fmt.Printf("============================\n\n")

	fmt.Println("==== pkg errors.New ====")
	fmt.Println(baz("pkg-new"))
	fmt.Printf("=========================\n\n")

	fmt.Println("==== pkg errors.Wrap ====")
	fmt.Println(baz("pkg-wrap"))
	fmt.Printf("==========================\n\n")

	fmt.Println("==== pkg errors.Wrap(stack trace) ====")
	if err, ok := baz("pkg-wrap").(stackTracer); ok {
		for _, f := range err.StackTrace() {
			fmt.Printf("%+s:%d\n", f, f)
		}
	}
	fmt.Printf("=====================================\n\n")

	fmt.Println("==== fioepq9 errors.New(normal style) ====")
	fmt.Println(baz("new"))
	fmt.Printf("===========================================\n\n")

	fmt.Println("==== fioepq9 errors.Wrap(normal style) ====")
	fmt.Println(baz("wrap"))
	fmt.Printf("============================================\n\n")

	errors.C.Style = errors.StyleStack
	fmt.Println("==== fioepq9 errors.New(stack style: default) ====")
	fmt.Println(baz("new"))
	fmt.Printf("===================================================\n\n")

	fmt.Println("==== fioepq9 errors.Wrap(stack style: default) ====")
	fmt.Println(baz("wrap"))
	fmt.Printf("====================================================\n\n")

	errors.C.StackFramesHandler = errors.JSONStackFramesHandler
	fmt.Println("==== fioepq9 errors.New(stack style: json) ====")
	fmt.Println(baz("new"))
	fmt.Printf("================================================\n\n")

	fmt.Println("==== fioepq9 errors.Wrap(stack style: json) ====")
	fmt.Println(baz("wrap"))
	fmt.Printf("=================================================\n\n")

}
