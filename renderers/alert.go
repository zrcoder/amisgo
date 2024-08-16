package renderers

// Alert 提示渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/alert
//
// @author slowlyo
// @version 6.6.0
type Alert struct {
	*BaseRenderer
}

// NewAlert 创建一个新的 Alert 实例
func NewAlert() *Alert {
	alert := &Alert{
		BaseRenderer: NewBaseRenderer(),
	}
	alert.set("type", "alert")
	return alert
}

func (a *Alert) set(key string, value interface{}) *Alert {
	a.BaseRenderer.set(key, value)
	return a
}

// 操作区域 (操作区域)
func (a *Alert) Actions(value string) *Alert {
	return a.set("actions", value)
}

// 内容区域 (内容区域)
func (a *Alert) Body(value string) *Alert {
	return a.set("body", value)
}

// 容器 css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (a *Alert) ClassName(value string) *Alert {
	return a.set("className", value)
}

// 关闭按钮CSS类名
func (a *Alert) CloseButtonClassName(value string) *Alert {
	return a.set("closeButtonClassName", value)
}

// 是否禁用
func (a *Alert) Disabled(value bool) *Alert {
	return a.set("disabled", value)
}

// 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (a *Alert) DisabledOn(value string) *Alert {
	return a.set("disabledOn", value)
}

// 编辑器配置，运行时可以忽略
func (a *Alert) EditorSetting(value string) *Alert {
	return a.set("editorSetting", value)
}

// 是否隐藏
func (a *Alert) Hidden(value bool) *Alert {
	return a.set("hidden", value)
}

// 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (a *Alert) HiddenOn(value string) *Alert {
	return a.set("hiddenOn", value)
}

// 左侧图标 (iconfont 里面的类名。)
func (a *Alert) Icon(value string) *Alert {
	return a.set("icon", value)
}

// 图标CSS类名
func (a *Alert) IconClassName(value string) *Alert {
	return a.set("iconClassName", value)
}

// 组件唯一 id，主要用于日志采集
func (a *Alert) ID(value string) *Alert {
	return a.set("id", value)
}

// 提示类型 可选值: info | warning | success | danger
func (a *Alert) Level(value string) *Alert {
	return a.set("level", value)
}

// 事件动作配置
func (a *Alert) OnEvent(value string) *Alert {
	return a.set("onEvent", value)
}

// 是否显示关闭按钮
func (a *Alert) ShowCloseButton(value bool) *Alert {
	return a.set("showCloseButton", value)
}

// 是否显示ICON
func (a *Alert) ShowIcon(value bool) *Alert {
	return a.set("showIcon", value)
}

// 是否静态展示
func (a *Alert) Static(value bool) *Alert {
	return a.set("static", value)
}

// 静态展示表单项类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (a *Alert) StaticClassName(value string) *Alert {
	return a.set("staticClassName", value)
}

// 静态展示表单项Value类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (a *Alert) StaticInputClassName(value string) *Alert {
	return a.set("staticInputClassName", value)
}

// 静态展示表单项Label类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (a *Alert) StaticLabelClassName(value string) *Alert {
	return a.set("staticLabelClassName", value)
}

// 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (a *Alert) StaticOn(value string) *Alert {
	return a.set("staticOn", value)
}

// 静态展示空值占位
func (a *Alert) StaticPlaceholder(value string) *Alert {
	return a.set("staticPlaceholder", value)
}

func (a *Alert) StaticSchema(value string) *Alert {
	return a.set("staticSchema", value)
}

// 组件样式
func (a *Alert) Style(value string) *Alert {
	return a.set("style", value)
}

func (a *Alert) TestIdBuilder(value string) *Alert {
	return a.set("testIdBuilder", value)
}

func (a *Alert) Testid(value string) *Alert {
	return a.set("testid", value)
}

// 提示框标题
func (a *Alert) Title(value string) *Alert {
	return a.set("title", value)
}

// 可以组件级别用来关闭移动端样式
func (a *Alert) UseMobileUI(value bool) *Alert {
	return a.set("useMobileUI", value)
}

// 是否显示
func (a *Alert) Visible(value bool) *Alert {
	return a.set("visible", value)
}

// 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (a *Alert) VisibleOn(value string) *Alert {
	return a.set("visibleOn", value)
}
