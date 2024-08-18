package comp

// ListRenderer 列表展示控件
//
// @version 6.7.0
type ListRenderer schema

// NewListRenderer 创建一个新的 ListRenderer 实例
func NewListRenderer() ListRenderer {
	lr := make(ListRenderer)
	lr.set("type", "list")
	return lr
}

// set 设置字段值
func (lr ListRenderer) set(key string, value interface{}) ListRenderer {
	lr[key] = value
	return lr
}

// AffixFooter 设置是否固底
func (lr ListRenderer) AffixFooter(value bool) ListRenderer {
	return lr.set("affixFooter", value)
}

// AffixHeader 设置是否固顶
func (lr ListRenderer) AffixHeader(value bool) ListRenderer {
	return lr.set("affixHeader", value)
}

// CheckOnItemClick 设置点击列表单行时是否选择
func (lr ListRenderer) CheckOnItemClick(value bool) ListRenderer {
	return lr.set("checkOnItemClick", value)
}

// ClassName 设置容器 css 类名
// 支持字符串或对象配置
func (lr ListRenderer) ClassName(value string) ListRenderer {
	return lr.set("className", value)
}

// Disabled 设置是否禁用
func (lr ListRenderer) Disabled(value bool) ListRenderer {
	return lr.set("disabled", value)
}

// DisabledOn 设置是否禁用表达式
// 表达式语法 `data.xxx > 5`
func (lr ListRenderer) DisabledOn(value string) ListRenderer {
	return lr.set("disabledOn", value)
}

// EditorSetting 设置编辑器配置
func (lr ListRenderer) EditorSetting(value string) ListRenderer {
	return lr.set("editorSetting", value)
}

// Footer 设置底部区域
func (lr ListRenderer) Footer(value string) ListRenderer {
	return lr.set("footer", value)
}

// FooterClassName 设置底部区域类名
// 支持字符串或对象配置
func (lr ListRenderer) FooterClassName(value string) ListRenderer {
	return lr.set("footerClassName", value)
}

// Header 设置顶部区域
func (lr ListRenderer) Header(value string) ListRenderer {
	return lr.set("header", value)
}

// HeaderClassName 设置顶部区域类名
// 支持字符串或对象配置
func (lr ListRenderer) HeaderClassName(value string) ListRenderer {
	return lr.set("headerClassName", value)
}

// Hidden 设置是否隐藏
func (lr ListRenderer) Hidden(value bool) ListRenderer {
	return lr.set("hidden", value)
}

// HiddenOn 设置是否隐藏表达式
// 表达式语法 `data.xxx > 5`
func (lr ListRenderer) HiddenOn(value string) ListRenderer {
	return lr.set("hiddenOn", value)
}

// HideCheckToggler 设置是否隐藏勾选框
func (lr ListRenderer) HideCheckToggler(value bool) ListRenderer {
	return lr.set("hideCheckToggler", value)
}

// Id 设置组件唯一 id
func (lr ListRenderer) Id(value string) ListRenderer {
	return lr.set("id", value)
}

// ItemAction 设置点击列表项的行为
func (lr ListRenderer) ItemAction(value string) ListRenderer {
	return lr.set("itemAction", value)
}

// ItemCheckableOn 设置配置某项是否可以点选
// 表达式语法 `data.xxx > 5`
func (lr ListRenderer) ItemCheckableOn(value string) ListRenderer {
	return lr.set("itemCheckableOn", value)
}

// ItemDraggableOn 设置配置某项是否可拖拽排序
// 表达式语法 `data.xxx > 5`
func (lr ListRenderer) ItemDraggableOn(value string) ListRenderer {
	return lr.set("itemDraggableOn", value)
}

// ListItem 设置单条数据展示内容配置
func (lr ListRenderer) ListItem(value string) ListRenderer {
	return lr.set("listItem", value)
}

// OnEvent 设置事件动作配置
func (lr ListRenderer) OnEvent(value string) ListRenderer {
	return lr.set("onEvent", value)
}

// Placeholder 设置无数据提示
// 支持两种语法：1. `${xxx}` 或 `${xxx|upperCase}` 2. `<%= data.xxx %>`
func (lr ListRenderer) Placeholder(value string) ListRenderer {
	return lr.set("placeholder", value)
}

// ShowFooter 设置是否显示底部
func (lr ListRenderer) ShowFooter(value bool) ListRenderer {
	return lr.set("showFooter", value)
}

// ShowHeader 设置是否显示头部
func (lr ListRenderer) ShowHeader(value bool) ListRenderer {
	return lr.set("showHeader", value)
}

// Size 设置大小
// 可选值: sm | base
func (lr ListRenderer) Size(value string) ListRenderer {
	return lr.set("size", value)
}

// Source 设置数据源
func (lr ListRenderer) Source(value string) ListRenderer {
	return lr.set("source", value)
}

// Static 设置是否静态展示
func (lr ListRenderer) Static(value bool) ListRenderer {
	return lr.set("static", value)
}

// StaticClassName 设置静态展示表单项类名
// 支持字符串或对象配置
func (lr ListRenderer) StaticClassName(value string) ListRenderer {
	return lr.set("staticClassName", value)
}

// StaticInputClassName 设置静态展示表单项Value类名
// 支持字符串或对象配置
func (lr ListRenderer) StaticInputClassName(value string) ListRenderer {
	return lr.set("staticInputClassName", value)
}

// StaticLabelClassName 设置静态展示表单项Label类名
// 支持字符串或对象配置
func (lr ListRenderer) StaticLabelClassName(value string) ListRenderer {
	return lr.set("staticLabelClassName", value)
}

// StaticOn 设置是否静态展示表达式
// 表达式语法 `data.xxx > 5`
func (lr ListRenderer) StaticOn(value string) ListRenderer {
	return lr.set("staticOn", value)
}

// StaticPlaceholder 设置静态展示空值占位
func (lr ListRenderer) StaticPlaceholder(value string) ListRenderer {
	return lr.set("staticPlaceholder", value)
}

// StaticSchema 设置静态展示模式的 schema
func (lr ListRenderer) StaticSchema(value string) ListRenderer {
	return lr.set("staticSchema", value)
}

// Style 设置组件样式
func (lr ListRenderer) Style(value string) ListRenderer {
	return lr.set("style", value)
}

// TestIdBuilder 设置测试 ID 构建器
func (lr ListRenderer) TestIdBuilder(value string) ListRenderer {
	return lr.set("testIdBuilder", value)
}

// Testid 设置测试 ID
func (lr ListRenderer) Testid(value string) ListRenderer {
	return lr.set("testid", value)
}

// Title 设置标题
// 支持两种语法：1. `${xxx}` 或 `${xxx|upperCase}` 2. `<%= data.xxx %>`
func (lr ListRenderer) Title(value string) ListRenderer {
	return lr.set("title", value)
}

// Type 设置为 List 列表展示控件
// 可选值: list | static-list
func (lr ListRenderer) Type(value string) ListRenderer {
	return lr.set("type", value)
}

// UseMobileUI 设置是否使用移动端样式
func (lr ListRenderer) UseMobileUI(value bool) ListRenderer {
	return lr.set("useMobileUI", value)
}

// ValueField 设置可以用来作为值的字段
func (lr ListRenderer) ValueField(value string) ListRenderer {
	return lr.set("valueField", value)
}

// Visible 设置是否显示
func (lr ListRenderer) Visible(value bool) ListRenderer {
	return lr.set("visible", value)
}

// VisibleOn 设置是否显示表达式
// 表达式语法 `data.xxx > 5`
func (lr ListRenderer) VisibleOn(value string) ListRenderer {
	return lr.set("visibleOn", value)
}
