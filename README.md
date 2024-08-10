# Amisgo

`amisgo` is a low code frontend framework for gophers.  
 The name stands for `amis go`, coincidentally means `friend` in Zulu.

> `amisgo` is based on [Amis](https://aisuda.bce.baidu.com/amis).

## Usage

Hello word with [Go+](https://goplus.org):

```c
index := page().title("Example").body(
	page().title("Inner"),
)

route "/", index
listenAndServe! ":9090"
```

> see detail in [classfile.md](classfile.md)

with Go:

```go
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
```

## Examples

TODO
