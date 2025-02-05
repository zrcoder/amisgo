package comp

import "github.com/zrcoder/amisgo/model"

// UrlAction represents a URL action button component
type UrlAction model.Schema

func NewUrlAction() UrlAction {
	return UrlAction{"type": "button", "actionType": "url"}
}

func (ua UrlAction) set(key string, value any) UrlAction {
	ua[key] = value
	return ua
}

// ActiveClassName sets the class name when active
func (ua UrlAction) ActiveClassName(value string) UrlAction {
	return ua.set("activeClassName", value)
}

// ActiveLevel sets the style when active
func (ua UrlAction) ActiveLevel(value string) UrlAction {
	return ua.set("activeLevel", value)
}

// Badge sets the badge
func (ua UrlAction) Badge(value string) UrlAction {
	return ua.set("badge", value)
}

// Blank sets whether to open in a new window
func (ua UrlAction) Blank(value bool) UrlAction {
	return ua.set("blank", value)
}

// Block sets whether to display as a block
func (ua UrlAction) Block(value bool) UrlAction {
	return ua.set("block", value)
}

// Body sets the child content
func (ua UrlAction) Body(value ...any) UrlAction {
	return ua.set("body", value)
}

// ClassName sets the container CSS class name
func (ua UrlAction) ClassName(value string) UrlAction {
	return ua.set("className", value)
}

// Close sets whether to close the dialog after action
func (ua UrlAction) Close(value string) UrlAction {
	return ua.set("close", value)
}

// ConfirmText sets the confirmation text
func (ua UrlAction) ConfirmText(value string) UrlAction {
	return ua.set("confirmText", value)
}

// ConfirmTitle sets the confirmation dialog title
func (ua UrlAction) ConfirmTitle(value any) UrlAction {
	return ua.set("confirmTitle", value)
}

// CountDown sets the countdown after click (seconds)
func (ua UrlAction) CountDown(value string) UrlAction {
	return ua.set("countDown", value)
}

// CountDownTpl sets the custom countdown text
func (ua UrlAction) CountDownTpl(value string) UrlAction {
	return ua.set("countDownTpl", value)
}

// Disabled sets whether the button is disabled
func (ua UrlAction) Disabled(value bool) UrlAction {
	return ua.set("disabled", value)
}

// DisabledOn sets the expression to disable the button
func (ua UrlAction) DisabledOn(value string) UrlAction {
	return ua.set("disabledOn", value)
}

// DisabledTip sets the tooltip when disabled
func (ua UrlAction) DisabledTip(value string) UrlAction {
	return ua.set("disabledTip", value)
}

// EditorSetting sets the editor configuration
func (ua UrlAction) EditorSetting(value string) UrlAction {
	return ua.set("editorSetting", value)
}

// Hidden sets whether the button is hidden
func (ua UrlAction) Hidden(value bool) UrlAction {
	return ua.set("hidden", value)
}

// HiddenOn sets the expression to hide the button
func (ua UrlAction) HiddenOn(value string) UrlAction {
	return ua.set("hiddenOn", value)
}

// HotKey sets the keyboard shortcut
func (ua UrlAction) HotKey(value string) UrlAction {
	return ua.set("hotKey", value)
}

// Icon sets the button icon
func (ua UrlAction) Icon(value string) UrlAction {
	return ua.set("icon", value)
}

// IconClassName sets the CSS class name for the icon
func (ua UrlAction) IconClassName(value string) UrlAction {
	return ua.set("iconClassName", value)
}

// ID sets the button ID for tracking
func (ua UrlAction) ID(value string) UrlAction {
	return ua.set("id", value)
}

// Label sets the button label
func (ua UrlAction) Label(value string) UrlAction {
	return ua.set("label", value)
}

// Level sets the button style level
func (ua UrlAction) Level(value string) UrlAction {
	return ua.set("level", value)
}

// Link sets the link URL
func (ua UrlAction) Link(value string) UrlAction {
	return ua.set("link", value)
}

// LoadingClassName sets the CSS class name for loading state
func (ua UrlAction) LoadingClassName(value string) UrlAction {
	return ua.set("loadingClassName", value)
}

// LoadingOn sets the expression to show loading state
func (ua UrlAction) LoadingOn(value string) UrlAction {
	return ua.set("loadingOn", value)
}

// MergeData sets whether to merge data into the parent scope
func (ua UrlAction) MergeData(value bool) UrlAction {
	return ua.set("mergeData", value)
}

// OnClick sets the custom event handler
func (ua UrlAction) OnClick(value string) UrlAction {
	return ua.set("onClick", value)
}

// OnEvent sets the event action configuration
func (ua UrlAction) OnEvent(value any) UrlAction {
	return ua.set("onEvent", value)
}

// Primary sets whether the button is primary
func (ua UrlAction) Primary(value bool) UrlAction {
	return ua.set("primary", value)
}

// RequireSelected sets whether selection is required for batch actions
func (ua UrlAction) RequireSelected(value bool) UrlAction {
	return ua.set("requireSelected", value)
}

// Required sets the required fields for form validation
func (ua UrlAction) Required(value string) UrlAction {
	return ua.set("required", value)
}

// RightIcon sets the right icon
func (ua UrlAction) RightIcon(value string) UrlAction {
	return ua.set("rightIcon", value)
}

// RightIconClassName sets the CSS class name for the right icon
func (ua UrlAction) RightIconClassName(value string) UrlAction {
	return ua.set("rightIconClassName", value)
}

// Size sets the button size
func (ua UrlAction) Size(value string) UrlAction {
	return ua.set("size", value)
}

// Static sets whether to display statically
func (ua UrlAction) Static(value bool) UrlAction {
	return ua.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (ua UrlAction) StaticClassName(value string) UrlAction {
	return ua.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input
func (ua UrlAction) StaticInputClassName(value string) UrlAction {
	return ua.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label
func (ua UrlAction) StaticLabelClassName(value string) UrlAction {
	return ua.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display
func (ua UrlAction) StaticOn(value string) UrlAction {
	return ua.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (ua UrlAction) StaticPlaceholder(value string) UrlAction {
	return ua.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema
func (ua UrlAction) StaticSchema(value string) UrlAction {
	return ua.set("staticSchema", value)
}

// Style sets the component style
func (ua UrlAction) Style(value any) UrlAction {
	return ua.set("style", value)
}

// Target sets the target for the action
func (ua UrlAction) Target(value string) UrlAction {
	return ua.set("target", value)
}

// TestIdBuilder sets the test ID builder
func (ua UrlAction) TestIdBuilder(value string) UrlAction {
	return ua.set("testIdBuilder", value)
}

// Testid sets the test ID
func (ua UrlAction) Testid(value string) UrlAction {
	return ua.set("testid", value)
}

// Tooltip sets the tooltip
func (ua UrlAction) Tooltip(value string) UrlAction {
	return ua.set("tooltip", value)
}

// TooltipPlacement sets the tooltip placement
func (ua UrlAction) TooltipPlacement(value string) UrlAction {
	return ua.set("tooltipPlacement", value)
}

// Type sets the button type
func (ua UrlAction) Type(value string) UrlAction {
	return ua.set("type", value)
}

// Url sets the target URL
func (ua UrlAction) Url(value string) UrlAction {
	return ua.set("url", value)
}

// UseMobileUI sets whether to use mobile UI
func (ua UrlAction) UseMobileUI(value bool) UrlAction {
	return ua.set("useMobileUI", value)
}

// Visible sets whether the button is visible
func (ua UrlAction) Visible(value bool) UrlAction {
	return ua.set("visible", value)
}

// VisibleOn sets the expression for visibility
func (ua UrlAction) VisibleOn(value string) UrlAction {
	return ua.set("visibleOn", value)
}
