package comp

type Dialog BaseRenderer

func NewDialog() Dialog {
	d := Dialog(make(BaseRenderer))
	d.set("type", "dialog")
	return d
}

func (d Dialog) set(key string, value interface{}) Dialog {
	d[key] = value
	return d
}

// Actions 默认不用填写，自动会创建确认和取消按钮。
func (d Dialog) Actions(value string) Dialog {
	return d.set("actions", value)
}

// Body 内容区域 (内容区域)
func (d Dialog) Body(value ...interface{}) Dialog {
	return d.set("body", value)
}

// BodyClassName 配置 Body 容器 className (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (d Dialog) BodyClassName(value string) Dialog {
	return d.set("bodyClassName", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (d Dialog) ClassName(value string) Dialog {
	return d.set("className", value)
}

// CloseOnEsc 是否支持按 ESC 关闭 Dialog
func (d Dialog) CloseOnEsc(value bool) Dialog {
	return d.set("closeOnEsc", value)
}

// CloseOnOutside 是否支持点其它区域关闭 Dialog
func (d Dialog) CloseOnOutside(value bool) Dialog {
	return d.set("closeOnOutside", value)
}

// Confirm 影响自动生成的按钮，如果自己配置了按钮这个配置无效。
func (d Dialog) Confirm(value bool) Dialog {
	return d.set("confirm", value)
}

// Data 数据映射
func (d Dialog) Data(value string) Dialog {
	return d.set("data", value)
}

// DialogType 弹框类型 confirm 确认弹框
func (d Dialog) DialogType(value string) Dialog {
	return d.set("dialogType", value)
}

// Disabled 是否禁用
func (d Dialog) Disabled(value bool) Dialog {
	return d.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (d Dialog) DisabledOn(value string) Dialog {
	return d.set("disabledOn", value)
}

// Draggable 可拖拽
func (d Dialog) Draggable(value bool) Dialog {
	return d.set("draggable", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (d Dialog) EditorSetting(value string) Dialog {
	return d.set("editorSetting", value)
}

// Footer
func (d Dialog) Footer(value string) Dialog {
	return d.set("footer", value)
}

// Header
func (d Dialog) Header(value string) Dialog {
	return d.set("header", value)
}

// HeaderClassName css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" }
func (d Dialog) HeaderClassName(value string) Dialog {
	return d.set("headerClassName", value)
}

// Height Dialog 高度
func (d Dialog) Height(value string) Dialog {
	return d.set("height", value)
}

// Hidden 是否隐藏
func (d Dialog) Hidden(value bool) Dialog {
	return d.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (d Dialog) HiddenOn(value string) Dialog {
	return d.set("hiddenOn", value)
}

// Id 组件唯一 id，主要用于日志采集
func (d Dialog) Id(value string) Dialog {
	return d.set("id", value)
}

// InputParams 弹窗参数说明，值格式为 JSONSchema。
func (d Dialog) InputParams(value string) Dialog {
	return d.set("inputParams", value)
}

// Name 组件名字，这个名字可以用来定位，用于组件通信
func (d Dialog) Name(value string) Dialog {
	return d.set("name", value)
}

// OnEvent 事件动作配置
func (d Dialog) OnEvent(value string) Dialog {
	return d.set("onEvent", value)
}

// Overlay 是否显示蒙层
func (d Dialog) Overlay(value bool) Dialog {
	return d.set("overlay", value)
}

// ShowCloseButton 是否显示关闭按钮
func (d Dialog) ShowCloseButton(value bool) Dialog {
	return d.set("showCloseButton", value)
}

// ShowErrorMsg 是否显示错误信息
func (d Dialog) ShowErrorMsg(value bool) Dialog {
	return d.set("showErrorMsg", value)
}

// ShowLoading 是否显示 spinner
func (d Dialog) ShowLoading(value bool) Dialog {
	return d.set("showLoading", value)
}

// Size Dialog 大小 可选值: xs | sm | md | lg | xl | full
func (d Dialog) Size(value string) Dialog {
	return d.set("size", value)
}

// Static 是否静态展示
func (d Dialog) Static(value bool) Dialog {
	return d.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (d Dialog) StaticClassName(value string) Dialog {
	return d.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (d Dialog) StaticInputClassName(value string) Dialog {
	return d.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (d Dialog) StaticLabelClassName(value string) Dialog {
	return d.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (d Dialog) StaticOn(value string) Dialog {
	return d.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (d Dialog) StaticPlaceholder(value string) Dialog {
	return d.set("staticPlaceholder", value)
}

// StaticSchema
func (d Dialog) StaticSchema(value string) Dialog {
	return d.set("staticSchema", value)
}

// Style 组件样式
func (d Dialog) Style(value string) Dialog {
	return d.set("style", value)
}

// TestIdBuilder
func (d Dialog) TestIdBuilder(value string) Dialog {
	return d.set("testIdBuilder", value)
}

// Testid
func (d Dialog) Testid(value string) Dialog {
	return d.set("testid", value)
}

// Title 请通过配置 title 设置标题 (请通过配置 title 设置标题)
func (d Dialog) Title(value string) Dialog {
	return d.set("title", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (d Dialog) UseMobileUI(value bool) Dialog {
	return d.set("useMobileUI", value)
}

// Visible 是否显示
func (d Dialog) Visible(value bool) Dialog {
	return d.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (d Dialog) VisibleOn(value string) Dialog {
	return d.set("visibleOn", value)
}

// Width Dialog 宽度
func (d Dialog) Width(value string) Dialog {
	return d.set("width", value)
}
