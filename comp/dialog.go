package comp

type dialog schema

func Dialog() dialog {
	return make(dialog).set("type", "dialog")
}

func (d dialog) set(key string, value interface{}) dialog {
	d[key] = value
	return d
}

// Actions 默认不用填写，自动会创建确认和取消按钮。
func (d dialog) Actions(value string) dialog {
	return d.set("actions", value)
}

// Body 内容区域 (内容区域)
func (d dialog) Body(value ...interface{}) dialog {
	return d.set("body", value)
}

// BodyClassName 配置 Body 容器 className (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (d dialog) BodyClassName(value string) dialog {
	return d.set("bodyClassName", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (d dialog) ClassName(value string) dialog {
	return d.set("className", value)
}

// CloseOnEsc 是否支持按 ESC 关闭 Dialog
func (d dialog) CloseOnEsc(value bool) dialog {
	return d.set("closeOnEsc", value)
}

// CloseOnOutside 是否支持点其它区域关闭 Dialog
func (d dialog) CloseOnOutside(value bool) dialog {
	return d.set("closeOnOutside", value)
}

// Confirm 影响自动生成的按钮，如果自己配置了按钮这个配置无效。
func (d dialog) Confirm(value bool) dialog {
	return d.set("confirm", value)
}

// Data 数据映射
func (d dialog) Data(value string) dialog {
	return d.set("data", value)
}

// DialogType 弹框类型 confirm 确认弹框
func (d dialog) DialogType(value string) dialog {
	return d.set("dialogType", value)
}

// Disabled 是否禁用
func (d dialog) Disabled(value bool) dialog {
	return d.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (d dialog) DisabledOn(value string) dialog {
	return d.set("disabledOn", value)
}

// Draggable 可拖拽
func (d dialog) Draggable(value bool) dialog {
	return d.set("draggable", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (d dialog) EditorSetting(value string) dialog {
	return d.set("editorSetting", value)
}

// Footer
func (d dialog) Footer(value string) dialog {
	return d.set("footer", value)
}

// Header
func (d dialog) Header(value string) dialog {
	return d.set("header", value)
}

// HeaderClassName css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" }
func (d dialog) HeaderClassName(value string) dialog {
	return d.set("headerClassName", value)
}

// Height Dialog 高度
func (d dialog) Height(value string) dialog {
	return d.set("height", value)
}

// Hidden 是否隐藏
func (d dialog) Hidden(value bool) dialog {
	return d.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (d dialog) HiddenOn(value string) dialog {
	return d.set("hiddenOn", value)
}

// Id 组件唯一 id，主要用于日志采集
func (d dialog) Id(value string) dialog {
	return d.set("id", value)
}

// InputParams 弹窗参数说明，值格式为 JSONSchema。
func (d dialog) InputParams(value string) dialog {
	return d.set("inputParams", value)
}

// Name 组件名字，这个名字可以用来定位，用于组件通信
func (d dialog) Name(value string) dialog {
	return d.set("name", value)
}

// OnEvent 事件动作配置
func (d dialog) OnEvent(value string) dialog {
	return d.set("onEvent", value)
}

// Overlay 是否显示蒙层
func (d dialog) Overlay(value bool) dialog {
	return d.set("overlay", value)
}

// ShowCloseButton 是否显示关闭按钮
func (d dialog) ShowCloseButton(value bool) dialog {
	return d.set("showCloseButton", value)
}

// ShowErrorMsg 是否显示错误信息
func (d dialog) ShowErrorMsg(value bool) dialog {
	return d.set("showErrorMsg", value)
}

// ShowLoading 是否显示 spinner
func (d dialog) ShowLoading(value bool) dialog {
	return d.set("showLoading", value)
}

// Size Dialog 大小 可选值: xs | sm | md | lg | xl | full
func (d dialog) Size(value string) dialog {
	return d.set("size", value)
}

// Static 是否静态展示
func (d dialog) Static(value bool) dialog {
	return d.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (d dialog) StaticClassName(value string) dialog {
	return d.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (d dialog) StaticInputClassName(value string) dialog {
	return d.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (d dialog) StaticLabelClassName(value string) dialog {
	return d.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (d dialog) StaticOn(value string) dialog {
	return d.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (d dialog) StaticPlaceholder(value string) dialog {
	return d.set("staticPlaceholder", value)
}

// StaticSchema
func (d dialog) StaticSchema(value string) dialog {
	return d.set("staticSchema", value)
}

// Style 组件样式
func (d dialog) Style(value string) dialog {
	return d.set("style", value)
}

// TestIdBuilder
func (d dialog) TestIdBuilder(value string) dialog {
	return d.set("testIdBuilder", value)
}

// Testid
func (d dialog) Testid(value string) dialog {
	return d.set("testid", value)
}

// Title 请通过配置 title 设置标题 (请通过配置 title 设置标题)
func (d dialog) Title(value string) dialog {
	return d.set("title", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (d dialog) UseMobileUI(value bool) dialog {
	return d.set("useMobileUI", value)
}

// Visible 是否显示
func (d dialog) Visible(value bool) dialog {
	return d.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (d dialog) VisibleOn(value string) dialog {
	return d.set("visibleOn", value)
}

// Width Dialog 宽度
func (d dialog) Width(value string) dialog {
	return d.set("width", value)
}
