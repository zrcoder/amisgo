package main

import (
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
	"github.com/zrcoder/amisgo/examples/all/pages"
	"github.com/zrcoder/amisgo/model"
)

func main() {
	app := comp.App().BrandName("Amisgo 示例").Pages(
		model.Page().Label("示例").Children(
			model.Page().Label("页面").Children(model.Page().Label("简单页面").Schema(pages.SimplePage)),
			model.Page().Label("表单").Children(),
			model.Page().Label("增删改查").Children(),
			model.Page().Label("弹框").Children(),
			model.Page().Label("选项卡").Children(),
			model.Page().Label("联动").Children(),
			model.Page().Label("事件动作机制").Children(),
			model.Page().Label("动态加载").Children(),
			model.Page().Label("主题编辑器").Children(),
			model.Page().Label("代码高亮").Children(),
			model.Page().Label("向导").Children(),
			model.Page().Label("排版").Children(),
			model.Page().Label("图表").Children(),
			model.Page().Label("ECharts编辑器").Children(),
			model.Page().Label("轮播图").Children(),
			model.Page().Label("音频").Children(),
			model.Page().Label("视频").Children(),
			model.Page().Label("异步任务").Children(),
			model.Page().Label("ifram").Children(),
			model.Page().Label("SDK").Children(),
			model.Page().Label("Office 文档预览").Children(),
			model.Page().Label("Pdf 预览").Children(),
			model.Page().Label("多 loading").Children(),
			model.Page().Label("App 多页应用").Children(),
			model.Page().Label("wizard页面").Children(),
		))
	panic(amisgo.ListenAndServe(app))
}
