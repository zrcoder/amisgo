package comp

// Form 表单渲染器。说明：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/index
//
// @version 6.7.0
type Form Schema

// NewForm 创建一个新的 Form 实例
func NewForm() Form {
	return make(Form).set("type", "form")
}

func (f Form) set(key string, value interface{}) Form {
	f[key] = value
	return f
}

// Reload 配置表单重新加载
func (f Form) Reload(value string) Form {
	return f.set("reload", value)
}

// ResetAfterSubmit 提交后重置表单
func (f Form) ResetAfterSubmit(value bool) Form {
	return f.set("resetAfterSubmit", value)
}

// Rules 组合校验规则，选填
func (f Form) Rules(value string) Form {
	return f.set("rules", value)
}

// SilentPolling 是否静默拉取
func (f Form) SilentPolling(value bool) Form {
	return f.set("silentPolling", value)
}

// Static 展示态时的className
func (f Form) Static(value bool) Form {
	return f.set("static", value)
}

// StaticClassName css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" }
func (f Form) StaticClassName(value string) Form {
	return f.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (f Form) StaticInputClassName(value string) Form {
	return f.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (f Form) StaticLabelClassName(value string) Form {
	return f.set("staticLabelClassName", value)
}

// StaticOn 表达式，语法 `data.xxx > 5`。
func (f Form) StaticOn(value string) Form {
	return f.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (f Form) StaticPlaceholder(value string) Form {
	return f.set("staticPlaceholder", value)
}

// StaticSchema
func (f Form) StaticSchema(value string) Form {
	return f.set("staticSchema", value)
}

// StopAutoRefreshWhen 配置停止轮询的条件
func (f Form) StopAutoRefreshWhen(value string) Form {
	return f.set("stopAutoRefreshWhen", value)
}

// Style 组件样式
func (f Form) Style(value string) Form {
	return f.set("style", value)
}

// SubmitOnChange 修改的时候是否直接提交表单
func (f Form) SubmitOnChange(value bool) Form {
	return f.set("submitOnChange", value)
}

// SubmitOnInit 表单初始先提交一次，联动的时候有用
func (f Form) SubmitOnInit(value bool) Form {
	return f.set("submitOnInit", value)
}

// SubmitText 默认的提交按钮名称，如果设置成空，则可以把默认按钮去掉
func (f Form) SubmitText(value string) Form {
	return f.set("submitText", value)
}

// Tabs 表单选项卡
func (f Form) Tabs(value string) Form {
	return f.set("tabs", value)
}

// Target 默认表单提交自己会通过发送 api 保存数据，但是也可以设定另外一个 form 的 name 值，或者另外一个 `CRUD` 模型的 name 值。如果 target 目标是一个 `Form`，则目标 `Form` 会重新触发 `initApi` 和 `schemaApi`，api 可以拿到当前 form 数据。如果目标是一个 `CRUD` 模型，则目标模型会重新触发搜索，参数为当前 Form 数据
func (f Form) Target(value string) Form {
	return f.set("target", value)
}

// TestIdBuilder
func (f Form) TestIdBuilder(value string) Form {
	return f.set("testIdBuilder", value)
}

// Testid
func (f Form) Testid(value string) Form {
	return f.set("testid", value)
}

// Title 表单标题
func (f Form) Title(value string) Form {
	return f.set("title", value)
}

// Type 指定为表单渲染器
func (f Form) Type(value string) Form {
	return f.set("type", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (f Form) UseMobileUI(value bool) Form {
	return f.set("useMobileUI", value)
}

// Visible 是否显示
func (f Form) Visible(value bool) Form {
	return f.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (f Form) VisibleOn(value string) Form {
	return f.set("visibleOn", value)
}

// WrapWithPanel 是否用 panel 包裹起来
func (f Form) WrapWithPanel(value bool) Form {
	return f.set("wrapWithPanel", value)
}
