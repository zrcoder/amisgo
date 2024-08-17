package comp

// Divider 分割线渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/divider
type Divider BaseRenderer

func NewDivider() Divider {
	d := Divider(make(BaseRenderer))
	return d.set("type", "divider")
}

func (d Divider) set(key string, value interface{}) Divider {
	d[key] = value
	return d
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (d Divider) ClassName(value string) Divider {
	return d.set("className", value)
}

// Color 颜色
func (d Divider) Color(value string) Divider {
	return d.set("color", value)
}

// Direction 可选值: horizontal | vertical
func (d Divider) Direction(value string) Divider {
	return d.set("direction", value)
}

// Disabled 是否禁用
func (d Divider) Disabled(value bool) Divider {
	return d.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (d Divider) DisabledOn(value string) Divider {
	return d.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (d Divider) EditorSetting(value string) Divider {
	return d.set("editorSetting", value)
}

// Hidden 是否隐藏
func (d Divider) Hidden(value bool) Divider {
	return d.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (d Divider) HiddenOn(value string) Divider {
	return d.set("hiddenOn", value)
}

// ID 组件唯一 id，主要用于日志采集
func (d Divider) ID(value string) Divider {
	return d.set("id", value)
}

// LineStyle 可选值: dashed | solid
func (d Divider) LineStyle(value string) Divider {
	return d.set("lineStyle", value)
}

// OnEvent 事件动作配置
func (d Divider) OnEvent(value string) Divider {
	return d.set("onEvent", value)
}

// Rotate 旋转角度
func (d Divider) Rotate(value string) Divider {
	return d.set("rotate", value)
}

// Static 是否静态展示
func (d Divider) Static(value bool) Divider {
	return d.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (d Divider) StaticClassName(value string) Divider {
	return d.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (d Divider) StaticInputClassName(value string) Divider {
	return d.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (d Divider) StaticLabelClassName(value string) Divider {
	return d.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (d Divider) StaticOn(value string) Divider {
	return d.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (d Divider) StaticPlaceholder(value string) Divider {
	return d.set("staticPlaceholder", value)
}

// StaticSchema 静态展示 schema
func (d Divider) StaticSchema(value string) Divider {
	return d.set("staticSchema", value)
}

// Style 组件样式
func (d Divider) Style(value string) Divider {
	return d.set("style", value)
}

// TestIdBuilder 测试 ID 构建器
func (d Divider) TestIdBuilder(value string) Divider {
	return d.set("testIdBuilder", value)
}

// Testid 测试 id
func (d Divider) Testid(value string) Divider {
	return d.set("testid", value)
}

// Title 标题
func (d Divider) Title(value string) Divider {
	return d.set("title", value)
}

// TitleClassName 标题类名
func (d Divider) TitleClassName(value string) Divider {
	return d.set("titleClassName", value)
}

// TitlePosition 可选值: left | center | right
func (d Divider) TitlePosition(value string) Divider {
	return d.set("titlePosition", value)
}

// Type 指定为 divider
func (d Divider) Type(value string) Divider {
	return d.set("type", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (d Divider) UseMobileUI(value bool) Divider {
	return d.set("useMobileUI", value)
}

// Visible 是否显示
func (d Divider) Visible(value bool) Divider {
	return d.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (d Divider) VisibleOn(value string) Divider {
	return d.set("visibleOn", value)
}
