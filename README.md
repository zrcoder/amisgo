# ![Amisgo](https://raw.githubusercontent.com/zrcoder/amisgo-assets/refs/heads/main/logo-with-text.svg)

A low-code frontend framework for Go developers, powered by Baidu's [Amis](https://aisuda.bce.baidu.com/amis).  
The name combines `amis` and `go`, which coincidentally means "friend" in Zulu.

For background information and detailed documentation, visit our [wiki](https://github.com/zrcoder/amisgo/wiki) (in Chinese).

## Quick Start

```go
package main

import (
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
)

func main() {
	index := comp.Page().Title("Amisgo").Body(
		comp.Form().Body(
			comp.InputText().Label("Name").Name("name"),
			comp.InputEmail().Label("Email").Name("email"),
		),
	)

	ag := amisgo.New().Mount("/", index)

	panic(ag.Run(":8080"))
}
```

Visit http://localhost:8080 after running the code.

## Examples

Check out our [examples repository](https://github.com/zrcoder/amisgo-examples) or try the live demo: [Dev-Toys](https://amisgo-examples.onrender.com).

## Project Status

- [x] Auto-generated components based on latest Amis schema
- [ ] Tests and examples (60% complete)

> **Note**: API is not stable yet and subject to change.
