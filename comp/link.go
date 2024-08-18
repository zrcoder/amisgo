package comp

// Link 链接展示控件
//
// 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/link
// @version 6.7.0
type Link Schema

// NewLink 创建一个新的 Link 实例，并设置默认的 type
func NewLink() Link {
	return make(Link).set("type", "link")
}

// Badge 角标
func (l Link) Badge(value string) Link {
	return l.set("badge", value)
}

// Blank 是否新窗口打开
func (l Link) Blank(value bool) Link {
	return l.set("blank", value)
}

// Body 链接内容
func (l Link) Body(value ...interface{}) Link {
	return l.set("body", value)
}

// ClassName 容器 css 类名
func (l Link) ClassName(value string) Link {
	return l.set("className", value)
}

// Disabled 是否禁用
func (l Link) Disabled(value bool) Link {
	return l.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (l Link) DisabledOn(value string) Link {
	return l.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (l Link) EditorSetting(value string) Link {
	return l.set("editorSetting", value)
}

// Hidden 是否隐藏
func (l Link) Hidden(value bool) Link {
	return l.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (l Link) HiddenOn(value string) Link {
	return l.set("hiddenOn", value)
}

// Href 链接地址
func (l Link) Href(value string) Link {
	return l.set("href", value)
}

// HtmlTarget a标签原生target属性
func (l Link) HtmlTarget(value string) Link {
	return l.set("htmlTarget", value)
}

// Icon 图标
func (l Link) Icon(value string) Link {
	return l.set("icon", value)
}

// ID 组件唯一 id
func (l Link) ID(value string) Link {
	return l.set("id", value)
}

// OnEvent 事件动作配置
func (l Link) OnEvent(value string) Link {
	return l.set("onEvent", value)
}

// RightIcon 右侧图标
func (l Link) RightIcon(value string) Link {
	return l.set("rightIcon", value)
}

// Static 是否静态展示
func (l Link) Static(value bool) Link {
	return l.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (l Link) StaticClassName(value string) Link {
	return l.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (l Link) StaticInputClassName(value string) Link {
	return l.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (l Link) StaticLabelClassName(value string) Link {
	return l.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (l Link) StaticOn(value string) Link {
	return l.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (l Link) StaticPlaceholder(value string) Link {
	return l.set("staticPlaceholder", value)
}

// StaticSchema 静态展示 Schema
func (l Link) StaticSchema(value string) Link {
	return l.set("staticSchema", value)
}

// Style 组件样式
func (l Link) Style(value string) Link {
	return l.set("style", value)
}

// TestIdBuilder 测试 ID 构建器
func (l Link) TestIdBuilder(value string) Link {
	return l.set("testIdBuilder", value)
}

// Testid 测试 id
func (l Link) Testid(value string) Link {
	return l.set("testid", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (l Link) UseMobileUI(value bool) Link {
	return l.set("useMobileUI", value)
}

// Visible 是否显示
func (l Link) Visible(value bool) Link {
	return l.set("visible", value)
}

// VisibleOn 是否显示表达式
func (l Link) VisibleOn(value string) Link {
	return l.set("visibleOn", value)
}

// set 设置键值对
func (l Link) set(key string, value interface{}) Link {
	l[key] = value
	return l
}
