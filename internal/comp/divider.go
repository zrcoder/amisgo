package comp

import "github.com/zrcoder/amisgo/schema"

// Divider Divider renderer. Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/Divider
type Divider schema.Schema

func NewDivider() Divider {
	return Divider{"type": "divider"}
}

func (d Divider) set(key string, value any) Divider {
	d[key] = value
	return d
}

// ClassName Container CSS class name (CSS class name, configuration string, or object. className: "red" When configured with an object, it means you can use it with expressions, such as: className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (d Divider) ClassName(value string) Divider {
	return d.set("className", value)
}

// Color Color
func (d Divider) Color(value string) Divider {
	return d.set("color", value)
}

// Direction Optional values: horizontal | vertical
func (d Divider) Direction(value string) Divider {
	return d.set("direction", value)
}

// Disabled Whether to disable
func (d Divider) Disabled(value bool) Divider {
	return d.set("disabled", value)
}

// DisabledOn Disable expression (expression, syntax `data.xxx > 5`.)
func (d Divider) DisabledOn(value string) Divider {
	return d.set("disabledOn", value)
}

// EditorSetting Editor configuration, can be ignored at runtime
func (d Divider) EditorSetting(value string) Divider {
	return d.set("editorSetting", value)
}

// Hidden Whether to hide
func (d Divider) Hidden(value bool) Divider {
	return d.set("hidden", value)
}

// HiddenOn Hide expression (expression, syntax `data.xxx > 5`.)
func (d Divider) HiddenOn(value string) Divider {
	return d.set("hiddenOn", value)
}

// ID Unique component ID, mainly used for log collection
func (d Divider) ID(value string) Divider {
	return d.set("id", value)
}

// LineStyle Optional values: dashed | solid
func (d Divider) LineStyle(value string) Divider {
	return d.set("lineStyle", value)
}

// OnEvent Event action configuration
func (d Divider) OnEvent(value any) Divider {
	return d.set("onEvent", value)
}

// Rotate Rotation angle
func (d Divider) Rotate(value int) Divider {
	return d.set("rotate", value)
}

// Static Whether to display statically
func (d Divider) Static(value bool) Divider {
	return d.set("static", value)
}

// StaticClassName Static display form item class name (CSS class name, configuration string, or object. className: "red" When configured with an object, it means you can use it with expressions, such as: className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (d Divider) StaticClassName(value string) Divider {
	return d.set("staticClassName", value)
}

// StaticInputClassName Static display form item value class name (CSS class name, configuration string, or object. className: "red" When configured with an object, it means you can use it with expressions, such as: className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (d Divider) StaticInputClassName(value string) Divider {
	return d.set("staticInputClassName", value)
}

// StaticLabelClassName Static display form item label class name (CSS class name, configuration string, or object. className: "red" When configured with an object, it means you can use it with expressions, such as: className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (d Divider) StaticLabelClassName(value string) Divider {
	return d.set("staticLabelClassName", value)
}

// StaticOn Static display expression (expression, syntax `data.xxx > 5`.)
func (d Divider) StaticOn(value string) Divider {
	return d.set("staticOn", value)
}

// StaticPlaceholder Static display placeholder for empty values
func (d Divider) StaticPlaceholder(value string) Divider {
	return d.set("staticPlaceholder", value)
}

// StaticSchema Static display schema.Schema
func (d Divider) StaticSchema(value string) Divider {
	return d.set("staticSchema", value)
}

// Style Component style
func (d Divider) Style(value any) Divider {
	return d.set("style", value)
}

// TestIdBuilder Test ID builder
func (d Divider) TestIdBuilder(value string) Divider {
	return d.set("testIdBuilder", value)
}

// Testid Test ID
func (d Divider) Testid(value string) Divider {
	return d.set("testid", value)
}

// Title Title
func (d Divider) Title(value any) Divider {
	return d.set("title", value)
}

// TitleClassName Title class name
func (d Divider) TitleClassName(value string) Divider {
	return d.set("titleClassName", value)
}

// TitlePosition Optional values: left | center | right
func (d Divider) TitlePosition(value string) Divider {
	return d.set("titlePosition", value)
}

// UseMobileUI Can be used at the component level to disable mobile styles
func (d Divider) UseMobileUI(value bool) Divider {
	return d.set("useMobileUI", value)
}

// Visible Whether to display
func (d Divider) Visible(value bool) Divider {
	return d.set("visible", value)
}

// VisibleOn Display expression (expression, syntax `data.xxx > 5`.)
func (d Divider) VisibleOn(value string) Divider {
	return d.set("visibleOn", value)
}
