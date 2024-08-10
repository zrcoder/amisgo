package main

import (
	"github.com/zrcoder/amisgo"
)

func main() {
	app := amisgo.New()

	app.Page().Title("Example").Body(
		app.Page().Title("Inner"),
	)

	app.HandleRouter()

	if err := app.ListenAndServe(":9090"); err != nil {
		panic(err)
	}
}
