package comp

// alert 提示渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/alert

type alert schema

// Alert 创建一个新的 Alert 实例
func Alert() alert {
	return make(alert).set("type", "alert")
}

func (a alert) set(key string, value any) alert {
	a[key] = value
	return a
}

// 操作区域 (操作区域)
func (a alert) Actions(value string) alert {
	return a.set("actions", value)
}

// 内容区域 (内容区域)
func (a alert) Body(value ...any) alert {
	return a.set("body", value)
}

// 容器 css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (a alert) ClassName(value string) alert {
	return a.set("className", value)
}

// 关闭按钮CSS类名
func (a alert) CloseButtonClassName(value string) alert {
	return a.set("closeButtonClassName", value)
}

// 是否禁用
func (a alert) Disabled(value bool) alert {
	return a.set("disabled", value)
}

// 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (a alert) DisabledOn(value string) alert {
	return a.set("disabledOn", value)
}

// 编辑器配置，运行时可以忽略
func (a alert) EditorSetting(value string) alert {
	return a.set("editorSetting", value)
}

// 是否隐藏
func (a alert) Hidden(value bool) alert {
	return a.set("hidden", value)
}

// 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (a alert) HiddenOn(value string) alert {
	return a.set("hiddenOn", value)
}

// 左侧图标 (iconfont 里面的类名。)
func (a alert) Icon(value string) alert {
	return a.set("icon", value)
}

// 图标CSS类名
func (a alert) IconClassName(value string) alert {
	return a.set("iconClassName", value)
}

// 组件唯一 id，主要用于日志采集
func (a alert) ID(value string) alert {
	return a.set("id", value)
}

// 提示类型 可选值: info | warning | success | danger
func (a alert) Level(value string) alert {
	return a.set("level", value)
}

// 事件动作配置
func (a alert) OnEvent(value any) alert {
	return a.set("onEvent", value)
}

// 是否显示关闭按钮
func (a alert) ShowCloseButton(value bool) alert {
	return a.set("showCloseButton", value)
}

// 是否显示ICON
func (a alert) ShowIcon(value bool) alert {
	return a.set("showIcon", value)
}

// 是否静态展示
func (a alert) Static(value bool) alert {
	return a.set("static", value)
}

// 静态展示表单项类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (a alert) StaticClassName(value string) alert {
	return a.set("staticClassName", value)
}

// 静态展示表单项Value类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (a alert) StaticInputClassName(value string) alert {
	return a.set("staticInputClassName", value)
}

// 静态展示表单项Label类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (a alert) StaticLabelClassName(value string) alert {
	return a.set("staticLabelClassName", value)
}

// 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (a alert) StaticOn(value string) alert {
	return a.set("staticOn", value)
}

// 静态展示空值占位
func (a alert) StaticPlaceholder(value string) alert {
	return a.set("staticPlaceholder", value)
}

func (a alert) StaticSchema(value string) alert {
	return a.set("staticSchema", value)
}

// 组件样式
func (a alert) Style(value any) alert {
	return a.set("style", value)
}

func (a alert) TestIdBuilder(value string) alert {
	return a.set("testIdBuilder", value)
}

func (a alert) Testid(value string) alert {
	return a.set("testid", value)
}

// 提示框标题
func (a alert) Title(value any) alert {
	return a.set("title", value)
}

// 可以组件级别用来关闭移动端样式
func (a alert) UseMobileUI(value bool) alert {
	return a.set("useMobileUI", value)
}

// 是否显示
func (a alert) Visible(value bool) alert {
	return a.set("visible", value)
}

// 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (a alert) VisibleOn(value string) alert {
	return a.set("visibleOn", value)
}
