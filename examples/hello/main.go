package main

import (
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
)

func main() {
	page := comp.Page().Body("Hello World!")
	amisgo.Route("/", page)
	panic(amisgo.ListenAndServe(""))
}
