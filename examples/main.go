package main

import (
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
)

func main() {
	index := comp.Page().Title("Amisgo").Body(
		comp.Form().Mode("horizontal").Body(
			comp.InputText().Label("Name").Name("name"),
			comp.InputEmail().Label("Email").Name("email"),
		),
	)

	cfg := &amisgo.Config{
		Theme: amisgo.ThemeAntd,
		Lang:  amisgo.LangEn,
		Title: "test",
	}

	amisgo.Route("/", index)
	panic(amisgo.ListenAndServe(":9090", cfg))
}
