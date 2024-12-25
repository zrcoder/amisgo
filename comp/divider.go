package comp

// divider 分割线渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/divider
type divider Schema

func Divider() divider {
	return make(divider).set("type", "divider")
}

func (d divider) set(key string, value any) divider {
	d[key] = value
	return d
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (d divider) ClassName(value string) divider {
	return d.set("className", value)
}

// Color 颜色
func (d divider) Color(value string) divider {
	return d.set("color", value)
}

// Direction 可选值: horizontal | vertical
func (d divider) Direction(value string) divider {
	return d.set("direction", value)
}

// Disabled 是否禁用
func (d divider) Disabled(value bool) divider {
	return d.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (d divider) DisabledOn(value string) divider {
	return d.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (d divider) EditorSetting(value string) divider {
	return d.set("editorSetting", value)
}

// Hidden 是否隐藏
func (d divider) Hidden(value bool) divider {
	return d.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (d divider) HiddenOn(value string) divider {
	return d.set("hiddenOn", value)
}

// ID 组件唯一 id，主要用于日志采集
func (d divider) ID(value string) divider {
	return d.set("id", value)
}

// LineStyle 可选值: dashed | solid
func (d divider) LineStyle(value string) divider {
	return d.set("lineStyle", value)
}

// OnEvent 事件动作配置
func (d divider) OnEvent(value any) divider {
	return d.set("onEvent", value)
}

// Rotate 旋转角度
func (d divider) Rotate(value int) divider {
	return d.set("rotate", value)
}

// Static 是否静态展示
func (d divider) Static(value bool) divider {
	return d.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (d divider) StaticClassName(value string) divider {
	return d.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (d divider) StaticInputClassName(value string) divider {
	return d.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (d divider) StaticLabelClassName(value string) divider {
	return d.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (d divider) StaticOn(value string) divider {
	return d.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (d divider) StaticPlaceholder(value string) divider {
	return d.set("staticPlaceholder", value)
}

// StaticSchema 静态展示 schema
func (d divider) StaticSchema(value string) divider {
	return d.set("staticSchema", value)
}

// Style 组件样式
func (d divider) Style(value any) divider {
	return d.set("style", value)
}

// TestIdBuilder 测试 ID 构建器
func (d divider) TestIdBuilder(value string) divider {
	return d.set("testIdBuilder", value)
}

// Testid 测试 id
func (d divider) Testid(value string) divider {
	return d.set("testid", value)
}

// Title 标题
func (d divider) Title(value any) divider {
	return d.set("title", value)
}

// TitleClassName 标题类名
func (d divider) TitleClassName(value string) divider {
	return d.set("titleClassName", value)
}

// TitlePosition 可选值: left | center | right
func (d divider) TitlePosition(value string) divider {
	return d.set("titlePosition", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (d divider) UseMobileUI(value bool) divider {
	return d.set("useMobileUI", value)
}

// Visible 是否显示
func (d divider) Visible(value bool) divider {
	return d.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (d divider) VisibleOn(value string) divider {
	return d.set("visibleOn", value)
}
