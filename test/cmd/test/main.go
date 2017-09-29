package main

import (
	"fmt"

	"github.com/JacobSieberg/golang-CI-tests/test/pack1"
)

func main() {
	pack1.Test("Value")
	fmt.Println("Hello World")
}
