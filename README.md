# Amisgo

Amisgo is a low code frontend framework for gophers.  
The name stands for `amis go`, coincidentally means `friend` in Zulu.

> Amisgo is based on baidu [Amis](https://aisuda.bce.baidu.com/amis).  
> See the [wiki](https://github.com/zrcoder/amisgo/wiki) for more background and details. 

## Usage

Hello world with [Go+](https://goplus.org):

```c
page().title("Example").body(
	page().title("Inner"),
)

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

## TODO

- [ ] Autoly generate codes for components.go and the comp directory, based on schema.json in Amis' last release(or Amis' docs)

- [ ] Go+ classfile improvement

- [ ] Tests and examples
