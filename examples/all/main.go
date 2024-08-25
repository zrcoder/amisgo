package main

import (
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
	"github.com/zrcoder/amisgo/examples/all/pages/form"
	"github.com/zrcoder/amisgo/examples/all/pages/page"
)

func main() {
	app := comp.App().BrandName("Amisgo 示例").Pages(
		comp.PageItem().Children(
			comp.PageItem().Icon("fa fa-th").Label("页面").Url("/pages").Children(
				comp.PageItem().Url("simple").Label("简单页面").Schema(page.Simple),
				comp.PageItem().Url("error").Label("初始化出错").Schema(page.InitError),
				comp.PageItem().Url("form").Label("表单页面").Schema(page.Form),
			),
		),
		comp.PageItem().Children(
			comp.PageItem().Icon("fa fa-list-alt").Label("表单").Url("/form").Children(
				comp.PageItem().Url("mode").Label("模式").Schema(form.Modes),
				comp.PageItem().Url("all").Label("所有组件").Schema(form.All),
			),
		),
		comp.PageItem().Label("其他").Url("other").Children(
			comp.PageItem().Icon("fa fa-table").Label("增删改查").Children(),
			comp.PageItem().Icon("fa fa-bomb").Label("弹框").Children(),
			comp.PageItem().Icon("fa fa-clone").Label("选项卡").Children(),
			comp.PageItem().Icon("fa fa-bolt").Label("联动").Children(),
			comp.PageItem().Icon("fa fa-bullhorn").Label("事件动作机制").Children(),
			comp.PageItem().Icon("fa fa-magic").Label("动态加载").Children(),
			comp.PageItem().Icon("fa fa-glasses").Label("主题编辑器").Children(),
			comp.PageItem().Icon("fa fa-code").Label("代码高亮").Children(),
			comp.PageItem().Icon("fa fa-desktop").Label("向导").Children(),
			comp.PageItem().Icon("fa fa-columns").Label("排版").Children(),
			comp.PageItem().Icon("fa fa-bar-chart").Label("图表").Children(),
			comp.PageItem().Icon("fa fa-bar-chart").Label("ECharts编辑器").Children(),
			comp.PageItem().Icon("fa fa-pause").Label("轮播图").Children(),
			comp.PageItem().Icon("fa fa-volume-up").Label("音频").Children(),
			comp.PageItem().Icon("fa fa-video-camera").Label("视频").Children(),
			comp.PageItem().Icon("fa fa-tasks").Label("异步任务").Children(),
			comp.PageItem().Icon("fa fa-cloud").Label("ifram").Children(),
			comp.PageItem().Icon("fa fa-rocket").Label("SDK").Children(),
			comp.PageItem().Icon("fa fa-file-word").Label("Office 文档预览").Children(),
			comp.PageItem().Icon("fa fa-file-pdf").Label("Pdf 预览").Children(),
			comp.PageItem().Icon("fa fa-spinner").Label("多 loading").Children(),
			comp.PageItem().Icon("fa fa-cubes").Label("App 多页应用").Children(),
			comp.PageItem().Icon("fa fa-desktop").Label("wizard页面").Children(),
		))

	panic(amisgo.ListenAndServe(app))
}
