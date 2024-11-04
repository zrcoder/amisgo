# Amisgo

Amisgo is a low code frontend framework for gophers.  
The name stands for `amis go`, coincidentally means `friend` in Zulu.

## Inspired by

Amisgo is inspired by and based on baidu [Amis](https://aisuda.bce.baidu.com/amis).  
Read the [wiki](https://github.com/zrcoder/amisgo/wiki)(Chinese) for more background and details.

## Hello world

```go
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
	ag := amisgo.New().Register("/", index)
	panic(ag.Run(":80"))
}
```

and then visit http://localhost to view the served site.

## Examples

See our [examples repo](https://github.com/zrcoder/amisgo-examples), or a live demo: [Dev-Toys](https://amisgo-examples.onrender.com)

## TODO

- [x] both ListenAndServe, GenerateStaticWebsite api
- [x] Generate codes in the comp directory, based on schema.json in Amis' last release(or Amis' docs)
- [ ] Tests and examples(40%)
- [ ] Go+ classfile improvement
