package comp

import "github.com/zrcoder/amisgo/schema"

// Flex layout documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/flex
type Flex schema.Schema

func NewFlex() Flex {
	return Flex{"type": "flex"}
}

func (f Flex) set(key string, value any) Flex {
	f[key] = value
	return f
}

// AlignContent sets vertical distribution for multiple lines. Options: normal | flex-start | flex-end | center | space-between | space-around | space-evenly | stretch
func (f Flex) AlignContent(value string) Flex {
	return f.set("alignContent", value)
}

// AlignItems sets vertical alignment. Options: stretch | start | flex-start | flex-end | end | center | baseline
func (f Flex) AlignItems(value string) Flex {
	return f.set("alignItems", value)
}

// ClassName sets the container CSS class name
func (f Flex) ClassName(value string) Flex {
	return f.set("className", value)
}

// Direction sets the direction. Options: row | column | row-reverse | column-reverse
func (f Flex) Direction(value string) Flex {
	return f.set("direction", value)
}

// Disabled sets whether the component is disabled
func (f Flex) Disabled(value bool) Flex {
	return f.set("disabled", value)
}

// DisabledOn sets the expression to determine if the component is disabled
func (f Flex) DisabledOn(value string) Flex {
	return f.set("disabledOn", value)
}

// EditorSetting sets the editor configuration, ignored at runtime
func (f Flex) EditorSetting(value string) Flex {
	return f.set("editorSetting", value)
}

// Hidden sets whether the component is hidden
func (f Flex) Hidden(value bool) Flex {
	return f.set("hidden", value)
}

// HiddenOn sets the expression to determine if the component is hidden
func (f Flex) HiddenOn(value string) Flex {
	return f.set("hiddenOn", value)
}

// ID sets the unique component ID, mainly for logging
func (f Flex) ID(value string) Flex {
	return f.set("id", value)
}

// Items sets the configuration for each flex item
func (f Flex) Items(value ...any) Flex {
	return f.set("items", value)
}

// Justify sets horizontal distribution. Options: start | flex-start | center | end | flex-end | space-around | space-between | space-evenly
func (f Flex) Justify(value string) Flex {
	return f.set("justify", value)
}

// OnEvent sets the event action configuration
func (f Flex) OnEvent(value any) Flex {
	return f.set("onEvent", value)
}

// Static sets whether the component is statically displayed
func (f Flex) Static(value bool) Flex {
	return f.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (f Flex) StaticClassName(value string) Flex {
	return f.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (f Flex) StaticInputClassName(value string) Flex {
	return f.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (f Flex) StaticLabelClassName(value string) Flex {
	return f.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the component is statically displayed
func (f Flex) StaticOn(value string) Flex {
	return f.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (f Flex) StaticPlaceholder(value string) Flex {
	return f.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema.Schema
func (f Flex) StaticSchema(value string) Flex {
	return f.set("staticSchema", value)
}

// Style sets custom styles
func (f Flex) Style(value any) Flex {
	return f.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (f Flex) TestIdBuilder(value string) Flex {
	return f.set("testIdBuilder", value)
}

// TestID sets the test ID
func (f Flex) TestID(value string) Flex {
	return f.set("testid", value)
}

// UseMobileUI sets whether to disable mobile styles at the component level
func (f Flex) UseMobileUI(value bool) Flex {
	return f.set("useMobileUI", value)
}

// Visible sets whether the component is visible
func (f Flex) Visible(value bool) Flex {
	return f.set("visible", value)
}

// VisibleOn sets the expression to determine if the component is visible
func (f Flex) VisibleOn(value string) Flex {
	return f.set("visibleOn", value)
}
