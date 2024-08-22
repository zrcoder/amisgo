package main

import (
	"fmt"

	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
)

func main() {
	index := comp.Page().Body(
		comp.Form().Body(
			comp.InputText().Label("姓名").Name("name"),
			comp.InputEmail().Label("邮箱").Name("email"),
		).Go(func(m map[string]any) {
			fmt.Println(m["name"])
			fmt.Println(m["email"])
			// save name and email to db
			// ...
		}),
	)

	amisgo.Route("/", index)
	panic(amisgo.ListenAndServe(":9090"))
}
