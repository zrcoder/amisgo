package comp

// progress
//
// @version 6.7.0
type progress schema

// Progress 创建一个新的 Progress 实例
func Progress() progress {
	return progress{}.set("type", "progress").set("mode", "line")
}

func (p progress) set(key string, value interface{}) progress {
	p[key] = value
	return p
}

// Animate 是否显示动画
func (p progress) Animate(value bool) progress {
	return p.set("animate", value)
}

// ClassName 容器 css 类名
func (p progress) ClassName(value string) progress {
	return p.set("className", value)
}

// Disabled 是否禁用
func (p progress) Disabled(value bool) progress {
	return p.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (p progress) DisabledOn(value string) progress {
	return p.set("disabledOn", value)
}

// EditorSetting 编辑器配置
func (p progress) EditorSetting(value string) progress {
	return p.set("editorSetting", value)
}

// GapDegree 仪表盘进度条缺口角度
func (p progress) GapDegree(value string) progress {
	return p.set("gapDegree", value)
}

// GapPosition 仪表盘进度条缺口位置
func (p progress) GapPosition(value string) progress {
	return p.set("gapPosition", value)
}

// Hidden 是否隐藏
func (p progress) Hidden(value bool) progress {
	return p.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (p progress) HiddenOn(value string) progress {
	return p.set("hiddenOn", value)
}

// ID 组件唯一 id
func (p progress) ID(value string) progress {
	return p.set("id", value)
}

// Map 配置不同的值段
func (p progress) Map(value string) progress {
	return p.set("map", value)
}

// Mode 进度条类型
func (p progress) Mode(value string) progress {
	return p.set("mode", value)
}

// Name 关联字段名
func (p progress) Name(value string) progress {
	return p.set("name", value)
}

// OnEvent 事件动作配置
func (p progress) OnEvent(value string) progress {
	return p.set("onEvent", value)
}

// Placeholder 占位符
func (p progress) Placeholder(value string) progress {
	return p.set("placeholder", value)
}

// ProgressClassName 进度条 CSS 类名
func (p progress) ProgressClassName(value string) progress {
	return p.set("progressClassName", value)
}

// ShowLabel 是否显示值
func (p progress) ShowLabel(value bool) progress {
	return p.set("showLabel", value)
}

// ShowThresholdText 是否显示阈值数值
func (p progress) ShowThresholdText(value bool) progress {
	return p.set("showThresholdText", value)
}

// Static 是否静态展示
func (p progress) Static(value bool) progress {
	return p.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (p progress) StaticClassName(value string) progress {
	return p.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (p progress) StaticInputClassName(value string) progress {
	return p.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (p progress) StaticLabelClassName(value string) progress {
	return p.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (p progress) StaticOn(value string) progress {
	return p.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (p progress) StaticPlaceholder(value string) progress {
	return p.set("staticPlaceholder", value)
}

// StaticSchema 静态展示 schema
func (p progress) StaticSchema(value string) progress {
	return p.set("staticSchema", value)
}

// Stripe 是否显示背景间隔
func (p progress) Stripe(value bool) progress {
	return p.set("stripe", value)
}

// StrokeWidth 进度条线的宽度
func (p progress) StrokeWidth(value string) progress {
	return p.set("strokeWidth", value)
}

// Style 组件样式
func (p progress) Style(value string) progress {
	return p.set("style", value)
}

// TestIdBuilder 测试 id 构造器
func (p progress) TestIdBuilder(value string) progress {
	return p.set("testIdBuilder", value)
}

// Testid 测试 id
func (p progress) Testid(value string) progress {
	return p.set("testid", value)
}

// Threshold 阈值
func (p progress) Threshold(value string) progress {
	return p.set("threshold", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (p progress) UseMobileUI(value bool) progress {
	return p.set("useMobileUI", value)
}

// Value 进度值
func (p progress) Value(value string) progress {
	return p.set("value", value)
}

// ValueTpl 内容的模板函数
func (p progress) ValueTpl(value string) progress {
	return p.set("valueTpl", value)
}

// Visible 是否显示
func (p progress) Visible(value bool) progress {
	return p.set("visible", value)
}

// VisibleOn 是否显示表达式
func (p progress) VisibleOn(value string) progress {
	return p.set("visibleOn", value)
}
