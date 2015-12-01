package main

import (
	"github.com/elago/ela"
	_ "github.com/elago/webapp/model"
	_ "github.com/elago/webapp/router"
)

func main() {
	ela.Run()
}
