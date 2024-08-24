package pages

import (
	"net/http"

	"github.com/zrcoder/amisgo/comp"
	"github.com/zrcoder/amisgo/model"
)

var InitError = comp.Page().
	Title("标题").
	Remark("提示 Tip").
	Body(
		"`initApi` 返回 status 非 0 时，页面内容区会显示对应的错误信息。<br />",
		"其他提示示例:<br />",
		comp.Alert().Level("success").Body("温馨提示：对页面功能的提示说明，绿色为正向类的消息提示"),
		comp.Alert().Level("warning").Body("您的私有网络已达到配额，如需更多私有网络，可以通过<a>工单</a>申请"),
	).
	Aside("边栏").
	Toolbar("工具栏").
	InitApi("/initDateError")

func init() {
	http.HandleFunc("/initDateError", func(w http.ResponseWriter, r *http.Request) {
		resp := model.Response{
			Status: http.StatusInternalServerError,
			Msg:    "测试返回错误",
		}
		w.Write(resp.Json())
	})
}
