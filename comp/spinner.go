package comp

// Spinner
//
// @version 6.7.0
type Spinner Schema

// NewSpinner 创建一个新的 Spinner 实例
func NewSpinner() Spinner {
	return Spinner{}.set("type", "spinner")
}

func (s Spinner) set(key string, value interface{}) Spinner {
	s[key] = value
	return s
}

// Body 作为容器使用时内容 (作为容器使用时内容)
func (s Spinner) Body(value ...interface{}) Spinner {
	return s.set("body", value)
}

// ClassName 自定义spinner的class
func (s Spinner) ClassName(value string) Spinner {
	return s.set("className", value)
}

// Delay 延迟显示
func (s Spinner) Delay(value string) Spinner {
	return s.set("delay", value)
}

// Disabled 是否禁用
func (s Spinner) Disabled(value bool) Spinner {
	return s.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (s Spinner) DisabledOn(value string) Spinner {
	return s.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (s Spinner) EditorSetting(value string) Spinner {
	return s.set("editorSetting", value)
}

// Hidden 是否隐藏
func (s Spinner) Hidden(value bool) Spinner {
	return s.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (s Spinner) HiddenOn(value string) Spinner {
	return s.set("hiddenOn", value)
}

// Icon 自定义icon
func (s Spinner) Icon(value string) Spinner {
	return s.set("icon", value)
}

// Id 组件唯一 id，主要用于日志采集
func (s Spinner) Id(value string) Spinner {
	return s.set("id", value)
}

// LoadingConfig
func (s Spinner) LoadingConfig(value string) Spinner {
	return s.set("loadingConfig", value)
}

// Mode
func (s Spinner) Mode(value string) Spinner {
	return s.set("mode", value)
}

// OnEvent 事件动作配置
func (s Spinner) OnEvent(value string) Spinner {
	return s.set("onEvent", value)
}

// Overlay 是否显示遮罩层
func (s Spinner) Overlay(value bool) Spinner {
	return s.set("overlay", value)
}

// Show 控制Spinner显示与隐藏
func (s Spinner) Show(value bool) Spinner {
	return s.set("show", value)
}

// Size spinner Icon 大小 可选值: sm | lg |
func (s Spinner) Size(value string) Spinner {
	return s.set("size", value)
}

// SpinnerClassName spin图标位置包裹元素的自定义class
func (s Spinner) SpinnerClassName(value string) Spinner {
	return s.set("spinnerClassName", value)
}

// SpinnerWrapClassName 作为容器使用时最外层元素的class
func (s Spinner) SpinnerWrapClassName(value string) Spinner {
	return s.set("spinnerWrapClassName", value)
}

// Static 是否静态展示
func (s Spinner) Static(value bool) Spinner {
	return s.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (s Spinner) StaticClassName(value string) Spinner {
	return s.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (s Spinner) StaticInputClassName(value string) Spinner {
	return s.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (s Spinner) StaticLabelClassName(value string) Spinner {
	return s.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (s Spinner) StaticOn(value string) Spinner {
	return s.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (s Spinner) StaticPlaceholder(value string) Spinner {
	return s.set("staticPlaceholder", value)
}

// StaticSchema
func (s Spinner) StaticSchema(value string) Spinner {
	return s.set("staticSchema", value)
}

// Style 组件样式
func (s Spinner) Style(value string) Spinner {
	return s.set("style", value)
}

// TestIdBuilder
func (s Spinner) TestIdBuilder(value string) Spinner {
	return s.set("testIdBuilder", value)
}

// Testid
func (s Spinner) Testid(value string) Spinner {
	return s.set("testid", value)
}

// Tip spinner文案
func (s Spinner) Tip(value string) Spinner {
	return s.set("tip", value)
}

// TipPlacement spinner文案位置 可选值: top | right | bottom | left
func (s Spinner) TipPlacement(value string) Spinner {
	return s.set("tipPlacement", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (s Spinner) UseMobileUI(value bool) Spinner {
	return s.set("useMobileUI", value)
}

// Visible 是否显示
func (s Spinner) Visible(value bool) Spinner {
	return s.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (s Spinner) VisibleOn(value string) Spinner {
	return s.set("visibleOn", value)
}
