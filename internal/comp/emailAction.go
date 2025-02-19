package comp

import "github.com/zrcoder/amisgo/schema"

// EmailAction represents an email action button
type EmailAction schema.Schema

// NewEmailAction creates a new NewEmailAction instance with default type and actionType
func NewEmailAction() EmailAction {
	return EmailAction{"type": "button", "actionType": "email"}
}

func (ea EmailAction) set(key string, value any) EmailAction {
	ea[key] = value
	return ea
}

// ActiveClassName sets the class name when active
func (ea EmailAction) ActiveClassName(value string) EmailAction {
	return ea.set("activeClassName", value)
}

// ActiveLevel sets the style when active
func (ea EmailAction) ActiveLevel(value string) EmailAction {
	return ea.set("activeLevel", value)
}

// Badge sets the badge (documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/badge)
func (ea EmailAction) Badge(value string) EmailAction {
	return ea.set("badge", value)
}

// Bcc sets the Bcc email address
func (ea EmailAction) Bcc(value string) EmailAction {
	return ea.set("bcc", value)
}

// Block sets whether to display as block, default is inline
func (ea EmailAction) Block(value bool) EmailAction {
	return ea.set("block", value)
}

// Body sets the email body
func (ea EmailAction) Body(value ...any) EmailAction {
	return ea.set("body", value)
}

// Cc sets the Cc email address
func (ea EmailAction) Cc(value string) EmailAction {
	return ea.set("cc", value)
}

// ClassName sets the container CSS class name
func (ea EmailAction) ClassName(value string) EmailAction {
	return ea.set("className", value)
}

// Close sets whether to close the dialog after action, or specify the target dialog to close
func (ea EmailAction) Close(value string) EmailAction {
	return ea.set("close", value)
}

// ConfirmText sets the confirmation text
func (ea EmailAction) ConfirmText(value string) EmailAction {
	return ea.set("confirmText", value)
}

// ConfirmTitle sets the confirmation dialog title
func (ea EmailAction) ConfirmTitle(value any) EmailAction {
	return ea.set("confirmTitle", value)
}

// CountDown sets the countdown timer (in seconds) after click
func (ea EmailAction) CountDown(value string) EmailAction {
	return ea.set("countDown", value)
}

// CountDownTpl sets the custom countdown text
func (ea EmailAction) CountDownTpl(value string) EmailAction {
	return ea.set("countDownTpl", value)
}

// Disabled sets whether the button is disabled
func (ea EmailAction) Disabled(value bool) EmailAction {
	return ea.set("disabled", value)
}

// DisabledOn sets the expression to disable the button
func (ea EmailAction) DisabledOn(value string) EmailAction {
	return ea.set("disabledOn", value)
}

// DisabledTip sets the tooltip text when disabled
func (ea EmailAction) DisabledTip(value string) EmailAction {
	return ea.set("disabledTip", value)
}

// EditorSetting sets the editor configuration, ignored at runtime
func (ea EmailAction) EditorSetting(value string) EmailAction {
	return ea.set("editorSetting", value)
}

// Hidden sets whether the button is hidden
func (ea EmailAction) Hidden(value bool) EmailAction {
	return ea.set("hidden", value)
}

// HiddenOn sets the expression to hide the button
func (ea EmailAction) HiddenOn(value string) EmailAction {
	return ea.set("hiddenOn", value)
}

// HotKey sets the keyboard shortcut
func (ea EmailAction) HotKey(value string) EmailAction {
	return ea.set("hotKey", value)
}

// Icon sets the button icon, iconfont class name
func (ea EmailAction) Icon(value string) EmailAction {
	return ea.set("icon", value)
}

// IconClassName sets the CSS class name for the icon
func (ea EmailAction) IconClassName(value string) EmailAction {
	return ea.set("iconClassName", value)
}

// ID sets the ID for user behavior tracking
func (ea EmailAction) ID(value string) EmailAction {
	return ea.set("id", value)
}

// Label sets the button text
func (ea EmailAction) Label(value string) EmailAction {
	return ea.set("label", value)
}

// Level sets the button style
func (ea EmailAction) Level(value string) EmailAction {
	return ea.set("level", value)
}

// LoadingClassName sets the CSS class name for loading state
func (ea EmailAction) LoadingClassName(value string) EmailAction {
	return ea.set("loadingClassName", value)
}

// LoadingOn sets the expression to show loading effect
func (ea EmailAction) LoadingOn(value string) EmailAction {
	return ea.set("loadingOn", value)
}

// MergeData sets whether to merge data from the dialog to the parent scope
func (ea EmailAction) MergeData(value bool) EmailAction {
	return ea.set("mergeData", value)
}

// OnClick sets the custom event handler
func (ea EmailAction) OnClick(value string) EmailAction {
	return ea.set("onClick", value)
}

// OnEvent sets the event action configuration
func (ea EmailAction) OnEvent(value any) EmailAction {
	return ea.set("onEvent", value)
}

// Primary sets the button as primary
func (ea EmailAction) Primary(value bool) EmailAction {
	return ea.set("primary", value)
}

// RequireSelected sets whether the button requires selected items to be clickable
func (ea EmailAction) RequireSelected(value bool) EmailAction {
	return ea.set("requireSelected", value)
}

// Required sets the required fields for validation before triggering the action
func (ea EmailAction) Required(value string) EmailAction {
	return ea.set("required", value)
}

// RightIcon sets the right icon, iconfont class name
func (ea EmailAction) RightIcon(value string) EmailAction {
	return ea.set("rightIcon", value)
}

// RightIconClassName sets the CSS class name for the right icon
func (ea EmailAction) RightIconClassName(value string) EmailAction {
	return ea.set("rightIconClassName", value)
}

// Size sets the button size
func (ea EmailAction) Size(value string) EmailAction {
	return ea.set("size", value)
}

// Static sets whether to display statically
func (ea EmailAction) Static(value bool) EmailAction {
	return ea.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (ea EmailAction) StaticClassName(value string) EmailAction {
	return ea.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (ea EmailAction) StaticInputClassName(value string) EmailAction {
	return ea.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (ea EmailAction) StaticLabelClassName(value string) EmailAction {
	return ea.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display
func (ea EmailAction) StaticOn(value string) EmailAction {
	return ea.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (ea EmailAction) StaticPlaceholder(value string) EmailAction {
	return ea.set("staticPlaceholder", value)
}

// StaticSchema sets the schema.Schema for static display
func (ea EmailAction) StaticSchema(value string) EmailAction {
	return ea.set("staticSchema", value)
}

// Style sets the component style
func (ea EmailAction) Style(value any) EmailAction {
	return ea.set("style", value)
}

// Subject sets the email subject
func (ea EmailAction) Subject(value string) EmailAction {
	return ea.set("subject", value)
}

// Target sets the target to trigger the action
func (ea EmailAction) Target(value string) EmailAction {
	return ea.set("target", value)
}

// TestIdBuilder sets the test ID builder
func (ea EmailAction) TestIdBuilder(value string) EmailAction {
	return ea.set("testIdBuilder", value)
}

// Testid sets the test ID
func (ea EmailAction) Testid(value string) EmailAction {
	return ea.set("testid", value)
}

// To sets the recipient email address
func (ea EmailAction) To(value string) EmailAction {
	return ea.set("to", value)
}

// Tooltip sets the tooltip text
func (ea EmailAction) Tooltip(value string) EmailAction {
	return ea.set("tooltip", value)
}

// TooltipPlacement sets the tooltip placement
func (ea EmailAction) TooltipPlacement(value string) EmailAction {
	return ea.set("tooltipPlacement", value)
}

// UseMobileUI sets whether to disable mobile UI styles
func (ea EmailAction) UseMobileUI(value bool) EmailAction {
	return ea.set("useMobileUI", value)
}

// Visible sets whether the button is visible
func (ea EmailAction) Visible(value bool) EmailAction {
	return ea.set("visible", value)
}

// VisibleOn sets the expression to show the button
func (ea EmailAction) VisibleOn(value string) EmailAction {
	return ea.set("visibleOn", value)
}
