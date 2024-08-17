package comp

// FeedbackDialog 表示反馈对话框
type FeedbackDialog Schema

// NewFeedbackDialog 创建一个新的 FeedbackDialog 实例
func NewFeedbackDialog() FeedbackDialog {
	return make(FeedbackDialog)
}

func (f FeedbackDialog) set(key string, value interface{}) FeedbackDialog {
	f[key] = value
	return f
}

// Actions 默认不用填写，自动会创建确认和取消按钮
func (f FeedbackDialog) Actions(value string) FeedbackDialog {
	return f.set("actions", value)
}

// Body 内容区域
func (f FeedbackDialog) Body(value ...interface{}) FeedbackDialog {
	return f.set("body", value)
}

// BodyClassName 配置 Body 容器 className
func (f FeedbackDialog) BodyClassName(value string) FeedbackDialog {
	return f.set("bodyClassName", value)
}

// ClassName 容器 css 类名
func (f FeedbackDialog) ClassName(value string) FeedbackDialog {
	return f.set("className", value)
}

// CloseOnEsc 是否支持按 ESC 关闭 Dialog
func (f FeedbackDialog) CloseOnEsc(value bool) FeedbackDialog {
	return f.set("closeOnEsc", value)
}

// CloseOnOutside 是否支持点其它区域关闭 Dialog
func (f FeedbackDialog) CloseOnOutside(value bool) FeedbackDialog {
	return f.set("closeOnOutside", value)
}

// Confirm 影响自动生成的按钮，如果自己配置了按钮这个配置无效
func (f FeedbackDialog) Confirm(value bool) FeedbackDialog {
	return f.set("confirm", value)
}

// Data 数据映射
func (f FeedbackDialog) Data(value string) FeedbackDialog {
	return f.set("data", value)
}

// DialogType 弹框类型 confirm 确认弹框
func (f FeedbackDialog) DialogType(value string) FeedbackDialog {
	return f.set("dialogType", value)
}

// Disabled 是否禁用
func (f FeedbackDialog) Disabled(value bool) FeedbackDialog {
	return f.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (f FeedbackDialog) DisabledOn(value string) FeedbackDialog {
	return f.set("disabledOn", value)
}

// Draggable 可拖拽
func (f FeedbackDialog) Draggable(value bool) FeedbackDialog {
	return f.set("draggable", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (f FeedbackDialog) EditorSetting(value string) FeedbackDialog {
	return f.set("editorSetting", value)
}

// Footer
func (f FeedbackDialog) Footer(value string) FeedbackDialog {
	return f.set("footer", value)
}

// Header
func (f FeedbackDialog) Header(value string) FeedbackDialog {
	return f.set("header", value)
}

// HeaderClassName css类名，配置字符串，或者对象
func (f FeedbackDialog) HeaderClassName(value string) FeedbackDialog {
	return f.set("headerClassName", value)
}

// Height Dialog 高度
func (f FeedbackDialog) Height(value string) FeedbackDialog {
	return f.set("height", value)
}

// Hidden 是否隐藏
func (f FeedbackDialog) Hidden(value bool) FeedbackDialog {
	return f.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (f FeedbackDialog) HiddenOn(value string) FeedbackDialog {
	return f.set("hiddenOn", value)
}

// Id 组件唯一 id，主要用于日志采集
func (f FeedbackDialog) Id(value string) FeedbackDialog {
	return f.set("id", value)
}

// InputParams 弹窗参数说明，值格式为 JSONSchema
func (f FeedbackDialog) InputParams(value string) FeedbackDialog {
	return f.set("inputParams", value)
}

// Name 组件名字，这个名字可以用来定位，用于组件通信
func (f FeedbackDialog) Name(value string) FeedbackDialog {
	return f.set("name", value)
}

// OnEvent 事件动作配置
func (f FeedbackDialog) OnEvent(value string) FeedbackDialog {
	return f.set("onEvent", value)
}

// Overlay 是否显示蒙层
func (f FeedbackDialog) Overlay(value bool) FeedbackDialog {
	return f.set("overlay", value)
}

// ShowCloseButton 是否显示关闭按钮
func (f FeedbackDialog) ShowCloseButton(value bool) FeedbackDialog {
	return f.set("showCloseButton", value)
}

// ShowErrorMsg 是否显示错误信息
func (f FeedbackDialog) ShowErrorMsg(value bool) FeedbackDialog {
	return f.set("showErrorMsg", value)
}

// ShowLoading 是否显示 spinner
func (f FeedbackDialog) ShowLoading(value bool) FeedbackDialog {
	return f.set("showLoading", value)
}

// Size Dialog 大小
func (f FeedbackDialog) Size(value string) FeedbackDialog {
	return f.set("size", value)
}

// SkipRestOnCancel feedback 弹框取消是否中断后续操作
func (f FeedbackDialog) SkipRestOnCancel(value bool) FeedbackDialog {
	return f.set("skipRestOnCancel", value)
}

// SkipRestOnConfirm feedback 弹框确认是否中断后续操作
func (f FeedbackDialog) SkipRestOnConfirm(value bool) FeedbackDialog {
	return f.set("skipRestOnConfirm", value)
}

// Static 是否静态展示
func (f FeedbackDialog) Static(value bool) FeedbackDialog {
	return f.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (f FeedbackDialog) StaticClassName(value string) FeedbackDialog {
	return f.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (f FeedbackDialog) StaticInputClassName(value string) FeedbackDialog {
	return f.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (f FeedbackDialog) StaticLabelClassName(value string) FeedbackDialog {
	return f.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (f FeedbackDialog) StaticOn(value string) FeedbackDialog {
	return f.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (f FeedbackDialog) StaticPlaceholder(value string) FeedbackDialog {
	return f.set("staticPlaceholder", value)
}

// StaticSchema
func (f FeedbackDialog) StaticSchema(value string) FeedbackDialog {
	return f.set("staticSchema", value)
}

// Style 组件样式
func (f FeedbackDialog) Style(value string) FeedbackDialog {
	return f.set("style", value)
}

// TestIdBuilder
func (f FeedbackDialog) TestIdBuilder(value string) FeedbackDialog {
	return f.set("testIdBuilder", value)
}

// Testid
func (f FeedbackDialog) Testid(value string) FeedbackDialog {
	return f.set("testid", value)
}

// Title 请通过配置 title 设置标题
func (f FeedbackDialog) Title(value string) FeedbackDialog {
	return f.set("title", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (f FeedbackDialog) UseMobileUI(value bool) FeedbackDialog {
	return f.set("useMobileUI", value)
}

// Visible 是否显示
func (f FeedbackDialog) Visible(value bool) FeedbackDialog {
	return f.set("visible", value)
}

// VisibleOn 可以用来配置 feedback 的出现条件
func (f FeedbackDialog) VisibleOn(value string) FeedbackDialog {
	return f.set("visibleOn", value)
}

// Width Dialog 宽度
func (f FeedbackDialog) Width(value string) FeedbackDialog {
	return f.set("width", value)
}
