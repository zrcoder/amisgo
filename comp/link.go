package comp

// link 链接展示控件
// 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/link

type link Schema

// Link 创建一个新的 Link 实例，并设置默认的 type
func Link() link {
	return make(link).set("type", "link")
}

// Badge 角标
func (l link) Badge(value string) link {
	return l.set("badge", value)
}

// Blank 是否新窗口打开
func (l link) Blank(value bool) link {
	return l.set("blank", value)
}

// Body 链接内容
func (l link) Body(value ...any) link {
	return l.set("body", value)
}

// ClassName 容器 css 类名
func (l link) ClassName(value string) link {
	return l.set("className", value)
}

// Disabled 是否禁用
func (l link) Disabled(value bool) link {
	return l.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (l link) DisabledOn(value string) link {
	return l.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (l link) EditorSetting(value string) link {
	return l.set("editorSetting", value)
}

// Hidden 是否隐藏
func (l link) Hidden(value bool) link {
	return l.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (l link) HiddenOn(value string) link {
	return l.set("hiddenOn", value)
}

// Href 链接地址
func (l link) Href(value string) link {
	return l.set("href", value)
}

// HtmlTarget a标签原生target属性
func (l link) HtmlTarget(value string) link {
	return l.set("htmlTarget", value)
}

// Icon 图标
func (l link) Icon(value string) link {
	return l.set("icon", value)
}

// ID 组件唯一 id
func (l link) ID(value string) link {
	return l.set("id", value)
}

// OnEvent 事件动作配置
func (l link) OnEvent(value any) link {
	return l.set("onEvent", value)
}

// RightIcon 右侧图标
func (l link) RightIcon(value string) link {
	return l.set("rightIcon", value)
}

// Static 是否静态展示
func (l link) Static(value bool) link {
	return l.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (l link) StaticClassName(value string) link {
	return l.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (l link) StaticInputClassName(value string) link {
	return l.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (l link) StaticLabelClassName(value string) link {
	return l.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (l link) StaticOn(value string) link {
	return l.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (l link) StaticPlaceholder(value string) link {
	return l.set("staticPlaceholder", value)
}

// StaticSchema 静态展示 Schema
func (l link) StaticSchema(value string) link {
	return l.set("staticSchema", value)
}

// Style 组件样式
func (l link) Style(value any) link {
	return l.set("style", value)
}

// TestIdBuilder 测试 ID 构建器
func (l link) TestIdBuilder(value string) link {
	return l.set("testIdBuilder", value)
}

// Testid 测试 id
func (l link) Testid(value string) link {
	return l.set("testid", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (l link) UseMobileUI(value bool) link {
	return l.set("useMobileUI", value)
}

// Visible 是否显示
func (l link) Visible(value bool) link {
	return l.set("visible", value)
}

// VisibleOn 是否显示表达式
func (l link) VisibleOn(value string) link {
	return l.set("visibleOn", value)
}

// set 设置键值对
func (l link) set(key string, value any) link {
	l[key] = value
	return l
}
