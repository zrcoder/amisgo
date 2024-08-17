package comp

// Drawer 抽出式弹框。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/drawer
type Drawer BaseRenderer

// NewDrawer 创建一个新的 Drawer 实例，并设置默认的 type
func NewDrawer() Drawer {
	d := Drawer(make(BaseRenderer))
	return d.set("type", "drawer")
}

func (d Drawer) set(key string, value interface{}) Drawer {
	d[key] = value
	return d
}

// Actions 默认不用填写，自动会创建确认和取消按钮。
func (d Drawer) Actions(value string) Drawer {
	return d.set("actions", value)
}

// Body 内容区域 (内容区域)
func (d Drawer) Body(value ...interface{}) Drawer {
	return d.set("body", value)
}

// BodyClassName 配置 Body 容器 className (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (d Drawer) BodyClassName(value string) Drawer {
	return d.set("bodyClassName", value)
}

// ClassName 配置 外层 className (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (d Drawer) ClassName(value string) Drawer {
	return d.set("className", value)
}

// CloseOnEsc 是否支持按 ESC 关闭 Dialog
func (d Drawer) CloseOnEsc(value bool) Drawer {
	return d.set("closeOnEsc", value)
}

// CloseOnOutside 点击外部的时候是否关闭弹框。
func (d Drawer) CloseOnOutside(value bool) Drawer {
	return d.set("closeOnOutside", value)
}

// Confirm 影响自动生成的按钮，如果自己配置了按钮这个配置无效。
func (d Drawer) Confirm(value bool) Drawer {
	return d.set("confirm", value)
}

// Data 数据映射
func (d Drawer) Data(value string) Drawer {
	return d.set("data", value)
}

// Disabled 是否禁用
func (d Drawer) Disabled(value bool) Drawer {
	return d.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (d Drawer) DisabledOn(value string) Drawer {
	return d.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (d Drawer) EditorSetting(value string) Drawer {
	return d.set("editorSetting", value)
}

// Footer 底部 (底部)
func (d Drawer) Footer(value string) Drawer {
	return d.set("footer", value)
}

// FooterClassName 配置 头部 容器 className (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (d Drawer) FooterClassName(value string) Drawer {
	return d.set("footerClassName", value)
}

// Header 头部 (头部)
func (d Drawer) Header(value string) Drawer {
	return d.set("header", value)
}

// HeaderClassName 配置 头部 容器 className (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (d Drawer) HeaderClassName(value string) Drawer {
	return d.set("headerClassName", value)
}

// Height 抽屉的高度 （当position为top | bottom时生效）
func (d Drawer) Height(value string) Drawer {
	return d.set("height", value)
}

// Hidden 是否隐藏
func (d Drawer) Hidden(value bool) Drawer {
	return d.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (d Drawer) HiddenOn(value string) Drawer {
	return d.set("hiddenOn", value)
}

// ID 组件唯一 id，主要用于日志采集
func (d Drawer) ID(value string) Drawer {
	return d.set("id", value)
}

// InputParams 弹窗参数说明，值格式为 JSONSchema。
func (d Drawer) InputParams(value string) Drawer {
	return d.set("inputParams", value)
}

// Name 组件名字，这个名字可以用来定位，用于组件通信
func (d Drawer) Name(value string) Drawer {
	return d.set("name", value)
}

// OnEvent 事件动作配置
func (d Drawer) OnEvent(value string) Drawer {
	return d.set("onEvent", value)
}

// Overlay 是否显示蒙层
func (d Drawer) Overlay(value bool) Drawer {
	return d.set("overlay", value)
}

// Position 从什么位置弹出 可选值: left | right | top | bottom
func (d Drawer) Position(value string) Drawer {
	return d.set("position", value)
}

// Resizable 是否可以拖动弹窗大小
func (d Drawer) Resizable(value bool) Drawer {
	return d.set("resizable", value)
}

// ShowCloseButton 是否展示关闭按钮 当值为false时，默认开启closeOnOutside
func (d Drawer) ShowCloseButton(value bool) Drawer {
	return d.set("showCloseButton", value)
}

// ShowErrorMsg 是否显示错误信息
func (d Drawer) ShowErrorMsg(value bool) Drawer {
	return d.set("showErrorMsg", value)
}

// Size Dialog 大小 可选值: xs | sm | md | lg | full
func (d Drawer) Size(value string) Drawer {
	return d.set("size", value)
}

// Static 是否静态展示
func (d Drawer) Static(value bool) Drawer {
	return d.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (d Drawer) StaticClassName(value string) Drawer {
	return d.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (d Drawer) StaticInputClassName(value string) Drawer {
	return d.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (d Drawer) StaticLabelClassName(value string) Drawer {
	return d.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (d Drawer) StaticOn(value string) Drawer {
	return d.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (d Drawer) StaticPlaceholder(value string) Drawer {
	return d.set("staticPlaceholder", value)
}

// StaticSchema 静态展示 schema
func (d Drawer) StaticSchema(value string) Drawer {
	return d.set("staticSchema", value)
}

// Style 组件样式
func (d Drawer) Style(value string) Drawer {
	return d.set("style", value)
}

// TestIdBuilder 测试 ID 构建器
func (d Drawer) TestIdBuilder(value string) Drawer {
	return d.set("testIdBuilder", value)
}

// Testid 测试 id
func (d Drawer) Testid(value string) Drawer {
	return d.set("testid", value)
}

// Title 请通过配置 title 设置标题 (请通过配置 title 设置标题)
func (d Drawer) Title(value string) Drawer {
	return d.set("title", value)
}

// Type 指定为 drawer
func (d Drawer) Type(value string) Drawer {
	return d.set("type", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (d Drawer) UseMobileUI(value bool) Drawer {
	return d.set("useMobileUI", value)
}

// Visible 是否显示
func (d Drawer) Visible(value bool) Drawer {
	return d.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (d Drawer) VisibleOn(value string) Drawer {
	return d.set("visibleOn", value)
}

// Width 抽屉的宽度 （当position为left | right时生效）
func (d Drawer) Width(value string) Drawer {
	return d.set("width", value)
}
