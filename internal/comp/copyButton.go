package comp

import "github.com/zrcoder/amisgo/schema"

// CopyButton represents a button with copy behavior
type CopyButton schema.Schema

func NewCopyButton() CopyButton {
	return CopyButton{"type": "button", "actionType": "copy"}
}

func (c CopyButton) set(key string, value any) CopyButton {
	c[key] = value
	return c
}

// ActiveClassName sets the CSS class name for the active state
func (c CopyButton) ActiveClassName(value string) CopyButton {
	return c.set("activeClassName", value)
}

// ActiveLevel sets the style for the active state
func (c CopyButton) ActiveLevel(value string) CopyButton {
	return c.set("activeLevel", value)
}

// Badge configures a badge for the button
func (c CopyButton) Badge(value string) CopyButton {
	return c.set("badge", value)
}

// Block enables or disables block-level display
func (c CopyButton) Block(value bool) CopyButton {
	return c.set("block", value)
}

// Body sets the child content for the button
func (c CopyButton) Body(value ...any) CopyButton {
	return c.set("body", value)
}

// ClassName sets the CSS class name for the container
func (c CopyButton) ClassName(value string) CopyButton {
	return c.set("className", value)
}

// Close configures whether to close the popup after action completion
func (c CopyButton) Close(value string) CopyButton {
	return c.set("close", value)
}

// ConfirmText sets the confirmation prompt text
func (c CopyButton) ConfirmText(value string) CopyButton {
	return c.set("confirmText", value)
}

// ConfirmTitle sets the title for the confirmation popup
func (c CopyButton) ConfirmTitle(value any) CopyButton {
	return c.set("confirmTitle", value)
}

// Content sets the content to be copied
func (c CopyButton) Content(value string) CopyButton {
	return c.set("content", value)
}

// Copy configures the copy content settings
func (c CopyButton) Copy(value string) CopyButton {
	return c.set("copy", value)
}

// CountDown sets the disabled countdown duration after clicking (in seconds)
func (c CopyButton) CountDown(value string) CopyButton {
	return c.set("countDown", value)
}

// CountDownTpl customizes the countdown text
func (c CopyButton) CountDownTpl(value string) CopyButton {
	return c.set("countDownTpl", value)
}

// Disabled enables or disables the button
func (c CopyButton) Disabled(value bool) CopyButton {
	return c.set("disabled", value)
}

// DisabledOn sets a conditional expression for disabling the button
func (c CopyButton) DisabledOn(value string) CopyButton {
	return c.set("disabledOn", value)
}

// DisabledTip sets the tooltip text when the button is disabled
func (c CopyButton) DisabledTip(value string) CopyButton {
	return c.set("disabledTip", value)
}

// EditorSetting configures editor-specific settings
func (c CopyButton) EditorSetting(value string) CopyButton {
	return c.set("editorSetting", value)
}

// Hidden controls the overall visibility of the button
func (c CopyButton) Hidden(value bool) CopyButton {
	return c.set("hidden", value)
}

// HiddenOn sets a conditional expression for hiding the button
func (c CopyButton) HiddenOn(value string) CopyButton {
	return c.set("hiddenOn", value)
}

// HotKey sets the keyboard shortcut for the button
func (c CopyButton) HotKey(value string) CopyButton {
	return c.set("hotKey", value)
}

// Icon sets the button icon
func (c CopyButton) Icon(value string) CopyButton {
	return c.set("icon", value)
}

// IconClassName sets the CSS class name for the icon
func (c CopyButton) IconClassName(value string) CopyButton {
	return c.set("iconClassName", value)
}

// ID sets a unique identifier for user behavior tracking
func (c CopyButton) ID(value string) CopyButton {
	return c.set("id", value)
}

// Label sets the button text
func (c CopyButton) Label(value string) CopyButton {
	return c.set("label", value)
}

// Level sets the button style
func (c CopyButton) Level(value string) CopyButton {
	return c.set("level", value)
}

// LoadingClassName sets the CSS class name for the loading state
func (c CopyButton) LoadingClassName(value string) CopyButton {
	return c.set("loadingClassName", value)
}

// LoadingOn configures the condition for displaying the loading effect
func (c CopyButton) LoadingOn(value string) CopyButton {
	return c.set("loadingOn", value)
}

// MergeData determines whether to merge popup data into the parent scope
func (c CopyButton) MergeData(value bool) CopyButton {
	return c.set("mergeData", value)
}

// OnClick sets a custom event handler function
func (c CopyButton) OnClick(value string) CopyButton {
	return c.set("onClick", value)
}

// OnEvent configures event-driven actions for the component
func (c CopyButton) OnEvent(value any) CopyButton {
	return c.set("onEvent", value)
}

// Primary determines whether the button is a primary action button
func (c CopyButton) Primary(value bool) CopyButton {
	return c.set("primary", value)
}

// RequireSelected requires element selection before button can be clicked
func (c CopyButton) RequireSelected(value bool) CopyButton {
	return c.set("requireSelected", value)
}

// Required configures field validation settings
func (c CopyButton) Required(value string) CopyButton {
	return c.set("required", value)
}

// RightIcon sets the icon on the right side of the button
func (c CopyButton) RightIcon(value string) CopyButton {
	return c.set("rightIcon", value)
}

// RightIconClassName sets the CSS class name for the right-side icon
func (c CopyButton) RightIconClassName(value string) CopyButton {
	return c.set("rightIconClassName", value)
}

// Size sets the button size
func (c CopyButton) Size(value string) CopyButton {
	return c.set("size", value)
}

// Static enables or disables static display mode
func (c CopyButton) Static(value bool) CopyButton {
	return c.set("static", value)
}

// StaticClassName sets the CSS class name for static form item display
func (c CopyButton) StaticClassName(value string) CopyButton {
	return c.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input value display
func (c CopyButton) StaticInputClassName(value string) CopyButton {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (c CopyButton) StaticLabelClassName(value string) CopyButton {
	return c.set("staticLabelClassName", value)
}

// StaticOn sets a conditional expression for static display
func (c CopyButton) StaticOn(value string) CopyButton {
	return c.set("staticOn", value)
}

// StaticPlaceholder defines a placeholder for empty static values
func (c CopyButton) StaticPlaceholder(value string) CopyButton {
	return c.set("staticPlaceholder", value)
}

// StaticSchema sets the schema.Schema for static display
func (c CopyButton) StaticSchema(value string) CopyButton {
	return c.set("staticSchema", value)
}

// Style sets custom inline styles for the component
func (c CopyButton) Style(value any) CopyButton {
	return c.set("style", value)
}

// Target specifies the object to trigger the action
func (c CopyButton) Target(value string) CopyButton {
	return c.set("target", value)
}

// TestIdBuilder configures test ID generation
func (c CopyButton) TestIdBuilder(value string) CopyButton {
	return c.set("testIdBuilder", value)
}

// Testid sets a specific test identifier
func (c CopyButton) Testid(value string) CopyButton {
	return c.set("testid", value)
}

// Tooltip sets the tooltip information
func (c CopyButton) Tooltip(value string) CopyButton {
	return c.set("tooltip", value)
}

// TooltipPlacement sets the position of the tooltip
func (c CopyButton) TooltipPlacement(value string) CopyButton {
	return c.set("tooltipPlacement", value)
}

// UseMobileUI enables or disables mobile-specific UI styling
func (c CopyButton) UseMobileUI(value bool) CopyButton {
	return c.set("useMobileUI", value)
}

// Visible controls the overall visibility of the component
func (c CopyButton) Visible(value bool) CopyButton {
	return c.set("visible", value)
}

// VisibleOn sets a conditional expression for component visibility
func (c CopyButton) VisibleOn(value string) CopyButton {
	return c.set("visibleOn", value)
}
