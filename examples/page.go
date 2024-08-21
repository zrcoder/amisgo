package main

import (
	"github.com/zrcoder/amisgo/comp"
	"github.com/zrcoder/amisgo/model"
)

/*
Page

	├── Toolbar
	│  └─ Form 顶部表单项
	├── Grid // 用于水平布局
	│  ├─ Panel
	│  │  └─ Tabs
	│  │    └─ Chart
	│  └─ Panel
	│     └─ Chart
	└── CRUD
*/
var page = comp.Page().
	Toolbar(
		comp.Form().StaticLabelClassName("").
			PanelClassName("mb-0").
			Body(
				comp.Select().
					Label("区域").
					Name("businessLineId").
					SelectFirst(true).
					Mode("inline").
					Options(
						"北京",
						"上海",
					).
					CheckAll(false),
				comp.InputDateRange().
					Label("时间范围").
					Name("dateRange").
					Inline(true).
					Value("-1month,+0month").
					InputFormat("YYYY-MM-DD").
					CloseOnSelect(true).
					Clearable(false),
			).
			Mode("inline").
			Target("mainPage").
			SubmitOnChange(true).
			SubmitOnInit(true),
	).
	Body(
		comp.Grid().
			Columns(
				comp.Panel().
					ClassName("h-full").
					Body(
						comp.Tabs().
							Tabs(
								model.Tab().
									Title("消费趋势").
									Tab(
										comp.Chart().
											Config(trendChartCOnfig),
									),
								model.Tab().
									Title("账户余额").
									Tab("0.00"),
							),
					),
				comp.Panel().
					ClassName("h-full").
					Body("Test"),
			),
		comp.CrudTable().
			ClassName("m-t-sm").
			Api(itemsRouter).
			Columns(
				model.Column().Name("id").Label("ID"),
				model.Column().Name("engine").Label("Rendering engine"),
				model.Column().Name("browser").Label("Browser"),
				model.Column().Name("platform").Label("Platform(s)"),
				model.Column().Name("version").Label("Engine version"),
				model.Column().Name("grade").Label("CSS grade"),
			),
	)
