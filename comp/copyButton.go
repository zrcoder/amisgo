package comp

import "github.com/zrcoder/amisgo/model"

// CopyButton represents a button with copy behavior
type copyButton model.Schema

// CopyButton creates a new CopyButton instance with default copy action type
func CopyButton() copyButton {
	return make(copyButton).set("type", "button").set("actionType", "copy")
}

func (c copyButton) set(key string, value any) copyButton {
	c[key] = value
	return c
}

// ActiveClassName sets the CSS class name for the active state
func (c copyButton) ActiveClassName(value string) copyButton {
	return c.set("activeClassName", value)
}

// ActiveLevel sets the style for the active state
func (c copyButton) ActiveLevel(value string) copyButton {
	return c.set("activeLevel", value)
}

// Badge configures a badge for the button
func (c copyButton) Badge(value string) copyButton {
	return c.set("badge", value)
}

// Block enables or disables block-level display
func (c copyButton) Block(value bool) copyButton {
	return c.set("block", value)
}

// Body sets the child content for the button
func (c copyButton) Body(value ...any) copyButton {
	return c.set("body", value)
}

// ClassName sets the CSS class name for the container
func (c copyButton) ClassName(value string) copyButton {
	return c.set("className", value)
}

// Close configures whether to close the popup after action completion
func (c copyButton) Close(value string) copyButton {
	return c.set("close", value)
}

// ConfirmText sets the confirmation prompt text
func (c copyButton) ConfirmText(value string) copyButton {
	return c.set("confirmText", value)
}

// ConfirmTitle sets the title for the confirmation popup
func (c copyButton) ConfirmTitle(value any) copyButton {
	return c.set("confirmTitle", value)
}

// Content sets the content to be copied
func (c copyButton) Content(value string) copyButton {
	return c.set("content", value)
}

// Copy configures the copy content settings
func (c copyButton) Copy(value string) copyButton {
	return c.set("copy", value)
}

// CountDown sets the disabled countdown duration after clicking (in seconds)
func (c copyButton) CountDown(value string) copyButton {
	return c.set("countDown", value)
}

// CountDownTpl customizes the countdown text
func (c copyButton) CountDownTpl(value string) copyButton {
	return c.set("countDownTpl", value)
}

// Disabled enables or disables the button
func (c copyButton) Disabled(value bool) copyButton {
	return c.set("disabled", value)
}

// DisabledOn sets a conditional expression for disabling the button
func (c copyButton) DisabledOn(value string) copyButton {
	return c.set("disabledOn", value)
}

// DisabledTip sets the tooltip text when the button is disabled
func (c copyButton) DisabledTip(value string) copyButton {
	return c.set("disabledTip", value)
}

// EditorSetting configures editor-specific settings
func (c copyButton) EditorSetting(value string) copyButton {
	return c.set("editorSetting", value)
}

// Hidden controls the overall visibility of the button
func (c copyButton) Hidden(value bool) copyButton {
	return c.set("hidden", value)
}

// HiddenOn sets a conditional expression for hiding the button
func (c copyButton) HiddenOn(value string) copyButton {
	return c.set("hiddenOn", value)
}

// HotKey sets the keyboard shortcut for the button
func (c copyButton) HotKey(value string) copyButton {
	return c.set("hotKey", value)
}

// Icon sets the button icon
func (c copyButton) Icon(value string) copyButton {
	return c.set("icon", value)
}

// IconClassName sets the CSS class name for the icon
func (c copyButton) IconClassName(value string) copyButton {
	return c.set("iconClassName", value)
}

// ID sets a unique identifier for user behavior tracking
func (c copyButton) ID(value string) copyButton {
	return c.set("id", value)
}

// Label sets the button text
func (c copyButton) Label(value string) copyButton {
	return c.set("label", value)
}

// Level sets the button style
func (c copyButton) Level(value string) copyButton {
	return c.set("level", value)
}

// LoadingClassName sets the CSS class name for the loading state
func (c copyButton) LoadingClassName(value string) copyButton {
	return c.set("loadingClassName", value)
}

// LoadingOn configures the condition for displaying the loading effect
func (c copyButton) LoadingOn(value string) copyButton {
	return c.set("loadingOn", value)
}

// MergeData determines whether to merge popup data into the parent scope
func (c copyButton) MergeData(value bool) copyButton {
	return c.set("mergeData", value)
}

// OnClick sets a custom event handler function
func (c copyButton) OnClick(value string) copyButton {
	return c.set("onClick", value)
}

// OnEvent configures event-driven actions for the component
func (c copyButton) OnEvent(value any) copyButton {
	return c.set("onEvent", value)
}

// Primary determines whether the button is a primary action button
func (c copyButton) Primary(value bool) copyButton {
	return c.set("primary", value)
}

// RequireSelected requires element selection before button can be clicked
func (c copyButton) RequireSelected(value bool) copyButton {
	return c.set("requireSelected", value)
}

// Required configures field validation settings
func (c copyButton) Required(value string) copyButton {
	return c.set("required", value)
}

// RightIcon sets the icon on the right side of the button
func (c copyButton) RightIcon(value string) copyButton {
	return c.set("rightIcon", value)
}

// RightIconClassName sets the CSS class name for the right-side icon
func (c copyButton) RightIconClassName(value string) copyButton {
	return c.set("rightIconClassName", value)
}

// Size sets the button size
func (c copyButton) Size(value string) copyButton {
	return c.set("size", value)
}

// Static enables or disables static display mode
func (c copyButton) Static(value bool) copyButton {
	return c.set("static", value)
}

// StaticClassName sets the CSS class name for static form item display
func (c copyButton) StaticClassName(value string) copyButton {
	return c.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input value display
func (c copyButton) StaticInputClassName(value string) copyButton {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (c copyButton) StaticLabelClassName(value string) copyButton {
	return c.set("staticLabelClassName", value)
}

// StaticOn sets a conditional expression for static display
func (c copyButton) StaticOn(value string) copyButton {
	return c.set("staticOn", value)
}

// StaticPlaceholder defines a placeholder for empty static values
func (c copyButton) StaticPlaceholder(value string) copyButton {
	return c.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (c copyButton) StaticSchema(value string) copyButton {
	return c.set("staticSchema", value)
}

// Style sets custom inline styles for the component
func (c copyButton) Style(value any) copyButton {
	return c.set("style", value)
}

// Target specifies the object to trigger the action
func (c copyButton) Target(value string) copyButton {
	return c.set("target", value)
}

// TestIdBuilder configures test ID generation
func (c copyButton) TestIdBuilder(value string) copyButton {
	return c.set("testIdBuilder", value)
}

// Testid sets a specific test identifier
func (c copyButton) Testid(value string) copyButton {
	return c.set("testid", value)
}

// Tooltip sets the tooltip information
func (c copyButton) Tooltip(value string) copyButton {
	return c.set("tooltip", value)
}

// TooltipPlacement sets the position of the tooltip
func (c copyButton) TooltipPlacement(value string) copyButton {
	return c.set("tooltipPlacement", value)
}

// UseMobileUI enables or disables mobile-specific UI styling
func (c copyButton) UseMobileUI(value bool) copyButton {
	return c.set("useMobileUI", value)
}

// Visible controls the overall visibility of the component
func (c copyButton) Visible(value bool) copyButton {
	return c.set("visible", value)
}

// VisibleOn sets a conditional expression for component visibility
func (c copyButton) VisibleOn(value string) copyButton {
	return c.set("visibleOn", value)
}
