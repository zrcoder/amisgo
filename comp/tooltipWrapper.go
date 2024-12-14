package comp

// tooltipWrapper

type tooltipWrapper schema

// TooltipWrapper 创建一个新的 TooltipWrapper 实例
func TooltipWrapper() tooltipWrapper {
	return tooltipWrapper{}.set("type", "tooltip-wrapper")
}

func (tw tooltipWrapper) set(key string, value any) tooltipWrapper {
	tw[key] = value
	return tw
}

// Body 内容区域 (内容区域)
func (tw tooltipWrapper) Body(value ...any) tooltipWrapper {
	return tw.set("body", value)
}

// ClassName 内容区CSS类名
func (tw tooltipWrapper) ClassName(value string) tooltipWrapper {
	return tw.set("className", value)
}

// Content 文字提示内容，兼容 tooltip，但建议通过 content 来实现提示内容
func (tw tooltipWrapper) Content(value string) tooltipWrapper {
	return tw.set("content", value)
}

// Disabled 是否禁用提示
func (tw tooltipWrapper) Disabled(value bool) tooltipWrapper {
	return tw.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (tw tooltipWrapper) DisabledOn(value string) tooltipWrapper {
	return tw.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (tw tooltipWrapper) EditorSetting(value string) tooltipWrapper {
	return tw.set("editorSetting", value)
}

// Enterable 是否可以移入浮层中, 默认true
func (tw tooltipWrapper) Enterable(value bool) tooltipWrapper {
	return tw.set("enterable", value)
}

// Hidden 是否隐藏
func (tw tooltipWrapper) Hidden(value bool) tooltipWrapper {
	return tw.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (tw tooltipWrapper) HiddenOn(value string) tooltipWrapper {
	return tw.set("hiddenOn", value)
}

// Id 组件唯一 id，主要用于日志采集
func (tw tooltipWrapper) ID(value string) tooltipWrapper {
	return tw.set("id", value)
}

// Inline 内容区是否内联显示，默认为false
func (tw tooltipWrapper) Inline(value bool) tooltipWrapper {
	return tw.set("inline", value)
}

// MouseEnterDelay 浮层延迟显示时间, 单位 ms
func (tw tooltipWrapper) MouseEnterDelay(value string) tooltipWrapper {
	return tw.set("mouseEnterDelay", value)
}

// MouseLeaveDelay 浮层延迟隐藏时间, 单位 ms
func (tw tooltipWrapper) MouseLeaveDelay(value string) tooltipWrapper {
	return tw.set("mouseLeaveDelay", value)
}

// Offset 浮层位置相对偏移量
func (tw tooltipWrapper) Offset(value string) tooltipWrapper {
	return tw.set("offset", value)
}

// OnEvent 事件动作配置
func (tw tooltipWrapper) OnEvent(value any) tooltipWrapper {
	return tw.set("onEvent", value)
}

// Placement 文字提示浮层出现位置，默认为top 可选值: top | right | bottom | left
func (tw tooltipWrapper) Placement(value string) tooltipWrapper {
	return tw.set("placement", value)
}

// RootClose 是否点击非内容区域关闭提示，默认为true
func (tw tooltipWrapper) RootClose(value bool) tooltipWrapper {
	return tw.set("rootClose", value)
}

// ShowArrow 是否展示浮层指向箭头
func (tw tooltipWrapper) ShowArrow(value bool) tooltipWrapper {
	return tw.set("showArrow", value)
}

// Static 是否静态展示
func (tw tooltipWrapper) Static(value bool) tooltipWrapper {
	return tw.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (tw tooltipWrapper) StaticClassName(value string) tooltipWrapper {
	return tw.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (tw tooltipWrapper) StaticInputClassName(value string) tooltipWrapper {
	return tw.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (tw tooltipWrapper) StaticLabelClassName(value string) tooltipWrapper {
	return tw.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (tw tooltipWrapper) StaticOn(value string) tooltipWrapper {
	return tw.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (tw tooltipWrapper) StaticPlaceholder(value string) tooltipWrapper {
	return tw.set("staticPlaceholder", value)
}

// StaticSchema
func (tw tooltipWrapper) StaticSchema(value string) tooltipWrapper {
	return tw.set("staticSchema", value)
}

// Style 内容区自定义样式
func (tw tooltipWrapper) Style(value any) tooltipWrapper {
	return tw.set("style", value)
}

// TestIdBuilder
func (tw tooltipWrapper) TestIdBuilder(value string) tooltipWrapper {
	return tw.set("testIdBuilder", value)
}

// Testid
func (tw tooltipWrapper) Testid(value string) tooltipWrapper {
	return tw.set("testid", value)
}

// Title 文字提示标题
func (tw tooltipWrapper) Title(value any) tooltipWrapper {
	return tw.set("title", value)
}

// Tooltip
func (tw tooltipWrapper) Tooltip(value string) tooltipWrapper {
	return tw.set("tooltip", value)
}

// TooltipClassName 文字提示浮层CSS类名
func (tw tooltipWrapper) TooltipClassName(value string) tooltipWrapper {
	return tw.set("tooltipClassName", value)
}

// TooltipStyle 自定义提示浮层样式
func (tw tooltipWrapper) TooltipStyle(value any) tooltipWrapper {
	return tw.set("tooltipStyle", value)
}

// TooltipTheme 主题样式， 默认为light 可选值: light | dark
func (tw tooltipWrapper) TooltipTheme(value string) tooltipWrapper {
	return tw.set("tooltipTheme", value)
}

// Trigger 浮层触发方式，默认为hover
func (tw tooltipWrapper) Trigger(value string) tooltipWrapper {
	return tw.set("trigger", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (tw tooltipWrapper) UseMobileUI(value bool) tooltipWrapper {
	return tw.set("useMobileUI", value)
}

// Visible 是否显示
func (tw tooltipWrapper) Visible(value bool) tooltipWrapper {
	return tw.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (tw tooltipWrapper) VisibleOn(value string) tooltipWrapper {
	return tw.set("visibleOn", value)
}

// WrapperComponent 内容区包裹标签
func (tw tooltipWrapper) WrapperComponent(value string) tooltipWrapper {
	return tw.set("wrapperComponent", value)
}
