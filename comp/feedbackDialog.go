package comp

// feedbackDialog 表示反馈对话框
type feedbackDialog schema

// FeedbackDialog 创建一个新的 FeedbackDialog 实例
func FeedbackDialog() feedbackDialog {
	return make(feedbackDialog)
}

func (f feedbackDialog) set(key string, value any) feedbackDialog {
	f[key] = value
	return f
}

// Actions 默认不用填写，自动会创建确认和取消按钮
func (f feedbackDialog) Actions(value string) feedbackDialog {
	return f.set("actions", value)
}

// Body 内容区域
func (f feedbackDialog) Body(value ...any) feedbackDialog {
	return f.set("body", value)
}

// BodyClassName 配置 Body 容器 className
func (f feedbackDialog) BodyClassName(value string) feedbackDialog {
	return f.set("bodyClassName", value)
}

// ClassName 容器 css 类名
func (f feedbackDialog) ClassName(value string) feedbackDialog {
	return f.set("className", value)
}

// CloseOnEsc 是否支持按 ESC 关闭 Dialog
func (f feedbackDialog) CloseOnEsc(value bool) feedbackDialog {
	return f.set("closeOnEsc", value)
}

// CloseOnOutside 是否支持点其它区域关闭 Dialog
func (f feedbackDialog) CloseOnOutside(value bool) feedbackDialog {
	return f.set("closeOnOutside", value)
}

// Confirm 影响自动生成的按钮，如果自己配置了按钮这个配置无效
func (f feedbackDialog) Confirm(value bool) feedbackDialog {
	return f.set("confirm", value)
}

// Data 数据映射
func (f feedbackDialog) Data(value Data) feedbackDialog {
	return f.set("data", value)
}

// DialogType 弹框类型 confirm 确认弹框
func (f feedbackDialog) DialogType(value string) feedbackDialog {
	return f.set("dialogType", value)
}

// Disabled 是否禁用
func (f feedbackDialog) Disabled(value bool) feedbackDialog {
	return f.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (f feedbackDialog) DisabledOn(value string) feedbackDialog {
	return f.set("disabledOn", value)
}

// Draggable 可拖拽
func (f feedbackDialog) Draggable(value bool) feedbackDialog {
	return f.set("draggable", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (f feedbackDialog) EditorSetting(value string) feedbackDialog {
	return f.set("editorSetting", value)
}

// Footer
func (f feedbackDialog) Footer(value string) feedbackDialog {
	return f.set("footer", value)
}

// Header
func (f feedbackDialog) Header(value string) feedbackDialog {
	return f.set("header", value)
}

// HeaderClassName css类名，配置字符串，或者对象
func (f feedbackDialog) HeaderClassName(value string) feedbackDialog {
	return f.set("headerClassName", value)
}

// Height Dialog 高度
func (f feedbackDialog) Height(value string) feedbackDialog {
	return f.set("height", value)
}

// Hidden 是否隐藏
func (f feedbackDialog) Hidden(value bool) feedbackDialog {
	return f.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (f feedbackDialog) HiddenOn(value string) feedbackDialog {
	return f.set("hiddenOn", value)
}

// Id 组件唯一 id，主要用于日志采集
func (f feedbackDialog) Id(value string) feedbackDialog {
	return f.set("id", value)
}

// InputParams 弹窗参数说明，值格式为 JSONSchema
func (f feedbackDialog) InputParams(value string) feedbackDialog {
	return f.set("inputParams", value)
}

// Name 组件名字，这个名字可以用来定位，用于组件通信
func (f feedbackDialog) Name(value string) feedbackDialog {
	return f.set("name", value)
}

// OnEvent 事件动作配置
func (f feedbackDialog) OnEvent(value any) feedbackDialog {
	return f.set("onEvent", value)
}

// Overlay 是否显示蒙层
func (f feedbackDialog) Overlay(value bool) feedbackDialog {
	return f.set("overlay", value)
}

// ShowCloseButton 是否显示关闭按钮
func (f feedbackDialog) ShowCloseButton(value bool) feedbackDialog {
	return f.set("showCloseButton", value)
}

// ShowErrorMsg 是否显示错误信息
func (f feedbackDialog) ShowErrorMsg(value bool) feedbackDialog {
	return f.set("showErrorMsg", value)
}

// ShowLoading 是否显示 spinner
func (f feedbackDialog) ShowLoading(value bool) feedbackDialog {
	return f.set("showLoading", value)
}

// Size Dialog 大小
func (f feedbackDialog) Size(value string) feedbackDialog {
	return f.set("size", value)
}

// SkipRestOnCancel feedback 弹框取消是否中断后续操作
func (f feedbackDialog) SkipRestOnCancel(value bool) feedbackDialog {
	return f.set("skipRestOnCancel", value)
}

// SkipRestOnConfirm feedback 弹框确认是否中断后续操作
func (f feedbackDialog) SkipRestOnConfirm(value bool) feedbackDialog {
	return f.set("skipRestOnConfirm", value)
}

// Static 是否静态展示
func (f feedbackDialog) Static(value bool) feedbackDialog {
	return f.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (f feedbackDialog) StaticClassName(value string) feedbackDialog {
	return f.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (f feedbackDialog) StaticInputClassName(value string) feedbackDialog {
	return f.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (f feedbackDialog) StaticLabelClassName(value string) feedbackDialog {
	return f.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (f feedbackDialog) StaticOn(value string) feedbackDialog {
	return f.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (f feedbackDialog) StaticPlaceholder(value string) feedbackDialog {
	return f.set("staticPlaceholder", value)
}

// StaticSchema
func (f feedbackDialog) StaticSchema(value string) feedbackDialog {
	return f.set("staticSchema", value)
}

// Style 组件样式
func (f feedbackDialog) Style(value any) feedbackDialog {
	return f.set("style", value)
}

// TestIdBuilder
func (f feedbackDialog) TestIdBuilder(value string) feedbackDialog {
	return f.set("testIdBuilder", value)
}

// Testid
func (f feedbackDialog) Testid(value string) feedbackDialog {
	return f.set("testid", value)
}

// Title 请通过配置 title 设置标题
func (f feedbackDialog) Title(value any) feedbackDialog {
	return f.set("title", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (f feedbackDialog) UseMobileUI(value bool) feedbackDialog {
	return f.set("useMobileUI", value)
}

// Visible 是否显示
func (f feedbackDialog) Visible(value bool) feedbackDialog {
	return f.set("visible", value)
}

// VisibleOn 可以用来配置 feedback 的出现条件
func (f feedbackDialog) VisibleOn(value string) feedbackDialog {
	return f.set("visibleOn", value)
}

// Width Dialog 宽度
func (f feedbackDialog) Width(value string) feedbackDialog {
	return f.set("width", value)
}
