package comp

// ListItem 列表项

type ListItem schema

// NewListItem 创建一个新的 ListItem 实例
func NewListItem() ListItem {
	return make(ListItem)
}

// set 设置字段值
func (li ListItem) set(key string, value any) ListItem {
	li[key] = value
	return li
}

// Actions 设置操作
func (li ListItem) Actions(value string) ListItem {
	return li.set("actions", value)
}

// ActionsPosition 设置操作位置，默认在右侧，可以设置成左侧
// 可选值: left | right
func (li ListItem) ActionsPosition(value string) ListItem {
	return li.set("actionsPosition", value)
}

// Avatar 设置图片地址
func (li ListItem) Avatar(value string) ListItem {
	return li.set("avatar", value)
}

// Body 设置内容区域
func (li ListItem) Body(value ...any) ListItem {
	return li.set("body", value)
}

// ClassName 设置容器 css 类名
// 支持字符串或对象配置
func (li ListItem) ClassName(value string) ListItem {
	return li.set("className", value)
}

// Desc 设置描述
// 支持两种语法：1. `${xxx}` 或 `${xxx|upperCase}` 2. `<%= data.xxx %>`
func (li ListItem) Desc(value string) ListItem {
	return li.set("desc", value)
}

// Disabled 设置是否禁用
func (li ListItem) Disabled(value bool) ListItem {
	return li.set("disabled", value)
}

// DisabledOn 设置是否禁用表达式
// 表达式语法 `data.xxx > 5`
func (li ListItem) DisabledOn(value string) ListItem {
	return li.set("disabledOn", value)
}

// EditorSetting 设置编辑器配置
func (li ListItem) EditorSetting(value string) ListItem {
	return li.set("editorSetting", value)
}

// Hidden 设置是否隐藏
func (li ListItem) Hidden(value bool) ListItem {
	return li.set("hidden", value)
}

// HiddenOn 设置是否隐藏表达式
// 表达式语法 `data.xxx > 5`
func (li ListItem) HiddenOn(value string) ListItem {
	return li.set("hiddenOn", value)
}

// Id 设置组件唯一 id
func (li ListItem) Id(value string) ListItem {
	return li.set("id", value)
}

// OnEvent 设置事件动作配置
func (li ListItem) OnEvent(value any) ListItem {
	return li.set("onEvent", value)
}

// Remark 设置 tooltip 说明
func (li ListItem) Remark(value string) ListItem {
	return li.set("remark", value)
}

// Static 设置是否静态展示
func (li ListItem) Static(value bool) ListItem {
	return li.set("static", value)
}

// StaticClassName 设置静态展示表单项类名
// 支持字符串或对象配置
func (li ListItem) StaticClassName(value string) ListItem {
	return li.set("staticClassName", value)
}

// StaticInputClassName 设置静态展示表单项Value类名
// 支持字符串或对象配置
func (li ListItem) StaticInputClassName(value string) ListItem {
	return li.set("staticInputClassName", value)
}

// StaticLabelClassName 设置静态展示表单项Label类名
// 支持字符串或对象配置
func (li ListItem) StaticLabelClassName(value string) ListItem {
	return li.set("staticLabelClassName", value)
}

// StaticOn 设置是否静态展示表达式
// 表达式语法 `data.xxx > 5`
func (li ListItem) StaticOn(value string) ListItem {
	return li.set("staticOn", value)
}

// StaticPlaceholder 设置静态展示空值占位
func (li ListItem) StaticPlaceholder(value string) ListItem {
	return li.set("staticPlaceholder", value)
}

// StaticSchema 设置静态展示模式的 schema
func (li ListItem) StaticSchema(value string) ListItem {
	return li.set("staticSchema", value)
}

// Style 设置组件样式
func (li ListItem) Style(value any) ListItem {
	return li.set("style", value)
}

// SubTitle 设置副标题
// 支持两种语法：1. `${xxx}` 或 `${xxx|upperCase}` 2. `<%= data.xxx %>`
func (li ListItem) SubTitle(value any) ListItem {
	return li.set("subTitle", value)
}

// TestIdBuilder 设置测试 ID 构建器
func (li ListItem) TestIdBuilder(value string) ListItem {
	return li.set("testIdBuilder", value)
}

// Testid 设置测试 ID
func (li ListItem) Testid(value string) ListItem {
	return li.set("testid", value)
}

// Title 设置标题
// 支持两种语法：1. `${xxx}` 或 `${xxx|upperCase}` 2. `<%= data.xxx %>`
func (li ListItem) Title(value any) ListItem {
	return li.set("title", value)
}

// UseMobileUI 设置是否使用移动端样式
func (li ListItem) UseMobileUI(value bool) ListItem {
	return li.set("useMobileUI", value)
}

// Visible 设置是否显示
func (li ListItem) Visible(value bool) ListItem {
	return li.set("visible", value)
}

// VisibleOn 设置是否显示表达式
// 表达式语法 `data.xxx > 5`
func (li ListItem) VisibleOn(value string) ListItem {
	return li.set("visibleOn", value)
}
