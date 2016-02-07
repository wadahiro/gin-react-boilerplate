package main

import (
	"fmt"

	"github.com/go-errors/errors"
)

// Must raises an error if it not nil
func Must(e error) {
	if e != nil {
		fmt.Println(e.(*errors.Error).ErrorStack())
		panic(e)
	}
}
