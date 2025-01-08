package comp

import "github.com/zrcoder/amisgo/model"

// dropdownButton represents a dropdown button renderer.
type dropdownButton model.Schema

// DropdownButton creates a new DropdownButton instance with the default type 'dropdown-button'.
func DropdownButton() dropdownButton {
	return make(dropdownButton).set("type", "dropdown-button")
}

func (d dropdownButton) set(key string, value any) dropdownButton {
	d[key] = value
	return d
}

// Align sets the alignment. Options: left | right.
func (d dropdownButton) Align(value string) dropdownButton {
	return d.set("align", value)
}

// Block sets whether the button occupies a full line (display: block).
func (d dropdownButton) Block(value bool) dropdownButton {
	return d.set("block", value)
}

// Body sets the content area.
func (d dropdownButton) Body(value ...any) dropdownButton {
	return d.set("body", value)
}

// BtnClassName sets the CSS class name for the button.
func (d dropdownButton) BtnClassName(value string) dropdownButton {
	return d.set("btnClassName", value)
}

// Buttons sets the button collection, supporting grouping.
func (d dropdownButton) Buttons(value string) dropdownButton {
	return d.set("buttons", value)
}

// ClassName sets the container CSS class name.
func (d dropdownButton) ClassName(value string) dropdownButton {
	return d.set("className", value)
}

// CloseOnClick sets whether to close on content click.
func (d dropdownButton) CloseOnClick(value bool) dropdownButton {
	return d.set("closeOnClick", value)
}

// CloseOnOutside sets whether to close on outside click.
func (d dropdownButton) CloseOnOutside(value bool) dropdownButton {
	return d.set("closeOnOutside", value)
}

// Disabled sets whether the button is disabled.
func (d dropdownButton) Disabled(value bool) dropdownButton {
	return d.set("disabled", value)
}

// DisabledOn sets the expression to determine if the button is disabled.
func (d dropdownButton) DisabledOn(value string) dropdownButton {
	return d.set("disabledOn", value)
}

// EditorSetting sets the editor configuration, ignored at runtime.
func (d dropdownButton) EditorSetting(value string) dropdownButton {
	return d.set("editorSetting", value)
}

// Hidden sets whether the button is hidden.
func (d dropdownButton) Hidden(value bool) dropdownButton {
	return d.set("hidden", value)
}

// HiddenOn sets the expression to determine if the button is hidden.
func (d dropdownButton) HiddenOn(value string) dropdownButton {
	return d.set("hiddenOn", value)
}

// HideCaret sets whether to hide the dropdown button arrow.
func (d dropdownButton) HideCaret(value bool) dropdownButton {
	return d.set("hideCaret", value)
}

// IconOnly sets whether to display only the icon.
func (d dropdownButton) IconOnly(value bool) dropdownButton {
	return d.set("iconOnly", value)
}

// ID sets the unique component ID, mainly for logging.
func (d dropdownButton) ID(value string) dropdownButton {
	return d.set("id", value)
}

// Label sets the button text.
func (d dropdownButton) Label(value string) dropdownButton {
	return d.set("label", value)
}

// Level sets the button level/style. Options: info | success | danger | warning | primary | link.
func (d dropdownButton) Level(value string) dropdownButton {
	return d.set("level", value)
}

// MenuClassName sets the menu CSS class name.
func (d dropdownButton) MenuClassName(value string) dropdownButton {
	return d.set("menuClassName", value)
}

// OnEvent sets the event action configuration.
func (d dropdownButton) OnEvent(value any) dropdownButton {
	return d.set("onEvent", value)
}

// OverlayPlacement sets the overlay placement.
func (d dropdownButton) OverlayPlacement(value string) dropdownButton {
	return d.set("overlayPlacement", value)
}

// RightIcon sets the right-side icon.
func (d dropdownButton) RightIcon(value string) dropdownButton {
	return d.set("rightIcon", value)
}

// Size sets the button size. Options: xs | sm | md | lg.
func (d dropdownButton) Size(value string) dropdownButton {
	return d.set("size", value)
}

// Static sets whether to display statically.
func (d dropdownButton) Static(value bool) dropdownButton {
	return d.set("static", value)
}

// StaticClassName sets the static display form item class name.
func (d dropdownButton) StaticClassName(value string) dropdownButton {
	return d.set("staticClassName", value)
}

// StaticInputClassName sets the static display form item value class name.
func (d dropdownButton) StaticInputClassName(value string) dropdownButton {
	return d.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static display form item label class name.
func (d dropdownButton) StaticLabelClassName(value string) dropdownButton {
	return d.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the button is displayed statically.
func (d dropdownButton) StaticOn(value string) dropdownButton {
	return d.set("staticOn", value)
}

// StaticPlaceholder sets the static display placeholder for empty values.
func (d dropdownButton) StaticPlaceholder(value string) dropdownButton {
	return d.set("staticPlaceholder", value)
}

// StaticSchema sets the static display schema.
func (d dropdownButton) StaticSchema(value string) dropdownButton {
	return d.set("staticSchema", value)
}

// Style sets the component style.
func (d dropdownButton) Style(value any) dropdownButton {
	return d.set("style", value)
}

// TestIdBuilder sets the test ID builder function.
func (d dropdownButton) TestIdBuilder(value string) dropdownButton {
	return d.set("testIdBuilder", value)
}

// Testid sets the test ID.
func (d dropdownButton) Testid(value string) dropdownButton {
	return d.set("testid", value)
}

// Trigger sets the trigger condition. Default is click. Options: click | hover.
func (d dropdownButton) Trigger(value string) dropdownButton {
	return d.set("trigger", value)
}

// UseMobileUI sets whether to disable mobile UI styles.
func (d dropdownButton) UseMobileUI(value bool) dropdownButton {
	return d.set("useMobileUI", value)
}

// Visible sets whether the button is visible.
func (d dropdownButton) Visible(value bool) dropdownButton {
	return d.set("visible", value)
}

// VisibleOn sets the expression to determine if the button is visible.
func (d dropdownButton) VisibleOn(value string) dropdownButton {
	return d.set("visibleOn", value)
}
