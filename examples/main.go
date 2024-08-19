package main

import (
	"github.com/zrcoder/amisgo"
)

func main() {
	ServeApi()

	amisgo.Route("/", page)
	panic(amisgo.ListenAndServe(":8080"))
}
