package pages

import "github.com/zrcoder/amisgo/comp"

var SimplePage = comp.Page().
	Title("标题").
	Remark(comp.Remark().
		Title("标题").
		Body("这是一段描述问题，注意到了没，还可以设置标题。而且只有点击了才弹出来。").
		Icon("question-mark").
		Placement("right").
		Trigger("click").
		RootClose(true),
	).Body("内容部分. 可以使用 \\${var} 获取变量。如: `\\$date`: ${date}").
	Aside("边栏部分").
	Toolbar("工具栏")
	// InitApi("")
