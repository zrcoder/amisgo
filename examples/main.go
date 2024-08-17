package main

import (
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
)

func main() {
	index := comp.NewPage().Title("Amisgo").Body(
		comp.NewPage().Title("Hello world!"),
	)

	app := &amisgo.Amis{}
	app.Route("/", index)
	panic(app.ListenAndServe(":9090"))
}
