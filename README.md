# Amisgo

Amisgo is a low code frontend framework for gophers.  
The name stands for `amis go`, coincidentally means `friend` in Zulu.

> Amisgo is based on baidu [Amis](https://aisuda.bce.baidu.com/amis).  
> See the [wiki](https://github.com/zrcoder/amisgo/wiki) for more background and details.

## Usage

Hello world

```go
package main

import (
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
)

func main() {
	index := comp.NewPage().Title("Amisgo").Body(
		comp.NewPage().Title("Hello world!"),
	)

	amisgo.Route("/", index)
	panic(amisgo.ListenAndServe(":9090"))
}
```

## TODO

- [x] both ListenAndServe, GenerateStaticWebsite api

- [ ] Autoly generate codes for components.go and the comp directory, based on schema.json in Amis' last release(or Amis' docs)

- [ ] Go+ classfile improvement

- [ ] Tests and examples
