package main

import (
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
)

func main() {
	index := comp.NewPage().Title("Amisgo").Body(
		comp.NewPage().Title("Hello world!"),
	)

	amisgo.Route("/", index)
	panic(amisgo.ListenAndServe(":9090"))
}
