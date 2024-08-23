package main

import (
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
)

func main() {
	index := comp.Page().Body(
		comp.Form().Title("").ColumnCount(2).Body(
			comp.Editor().Language("json").Size("xxl"),
			comp.Editor().Language("yaml").Size("xxl").ReadOnly(true),
		),
	)

	amisgo.Route("/", index)
	panic(amisgo.ListenAndServe(":9090"))
}
