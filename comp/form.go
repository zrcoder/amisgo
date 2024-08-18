package comp

// form 表单渲染器。说明：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/index
//
// @version 6.7.0
type form schema

// Form 创建一个新的 Form 实例
func Form() form {
	return make(form).set("type", "form")
}

func (f form) set(key string, value interface{}) form {
	f[key] = value
	return f
}

// Reload 配置表单重新加载
func (f form) Reload(value string) form {
	return f.set("reload", value)
}

// Mode 配置当前表单项展示模式 可选值: normal | inline | horizontal
func (f form) Mode(value string) form {
	return f.set("mode", value)
}

// Body Form 内容
func (f form) Body(value ...interface{}) form {
	return f.set("body", value)
}

// ResetAfterSubmit 提交后重置表单
func (f form) ResetAfterSubmit(value bool) form {
	return f.set("resetAfterSubmit", value)
}

// Rules 组合校验规则，选填
func (f form) Rules(value string) form {
	return f.set("rules", value)
}

// SilentPolling 是否静默拉取
func (f form) SilentPolling(value bool) form {
	return f.set("silentPolling", value)
}

// Static 展示态时的className
func (f form) Static(value bool) form {
	return f.set("static", value)
}

// StaticClassName css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" }
func (f form) StaticClassName(value string) form {
	return f.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (f form) StaticInputClassName(value string) form {
	return f.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (f form) StaticLabelClassName(value string) form {
	return f.set("staticLabelClassName", value)
}

// StaticOn 表达式，语法 `data.xxx > 5`。
func (f form) StaticOn(value string) form {
	return f.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (f form) StaticPlaceholder(value string) form {
	return f.set("staticPlaceholder", value)
}

// StaticSchema
func (f form) StaticSchema(value string) form {
	return f.set("staticSchema", value)
}

// StopAutoRefreshWhen 配置停止轮询的条件
func (f form) StopAutoRefreshWhen(value string) form {
	return f.set("stopAutoRefreshWhen", value)
}

// Style 组件样式
func (f form) Style(value string) form {
	return f.set("style", value)
}

// SubmitOnChange 修改的时候是否直接提交表单
func (f form) SubmitOnChange(value bool) form {
	return f.set("submitOnChange", value)
}

// SubmitOnInit 表单初始先提交一次，联动的时候有用
func (f form) SubmitOnInit(value bool) form {
	return f.set("submitOnInit", value)
}

// SubmitText 默认的提交按钮名称，如果设置成空，则可以把默认按钮去掉
func (f form) SubmitText(value string) form {
	return f.set("submitText", value)
}

// Tabs 表单选项卡
func (f form) Tabs(value string) form {
	return f.set("tabs", value)
}

// Target 默认表单提交自己会通过发送 api 保存数据，但是也可以设定另外一个 form 的 name 值，或者另外一个 `CRUD` 模型的 name 值。如果 target 目标是一个 `Form`，则目标 `Form` 会重新触发 `initApi` 和 `schemaApi`，api 可以拿到当前 form 数据。如果目标是一个 `CRUD` 模型，则目标模型会重新触发搜索，参数为当前 Form 数据
func (f form) Target(value string) form {
	return f.set("target", value)
}

// TestIdBuilder
func (f form) TestIdBuilder(value string) form {
	return f.set("testIdBuilder", value)
}

// Testid
func (f form) Testid(value string) form {
	return f.set("testid", value)
}

// Title 表单标题
func (f form) Title(value string) form {
	return f.set("title", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (f form) UseMobileUI(value bool) form {
	return f.set("useMobileUI", value)
}

// Visible 是否显示
func (f form) Visible(value bool) form {
	return f.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (f form) VisibleOn(value string) form {
	return f.set("visibleOn", value)
}

// WrapWithPanel 是否用 panel 包裹起来
func (f form) WrapWithPanel(value bool) form {
	return f.set("wrapWithPanel", value)
}
