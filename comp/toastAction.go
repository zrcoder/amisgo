package comp

import "github.com/zrcoder/amisgo/model"

// toastAction

type toastAction model.Schema

// ToastAction creates a new ToastAction instance
func ToastAction() toastAction {
	return toastAction{}.set("type", "button").set("actionType", "toast")
}

func (ta toastAction) set(key string, value any) toastAction {
	ta[key] = value
	return ta
}

// ActiveClassName sets the active class name
func (ta toastAction) ActiveClassName(value string) toastAction {
	return ta.set("activeClassName", value)
}

// ActiveLevel sets the active level style
func (ta toastAction) ActiveLevel(value string) toastAction {
	return ta.set("activeLevel", value)
}

// Badge sets the badge
func (ta toastAction) Badge(value string) toastAction {
	return ta.set("badge", value)
}

// Block sets whether to display as block
func (ta toastAction) Block(value bool) toastAction {
	return ta.set("block", value)
}

// Body sets the body content
func (ta toastAction) Body(value ...any) toastAction {
	return ta.set("body", value)
}

// ClassName sets the container CSS class name
func (ta toastAction) ClassName(value string) toastAction {
	return ta.set("className", value)
}

// Close sets the close action
func (ta toastAction) Close(value string) toastAction {
	return ta.set("close", value)
}

// ConfirmText sets the confirmation text
func (ta toastAction) ConfirmText(value string) toastAction {
	return ta.set("confirmText", value)
}

// ConfirmTitle sets the confirmation dialog title
func (ta toastAction) ConfirmTitle(value any) toastAction {
	return ta.set("confirmTitle", value)
}

// CountDown sets the countdown timer
func (ta toastAction) CountDown(value string) toastAction {
	return ta.set("countDown", value)
}

// CountDownTpl sets the countdown template
func (ta toastAction) CountDownTpl(value string) toastAction {
	return ta.set("countDownTpl", value)
}

// Disabled sets whether the action is disabled
func (ta toastAction) Disabled(value bool) toastAction {
	return ta.set("disabled", value)
}

// DisabledOn sets the disabled expression
func (ta toastAction) DisabledOn(value string) toastAction {
	return ta.set("disabledOn", value)
}

// DisabledTip sets the disabled tooltip
func (ta toastAction) DisabledTip(value string) toastAction {
	return ta.set("disabledTip", value)
}

// EditorSetting sets the editor configuration
func (ta toastAction) EditorSetting(value string) toastAction {
	return ta.set("editorSetting", value)
}

// Hidden sets whether the action is hidden
func (ta toastAction) Hidden(value bool) toastAction {
	return ta.set("hidden", value)
}

// HiddenOn sets the hidden expression
func (ta toastAction) HiddenOn(value string) toastAction {
	return ta.set("hiddenOn", value)
}

// HotKey sets the keyboard shortcut
func (ta toastAction) HotKey(value string) toastAction {
	return ta.set("hotKey", value)
}

// Icon sets the button icon
func (ta toastAction) Icon(value string) toastAction {
	return ta.set("icon", value)
}

// IconClassName sets the icon CSS class name
func (ta toastAction) IconClassName(value string) toastAction {
	return ta.set("iconClassName", value)
}

// ID sets the button ID
func (ta toastAction) ID(value string) toastAction {
	return ta.set("id", value)
}

// Label sets the button label
func (ta toastAction) Label(value string) toastAction {
	return ta.set("label", value)
}

// Level sets the button style level
func (ta toastAction) Level(value string) toastAction {
	return ta.set("level", value)
}

// LoadingClassName sets the loading CSS class name
func (ta toastAction) LoadingClassName(value string) toastAction {
	return ta.set("loadingClassName", value)
}

// LoadingOn sets the loading expression
func (ta toastAction) LoadingOn(value string) toastAction {
	return ta.set("loadingOn", value)
}

// MergeData sets whether to merge data into the parent scope
func (ta toastAction) MergeData(value bool) toastAction {
	return ta.set("mergeData", value)
}

// OnClick sets the custom event handler
func (ta toastAction) OnClick(value string) toastAction {
	return ta.set("onClick", value)
}

// OnEvent sets the event action configuration
func (ta toastAction) OnEvent(value any) toastAction {
	return ta.set("onEvent", value)
}

// Primary sets whether the action is primary
func (ta toastAction) Primary(value bool) toastAction {
	return ta.set("primary", value)
}

// RequireSelected sets whether selection is required
func (ta toastAction) RequireSelected(value bool) toastAction {
	return ta.set("requireSelected", value)
}

// Required sets the required field for form validation
func (ta toastAction) Required(value string) toastAction {
	return ta.set("required", value)
}

// RightIcon sets the right icon
func (ta toastAction) RightIcon(value string) toastAction {
	return ta.set("rightIcon", value)
}

// RightIconClassName sets the right icon CSS class name
func (ta toastAction) RightIconClassName(value string) toastAction {
	return ta.set("rightIconClassName", value)
}

// Size sets the button size
func (ta toastAction) Size(value string) toastAction {
	return ta.set("size", value)
}

// Static sets whether to display statically
func (ta toastAction) Static(value bool) toastAction {
	return ta.set("static", value)
}

// StaticClassName sets the static display CSS class name
func (ta toastAction) StaticClassName(value string) toastAction {
	return ta.set("staticClassName", value)
}

// StaticInputClassName sets the static input CSS class name
func (ta toastAction) StaticInputClassName(value string) toastAction {
	return ta.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static label CSS class name
func (ta toastAction) StaticLabelClassName(value string) toastAction {
	return ta.set("staticLabelClassName", value)
}

// StaticOn sets the static display expression
func (ta toastAction) StaticOn(value string) toastAction {
	return ta.set("staticOn", value)
}

// StaticPlaceholder sets the static display placeholder
func (ta toastAction) StaticPlaceholder(value string) toastAction {
	return ta.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema
func (ta toastAction) StaticSchema(value string) toastAction {
	return ta.set("staticSchema", value)
}

// Style sets the component style
func (ta toastAction) Style(value any) toastAction {
	return ta.set("style", value)
}

// Target sets the target for the action
func (ta toastAction) Target(value string) toastAction {
	return ta.set("target", value)
}

// TestIdBuilder sets the test ID builder
func (ta toastAction) TestIdBuilder(value string) toastAction {
	return ta.set("testIdBuilder", value)
}

// Testid sets the test ID
func (ta toastAction) Testid(value string) toastAction {
	return ta.set("testid", value)
}

// Toast sets the toast details
func (ta toastAction) Toast(value string) toastAction {
	return ta.set("toast", value)
}

// Tooltip sets the tooltip
func (ta toastAction) Tooltip(value string) toastAction {
	return ta.set("tooltip", value)
}

// TooltipPlacement sets the tooltip placement
func (ta toastAction) TooltipPlacement(value string) toastAction {
	return ta.set("tooltipPlacement", value)
}

// UseMobileUI sets whether to use mobile UI
func (ta toastAction) UseMobileUI(value bool) toastAction {
	return ta.set("useMobileUI", value)
}

// Visible sets whether the action is visible
func (ta toastAction) Visible(value bool) toastAction {
	return ta.set("visible", value)
}

// VisibleOn sets the visible expression
func (ta toastAction) VisibleOn(value string) toastAction {
	return ta.set("visibleOn", value)
}
