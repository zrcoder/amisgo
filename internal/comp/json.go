package comp

import "github.com/zrcoder/amisgo/schema"

// Json JSON data display component
// Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/Json
type Json schema.Schema

func NewJson() Json {
	return Json{"type": "json"}
}

// ClassName sets the container CSS class name
func (j Json) ClassName(value string) Json {
	j.set("className", value)
	return j
}

func (j Json) set(key string, value any) Json {
	j[key] = value
	return j
}

// Disabled sets whether the component is disabled
func (j Json) Disabled(value bool) Json {
	j.set("disabled", value)
	return j
}

// DisabledOn sets the expression to disable the component
func (j Json) DisabledOn(value string) Json {
	j.set("disabledOn", value)
	return j
}

// DisplayDataTypes sets whether to display data types
func (j Json) DisplayDataTypes(value bool) Json {
	j.set("displayDataTypes", value)
	return j
}

// EditorSetting sets the editor configuration
func (j Json) EditorSetting(value string) Json {
	j.set("editorSetting", value)
	return j
}

// EllipsisThreshold sets the maximum display length of strings
func (j Json) EllipsisThreshold(value string) Json {
	j.set("ellipsisThreshold", value)
	return j
}

// EnableClipboard sets whether the content is copyable
func (j Json) EnableClipboard(value bool) Json {
	j.set("enableClipboard", value)
	return j
}

// Hidden sets whether the component is hidden
func (j Json) Hidden(value bool) Json {
	j.set("hidden", value)
	return j
}

// HiddenOn sets the expression to hide the component
func (j Json) HiddenOn(value string) Json {
	j.set("hiddenOn", value)
	return j
}

// IconStyle sets the icon style
func (j Json) IconStyle(value any) Json {
	j.set("iconStyle", value)
	return j
}

// ID sets the unique component ID
func (j Json) ID(value string) Json {
	j.set("id", value)
	return j
}

// LevelExpand sets the default expand level
func (j Json) LevelExpand(value string) Json {
	j.set("levelExpand", value)
	return j
}

// Mutable sets whether the component is mutable
func (j Json) Mutable(value bool) Json {
	j.set("mutable", value)
	return j
}

// OnEvent sets the event action configuration
func (j Json) OnEvent(value Event) Json {
	j.set("onEvent", value)
	return j
}

// QuotesOnKeys sets whether to show quotes on keys
func (j Json) QuotesOnKeys(value bool) Json {
	j.set("quotesOnKeys", value)
	return j
}

// SortKeys sets whether to sort keys
func (j Json) SortKeys(value bool) Json {
	j.set("sortKeys", value)
	return j
}

// Source sets the data source
func (j Json) Source(value string) Json {
	j.set("source", value)
	return j
}

// Static sets whether the component is static
func (j Json) Static(value bool) Json {
	j.set("static", value)
	return j
}

// StaticClassName sets the static display form item class name
func (j Json) StaticClassName(value string) Json {
	j.set("staticClassName", value)
	return j
}

// StaticInputClassName sets the static display form item value class name
func (j Json) StaticInputClassName(value string) Json {
	j.set("staticInputClassName", value)
	return j
}

// StaticLabelClassName sets the static display form item label class name
func (j Json) StaticLabelClassName(value string) Json {
	j.set("staticLabelClassName", value)
	return j
}

// StaticOn sets the expression for static display
func (j Json) StaticOn(value string) Json {
	j.set("staticOn", value)
	return j
}

// StaticPlaceholder sets the placeholder for static display
func (j Json) StaticPlaceholder(value string) Json {
	j.set("staticPlaceholder", value)
	return j
}

// StaticSchema sets the static display schema.Schema
func (j Json) StaticSchema(value string) Json {
	j.set("staticSchema", value)
	return j
}

// Style sets the component style
func (j Json) Style(value any) Json {
	j.set("style", value)
	return j
}

// TestIdBuilder sets the test ID builder
func (j Json) TestIdBuilder(value string) Json {
	j.set("testIdBuilder", value)
	return j
}

// Testid sets the test ID
func (j Json) Testid(value string) Json {
	j.set("testid", value)
	return j
}

// UseMobileUI sets whether to use mobile UI
func (j Json) UseMobileUI(value bool) Json {
	j.set("useMobileUI", value)
	return j
}

// Value sets the JSON data to display
func (j Json) Value(value string) Json {
	j.set("value", value)
	return j
}

// Visible sets whether the component is visible
func (j Json) Visible(value bool) Json {
	j.set("visible", value)
	return j
}

// VisibleOn sets the expression to show the component
func (j Json) VisibleOn(value string) Json {
	j.set("visibleOn", value)
	return j
}
