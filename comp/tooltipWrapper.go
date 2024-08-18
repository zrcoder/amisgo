package comp

// TooltipWrapper
//
// @author slowlyo
// @version 6.7.0
type TooltipWrapper Schema

// NewTooltipWrapper 创建一个新的 TooltipWrapper 实例
func NewTooltipWrapper() TooltipWrapper {
	return TooltipWrapper{}.set("type", "tooltip-wrapper")
}

func (tw TooltipWrapper) set(key string, value interface{}) TooltipWrapper {
	tw[key] = value
	return tw
}

// Body 内容区域 (内容区域)
func (tw TooltipWrapper) Body(value ...interface{}) TooltipWrapper {
	return tw.set("body", value)
}

// ClassName 内容区CSS类名
func (tw TooltipWrapper) ClassName(value string) TooltipWrapper {
	return tw.set("className", value)
}

// Content 文字提示内容，兼容 tooltip，但建议通过 content 来实现提示内容
func (tw TooltipWrapper) Content(value string) TooltipWrapper {
	return tw.set("content", value)
}

// Disabled 是否禁用提示
func (tw TooltipWrapper) Disabled(value bool) TooltipWrapper {
	return tw.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (tw TooltipWrapper) DisabledOn(value string) TooltipWrapper {
	return tw.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (tw TooltipWrapper) EditorSetting(value string) TooltipWrapper {
	return tw.set("editorSetting", value)
}

// Enterable 是否可以移入浮层中, 默认true
func (tw TooltipWrapper) Enterable(value bool) TooltipWrapper {
	return tw.set("enterable", value)
}

// Hidden 是否隐藏
func (tw TooltipWrapper) Hidden(value bool) TooltipWrapper {
	return tw.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (tw TooltipWrapper) HiddenOn(value string) TooltipWrapper {
	return tw.set("hiddenOn", value)
}

// Id 组件唯一 id，主要用于日志采集
func (tw TooltipWrapper) Id(value string) TooltipWrapper {
	return tw.set("id", value)
}

// Inline 内容区是否内联显示，默认为false
func (tw TooltipWrapper) Inline(value bool) TooltipWrapper {
	return tw.set("inline", value)
}

// MouseEnterDelay 浮层延迟显示时间, 单位 ms
func (tw TooltipWrapper) MouseEnterDelay(value string) TooltipWrapper {
	return tw.set("mouseEnterDelay", value)
}

// MouseLeaveDelay 浮层延迟隐藏时间, 单位 ms
func (tw TooltipWrapper) MouseLeaveDelay(value string) TooltipWrapper {
	return tw.set("mouseLeaveDelay", value)
}

// Offset 浮层位置相对偏移量
func (tw TooltipWrapper) Offset(value string) TooltipWrapper {
	return tw.set("offset", value)
}

// OnEvent 事件动作配置
func (tw TooltipWrapper) OnEvent(value string) TooltipWrapper {
	return tw.set("onEvent", value)
}

// Placement 文字提示浮层出现位置，默认为top 可选值: top | right | bottom | left
func (tw TooltipWrapper) Placement(value string) TooltipWrapper {
	return tw.set("placement", value)
}

// RootClose 是否点击非内容区域关闭提示，默认为true
func (tw TooltipWrapper) RootClose(value bool) TooltipWrapper {
	return tw.set("rootClose", value)
}

// ShowArrow 是否展示浮层指向箭头
func (tw TooltipWrapper) ShowArrow(value bool) TooltipWrapper {
	return tw.set("showArrow", value)
}

// Static 是否静态展示
func (tw TooltipWrapper) Static(value bool) TooltipWrapper {
	return tw.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (tw TooltipWrapper) StaticClassName(value string) TooltipWrapper {
	return tw.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (tw TooltipWrapper) StaticInputClassName(value string) TooltipWrapper {
	return tw.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (tw TooltipWrapper) StaticLabelClassName(value string) TooltipWrapper {
	return tw.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (tw TooltipWrapper) StaticOn(value string) TooltipWrapper {
	return tw.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (tw TooltipWrapper) StaticPlaceholder(value string) TooltipWrapper {
	return tw.set("staticPlaceholder", value)
}

// StaticSchema
func (tw TooltipWrapper) StaticSchema(value string) TooltipWrapper {
	return tw.set("staticSchema", value)
}

// Style 内容区自定义样式
func (tw TooltipWrapper) Style(value string) TooltipWrapper {
	return tw.set("style", value)
}

// TestIdBuilder
func (tw TooltipWrapper) TestIdBuilder(value string) TooltipWrapper {
	return tw.set("testIdBuilder", value)
}

// Testid
func (tw TooltipWrapper) Testid(value string) TooltipWrapper {
	return tw.set("testid", value)
}

// Title 文字提示标题
func (tw TooltipWrapper) Title(value string) TooltipWrapper {
	return tw.set("title", value)
}

// Tooltip
func (tw TooltipWrapper) Tooltip(value string) TooltipWrapper {
	return tw.set("tooltip", value)
}

// TooltipClassName 文字提示浮层CSS类名
func (tw TooltipWrapper) TooltipClassName(value string) TooltipWrapper {
	return tw.set("tooltipClassName", value)
}

// TooltipStyle 自定义提示浮层样式
func (tw TooltipWrapper) TooltipStyle(value string) TooltipWrapper {
	return tw.set("tooltipStyle", value)
}

// TooltipTheme 主题样式， 默认为light 可选值: light | dark
func (tw TooltipWrapper) TooltipTheme(value string) TooltipWrapper {
	return tw.set("tooltipTheme", value)
}

// Trigger 浮层触发方式，默认为hover
func (tw TooltipWrapper) Trigger(value string) TooltipWrapper {
	return tw.set("trigger", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (tw TooltipWrapper) UseMobileUI(value bool) TooltipWrapper {
	return tw.set("useMobileUI", value)
}

// Visible 是否显示
func (tw TooltipWrapper) Visible(value bool) TooltipWrapper {
	return tw.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (tw TooltipWrapper) VisibleOn(value string) TooltipWrapper {
	return tw.set("visibleOn", value)
}

// WrapperComponent 内容区包裹标签
func (tw TooltipWrapper) WrapperComponent(value string) TooltipWrapper {
	return tw.set("wrapperComponent", value)
}
