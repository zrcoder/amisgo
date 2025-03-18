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

## Use Cases

### [podFiles](https://github.com/zrcoder/podFiles)

PodFiles is a tool based on amisgo, designed to manage files within Kubernetes pods.

### [amisgo-examples](https://github.com/zrcoder/amisgo-examples)

amisgo-examples is a sample repository that demonstrates how to use amisgo, containing multiple examples, each of which has practical value.

## Issues and Contributions

This project is primarily maintained on Gitee. For issues or pull requests, please visit https://gitee.com/rdor/amisgo
