package comp

// divider Divider renderer. Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/divider
type divider Schema

func Divider() divider {
	return make(divider).set("type", "divider")
}

func (d divider) set(key string, value any) divider {
	d[key] = value
	return d
}

// ClassName Container CSS class name (CSS class name, configuration string, or object. className: "red" When configured with an object, it means you can use it with expressions, such as: className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (d divider) ClassName(value string) divider {
	return d.set("className", value)
}

// Color Color
func (d divider) Color(value string) divider {
	return d.set("color", value)
}

// Direction Optional values: horizontal | vertical
func (d divider) Direction(value string) divider {
	return d.set("direction", value)
}

// Disabled Whether to disable
func (d divider) Disabled(value bool) divider {
	return d.set("disabled", value)
}

// DisabledOn Disable expression (expression, syntax `data.xxx > 5`.)
func (d divider) DisabledOn(value string) divider {
	return d.set("disabledOn", value)
}

// EditorSetting Editor configuration, can be ignored at runtime
func (d divider) EditorSetting(value string) divider {
	return d.set("editorSetting", value)
}

// Hidden Whether to hide
func (d divider) Hidden(value bool) divider {
	return d.set("hidden", value)
}

// HiddenOn Hide expression (expression, syntax `data.xxx > 5`.)
func (d divider) HiddenOn(value string) divider {
	return d.set("hiddenOn", value)
}

// ID Unique component ID, mainly used for log collection
func (d divider) ID(value string) divider {
	return d.set("id", value)
}

// LineStyle Optional values: dashed | solid
func (d divider) LineStyle(value string) divider {
	return d.set("lineStyle", value)
}

// OnEvent Event action configuration
func (d divider) OnEvent(value any) divider {
	return d.set("onEvent", value)
}

// Rotate Rotation angle
func (d divider) Rotate(value int) divider {
	return d.set("rotate", value)
}

// Static Whether to display statically
func (d divider) Static(value bool) divider {
	return d.set("static", value)
}

// StaticClassName Static display form item class name (CSS class name, configuration string, or object. className: "red" When configured with an object, it means you can use it with expressions, such as: className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (d divider) StaticClassName(value string) divider {
	return d.set("staticClassName", value)
}

// StaticInputClassName Static display form item value class name (CSS class name, configuration string, or object. className: "red" When configured with an object, it means you can use it with expressions, such as: className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (d divider) StaticInputClassName(value string) divider {
	return d.set("staticInputClassName", value)
}

// StaticLabelClassName Static display form item label class name (CSS class name, configuration string, or object. className: "red" When configured with an object, it means you can use it with expressions, such as: className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (d divider) StaticLabelClassName(value string) divider {
	return d.set("staticLabelClassName", value)
}

// StaticOn Static display expression (expression, syntax `data.xxx > 5`.)
func (d divider) StaticOn(value string) divider {
	return d.set("staticOn", value)
}

// StaticPlaceholder Static display placeholder for empty values
func (d divider) StaticPlaceholder(value string) divider {
	return d.set("staticPlaceholder", value)
}

// StaticSchema Static display schema
func (d divider) StaticSchema(value string) divider {
	return d.set("staticSchema", value)
}

// Style Component style
func (d divider) Style(value any) divider {
	return d.set("style", value)
}

// TestIdBuilder Test ID builder
func (d divider) TestIdBuilder(value string) divider {
	return d.set("testIdBuilder", value)
}

// Testid Test ID
func (d divider) Testid(value string) divider {
	return d.set("testid", value)
}

// Title Title
func (d divider) Title(value any) divider {
	return d.set("title", value)
}

// TitleClassName Title class name
func (d divider) TitleClassName(value string) divider {
	return d.set("titleClassName", value)
}

// TitlePosition Optional values: left | center | right
func (d divider) TitlePosition(value string) divider {
	return d.set("titlePosition", value)
}

// UseMobileUI Can be used at the component level to disable mobile styles
func (d divider) UseMobileUI(value bool) divider {
	return d.set("useMobileUI", value)
}

// Visible Whether to display
func (d divider) Visible(value bool) divider {
	return d.set("visible", value)
}

// VisibleOn Display expression (expression, syntax `data.xxx > 5`.)
func (d divider) VisibleOn(value string) divider {
	return d.set("visibleOn", value)
}
