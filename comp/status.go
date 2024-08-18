package comp

// Status 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/status
//
// @version 6.7.0
type Status Schema

// NewStatus 创建一个新的 Status 实例
func NewStatus() Status {
	return Status{}.set("type", "status")
}

func (s Status) set(key string, value interface{}) Status {
	s[key] = value
	return s
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (s Status) ClassName(value string) Status {
	return s.set("className", value)
}

// Disabled 是否禁用
func (s Status) Disabled(value bool) Status {
	return s.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (s Status) DisabledOn(value string) Status {
	return s.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (s Status) EditorSetting(value string) Status {
	return s.set("editorSetting", value)
}

// Hidden 是否隐藏
func (s Status) Hidden(value bool) Status {
	return s.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (s Status) HiddenOn(value string) Status {
	return s.set("hiddenOn", value)
}

// Id 组件唯一 id，主要用于日志采集
func (s Status) Id(value string) Status {
	return s.set("id", value)
}

// LabelMap 文字映射关系
func (s Status) LabelMap(value string) Status {
	return s.set("labelMap", value)
}

// Map 状态图标映射关系
func (s Status) Map(value string) Status {
	return s.set("map", value)
}

// OnEvent 事件动作配置
func (s Status) OnEvent(value string) Status {
	return s.set("onEvent", value)
}

// Placeholder 占位符
func (s Status) Placeholder(value string) Status {
	return s.set("placeholder", value)
}

// Source 新版配置映射源的字段 可以兼容新版icon并且配置颜色 2.8.0 新增
func (s Status) Source(value string) Status {
	return s.set("source", value)
}

// Static 是否静态展示
func (s Status) Static(value bool) Status {
	return s.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (s Status) StaticClassName(value string) Status {
	return s.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (s Status) StaticInputClassName(value string) Status {
	return s.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (s Status) StaticLabelClassName(value string) Status {
	return s.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (s Status) StaticOn(value string) Status {
	return s.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (s Status) StaticPlaceholder(value string) Status {
	return s.set("staticPlaceholder", value)
}

// StaticSchema
func (s Status) StaticSchema(value string) Status {
	return s.set("staticSchema", value)
}

// Style 组件样式
func (s Status) Style(value string) Status {
	return s.set("style", value)
}

// TestIdBuilder
func (s Status) TestIdBuilder(value string) Status {
	return s.set("testIdBuilder", value)
}

// Testid
func (s Status) Testid(value string) Status {
	return s.set("testid", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (s Status) UseMobileUI(value bool) Status {
	return s.set("useMobileUI", value)
}

// Visible 是否显示
func (s Status) Visible(value bool) Status {
	return s.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (s Status) VisibleOn(value string) Status {
	return s.set("visibleOn", value)
}
