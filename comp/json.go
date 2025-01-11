package comp

import "github.com/zrcoder/amisgo/model"

// json JSON data display component
// Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/json

type json model.Schema

// Json creates a new Json instance with default type
func Json() json {
	return json{"type": "json"}
}

// ClassName sets the container CSS class name
func (j json) ClassName(value string) json {
	j.set("className", value)
	return j
}

func (j json) set(key string, value any) json {
	j[key] = value
	return j
}

// Disabled sets whether the component is disabled
func (j json) Disabled(value bool) json {
	j.set("disabled", value)
	return j
}

// DisabledOn sets the expression to disable the component
func (j json) DisabledOn(value string) json {
	j.set("disabledOn", value)
	return j
}

// DisplayDataTypes sets whether to display data types
func (j json) DisplayDataTypes(value bool) json {
	j.set("displayDataTypes", value)
	return j
}

// EditorSetting sets the editor configuration
func (j json) EditorSetting(value string) json {
	j.set("editorSetting", value)
	return j
}

// EllipsisThreshold sets the maximum display length of strings
func (j json) EllipsisThreshold(value string) json {
	j.set("ellipsisThreshold", value)
	return j
}

// EnableClipboard sets whether the content is copyable
func (j json) EnableClipboard(value bool) json {
	j.set("enableClipboard", value)
	return j
}

// Hidden sets whether the component is hidden
func (j json) Hidden(value bool) json {
	j.set("hidden", value)
	return j
}

// HiddenOn sets the expression to hide the component
func (j json) HiddenOn(value string) json {
	j.set("hiddenOn", value)
	return j
}

// IconStyle sets the icon style
func (j json) IconStyle(value any) json {
	j.set("iconStyle", value)
	return j
}

// ID sets the unique component ID
func (j json) ID(value string) json {
	j.set("id", value)
	return j
}

// LevelExpand sets the default expand level
func (j json) LevelExpand(value string) json {
	j.set("levelExpand", value)
	return j
}

// Mutable sets whether the component is mutable
func (j json) Mutable(value bool) json {
	j.set("mutable", value)
	return j
}

// OnEvent sets the event action configuration
func (j json) OnEvent(value any) json {
	j.set("onEvent", value)
	return j
}

// QuotesOnKeys sets whether to show quotes on keys
func (j json) QuotesOnKeys(value bool) json {
	j.set("quotesOnKeys", value)
	return j
}

// SortKeys sets whether to sort keys
func (j json) SortKeys(value bool) json {
	j.set("sortKeys", value)
	return j
}

// Source sets the data source
func (j json) Source(value string) json {
	j.set("source", value)
	return j
}

// Static sets whether the component is static
func (j json) Static(value bool) json {
	j.set("static", value)
	return j
}

// StaticClassName sets the static display form item class name
func (j json) StaticClassName(value string) json {
	j.set("staticClassName", value)
	return j
}

// StaticInputClassName sets the static display form item value class name
func (j json) StaticInputClassName(value string) json {
	j.set("staticInputClassName", value)
	return j
}

// StaticLabelClassName sets the static display form item label class name
func (j json) StaticLabelClassName(value string) json {
	j.set("staticLabelClassName", value)
	return j
}

// StaticOn sets the expression for static display
func (j json) StaticOn(value string) json {
	j.set("staticOn", value)
	return j
}

// StaticPlaceholder sets the placeholder for static display
func (j json) StaticPlaceholder(value string) json {
	j.set("staticPlaceholder", value)
	return j
}

// StaticSchema sets the static display schema
func (j json) StaticSchema(value string) json {
	j.set("staticSchema", value)
	return j
}

// Style sets the component style
func (j json) Style(value any) json {
	j.set("style", value)
	return j
}

// TestIdBuilder sets the test ID builder
func (j json) TestIdBuilder(value string) json {
	j.set("testIdBuilder", value)
	return j
}

// Testid sets the test ID
func (j json) Testid(value string) json {
	j.set("testid", value)
	return j
}

// UseMobileUI sets whether to use mobile UI
func (j json) UseMobileUI(value bool) json {
	j.set("useMobileUI", value)
	return j
}

// Value sets the JSON data to display
func (j json) Value(value string) json {
	j.set("value", value)
	return j
}

// Visible sets whether the component is visible
func (j json) Visible(value bool) json {
	j.set("visible", value)
	return j
}

// VisibleOn sets the expression to show the component
func (j json) VisibleOn(value string) json {
	j.set("visibleOn", value)
	return j
}
