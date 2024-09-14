package comp

// spinner

type spinner schema

// Spinner 创建一个 Spinner 实例
func Spinner() spinner {
	return spinner{}.set("type", "spinner")
}

func (s spinner) set(key string, value any) spinner {
	s[key] = value
	return s
}

// Body 作为容器使用时内容 (作为容器使用时内容)
func (s spinner) Body(value ...any) spinner {
	return s.set("body", value)
}

// ClassName 自定义spinner的class
func (s spinner) ClassName(value string) spinner {
	return s.set("className", value)
}

// Delay 延迟显示
func (s spinner) Delay(value string) spinner {
	return s.set("delay", value)
}

// Disabled 是否禁用
func (s spinner) Disabled(value bool) spinner {
	return s.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (s spinner) DisabledOn(value string) spinner {
	return s.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (s spinner) EditorSetting(value string) spinner {
	return s.set("editorSetting", value)
}

// Hidden 是否隐藏
func (s spinner) Hidden(value bool) spinner {
	return s.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (s spinner) HiddenOn(value string) spinner {
	return s.set("hiddenOn", value)
}

// Icon 自定义icon
func (s spinner) Icon(value string) spinner {
	return s.set("icon", value)
}

// Id 组件唯一 id，主要用于日志采集
func (s spinner) Id(value string) spinner {
	return s.set("id", value)
}

// LoadingConfig
func (s spinner) LoadingConfig(value string) spinner {
	return s.set("loadingConfig", value)
}

// Mode
func (s spinner) Mode(value string) spinner {
	return s.set("mode", value)
}

// OnEvent 事件动作配置
func (s spinner) OnEvent(value any) spinner {
	return s.set("onEvent", value)
}

// Overlay 是否显示遮罩层
func (s spinner) Overlay(value bool) spinner {
	return s.set("overlay", value)
}

// Show 控制Spinner显示与隐藏
func (s spinner) Show(value bool) spinner {
	return s.set("show", value)
}

// Size spinner Icon 大小 可选值: sm | lg |
func (s spinner) Size(value string) spinner {
	return s.set("size", value)
}

// SpinnerClassName spin图标位置包裹元素的自定义class
func (s spinner) SpinnerClassName(value string) spinner {
	return s.set("spinnerClassName", value)
}

// SpinnerWrapClassName 作为容器使用时最外层元素的class
func (s spinner) SpinnerWrapClassName(value string) spinner {
	return s.set("spinnerWrapClassName", value)
}

// Static 是否静态展示
func (s spinner) Static(value bool) spinner {
	return s.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (s spinner) StaticClassName(value string) spinner {
	return s.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (s spinner) StaticInputClassName(value string) spinner {
	return s.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (s spinner) StaticLabelClassName(value string) spinner {
	return s.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (s spinner) StaticOn(value string) spinner {
	return s.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (s spinner) StaticPlaceholder(value string) spinner {
	return s.set("staticPlaceholder", value)
}

// StaticSchema
func (s spinner) StaticSchema(value string) spinner {
	return s.set("staticSchema", value)
}

// Style 组件样式
func (s spinner) Style(value any) spinner {
	return s.set("style", value)
}

// TestIdBuilder
func (s spinner) TestIdBuilder(value string) spinner {
	return s.set("testIdBuilder", value)
}

// Testid
func (s spinner) Testid(value string) spinner {
	return s.set("testid", value)
}

// Tip spinner文案
func (s spinner) Tip(value string) spinner {
	return s.set("tip", value)
}

// TipPlacement spinner文案位置 可选值: top | right | bottom | left
func (s spinner) TipPlacement(value string) spinner {
	return s.set("tipPlacement", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (s spinner) UseMobileUI(value bool) spinner {
	return s.set("useMobileUI", value)
}

// Visible 是否显示
func (s spinner) Visible(value bool) spinner {
	return s.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (s spinner) VisibleOn(value string) spinner {
	return s.set("visibleOn", value)
}
