package comp

import (
	js "encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/zrcoder/amisgo/internal/servermux"
)

// action 行为按钮 https://aisuda.bce.baidu.com/amis/zh-CN/components/action
type action schema

// Action 创建一个新的 Action 实例
func Action() action {
	return action{}.set("type", "action")
}

// Button 实际为 Action 的 别名
func Button() action {
	return Action().set("type", "button")
}

// Submit submit行为按钮
func Submit() action {
	return Action().set("type", "submit")
}

// ActionType 【必填】这是 action 最核心的配置，来指定该 action 的作用类型
// 可选值:
// ajax | link | url | drawer | dialog | confirm | cancel | prev | next | copy | close
// button | reset | submit | clear
func (a action) ActionType(value string) action {
	return a.set("actionType", value)
}

// Dialog 配置按钮点击后的 dialog
func (a action) Dialog(value dialog) action {
	return a.set("dialog", value)
}

// Drawer 配置按钮点击后的 drawer
func (a action) Drawer(value any) action {
	return a.set("drawer", value)
}

// Toast 配置按钮点击后的 toast
func (a action) Toast(value toast) action {
	return a.set("toast", value)
}

// Transform transform the src value with transfor, and renderer the result to dst component
func (a action) Transform(src, dst, successMsg string, transfor func(input any) (any, error)) action {
	return a.transform(src, dst, successMsg, transfor)
}

// TransformMultipletransform the inputs(single or multiple) with transfor, and renderer the result to multiple destinates
func (a action) TransformMultiple(inputs any, successMsg string, transfor func(any) (any, error)) action {
	return a.transform(inputs, "", successMsg, transfor)
}

func (a action) transform(input any, dstKey, successMsg string, transfor func(any) (any, error)) action {
	route := fmt.Sprintf("/__amisgo_api_%d", getInnerApiID())
	servermux.Mux().HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		inputData, err := io.ReadAll(r.Body)
		if err != nil {
			resp := ErrorResponse(err.Error())
			w.Write(resp.Json())
			return
		}
		defer r.Body.Close()
		m := schema{}
		js.Unmarshal(inputData, &m)
		input := m["input"]
		output, err := transfor(input)
		if err != nil {
			resp := ErrorResponse(err.Error())
			w.Write(resp.Json())
			return
		}

		resp := Response{}
		if dstKey != "" {
			resp = Response{Msg: successMsg, Data: Data{dstKey: output}}
		} else {
			resp = Response{Msg: successMsg, Data: output.(Data)}
		}
		w.Write(resp.Json())
	})

	ipt := input
	if s, ok := input.(string); ok {
		ipt = fmt.Sprintf("${%s}", s)
	}
	return a.ActionType("ajax").Api(
		Schema{
			"url":  route,
			"data": Schema{"input": ipt},
			"responses": Schema{
				"200": Schema{
					"then": EventAction().ActionType("setValue").Args(Schema{"value": "${response}"}),
				},
			},
		},
	)
}

// ActiveClassName 给按钮高亮添加类名。
func (a action) ActiveClassName(value string) action {
	return a.set("activeClassName", value)
}

// ActiveLevel 按钮高亮时的样式，配置支持同level。
func (a action) ActiveLevel(value string) action {
	return a.set("activeLevel", value)
}

// Api 设置 API 地址/描述。
func (a action) Api(value any) action {
	return a.set("api", value)
}

// ClassName 添加类名。
func (a action) ClassName(value string) action {
	return a.set("className", value)
}

// Close 当action配置在dialog或drawer的actions中时，配置为true指定此次操作完后关闭当前dialog或drawer。
// 当值为字符串，并且是祖先层弹框的名字的时候，会把祖先弹框关闭掉。
func (a action) Close(value string) action {
	return a.set("close", value)
}

// ConfirmText 当设置后，操作在开始前会询问用户。可用 '$[xxx]' 取值。
func (a action) ConfirmText(value string) action {
	return a.set("confirmText", value)
}

// ConfirmTitle 确认弹窗标题
func (a action) ConfirmTitle(value any) action {
	return a.set("confirmTitle", value)
}

// DisabledTip 被禁用后鼠标停留时弹出该段文字，也可以配置对象类型：字段为title和content。可用 '$[xxx]' 取值。
func (a action) DisabledTip(value any) action {
	return a.set("disabledTip", value)
}

// Icon 设置图标，例如fa fa-plus。
func (a action) Icon(value string) action {
	return a.set("icon", value)
}

// IconClassName 给图标上添加类名。
func (a action) IconClassName(value string) action {
	return a.set("iconClassName", value)
}

// Redirect 配置相对路径，实现单页跳转
func (a action) Redirect(value string) action {
	return a.set("redirect", value)
}

// Label 按钮文本。可用 '$[xxx]' 取值。
func (a action) Label(value string) action {
	return a.set("label", value)
}

// Level 按钮样式
// 'link' | 'primary' | 'enhance' | 'secondary' | 'info'|'success' |
// 'warning' | 'danger' | 'light'| 'dark' | 'default'
func (a action) Level(value string) action {
	return a.set("level", value)
}

// Link 设置链接。
func (a action) Link(value string) action {
	return a.set("link", value)
}

// Reload 指定此次操作完后，需要刷新的目标组件名字（组件的name值，自己配置的），多个请用 "," 号隔开。
func (a action) Reload(value string) action {
	return a.set("reload", value)
}

// ReloadWindow 刷新当前页面
func (a action) ReloadWindow() action {
	return a.Reload("window")
}

// Required 配置字符串数组，指定在form中进行操作之前，需要指定的字段名的表单项通过验证
func (a action) Required(value ...string) action {
	return a.set("required", value)
}

// RightIcon 在按钮文本右侧设置图标，例如 "fa fa-plus"。
func (a action) RightIcon(value string) action {
	return a.set("rightIcon", value)
}

// RightIconClassName 给右侧图标上添加类名。
func (a action) RightIconClassName(value string) action {
	return a.set("rightIconClassName", value)
}

// Size 按钮大小，支持：xs、sm、md、lg。 可选值: xs | sm | md | lg
func (a action) Size(value string) action {
	return a.set("size", value)
}

// Tooltip 鼠标停留时弹出该段文字，也可以配置对象类型：字段为title和content。可用 '$[xxx]' 取值。
func (a action) Tooltip(value string) action {
	return a.set("tooltip", value)
}

// TooltipPlacement 如果配置了tooltip或者disabledTip，指定提示信息位置，可配置top、bottom、left、right。
func (a action) TooltipPlacement(value string) action {
	return a.set("tooltipPlacement", value)
}

// TooltipTrigger 'hover' | 'focus'
func (a action) TooltipTrigger(value string) action {
	return a.set("tooltipTrigger", value)
}

// Badge 角标 (Badge 角标。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/badge)
func (v action) Badge(value string) action {
	return v.set("badge", value)
}

// Block 是否为块状展示，默认为内联。
func (v action) Block(value bool) action {
	return v.set("block", value)
}

// Body 子内容 (子内容)
func (v action) Body(value ...any) action {
	return v.set("body", value)
}

// CountDown 点击后的禁止倒计时（秒）
func (v action) CountDown(value string) action {
	return v.set("countDown", value)
}

// CountDownTpl 倒计时文字自定义
func (v action) CountDownTpl(value string) action {
	return v.set("countDownTpl", value)
}

// Disabled 是否禁用
func (v action) Disabled(value bool) action {
	return v.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (v action) DisabledOn(value string) action {
	return v.set("disabledOn", value)
}

// DownloadFileName
func (v action) DownloadFileName(value string) action {
	return v.set("downloadFileName", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (v action) EditorSetting(value string) action {
	return v.set("editorSetting", value)
}

// Hidden 是否隐藏
func (v action) Hidden(value bool) action {
	return v.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (v action) HiddenOn(value string) action {
	return v.set("hiddenOn", value)
}

// HotKey 键盘快捷键
func (v action) HotKey(value string) action {
	return v.set("hotKey", value)
}

// ID 主要用于用户行为跟踪里区分是哪个按钮
func (v action) ID(value string) action {
	return v.set("id", value)
}

// Url 按钮链接
func (v action) Url(value string) action {
	return v.set("url", value)
}

// Loading 是否显示按钮加载效果
func (a action) Loading(value bool) action {
	return a.set("loading", value)
}

// LoadingClassName loading 上的css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (v action) LoadingClassName(value string) action {
	return v.set("loadingClassName", value)
}

// LoadingOn 是否显示loading效果
func (v action) LoadingOn(value string) action {
	return v.set("loadingOn", value)
}

// MergeData 是否将弹框中数据 merge 到父级作用域。
func (v action) MergeData(value bool) action {
	return v.set("mergeData", value)
}

// OnClick 自定义事件处理函数
func (v action) OnClick(value string) action {
	return v.set("onClick", value)
}

// OnEvent 事件动作配置
func (v action) OnEvent(value any) action {
	return v.set("onEvent", value)
}

// Primary
func (v action) Primary(value bool) action {
	return v.set("primary", value)
}

// RequireSelected 当按钮时批量操作按钮时，默认必须有勾选元素才能可点击，如果此属性配置成 false，则没有点选成员也能点击。
func (v action) RequireSelected(value bool) action {
	return v.set("requireSelected", value)
}

// Static 是否静态展示
func (v action) Static(value bool) action {
	return v.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (v action) StaticClassName(value string) action {
	return v.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (v action) StaticInputClassName(value string) action {
	return v.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (v action) StaticLabelClassName(value string) action {
	return v.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (v action) StaticOn(value string) action {
	return v.set("staticOn", value)
}

// StaticPlaceholder 静态展示为空时的占位符。
func (v action) StaticPlaceholder(value string) action {
	return v.set("staticPlaceholder", value)
}

// StaticSchema 静态展示模式的 schema
func (v action) StaticSchema(value string) action {
	return v.set("staticSchema", value)
}

// Style 组件样式
func (v action) Style(value any) action {
	return v.set("style", value)
}

// Target 谁能触发这个动作。
func (v action) Target(value string) action {
	return v.set("target", value)
}

// TestIdBuilder 设置测试ID生成函数
func (v action) TestIdBuilder(value string) action {
	return v.set("testIdBuilder", value)
}

// Testid 测试ID
func (v action) Testid(value string) action {
	return v.set("testid", value)
}

// UseMobileUI 是否禁用移动端样式
func (v action) UseMobileUI(value bool) action {
	return v.set("useMobileUI", value)
}

// Visible 是否可见
func (v action) Visible(value bool) action {
	return v.set("visible", value)
}

// VisibleOn 是否可见表达式 (表达式，语法 `data.xxx > 5`。)
func (v action) VisibleOn(value string) action {
	return v.set("visibleOn", value)
}

func (a action) set(key string, value any) action {
	a[key] = value
	return a
}
