package comp

// hBox 水平布局渲染器
type hBox schema

// HBox 创建一个新的 HBox 实例
func HBox() hBox {
	return make(hBox).set("type", "hbox")
}

func (h hBox) set(key string, value interface{}) hBox {
	h[key] = value
	return h
}

// align 水平对齐方式 可选值: left | right | between | center
func (h hBox) Align(value string) hBox {
	return h.set("align", value)
}

// className 容器 css 类名
func (h hBox) ClassName(value string) hBox {
	return h.set("className", value)
}

// columns
func (h hBox) Columns(value string) hBox {
	return h.set("columns", value)
}

// disabled 是否禁用
func (h hBox) Disabled(value bool) hBox {
	return h.set("disabled", value)
}

// disabledOn 是否禁用表达式
func (h hBox) DisabledOn(value string) hBox {
	return h.set("disabledOn", value)
}

// editorSetting 编辑器配置
func (h hBox) EditorSetting(value string) hBox {
	return h.set("editorSetting", value)
}

// gap 水平间距 可选值: xs | sm | base | none | md | lg
func (h hBox) Gap(value string) hBox {
	return h.set("gap", value)
}

// hidden 是否隐藏
func (h hBox) Hidden(value bool) hBox {
	return h.set("hidden", value)
}

// hiddenOn 是否隐藏表达式
func (h hBox) HiddenOn(value string) hBox {
	return h.set("hiddenOn", value)
}

// id 组件唯一 id
func (h hBox) Id(value string) hBox {
	return h.set("id", value)
}

// onEvent 事件动作配置
func (h hBox) OnEvent(value string) hBox {
	return h.set("onEvent", value)
}

// static 是否静态展示
func (h hBox) Static(value bool) hBox {
	return h.set("static", value)
}

// staticClassName 静态展示表单项类名
func (h hBox) StaticClassName(value string) hBox {
	return h.set("staticClassName", value)
}

// staticInputClassName 静态展示表单项Value类名
func (h hBox) StaticInputClassName(value string) hBox {
	return h.set("staticInputClassName", value)
}

// staticLabelClassName 静态展示表单项Label类名
func (h hBox) StaticLabelClassName(value string) hBox {
	return h.set("staticLabelClassName", value)
}

// staticOn 是否静态展示表达式
func (h hBox) StaticOn(value string) hBox {
	return h.set("staticOn", value)
}

// staticPlaceholder 静态展示空值占位
func (h hBox) StaticPlaceholder(value string) hBox {
	return h.set("staticPlaceholder", value)
}

// staticSchema
func (h hBox) StaticSchema(value string) hBox {
	return h.set("staticSchema", value)
}

// style 组件样式
func (h hBox) Style(value string) hBox {
	return h.set("style", value)
}

// subFormHorizontal 水平排版左右宽度占比
func (h hBox) SubFormHorizontal(value string) hBox {
	return h.set("subFormHorizontal", value)
}

// subFormMode 子表单项展示方式
func (h hBox) SubFormMode(value string) hBox {
	return h.set("subFormMode", value)
}

// testIdBuilder
func (h hBox) TestIdBuilder(value string) hBox {
	return h.set("testIdBuilder", value)
}

// testid
func (h hBox) Testid(value string) hBox {
	return h.set("testid", value)
}

// useMobileUI 关闭移动端样式
func (h hBox) UseMobileUI(value bool) hBox {
	return h.set("useMobileUI", value)
}

// valign 垂直对齐方式
func (h hBox) Valign(value string) hBox {
	return h.set("valign", value)
}

// visible 是否显示
func (h hBox) Visible(value bool) hBox {
	return h.set("visible", value)
}

// visibleOn 是否显示表达式
func (h hBox) VisibleOn(value string) hBox {
	return h.set("visibleOn", value)
}
