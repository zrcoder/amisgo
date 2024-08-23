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

	cfg := amisgo.GetDefaultConfig()
	cfg.Theme = amisgo.ThemeAntd
	cfg.Lang = amisgo.LangEn
	cfg.Title = "test"

	panic(amisgo.ListenAndServe(":9090", index, cfg))
}
