package comp

// SearchBox
//
// @author  slowlyo
// @version 6.7.0
type SearchBox Schema

// NewSearchBox 创建一个新的 SearchBox 实例
func NewSearchBox() SearchBox {
	return SearchBox{}.set("type", "search-box")
}

func (s SearchBox) set(key string, value interface{}) SearchBox {
	s[key] = value
	return s
}

// ClassName 外层 css 类名
func (s SearchBox) ClassName(value string) SearchBox {
	return s.set("className", value)
}

// ClearAndSubmit 是否开启清空内容后立即重新搜索
func (s SearchBox) ClearAndSubmit(value bool) SearchBox {
	return s.set("clearAndSubmit", value)
}

// Clearable 是否可清除
func (s SearchBox) Clearable(value bool) SearchBox {
	return s.set("clearable", value)
}

// Disabled 是否禁用
func (s SearchBox) Disabled(value bool) SearchBox {
	return s.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (s SearchBox) DisabledOn(value string) SearchBox {
	return s.set("disabledOn", value)
}

// EditorSetting 编辑器配置
func (s SearchBox) EditorSetting(value string) SearchBox {
	return s.set("editorSetting", value)
}

// Enhance 是否为加强样式
func (s SearchBox) Enhance(value bool) SearchBox {
	return s.set("enhance", value)
}

// Hidden 是否隐藏
func (s SearchBox) Hidden(value bool) SearchBox {
	return s.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (s SearchBox) HiddenOn(value string) SearchBox {
	return s.set("hiddenOn", value)
}

// Id 组件唯一 id
func (s SearchBox) Id(value string) SearchBox {
	return s.set("id", value)
}

// Loading 是否处于加载状态
func (s SearchBox) Loading(value bool) SearchBox {
	return s.set("loading", value)
}

// Mini 是否为 Mini 样式
func (s SearchBox) Mini(value bool) SearchBox {
	return s.set("mini", value)
}

// Name 关键字名字
func (s SearchBox) Name(value string) SearchBox {
	return s.set("name", value)
}

// OnEvent 事件动作配置
func (s SearchBox) OnEvent(value string) SearchBox {
	return s.set("onEvent", value)
}

// Placeholder 占位符
func (s SearchBox) Placeholder(value string) SearchBox {
	return s.set("placeholder", value)
}

// SearchImediately 是否立马搜索
func (s SearchBox) SearchImediately(value bool) SearchBox {
	return s.set("searchImediately", value)
}

// Static 是否静态展示
func (s SearchBox) Static(value bool) SearchBox {
	return s.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (s SearchBox) StaticClassName(value string) SearchBox {
	return s.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (s SearchBox) StaticInputClassName(value string) SearchBox {
	return s.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (s SearchBox) StaticLabelClassName(value string) SearchBox {
	return s.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (s SearchBox) StaticOn(value string) SearchBox {
	return s.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (s SearchBox) StaticPlaceholder(value string) SearchBox {
	return s.set("staticPlaceholder", value)
}

// StaticSchema
func (s SearchBox) StaticSchema(value string) SearchBox {
	return s.set("staticSchema", value)
}

// Style 组件样式
func (s SearchBox) Style(value string) SearchBox {
	return s.set("style", value)
}

// TestIdBuilder
func (s SearchBox) TestIdBuilder(value string) SearchBox {
	return s.set("testIdBuilder", value)
}

// Testid
func (s SearchBox) Testid(value string) SearchBox {
	return s.set("testid", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (s SearchBox) UseMobileUI(value bool) SearchBox {
	return s.set("useMobileUI", value)
}

// Visible 是否显示
func (s SearchBox) Visible(value bool) SearchBox {
	return s.set("visible", value)
}

// VisibleOn 是否显示表达式
func (s SearchBox) VisibleOn(value string) SearchBox {
	return s.set("visibleOn", value)
}
