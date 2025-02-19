package comp

import "github.com/zrcoder/amisgo/schema"

// ToastAction
type ToastAction schema.Schema

func NewToastAction() ToastAction {
	return ToastAction{"type": "button", "actionType": "toast"}
}

func (ta ToastAction) set(key string, value any) ToastAction {
	ta[key] = value
	return ta
}

// ActiveClassName sets the active class name
func (ta ToastAction) ActiveClassName(value string) ToastAction {
	return ta.set("activeClassName", value)
}

// ActiveLevel sets the active level style
func (ta ToastAction) ActiveLevel(value string) ToastAction {
	return ta.set("activeLevel", value)
}

// Badge sets the badge
func (ta ToastAction) Badge(value string) ToastAction {
	return ta.set("badge", value)
}

// Block sets whether to display as block
func (ta ToastAction) Block(value bool) ToastAction {
	return ta.set("block", value)
}

// Body sets the body content
func (ta ToastAction) Body(value ...any) ToastAction {
	return ta.set("body", value)
}

// ClassName sets the container CSS class name
func (ta ToastAction) ClassName(value string) ToastAction {
	return ta.set("className", value)
}

// Close sets the close action
func (ta ToastAction) Close(value string) ToastAction {
	return ta.set("close", value)
}

// ConfirmText sets the confirmation text
func (ta ToastAction) ConfirmText(value string) ToastAction {
	return ta.set("confirmText", value)
}

// ConfirmTitle sets the confirmation dialog title
func (ta ToastAction) ConfirmTitle(value any) ToastAction {
	return ta.set("confirmTitle", value)
}

// CountDown sets the countdown timer
func (ta ToastAction) CountDown(value string) ToastAction {
	return ta.set("countDown", value)
}

// CountDownTpl sets the countdown template
func (ta ToastAction) CountDownTpl(value string) ToastAction {
	return ta.set("countDownTpl", value)
}

// Disabled sets whether the action is disabled
func (ta ToastAction) Disabled(value bool) ToastAction {
	return ta.set("disabled", value)
}

// DisabledOn sets the disabled expression
func (ta ToastAction) DisabledOn(value string) ToastAction {
	return ta.set("disabledOn", value)
}

// DisabledTip sets the disabled tooltip
func (ta ToastAction) DisabledTip(value string) ToastAction {
	return ta.set("disabledTip", value)
}

// EditorSetting sets the editor configuration
func (ta ToastAction) EditorSetting(value string) ToastAction {
	return ta.set("editorSetting", value)
}

// Hidden sets whether the action is hidden
func (ta ToastAction) Hidden(value bool) ToastAction {
	return ta.set("hidden", value)
}

// HiddenOn sets the hidden expression
func (ta ToastAction) HiddenOn(value string) ToastAction {
	return ta.set("hiddenOn", value)
}

// HotKey sets the keyboard shortcut
func (ta ToastAction) HotKey(value string) ToastAction {
	return ta.set("hotKey", value)
}

// Icon sets the button icon
func (ta ToastAction) Icon(value string) ToastAction {
	return ta.set("icon", value)
}

// IconClassName sets the icon CSS class name
func (ta ToastAction) IconClassName(value string) ToastAction {
	return ta.set("iconClassName", value)
}

// ID sets the button ID
func (ta ToastAction) ID(value string) ToastAction {
	return ta.set("id", value)
}

// Label sets the button label
func (ta ToastAction) Label(value string) ToastAction {
	return ta.set("label", value)
}

// Level sets the button style level
func (ta ToastAction) Level(value string) ToastAction {
	return ta.set("level", value)
}

// LoadingClassName sets the loading CSS class name
func (ta ToastAction) LoadingClassName(value string) ToastAction {
	return ta.set("loadingClassName", value)
}

// LoadingOn sets the loading expression
func (ta ToastAction) LoadingOn(value string) ToastAction {
	return ta.set("loadingOn", value)
}

// MergeData sets whether to merge data into the parent scope
func (ta ToastAction) MergeData(value bool) ToastAction {
	return ta.set("mergeData", value)
}

// OnClick sets the custom event handler
func (ta ToastAction) OnClick(value string) ToastAction {
	return ta.set("onClick", value)
}

// OnEvent sets the event action configuration
func (ta ToastAction) OnEvent(value any) ToastAction {
	return ta.set("onEvent", value)
}

// Primary sets whether the action is primary
func (ta ToastAction) Primary(value bool) ToastAction {
	return ta.set("primary", value)
}

// RequireSelected sets whether selection is required
func (ta ToastAction) RequireSelected(value bool) ToastAction {
	return ta.set("requireSelected", value)
}

// Required sets the required field for form validation
func (ta ToastAction) Required(value string) ToastAction {
	return ta.set("required", value)
}

// RightIcon sets the right icon
func (ta ToastAction) RightIcon(value string) ToastAction {
	return ta.set("rightIcon", value)
}

// RightIconClassName sets the right icon CSS class name
func (ta ToastAction) RightIconClassName(value string) ToastAction {
	return ta.set("rightIconClassName", value)
}

// Size sets the button size
func (ta ToastAction) Size(value string) ToastAction {
	return ta.set("size", value)
}

// Static sets whether to display statically
func (ta ToastAction) Static(value bool) ToastAction {
	return ta.set("static", value)
}

// StaticClassName sets the static display CSS class name
func (ta ToastAction) StaticClassName(value string) ToastAction {
	return ta.set("staticClassName", value)
}

// StaticInputClassName sets the static input CSS class name
func (ta ToastAction) StaticInputClassName(value string) ToastAction {
	return ta.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static label CSS class name
func (ta ToastAction) StaticLabelClassName(value string) ToastAction {
	return ta.set("staticLabelClassName", value)
}

// StaticOn sets the static display expression
func (ta ToastAction) StaticOn(value string) ToastAction {
	return ta.set("staticOn", value)
}

// StaticPlaceholder sets the static display placeholder
func (ta ToastAction) StaticPlaceholder(value string) ToastAction {
	return ta.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema.Schema
func (ta ToastAction) StaticSchema(value string) ToastAction {
	return ta.set("staticSchema", value)
}

// Style sets the component style
func (ta ToastAction) Style(value any) ToastAction {
	return ta.set("style", value)
}

// Target sets the target for the action
func (ta ToastAction) Target(value string) ToastAction {
	return ta.set("target", value)
}

// TestIdBuilder sets the test ID builder
func (ta ToastAction) TestIdBuilder(value string) ToastAction {
	return ta.set("testIdBuilder", value)
}

// Testid sets the test ID
func (ta ToastAction) Testid(value string) ToastAction {
	return ta.set("testid", value)
}

// Toast sets the toast details
func (ta ToastAction) Toast(value any) ToastAction {
	return ta.set("toast", value)
}

// Tooltip sets the tooltip
func (ta ToastAction) Tooltip(value string) ToastAction {
	return ta.set("tooltip", value)
}

// TooltipPlacement sets the tooltip placement
func (ta ToastAction) TooltipPlacement(value string) ToastAction {
	return ta.set("tooltipPlacement", value)
}

// UseMobileUI sets whether to use mobile UI
func (ta ToastAction) UseMobileUI(value bool) ToastAction {
	return ta.set("useMobileUI", value)
}

// Visible sets whether the action is visible
func (ta ToastAction) Visible(value bool) ToastAction {
	return ta.set("visible", value)
}

// VisibleOn sets the visible expression
func (ta ToastAction) VisibleOn(value string) ToastAction {
	return ta.set("visibleOn", value)
}
