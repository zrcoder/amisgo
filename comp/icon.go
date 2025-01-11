package comp

import "github.com/zrcoder/amisgo/model"

// icon represents an icon renderer
type icon model.Schema

// Icon creates a new Icon instance with default type
func Icon() icon {
	return icon{"type": "icon"}
}

func (i icon) set(key string, value any) icon {
	i[key] = value
	return i
}

// AddOnclassName sets the add-on class name
func (i icon) AddOnclassName(value string) icon {
	return i.set("addOnclassName", value)
}

// Badge sets the badge
func (i icon) Badge(value string) icon {
	return i.set("badge", value)
}

// ClassName sets the container CSS class name
func (i icon) ClassName(value string) icon {
	return i.set("className", value)
}

// Disabled sets whether the icon is disabled
func (i icon) Disabled(value bool) icon {
	return i.set("disabled", value)
}

// DisabledOn sets the expression for disabling the icon
func (i icon) DisabledOn(value string) icon {
	return i.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (i icon) EditorSetting(value string) icon {
	return i.set("editorSetting", value)
}

// Hidden sets whether the icon is hidden
func (i icon) Hidden(value bool) icon {
	return i.set("hidden", value)
}

// HiddenOn sets the expression for hiding the icon
func (i icon) HiddenOn(value string) icon {
	return i.set("hiddenOn", value)
}

// Icon sets the icon type
func (i icon) Icon(value string) icon {
	return i.set("icon", value)
}

// ID sets the unique component ID
func (i icon) ID(value string) icon {
	return i.set("id", value)
}

// OnEvent sets the event action configuration
func (i icon) OnEvent(value any) icon {
	return i.set("onEvent", value)
}

// Static sets whether the icon is statically displayed
func (i icon) Static(value bool) icon {
	return i.set("static", value)
}

// StaticClassName sets the static display form item class name
func (i icon) StaticClassName(value string) icon {
	return i.set("staticClassName", value)
}

// StaticInputClassName sets the static display form item value class name
func (i icon) StaticInputClassName(value string) icon {
	return i.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static display form item label class name
func (i icon) StaticLabelClassName(value string) icon {
	return i.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display
func (i icon) StaticOn(value string) icon {
	return i.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (i icon) StaticPlaceholder(value string) icon {
	return i.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (i icon) StaticSchema(value string) icon {
	return i.set("staticSchema", value)
}

// Style sets the component style
func (i icon) Style(value any) icon {
	return i.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (i icon) TestIdBuilder(value string) icon {
	return i.set("testIdBuilder", value)
}

// Testid sets the test ID
func (i icon) Testid(value string) icon {
	return i.set("testid", value)
}

// UseMobileUI sets whether to disable mobile UI styles
func (i icon) UseMobileUI(value bool) icon {
	return i.set("useMobileUI", value)
}

// Vendor sets the vendor (e.g., iconfont, fa)
func (i icon) Vendor(value string) icon {
	return i.set("vendor", value)
}

// Visible sets whether the icon is visible
func (i icon) Visible(value bool) icon {
	return i.set("visible", value)
}

// VisibleOn sets the expression for visibility
func (i icon) VisibleOn(value string) icon {
	return i.set("visibleOn", value)
}
