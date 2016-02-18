package main

import (
	"github.com/elago/ela"
	_ "github.com/elago/elapen/model"
	_ "github.com/elago/elapen/router"
)

func main() {
	ela.Run()
}
