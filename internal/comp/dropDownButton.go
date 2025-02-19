package comp

import "github.com/zrcoder/amisgo/schema"

// DropdownButton represents a dropdown button renderer.
type DropdownButton schema.Schema

// NewDropdownButton creates a new NewDropdownButton instance with the default type 'dropdown-button'.
func NewDropdownButton() DropdownButton {
	return DropdownButton{"type": "dropdown-button"}
}

func (d DropdownButton) set(key string, value any) DropdownButton {
	d[key] = value
	return d
}

// Align sets the alignment. Options: left | right.
func (d DropdownButton) Align(value string) DropdownButton {
	return d.set("align", value)
}

// Block sets whether the button occupies a full line (display: block).
func (d DropdownButton) Block(value bool) DropdownButton {
	return d.set("block", value)
}

// Body sets the content area.
func (d DropdownButton) Body(value ...any) DropdownButton {
	return d.set("body", value)
}

// BtnClassName sets the CSS class name for the button.
func (d DropdownButton) BtnClassName(value string) DropdownButton {
	return d.set("btnClassName", value)
}

// Buttons sets the button collection, supporting grouping.
func (d DropdownButton) Buttons(value string) DropdownButton {
	return d.set("buttons", value)
}

// ClassName sets the container CSS class name.
func (d DropdownButton) ClassName(value string) DropdownButton {
	return d.set("className", value)
}

// CloseOnClick sets whether to close on content click.
func (d DropdownButton) CloseOnClick(value bool) DropdownButton {
	return d.set("closeOnClick", value)
}

// CloseOnOutside sets whether to close on outside click.
func (d DropdownButton) CloseOnOutside(value bool) DropdownButton {
	return d.set("closeOnOutside", value)
}

// Disabled sets whether the button is disabled.
func (d DropdownButton) Disabled(value bool) DropdownButton {
	return d.set("disabled", value)
}

// DisabledOn sets the expression to determine if the button is disabled.
func (d DropdownButton) DisabledOn(value string) DropdownButton {
	return d.set("disabledOn", value)
}

// EditorSetting sets the editor configuration, ignored at runtime.
func (d DropdownButton) EditorSetting(value string) DropdownButton {
	return d.set("editorSetting", value)
}

// Hidden sets whether the button is hidden.
func (d DropdownButton) Hidden(value bool) DropdownButton {
	return d.set("hidden", value)
}

// HiddenOn sets the expression to determine if the button is hidden.
func (d DropdownButton) HiddenOn(value string) DropdownButton {
	return d.set("hiddenOn", value)
}

// HideCaret sets whether to hide the dropdown button arrow.
func (d DropdownButton) HideCaret(value bool) DropdownButton {
	return d.set("hideCaret", value)
}

// IconOnly sets whether to display only the icon.
func (d DropdownButton) IconOnly(value bool) DropdownButton {
	return d.set("iconOnly", value)
}

// ID sets the unique component ID, mainly for logging.
func (d DropdownButton) ID(value string) DropdownButton {
	return d.set("id", value)
}

// Label sets the button text.
func (d DropdownButton) Label(value string) DropdownButton {
	return d.set("label", value)
}

// Level sets the button level/style. Options: info | success | danger | warning | primary | link.
func (d DropdownButton) Level(value string) DropdownButton {
	return d.set("level", value)
}

// MenuClassName sets the menu CSS class name.
func (d DropdownButton) MenuClassName(value string) DropdownButton {
	return d.set("menuClassName", value)
}

// OnEvent sets the event action configuration.
func (d DropdownButton) OnEvent(value any) DropdownButton {
	return d.set("onEvent", value)
}

// OverlayPlacement sets the overlay placement.
func (d DropdownButton) OverlayPlacement(value string) DropdownButton {
	return d.set("overlayPlacement", value)
}

// RightIcon sets the right-side icon.
func (d DropdownButton) RightIcon(value string) DropdownButton {
	return d.set("rightIcon", value)
}

// Size sets the button size. Options: xs | sm | md | lg.
func (d DropdownButton) Size(value string) DropdownButton {
	return d.set("size", value)
}

// Static sets whether to display statically.
func (d DropdownButton) Static(value bool) DropdownButton {
	return d.set("static", value)
}

// StaticClassName sets the static display form item class name.
func (d DropdownButton) StaticClassName(value string) DropdownButton {
	return d.set("staticClassName", value)
}

// StaticInputClassName sets the static display form item value class name.
func (d DropdownButton) StaticInputClassName(value string) DropdownButton {
	return d.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static display form item label class name.
func (d DropdownButton) StaticLabelClassName(value string) DropdownButton {
	return d.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the button is displayed statically.
func (d DropdownButton) StaticOn(value string) DropdownButton {
	return d.set("staticOn", value)
}

// StaticPlaceholder sets the static display placeholder for empty values.
func (d DropdownButton) StaticPlaceholder(value string) DropdownButton {
	return d.set("staticPlaceholder", value)
}

// StaticSchema sets the static display schema.Schema.
func (d DropdownButton) StaticSchema(value string) DropdownButton {
	return d.set("staticSchema", value)
}

// Style sets the component style.
func (d DropdownButton) Style(value any) DropdownButton {
	return d.set("style", value)
}

// TestIdBuilder sets the test ID builder function.
func (d DropdownButton) TestIdBuilder(value string) DropdownButton {
	return d.set("testIdBuilder", value)
}

// Testid sets the test ID.
func (d DropdownButton) Testid(value string) DropdownButton {
	return d.set("testid", value)
}

// Trigger sets the trigger condition. Default is click. Options: click | hover.
func (d DropdownButton) Trigger(value string) DropdownButton {
	return d.set("trigger", value)
}

// UseMobileUI sets whether to disable mobile UI styles.
func (d DropdownButton) UseMobileUI(value bool) DropdownButton {
	return d.set("useMobileUI", value)
}

// Visible sets whether the button is visible.
func (d DropdownButton) Visible(value bool) DropdownButton {
	return d.set("visible", value)
}

// VisibleOn sets the expression to determine if the button is visible.
func (d DropdownButton) VisibleOn(value string) DropdownButton {
	return d.set("visibleOn", value)
}
