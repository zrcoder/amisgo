# ![amisgo](https://raw.githubusercontent.com/zrcoder/amisgo-assets/refs/heads/main/logo-with-text.svg)

A low-code frontend framework for Go developers, powered by Baidu's [amis](https://aisuda.bce.baidu.com/amis).  
The name combines `amis` and `go`, which coincidentally means "friend" in Zulu.

For background info and deeper insights, check out our [documentation](https://amisgo.pages.dev) (in Chinese).

## Quick Start

```go
package main

import (
	"github.com/zrcoder/amisgo"
)

func main() {
	app := amisgo.New()
	index := app.Page().Title("amisgo").Body(
		app.Form().
			Api("https://xxx/api/saveForm").
			Body(
				app.InputText().Label("姓名").Name("name"),
				app.InputEmail().Label("邮箱").Name("email"),
			),
	)
	app.Mount("/", index)

	panic(app.Run(":8080"))
}
```

Visit http://localhost:8080 after running the code.

## Tutorials

Explore our [documentation](https://amisgo.pages.dev) for in-depth tutorials and guides.

## Examples

Check out our [examples repository](https://github.com/zrcoder/amisgo-examples).