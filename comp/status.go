package comp

// status 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/status

type status schema

// Status 创建一个新的 Status 实例
func Status() status {
	return status{}.set("type", "status")
}

func (s status) set(key string, value any) status {
	s[key] = value
	return s
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (s status) ClassName(value string) status {
	return s.set("className", value)
}

// Disabled 是否禁用
func (s status) Disabled(value bool) status {
	return s.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (s status) DisabledOn(value string) status {
	return s.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (s status) EditorSetting(value string) status {
	return s.set("editorSetting", value)
}

// Hidden 是否隐藏
func (s status) Hidden(value bool) status {
	return s.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (s status) HiddenOn(value string) status {
	return s.set("hiddenOn", value)
}

// Id 组件唯一 id，主要用于日志采集
func (s status) ID(value string) status {
	return s.set("id", value)
}

// LabelMap 文字映射关系
func (s status) LabelMap(value string) status {
	return s.set("labelMap", value)
}

// Map 状态图标映射关系
func (s status) Map(value string) status {
	return s.set("map", value)
}

// OnEvent 事件动作配置
func (s status) OnEvent(value any) status {
	return s.set("onEvent", value)
}

// Placeholder 占位符
func (s status) Placeholder(value string) status {
	return s.set("placeholder", value)
}

// Source 新版配置映射源的字段 可以兼容新版icon并且配置颜色 2.8.0 新增
func (s status) Source(value string) status {
	return s.set("source", value)
}

// Static 是否静态展示
func (s status) Static(value bool) status {
	return s.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (s status) StaticClassName(value string) status {
	return s.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (s status) StaticInputClassName(value string) status {
	return s.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (s status) StaticLabelClassName(value string) status {
	return s.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (s status) StaticOn(value string) status {
	return s.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (s status) StaticPlaceholder(value string) status {
	return s.set("staticPlaceholder", value)
}

// StaticSchema
func (s status) StaticSchema(value string) status {
	return s.set("staticSchema", value)
}

// Style 组件样式
func (s status) Style(value any) status {
	return s.set("style", value)
}

// TestIdBuilder
func (s status) TestIdBuilder(value string) status {
	return s.set("testIdBuilder", value)
}

// Testid
func (s status) Testid(value string) status {
	return s.set("testid", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (s status) UseMobileUI(value bool) status {
	return s.set("useMobileUI", value)
}

// Visible 是否显示
func (s status) Visible(value bool) status {
	return s.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (s status) VisibleOn(value string) status {
	return s.set("visibleOn", value)
}
