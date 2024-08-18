package comp

// Progress
//
// @version 6.7.0
type Progress Schema

// NewProgress 创建一个新的 Progress 实例
func NewProgress() Progress {
	return Progress{}.set("type", "progress").set("mode", "line")
}

func (p Progress) set(key string, value interface{}) Progress {
	p[key] = value
	return p
}

// Animate 是否显示动画
func (p Progress) Animate(value bool) Progress {
	return p.set("animate", value)
}

// ClassName 容器 css 类名
func (p Progress) ClassName(value string) Progress {
	return p.set("className", value)
}

// Disabled 是否禁用
func (p Progress) Disabled(value bool) Progress {
	return p.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (p Progress) DisabledOn(value string) Progress {
	return p.set("disabledOn", value)
}

// EditorSetting 编辑器配置
func (p Progress) EditorSetting(value string) Progress {
	return p.set("editorSetting", value)
}

// GapDegree 仪表盘进度条缺口角度
func (p Progress) GapDegree(value string) Progress {
	return p.set("gapDegree", value)
}

// GapPosition 仪表盘进度条缺口位置
func (p Progress) GapPosition(value string) Progress {
	return p.set("gapPosition", value)
}

// Hidden 是否隐藏
func (p Progress) Hidden(value bool) Progress {
	return p.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (p Progress) HiddenOn(value string) Progress {
	return p.set("hiddenOn", value)
}

// ID 组件唯一 id
func (p Progress) ID(value string) Progress {
	return p.set("id", value)
}

// Map 配置不同的值段
func (p Progress) Map(value string) Progress {
	return p.set("map", value)
}

// Mode 进度条类型
func (p Progress) Mode(value string) Progress {
	return p.set("mode", value)
}

// Name 关联字段名
func (p Progress) Name(value string) Progress {
	return p.set("name", value)
}

// OnEvent 事件动作配置
func (p Progress) OnEvent(value string) Progress {
	return p.set("onEvent", value)
}

// Placeholder 占位符
func (p Progress) Placeholder(value string) Progress {
	return p.set("placeholder", value)
}

// ProgressClassName 进度条 CSS 类名
func (p Progress) ProgressClassName(value string) Progress {
	return p.set("progressClassName", value)
}

// ShowLabel 是否显示值
func (p Progress) ShowLabel(value bool) Progress {
	return p.set("showLabel", value)
}

// ShowThresholdText 是否显示阈值数值
func (p Progress) ShowThresholdText(value bool) Progress {
	return p.set("showThresholdText", value)
}

// Static 是否静态展示
func (p Progress) Static(value bool) Progress {
	return p.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (p Progress) StaticClassName(value string) Progress {
	return p.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (p Progress) StaticInputClassName(value string) Progress {
	return p.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (p Progress) StaticLabelClassName(value string) Progress {
	return p.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (p Progress) StaticOn(value string) Progress {
	return p.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (p Progress) StaticPlaceholder(value string) Progress {
	return p.set("staticPlaceholder", value)
}

// StaticSchema 静态展示 schema
func (p Progress) StaticSchema(value string) Progress {
	return p.set("staticSchema", value)
}

// Stripe 是否显示背景间隔
func (p Progress) Stripe(value bool) Progress {
	return p.set("stripe", value)
}

// StrokeWidth 进度条线的宽度
func (p Progress) StrokeWidth(value string) Progress {
	return p.set("strokeWidth", value)
}

// Style 组件样式
func (p Progress) Style(value string) Progress {
	return p.set("style", value)
}

// TestIdBuilder 测试 id 构造器
func (p Progress) TestIdBuilder(value string) Progress {
	return p.set("testIdBuilder", value)
}

// Testid 测试 id
func (p Progress) Testid(value string) Progress {
	return p.set("testid", value)
}

// Threshold 阈值
func (p Progress) Threshold(value string) Progress {
	return p.set("threshold", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (p Progress) UseMobileUI(value bool) Progress {
	return p.set("useMobileUI", value)
}

// Value 进度值
func (p Progress) Value(value string) Progress {
	return p.set("value", value)
}

// ValueTpl 内容的模板函数
func (p Progress) ValueTpl(value string) Progress {
	return p.set("valueTpl", value)
}

// Visible 是否显示
func (p Progress) Visible(value bool) Progress {
	return p.set("visible", value)
}

// VisibleOn 是否显示表达式
func (p Progress) VisibleOn(value string) Progress {
	return p.set("visibleOn", value)
}
