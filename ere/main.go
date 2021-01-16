package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	fmt.Println("vim-go")
	err := Z()
	if err != nil {
		fmt.Println(err)
	}

	Cause := errors.New("what")

	e2 := errors.Wrapf(Cause, "hi")
	if errors.Is(e2, Cause) {
		fmt.Println("hello")
	}

	e3 := Dao(Cause)
	fmt.Println(errors.Is(e3, Cause))
	if errors.Is(e3, Cause) {
		fmt.Println("hello")
	}
}

func Y() error {
	return fmt.Errorf("where")
}

func X() error {
	return Y()
}

func Z() error {
	return X()
}

func Dao(Cause error) error {
	return errors.Wrapf(Cause, "this is only a test.")
}
