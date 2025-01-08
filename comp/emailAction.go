package comp

import "github.com/zrcoder/amisgo/model"

// emailAction represents an email action button
type emailAction model.Schema

// EmailAction creates a new EmailAction instance with default type and actionType
func EmailAction() emailAction {
	return make(emailAction).set("type", "button").set("actionType", "email")
}

func (ea emailAction) set(key string, value any) emailAction {
	ea[key] = value
	return ea
}

// ActiveClassName sets the class name when active
func (ea emailAction) ActiveClassName(value string) emailAction {
	return ea.set("activeClassName", value)
}

// ActiveLevel sets the style when active
func (ea emailAction) ActiveLevel(value string) emailAction {
	return ea.set("activeLevel", value)
}

// Badge sets the badge (documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/badge)
func (ea emailAction) Badge(value string) emailAction {
	return ea.set("badge", value)
}

// Bcc sets the Bcc email address
func (ea emailAction) Bcc(value string) emailAction {
	return ea.set("bcc", value)
}

// Block sets whether to display as block, default is inline
func (ea emailAction) Block(value bool) emailAction {
	return ea.set("block", value)
}

// Body sets the email body
func (ea emailAction) Body(value ...any) emailAction {
	return ea.set("body", value)
}

// Cc sets the Cc email address
func (ea emailAction) Cc(value string) emailAction {
	return ea.set("cc", value)
}

// ClassName sets the container CSS class name
func (ea emailAction) ClassName(value string) emailAction {
	return ea.set("className", value)
}

// Close sets whether to close the dialog after action, or specify the target dialog to close
func (ea emailAction) Close(value string) emailAction {
	return ea.set("close", value)
}

// ConfirmText sets the confirmation text
func (ea emailAction) ConfirmText(value string) emailAction {
	return ea.set("confirmText", value)
}

// ConfirmTitle sets the confirmation dialog title
func (ea emailAction) ConfirmTitle(value any) emailAction {
	return ea.set("confirmTitle", value)
}

// CountDown sets the countdown timer (in seconds) after click
func (ea emailAction) CountDown(value string) emailAction {
	return ea.set("countDown", value)
}

// CountDownTpl sets the custom countdown text
func (ea emailAction) CountDownTpl(value string) emailAction {
	return ea.set("countDownTpl", value)
}

// Disabled sets whether the button is disabled
func (ea emailAction) Disabled(value bool) emailAction {
	return ea.set("disabled", value)
}

// DisabledOn sets the expression to disable the button
func (ea emailAction) DisabledOn(value string) emailAction {
	return ea.set("disabledOn", value)
}

// DisabledTip sets the tooltip text when disabled
func (ea emailAction) DisabledTip(value string) emailAction {
	return ea.set("disabledTip", value)
}

// EditorSetting sets the editor configuration, ignored at runtime
func (ea emailAction) EditorSetting(value string) emailAction {
	return ea.set("editorSetting", value)
}

// Hidden sets whether the button is hidden
func (ea emailAction) Hidden(value bool) emailAction {
	return ea.set("hidden", value)
}

// HiddenOn sets the expression to hide the button
func (ea emailAction) HiddenOn(value string) emailAction {
	return ea.set("hiddenOn", value)
}

// HotKey sets the keyboard shortcut
func (ea emailAction) HotKey(value string) emailAction {
	return ea.set("hotKey", value)
}

// Icon sets the button icon, iconfont class name
func (ea emailAction) Icon(value string) emailAction {
	return ea.set("icon", value)
}

// IconClassName sets the CSS class name for the icon
func (ea emailAction) IconClassName(value string) emailAction {
	return ea.set("iconClassName", value)
}

// ID sets the ID for user behavior tracking
func (ea emailAction) ID(value string) emailAction {
	return ea.set("id", value)
}

// Label sets the button text
func (ea emailAction) Label(value string) emailAction {
	return ea.set("label", value)
}

// Level sets the button style
func (ea emailAction) Level(value string) emailAction {
	return ea.set("level", value)
}

// LoadingClassName sets the CSS class name for loading state
func (ea emailAction) LoadingClassName(value string) emailAction {
	return ea.set("loadingClassName", value)
}

// LoadingOn sets the expression to show loading effect
func (ea emailAction) LoadingOn(value string) emailAction {
	return ea.set("loadingOn", value)
}

// MergeData sets whether to merge data from the dialog to the parent scope
func (ea emailAction) MergeData(value bool) emailAction {
	return ea.set("mergeData", value)
}

// OnClick sets the custom event handler
func (ea emailAction) OnClick(value string) emailAction {
	return ea.set("onClick", value)
}

// OnEvent sets the event action configuration
func (ea emailAction) OnEvent(value any) emailAction {
	return ea.set("onEvent", value)
}

// Primary sets the button as primary
func (ea emailAction) Primary(value bool) emailAction {
	return ea.set("primary", value)
}

// RequireSelected sets whether the button requires selected items to be clickable
func (ea emailAction) RequireSelected(value bool) emailAction {
	return ea.set("requireSelected", value)
}

// Required sets the required fields for validation before triggering the action
func (ea emailAction) Required(value string) emailAction {
	return ea.set("required", value)
}

// RightIcon sets the right icon, iconfont class name
func (ea emailAction) RightIcon(value string) emailAction {
	return ea.set("rightIcon", value)
}

// RightIconClassName sets the CSS class name for the right icon
func (ea emailAction) RightIconClassName(value string) emailAction {
	return ea.set("rightIconClassName", value)
}

// Size sets the button size
func (ea emailAction) Size(value string) emailAction {
	return ea.set("size", value)
}

// Static sets whether to display statically
func (ea emailAction) Static(value bool) emailAction {
	return ea.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (ea emailAction) StaticClassName(value string) emailAction {
	return ea.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (ea emailAction) StaticInputClassName(value string) emailAction {
	return ea.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (ea emailAction) StaticLabelClassName(value string) emailAction {
	return ea.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display
func (ea emailAction) StaticOn(value string) emailAction {
	return ea.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (ea emailAction) StaticPlaceholder(value string) emailAction {
	return ea.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (ea emailAction) StaticSchema(value string) emailAction {
	return ea.set("staticSchema", value)
}

// Style sets the component style
func (ea emailAction) Style(value any) emailAction {
	return ea.set("style", value)
}

// Subject sets the email subject
func (ea emailAction) Subject(value string) emailAction {
	return ea.set("subject", value)
}

// Target sets the target to trigger the action
func (ea emailAction) Target(value string) emailAction {
	return ea.set("target", value)
}

// TestIdBuilder sets the test ID builder
func (ea emailAction) TestIdBuilder(value string) emailAction {
	return ea.set("testIdBuilder", value)
}

// Testid sets the test ID
func (ea emailAction) Testid(value string) emailAction {
	return ea.set("testid", value)
}

// To sets the recipient email address
func (ea emailAction) To(value string) emailAction {
	return ea.set("to", value)
}

// Tooltip sets the tooltip text
func (ea emailAction) Tooltip(value string) emailAction {
	return ea.set("tooltip", value)
}

// TooltipPlacement sets the tooltip placement
func (ea emailAction) TooltipPlacement(value string) emailAction {
	return ea.set("tooltipPlacement", value)
}

// UseMobileUI sets whether to disable mobile UI styles
func (ea emailAction) UseMobileUI(value bool) emailAction {
	return ea.set("useMobileUI", value)
}

// Visible sets whether the button is visible
func (ea emailAction) Visible(value bool) emailAction {
	return ea.set("visible", value)
}

// VisibleOn sets the expression to show the button
func (ea emailAction) VisibleOn(value string) emailAction {
	return ea.set("visibleOn", value)
}
