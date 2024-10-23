package comp

// drawer 抽出式弹框。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/drawer
type drawer schema

func Drawer() drawer {
	return make(drawer).set("type", "drawer")
}

func (d drawer) set(key string, value any) drawer {
	d[key] = value
	return d
}

// Actions 默认不用填写，自动会创建确认和取消按钮。
func (d drawer) Actions(value ...any) drawer {
	return d.set("actions", value)
}

// Body 内容区域 (内容区域)
func (d drawer) Body(value ...any) drawer {
	return d.set("body", value)
}

// BodyClassName 配置 Body 容器 className (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (d drawer) BodyClassName(value string) drawer {
	return d.set("bodyClassName", value)
}

// ClassName 配置 外层 className (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (d drawer) ClassName(value string) drawer {
	return d.set("className", value)
}

// CloseOnEsc 是否支持按 ESC 关闭 Dialog
func (d drawer) CloseOnEsc(value bool) drawer {
	return d.set("closeOnEsc", value)
}

// CloseOnOutside 点击外部的时候是否关闭弹框。
func (d drawer) CloseOnOutside(value bool) drawer {
	return d.set("closeOnOutside", value)
}

// Confirm 影响自动生成的按钮，如果自己配置了按钮这个配置无效。
func (d drawer) Confirm(value bool) drawer {
	return d.set("confirm", value)
}

// Data 数据映射
func (d drawer) Data(value Data) drawer {
	return d.set("data", value)
}

// Disabled 是否禁用
func (d drawer) Disabled(value bool) drawer {
	return d.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (d drawer) DisabledOn(value string) drawer {
	return d.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (d drawer) EditorSetting(value string) drawer {
	return d.set("editorSetting", value)
}

// Footer 底部 (底部)
func (d drawer) Footer(value string) drawer {
	return d.set("footer", value)
}

// FooterClassName 配置 头部 容器 className (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (d drawer) FooterClassName(value string) drawer {
	return d.set("footerClassName", value)
}

// Header 头部 (头部)
func (d drawer) Header(value string) drawer {
	return d.set("header", value)
}

// HeaderClassName 配置 头部 容器 className (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (d drawer) HeaderClassName(value string) drawer {
	return d.set("headerClassName", value)
}

// Height 抽屉的高度 （string | number, 当position为top | bottom时生效）
func (d drawer) Height(value any) drawer {
	return d.set("height", value)
}

// Hidden 是否隐藏
func (d drawer) Hidden(value bool) drawer {
	return d.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (d drawer) HiddenOn(value string) drawer {
	return d.set("hiddenOn", value)
}

// ID 组件唯一 id，主要用于日志采集
func (d drawer) ID(value string) drawer {
	return d.set("id", value)
}

// InputParams 弹窗参数说明，值格式为 JSONSchema。
func (d drawer) InputParams(value string) drawer {
	return d.set("inputParams", value)
}

// Name 组件名字，这个名字可以用来定位，用于组件通信
func (d drawer) Name(value string) drawer {
	return d.set("name", value)
}

// OnEvent 事件动作配置
func (d drawer) OnEvent(value any) drawer {
	return d.set("onEvent", value)
}

// Overlay 是否显示蒙层
func (d drawer) Overlay(value bool) drawer {
	return d.set("overlay", value)
}

// Position 从什么位置弹出 可选值: left | right | top | bottom
func (d drawer) Position(value string) drawer {
	return d.set("position", value)
}

// Resizable 是否可以拖动弹窗大小
func (d drawer) Resizable(value bool) drawer {
	return d.set("resizable", value)
}

// ShowCloseButton 是否展示关闭按钮 当值为false时，默认开启closeOnOutside
func (d drawer) ShowCloseButton(value bool) drawer {
	return d.set("showCloseButton", value)
}

// ShowErrorMsg 是否显示错误信息
func (d drawer) ShowErrorMsg(value bool) drawer {
	return d.set("showErrorMsg", value)
}

// Size 大小 可选值: xs、sm、md、lg、xl
func (d drawer) Size(value string) drawer {
	return d.set("size", value)
}

// Static 是否静态展示
func (d drawer) Static(value bool) drawer {
	return d.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (d drawer) StaticClassName(value string) drawer {
	return d.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (d drawer) StaticInputClassName(value string) drawer {
	return d.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (d drawer) StaticLabelClassName(value string) drawer {
	return d.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (d drawer) StaticOn(value string) drawer {
	return d.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (d drawer) StaticPlaceholder(value string) drawer {
	return d.set("staticPlaceholder", value)
}

// StaticSchema 静态展示 schema
func (d drawer) StaticSchema(value string) drawer {
	return d.set("staticSchema", value)
}

// Style 组件样式
func (d drawer) Style(value any) drawer {
	return d.set("style", value)
}

// TestIdBuilder 测试 ID 构建器
func (d drawer) TestIdBuilder(value string) drawer {
	return d.set("testIdBuilder", value)
}

// Testid 测试 id
func (d drawer) Testid(value string) drawer {
	return d.set("testid", value)
}

// Title 请通过配置 title 设置标题 (请通过配置 title 设置标题)
func (d drawer) Title(value any) drawer {
	return d.set("title", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (d drawer) UseMobileUI(value bool) drawer {
	return d.set("useMobileUI", value)
}

// Visible 是否显示
func (d drawer) Visible(value bool) drawer {
	return d.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (d drawer) VisibleOn(value string) drawer {
	return d.set("visibleOn", value)
}

// Width 抽屉的宽度 （string | number, 当position为 left | right 时生效）
func (d drawer) Width(value any) drawer {
	return d.set("width", value)
}
