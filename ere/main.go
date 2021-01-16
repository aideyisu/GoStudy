package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	err := Z()
	if err != nil {
		fmt.Println(err)
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
