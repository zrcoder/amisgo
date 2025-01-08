package comp

import "github.com/zrcoder/amisgo/model"

// urlAction represents a URL action button component
type urlAction model.Schema

// UrlAction creates a new UrlAction instance
func UrlAction() urlAction {
	return urlAction{}.set("type", "button").set("actionType", "url")
}

func (ua urlAction) set(key string, value any) urlAction {
	ua[key] = value
	return ua
}

// ActiveClassName sets the class name when active
func (ua urlAction) ActiveClassName(value string) urlAction {
	return ua.set("activeClassName", value)
}

// ActiveLevel sets the style when active
func (ua urlAction) ActiveLevel(value string) urlAction {
	return ua.set("activeLevel", value)
}

// Badge sets the badge
func (ua urlAction) Badge(value string) urlAction {
	return ua.set("badge", value)
}

// Blank sets whether to open in a new window
func (ua urlAction) Blank(value bool) urlAction {
	return ua.set("blank", value)
}

// Block sets whether to display as a block
func (ua urlAction) Block(value bool) urlAction {
	return ua.set("block", value)
}

// Body sets the child content
func (ua urlAction) Body(value ...any) urlAction {
	return ua.set("body", value)
}

// ClassName sets the container CSS class name
func (ua urlAction) ClassName(value string) urlAction {
	return ua.set("className", value)
}

// Close sets whether to close the dialog after action
func (ua urlAction) Close(value string) urlAction {
	return ua.set("close", value)
}

// ConfirmText sets the confirmation text
func (ua urlAction) ConfirmText(value string) urlAction {
	return ua.set("confirmText", value)
}

// ConfirmTitle sets the confirmation dialog title
func (ua urlAction) ConfirmTitle(value any) urlAction {
	return ua.set("confirmTitle", value)
}

// CountDown sets the countdown after click (seconds)
func (ua urlAction) CountDown(value string) urlAction {
	return ua.set("countDown", value)
}

// CountDownTpl sets the custom countdown text
func (ua urlAction) CountDownTpl(value string) urlAction {
	return ua.set("countDownTpl", value)
}

// Disabled sets whether the button is disabled
func (ua urlAction) Disabled(value bool) urlAction {
	return ua.set("disabled", value)
}

// DisabledOn sets the expression to disable the button
func (ua urlAction) DisabledOn(value string) urlAction {
	return ua.set("disabledOn", value)
}

// DisabledTip sets the tooltip when disabled
func (ua urlAction) DisabledTip(value string) urlAction {
	return ua.set("disabledTip", value)
}

// EditorSetting sets the editor configuration
func (ua urlAction) EditorSetting(value string) urlAction {
	return ua.set("editorSetting", value)
}

// Hidden sets whether the button is hidden
func (ua urlAction) Hidden(value bool) urlAction {
	return ua.set("hidden", value)
}

// HiddenOn sets the expression to hide the button
func (ua urlAction) HiddenOn(value string) urlAction {
	return ua.set("hiddenOn", value)
}

// HotKey sets the keyboard shortcut
func (ua urlAction) HotKey(value string) urlAction {
	return ua.set("hotKey", value)
}

// Icon sets the button icon
func (ua urlAction) Icon(value string) urlAction {
	return ua.set("icon", value)
}

// IconClassName sets the CSS class name for the icon
func (ua urlAction) IconClassName(value string) urlAction {
	return ua.set("iconClassName", value)
}

// ID sets the button ID for tracking
func (ua urlAction) ID(value string) urlAction {
	return ua.set("id", value)
}

// Label sets the button label
func (ua urlAction) Label(value string) urlAction {
	return ua.set("label", value)
}

// Level sets the button style level
func (ua urlAction) Level(value string) urlAction {
	return ua.set("level", value)
}

// Link sets the link URL
func (ua urlAction) Link(value string) urlAction {
	return ua.set("link", value)
}

// LoadingClassName sets the CSS class name for loading state
func (ua urlAction) LoadingClassName(value string) urlAction {
	return ua.set("loadingClassName", value)
}

// LoadingOn sets the expression to show loading state
func (ua urlAction) LoadingOn(value string) urlAction {
	return ua.set("loadingOn", value)
}

// MergeData sets whether to merge data into the parent scope
func (ua urlAction) MergeData(value bool) urlAction {
	return ua.set("mergeData", value)
}

// OnClick sets the custom event handler
func (ua urlAction) OnClick(value string) urlAction {
	return ua.set("onClick", value)
}

// OnEvent sets the event action configuration
func (ua urlAction) OnEvent(value any) urlAction {
	return ua.set("onEvent", value)
}

// Primary sets whether the button is primary
func (ua urlAction) Primary(value bool) urlAction {
	return ua.set("primary", value)
}

// RequireSelected sets whether selection is required for batch actions
func (ua urlAction) RequireSelected(value bool) urlAction {
	return ua.set("requireSelected", value)
}

// Required sets the required fields for form validation
func (ua urlAction) Required(value string) urlAction {
	return ua.set("required", value)
}

// RightIcon sets the right icon
func (ua urlAction) RightIcon(value string) urlAction {
	return ua.set("rightIcon", value)
}

// RightIconClassName sets the CSS class name for the right icon
func (ua urlAction) RightIconClassName(value string) urlAction {
	return ua.set("rightIconClassName", value)
}

// Size sets the button size
func (ua urlAction) Size(value string) urlAction {
	return ua.set("size", value)
}

// Static sets whether to display statically
func (ua urlAction) Static(value bool) urlAction {
	return ua.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (ua urlAction) StaticClassName(value string) urlAction {
	return ua.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input
func (ua urlAction) StaticInputClassName(value string) urlAction {
	return ua.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label
func (ua urlAction) StaticLabelClassName(value string) urlAction {
	return ua.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display
func (ua urlAction) StaticOn(value string) urlAction {
	return ua.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (ua urlAction) StaticPlaceholder(value string) urlAction {
	return ua.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema
func (ua urlAction) StaticSchema(value string) urlAction {
	return ua.set("staticSchema", value)
}

// Style sets the component style
func (ua urlAction) Style(value any) urlAction {
	return ua.set("style", value)
}

// Target sets the target for the action
func (ua urlAction) Target(value string) urlAction {
	return ua.set("target", value)
}

// TestIdBuilder sets the test ID builder
func (ua urlAction) TestIdBuilder(value string) urlAction {
	return ua.set("testIdBuilder", value)
}

// Testid sets the test ID
func (ua urlAction) Testid(value string) urlAction {
	return ua.set("testid", value)
}

// Tooltip sets the tooltip
func (ua urlAction) Tooltip(value string) urlAction {
	return ua.set("tooltip", value)
}

// TooltipPlacement sets the tooltip placement
func (ua urlAction) TooltipPlacement(value string) urlAction {
	return ua.set("tooltipPlacement", value)
}

// Type sets the button type
func (ua urlAction) Type(value string) urlAction {
	return ua.set("type", value)
}

// Url sets the target URL
func (ua urlAction) Url(value string) urlAction {
	return ua.set("url", value)
}

// UseMobileUI sets whether to use mobile UI
func (ua urlAction) UseMobileUI(value bool) urlAction {
	return ua.set("useMobileUI", value)
}

// Visible sets whether the button is visible
func (ua urlAction) Visible(value bool) urlAction {
	return ua.set("visible", value)
}

// VisibleOn sets the expression for visibility
func (ua urlAction) VisibleOn(value string) urlAction {
	return ua.set("visibleOn", value)
}
