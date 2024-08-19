package comp

// searchBox

// @version 6.7.0
type searchBox schema

// SearchBox 创建一个新的 SearchBox 实例
func SearchBox() searchBox {
	return searchBox{}.set("type", "search-box")
}

func (s searchBox) set(key string, value any) searchBox {
	s[key] = value
	return s
}

// ClassName 外层 css 类名
func (s searchBox) ClassName(value string) searchBox {
	return s.set("className", value)
}

// ClearAndSubmit 是否开启清空内容后立即重新搜索
func (s searchBox) ClearAndSubmit(value bool) searchBox {
	return s.set("clearAndSubmit", value)
}

// Clearable 是否可清除
func (s searchBox) Clearable(value bool) searchBox {
	return s.set("clearable", value)
}

// Disabled 是否禁用
func (s searchBox) Disabled(value bool) searchBox {
	return s.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (s searchBox) DisabledOn(value string) searchBox {
	return s.set("disabledOn", value)
}

// EditorSetting 编辑器配置
func (s searchBox) EditorSetting(value string) searchBox {
	return s.set("editorSetting", value)
}

// Enhance 是否为加强样式
func (s searchBox) Enhance(value bool) searchBox {
	return s.set("enhance", value)
}

// Hidden 是否隐藏
func (s searchBox) Hidden(value bool) searchBox {
	return s.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (s searchBox) HiddenOn(value string) searchBox {
	return s.set("hiddenOn", value)
}

// Id 组件唯一 id
func (s searchBox) Id(value string) searchBox {
	return s.set("id", value)
}

// Loading 是否处于加载状态
func (s searchBox) Loading(value bool) searchBox {
	return s.set("loading", value)
}

// Mini 是否为 Mini 样式
func (s searchBox) Mini(value bool) searchBox {
	return s.set("mini", value)
}

// Name 关键字名字
func (s searchBox) Name(value string) searchBox {
	return s.set("name", value)
}

// OnEvent 事件动作配置
func (s searchBox) OnEvent(value any) searchBox {
	return s.set("onEvent", value)
}

// Placeholder 占位符
func (s searchBox) Placeholder(value string) searchBox {
	return s.set("placeholder", value)
}

// SearchImediately 是否立马搜索
func (s searchBox) SearchImediately(value bool) searchBox {
	return s.set("searchImediately", value)
}

// Static 是否静态展示
func (s searchBox) Static(value bool) searchBox {
	return s.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (s searchBox) StaticClassName(value string) searchBox {
	return s.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (s searchBox) StaticInputClassName(value string) searchBox {
	return s.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (s searchBox) StaticLabelClassName(value string) searchBox {
	return s.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (s searchBox) StaticOn(value string) searchBox {
	return s.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (s searchBox) StaticPlaceholder(value string) searchBox {
	return s.set("staticPlaceholder", value)
}

// StaticSchema
func (s searchBox) StaticSchema(value string) searchBox {
	return s.set("staticSchema", value)
}

// Style 组件样式
func (s searchBox) Style(value any) searchBox {
	return s.set("style", value)
}

// TestIdBuilder
func (s searchBox) TestIdBuilder(value string) searchBox {
	return s.set("testIdBuilder", value)
}

// Testid
func (s searchBox) Testid(value string) searchBox {
	return s.set("testid", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (s searchBox) UseMobileUI(value bool) searchBox {
	return s.set("useMobileUI", value)
}

// Visible 是否显示
func (s searchBox) Visible(value bool) searchBox {
	return s.set("visible", value)
}

// VisibleOn 是否显示表达式
func (s searchBox) VisibleOn(value string) searchBox {
	return s.set("visibleOn", value)
}
