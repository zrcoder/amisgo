# Amisgo

Amisgo is a low code frontend framework for gophers.  
The name stands for `amis go`, coincidentally means `friend` in Zulu.

## Inspired by

Amisgo is inspired by and based on baidu [Amis](https://aisuda.bce.baidu.com/amis).  
Read the [wiki](https://github.com/zrcoder/amisgo/wiki)(Chinese) for more background and details.

## Usage

Hello world

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

	panic(amisgo.ListenAndServe(index))
}
```

and then visit http://localhost, you will see a page like this:

![hello-amis](./hello-amis.png)

## Examples

See our [examples repo](https://github.com/zrcoder/amisgo-examples)

## TODO

- [x] both ListenAndServe, GenerateStaticWebsite api
- [x] Generate codes in the comp directory, based on schema.json in Amis' last release(or Amis' docs)
- [ ] Tests and examples(10%)
- [ ] Go+ classfile improvement
