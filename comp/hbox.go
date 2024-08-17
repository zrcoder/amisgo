package comp

// HBox 水平布局渲染器
type HBox Schema

// NewHBox 创建一个新的 HBox 实例
func NewHBox() HBox {
	return make(HBox).set("type", "hbox")
}

func (h HBox) set(key string, value interface{}) HBox {
	h[key] = value
	return h
}

// align 水平对齐方式 可选值: left | right | between | center
func (h HBox) Align(value string) HBox {
	return h.set("align", value)
}

// className 容器 css 类名
func (h HBox) ClassName(value string) HBox {
	return h.set("className", value)
}

// columns
func (h HBox) Columns(value string) HBox {
	return h.set("columns", value)
}

// disabled 是否禁用
func (h HBox) Disabled(value bool) HBox {
	return h.set("disabled", value)
}

// disabledOn 是否禁用表达式
func (h HBox) DisabledOn(value string) HBox {
	return h.set("disabledOn", value)
}

// editorSetting 编辑器配置
func (h HBox) EditorSetting(value string) HBox {
	return h.set("editorSetting", value)
}

// gap 水平间距 可选值: xs | sm | base | none | md | lg
func (h HBox) Gap(value string) HBox {
	return h.set("gap", value)
}

// hidden 是否隐藏
func (h HBox) Hidden(value bool) HBox {
	return h.set("hidden", value)
}

// hiddenOn 是否隐藏表达式
func (h HBox) HiddenOn(value string) HBox {
	return h.set("hiddenOn", value)
}

// id 组件唯一 id
func (h HBox) Id(value string) HBox {
	return h.set("id", value)
}

// onEvent 事件动作配置
func (h HBox) OnEvent(value string) HBox {
	return h.set("onEvent", value)
}

// static 是否静态展示
func (h HBox) Static(value bool) HBox {
	return h.set("static", value)
}

// staticClassName 静态展示表单项类名
func (h HBox) StaticClassName(value string) HBox {
	return h.set("staticClassName", value)
}

// staticInputClassName 静态展示表单项Value类名
func (h HBox) StaticInputClassName(value string) HBox {
	return h.set("staticInputClassName", value)
}

// staticLabelClassName 静态展示表单项Label类名
func (h HBox) StaticLabelClassName(value string) HBox {
	return h.set("staticLabelClassName", value)
}

// staticOn 是否静态展示表达式
func (h HBox) StaticOn(value string) HBox {
	return h.set("staticOn", value)
}

// staticPlaceholder 静态展示空值占位
func (h HBox) StaticPlaceholder(value string) HBox {
	return h.set("staticPlaceholder", value)
}

// staticSchema
func (h HBox) StaticSchema(value string) HBox {
	return h.set("staticSchema", value)
}

// style 组件样式
func (h HBox) Style(value string) HBox {
	return h.set("style", value)
}

// subFormHorizontal 水平排版左右宽度占比
func (h HBox) SubFormHorizontal(value string) HBox {
	return h.set("subFormHorizontal", value)
}

// subFormMode 子表单项展示方式
func (h HBox) SubFormMode(value string) HBox {
	return h.set("subFormMode", value)
}

// testIdBuilder
func (h HBox) TestIdBuilder(value string) HBox {
	return h.set("testIdBuilder", value)
}

// testid
func (h HBox) Testid(value string) HBox {
	return h.set("testid", value)
}

// type 指定为 each 展示类型
func (h HBox) Type(value string) HBox {
	return h.set("type", value)
}

// useMobileUI 关闭移动端样式
func (h HBox) UseMobileUI(value bool) HBox {
	return h.set("useMobileUI", value)
}

// valign 垂直对齐方式
func (h HBox) Valign(value string) HBox {
	return h.set("valign", value)
}

// visible 是否显示
func (h HBox) Visible(value bool) HBox {
	return h.set("visible", value)
}

// visibleOn 是否显示表达式
func (h HBox) VisibleOn(value string) HBox {
	return h.set("visibleOn", value)
}
