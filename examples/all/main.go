package main

import (
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
	"github.com/zrcoder/amisgo/examples/all/pages/form"
	"github.com/zrcoder/amisgo/examples/all/pages/page"
	"github.com/zrcoder/amisgo/model"
)

func main() {
	app := comp.App().BrandName("Amisgo 示例").Pages(
		model.Page().Children(
			model.Page().Icon("fa fa-th").Label("页面").Url("/pages").Children(
				model.Page().Url("simple").Label("简单页面").Schema(page.Simple),
				model.Page().Url("error").Label("初始化出错").Schema(page.InitError),
				model.Page().Url("form").Label("表单页面").Schema(page.Form),
			),
		),
		model.Page().Children(
			model.Page().Icon("fa fa-list-alt").Label("表单").Url("/form").Children(
				model.Page().Url("mode").Label("模式").Schema(form.Modes),
			),
		),
		model.Page().Label("其他").Url("other").Children(
			model.Page().Icon("fa fa-table").Label("增删改查").Children(),
			model.Page().Icon("fa fa-bomb").Label("弹框").Children(),
			model.Page().Icon("fa fa-clone").Label("选项卡").Children(),
			model.Page().Icon("fa fa-bolt").Label("联动").Children(),
			model.Page().Icon("fa fa-bullhorn").Label("事件动作机制").Children(),
			model.Page().Icon("fa fa-magic").Label("动态加载").Children(),
			model.Page().Icon("fa fa-glasses").Label("主题编辑器").Children(),
			model.Page().Icon("fa fa-code").Label("代码高亮").Children(),
			model.Page().Icon("fa fa-desktop").Label("向导").Children(),
			model.Page().Icon("fa fa-columns").Label("排版").Children(),
			model.Page().Icon("fa fa-bar-chart").Label("图表").Children(),
			model.Page().Icon("fa fa-bar-chart").Label("ECharts编辑器").Children(),
			model.Page().Icon("fa fa-pause").Label("轮播图").Children(),
			model.Page().Icon("fa fa-volume-up").Label("音频").Children(),
			model.Page().Icon("fa fa-video-camera").Label("视频").Children(),
			model.Page().Icon("fa fa-tasks").Label("异步任务").Children(),
			model.Page().Icon("fa fa-cloud").Label("ifram").Children(),
			model.Page().Icon("fa fa-rocket").Label("SDK").Children(),
			model.Page().Icon("fa fa-file-word").Label("Office 文档预览").Children(),
			model.Page().Icon("fa fa-file-pdf").Label("Pdf 预览").Children(),
			model.Page().Icon("fa fa-spinner").Label("多 loading").Children(),
			model.Page().Icon("fa fa-cubes").Label("App 多页应用").Children(),
			model.Page().Icon("fa fa-desktop").Label("wizard页面").Children(),
		))

	panic(amisgo.ListenAndServe(app))
}
