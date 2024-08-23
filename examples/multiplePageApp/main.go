package main

import (
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
	"github.com/zrcoder/amisgo/model"
)

func main() {
	app := comp.App().BrandName("Multiple page test").Pages(
		model.Page().Label("Group 1").Children(
			// model.Page().Children(
			model.Page().Label("Sub page").Url("pageA").Schema(
				comp.Page().Title("Page A"),
			),
			model.Page().Label("Sub page").Url("pageB").Schema(
				comp.Page().Title("Page B"),
			),
			// ),
		),
	)
	panic(amisgo.ListenAndServe(app))
}
