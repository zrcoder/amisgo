package comp

import "github.com/zrcoder/amisgo/schema"

// Icon represents an Icon renderer
type Icon schema.Schema

func NewIcon() Icon {
	return Icon{"type": "icon"}
}

func (i Icon) set(key string, value any) Icon {
	i[key] = value
	return i
}

// AddOnclassName sets the add-on class name
func (i Icon) AddOnclassName(value string) Icon {
	return i.set("addOnclassName", value)
}

// Badge sets the badge
func (i Icon) Badge(value string) Icon {
	return i.set("badge", value)
}

// ClassName sets the container CSS class name
func (i Icon) ClassName(value string) Icon {
	return i.set("className", value)
}

// Disabled sets whether the icon is disabled
func (i Icon) Disabled(value bool) Icon {
	return i.set("disabled", value)
}

// DisabledOn sets the expression for disabling the icon
func (i Icon) DisabledOn(value string) Icon {
	return i.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (i Icon) EditorSetting(value string) Icon {
	return i.set("editorSetting", value)
}

// Hidden sets whether the icon is hidden
func (i Icon) Hidden(value bool) Icon {
	return i.set("hidden", value)
}

// HiddenOn sets the expression for hiding the icon
func (i Icon) HiddenOn(value string) Icon {
	return i.set("hiddenOn", value)
}

// Icon sets the icon type
func (i Icon) Icon(value string) Icon {
	return i.set("icon", value)
}

// ID sets the unique component ID
func (i Icon) ID(value string) Icon {
	return i.set("id", value)
}

// OnEvent sets the event action configuration
func (i Icon) OnEvent(value any) Icon {
	return i.set("onEvent", value)
}

// Static sets whether the icon is statically displayed
func (i Icon) Static(value bool) Icon {
	return i.set("static", value)
}

// StaticClassName sets the static display form item class name
func (i Icon) StaticClassName(value string) Icon {
	return i.set("staticClassName", value)
}

// StaticInputClassName sets the static display form item value class name
func (i Icon) StaticInputClassName(value string) Icon {
	return i.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static display form item label class name
func (i Icon) StaticLabelClassName(value string) Icon {
	return i.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display
func (i Icon) StaticOn(value string) Icon {
	return i.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (i Icon) StaticPlaceholder(value string) Icon {
	return i.set("staticPlaceholder", value)
}

// StaticSchema sets the schema.Schema for static display
func (i Icon) StaticSchema(value string) Icon {
	return i.set("staticSchema", value)
}

// Style sets the component style
func (i Icon) Style(value any) Icon {
	return i.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (i Icon) TestIdBuilder(value string) Icon {
	return i.set("testIdBuilder", value)
}

// Testid sets the test ID
func (i Icon) Testid(value string) Icon {
	return i.set("testid", value)
}

// UseMobileUI sets whether to disable mobile UI styles
func (i Icon) UseMobileUI(value bool) Icon {
	return i.set("useMobileUI", value)
}

// Vendor sets the vendor (e.g., iconfont, fa)
func (i Icon) Vendor(value string) Icon {
	return i.set("vendor", value)
}

// Visible sets whether the icon is visible
func (i Icon) Visible(value bool) Icon {
	return i.set("visible", value)
}

// VisibleOn sets the expression for visibility
func (i Icon) VisibleOn(value string) Icon {
	return i.set("visibleOn", value)
}
