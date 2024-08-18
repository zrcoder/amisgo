package main

import (
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
)

func main() {
	index := comp.NewPage().Title("Amisgo").Body(
		comp.NewFormControl().Mode("horizontal").Body(
			comp.NewInputText().Label("Name").Name("name"),
			comp.NewInputText().Type("input-email").Label("Email").Name("email"),
		),
	)

	amisgo.Route("/", index)
	panic(amisgo.ListenAndServe(":9090"))
}
