package main

import (
	"fmt"
	"github.com/go-elaeagnus/elaeagnus"
)

func f1() {
	fmt.Println("hello f1")
}

func f2() {
	fmt.Println("hello f2")
}

func main() {
	elaeagnus.Router("/hello1", f1)
	elaeagnus.Router("/hello2", f2)
	elaeagnus.Run()
}
