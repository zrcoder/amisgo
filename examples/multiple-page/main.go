package main

import (
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
)

func main() {
	app := comp.App().BrandName("Multiple page test").Pages(
		comp.PageItem().Label("Group 1").Children(
			comp.PageItem().Label("pages part1").Url("/pages1").Children(
				comp.PageItem().Label("Sub page").Url("pageA").Schema(comp.Page().Title("Page A")),
				comp.PageItem().Label("Sub page").Url("pageB").Schema(comp.Page().Title("Page B")),
			),
			comp.PageItem().Label("pages part2").Url("/pages2").Children(
				comp.PageItem().Label("Sub page").Url("pageA").Schema(comp.Page().Title("Page A")),
				comp.PageItem().Label("Sub page").Url("pageB").Schema(comp.Page().Title("Page B")),
			),
		),
		comp.PageItem().Label("Group 2").Children(
			comp.PageItem().Label("pages part3").Url("/pages3").Children(
				comp.PageItem().Label("Sub page").Url("pageA").Schema(comp.Page().Title("Page A")),
				comp.PageItem().Label("Sub page").Url("pageB").Schema(comp.Page().Title("Page B")),
			),
		),
	)
	panic(amisgo.ListenAndServe(app))
}
