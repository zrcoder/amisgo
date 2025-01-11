package comp

import "github.com/zrcoder/amisgo/model"

// flex layout documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/flex

type flex model.Schema

// Flex creates a new Flex instance
func Flex() flex {
	return flex{"type": "flex"}
}

func (f flex) set(key string, value any) flex {
	f[key] = value
	return f
}

// AlignContent sets vertical distribution for multiple lines. Options: normal | flex-start | flex-end | center | space-between | space-around | space-evenly | stretch
func (f flex) AlignContent(value string) flex {
	return f.set("alignContent", value)
}

// AlignItems sets vertical alignment. Options: stretch | start | flex-start | flex-end | end | center | baseline
func (f flex) AlignItems(value string) flex {
	return f.set("alignItems", value)
}

// ClassName sets the container CSS class name
func (f flex) ClassName(value string) flex {
	return f.set("className", value)
}

// Direction sets the direction. Options: row | column | row-reverse | column-reverse
func (f flex) Direction(value string) flex {
	return f.set("direction", value)
}

// Disabled sets whether the component is disabled
func (f flex) Disabled(value bool) flex {
	return f.set("disabled", value)
}

// DisabledOn sets the expression to determine if the component is disabled
func (f flex) DisabledOn(value string) flex {
	return f.set("disabledOn", value)
}

// EditorSetting sets the editor configuration, ignored at runtime
func (f flex) EditorSetting(value string) flex {
	return f.set("editorSetting", value)
}

// Hidden sets whether the component is hidden
func (f flex) Hidden(value bool) flex {
	return f.set("hidden", value)
}

// HiddenOn sets the expression to determine if the component is hidden
func (f flex) HiddenOn(value string) flex {
	return f.set("hiddenOn", value)
}

// ID sets the unique component ID, mainly for logging
func (f flex) ID(value string) flex {
	return f.set("id", value)
}

// Items sets the configuration for each flex item
func (f flex) Items(value ...any) flex {
	return f.set("items", value)
}

// Justify sets horizontal distribution. Options: start | flex-start | center | end | flex-end | space-around | space-between | space-evenly
func (f flex) Justify(value string) flex {
	return f.set("justify", value)
}

// OnEvent sets the event action configuration
func (f flex) OnEvent(value any) flex {
	return f.set("onEvent", value)
}

// Static sets whether the component is statically displayed
func (f flex) Static(value bool) flex {
	return f.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (f flex) StaticClassName(value string) flex {
	return f.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (f flex) StaticInputClassName(value string) flex {
	return f.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (f flex) StaticLabelClassName(value string) flex {
	return f.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the component is statically displayed
func (f flex) StaticOn(value string) flex {
	return f.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (f flex) StaticPlaceholder(value string) flex {
	return f.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema
func (f flex) StaticSchema(value string) flex {
	return f.set("staticSchema", value)
}

// Style sets custom styles
func (f flex) Style(value any) flex {
	return f.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (f flex) TestIdBuilder(value string) flex {
	return f.set("testIdBuilder", value)
}

// TestID sets the test ID
func (f flex) TestID(value string) flex {
	return f.set("testid", value)
}

// UseMobileUI sets whether to disable mobile styles at the component level
func (f flex) UseMobileUI(value bool) flex {
	return f.set("useMobileUI", value)
}

// Visible sets whether the component is visible
func (f flex) Visible(value bool) flex {
	return f.set("visible", value)
}

// VisibleOn sets the expression to determine if the component is visible
func (f flex) VisibleOn(value string) flex {
	return f.set("visibleOn", value)
}
