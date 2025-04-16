# ![amisgo](https://raw.githubusercontent.com/zrcoder/amisgo-assets/refs/heads/main/logo-with-text.svg)
A Go package for building low-code frontend applications with Baidu [amis](https://aisuda.bce.baidu.com/amis).

"amisgo" combines "amis" and "go", meaning "friend" in Zulu.

For full documentation, visit [here](https://amisgo.pages.dev) (Chinese).

## Quick Start

```go
package main

import "github.com/zrcoder/amisgo"

func main() {
	app := amisgo.New()
	index := app.Page().Title("amisgo").Body(
		app.Form().Api("https://xxx/api/saveForm").Body(
			app.InputText().Label("姓名").Name("name"),
			app.InputEmail().Label("邮箱").Name("email"),
		),
	)
	app.Mount("/", index)
	panic(app.Run(":8080"))
}
```

Visit http://localhost:8080 to see the result.

## Use Cases

- [podFiles](https://github.com/zrcoder/podFiles) - Kubernetes pod file manager
- [amisgo-examples](https://github.com/zrcoder/amisgo-examples) - Sample projects
- [amisgo-games](https://github.com/zrcoder/agg) - Amisgo games for STEM education

## Contributing

For issues or pull requests, visit https://gitee.com/rdor/amisgo
