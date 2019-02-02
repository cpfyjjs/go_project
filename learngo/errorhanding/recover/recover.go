package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Printf("Error occurred: %s", err)
		} else {
			panic(r)
		}
	}()

	panic(errors.New("this is an error"))
}

func main() {
	tryRecover()

}
