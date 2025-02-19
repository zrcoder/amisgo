package comp

import "github.com/zrcoder/amisgo/schema"

// LinkAction represents a link action component.
type LinkAction schema.Schema

func NewLinkAction() LinkAction {
	return LinkAction{"type": "button", "actionType": "link"}
}

// ActiveClassName sets the active class name.
func (la LinkAction) ActiveClassName(value string) LinkAction {
	la.set("activeClassName", value)
	return la
}

// ActiveLevel sets the active level style.
func (la LinkAction) ActiveLevel(value string) LinkAction {
	la.set("activeLevel", value)
	return la
}

// Badge sets the badge.
func (la LinkAction) Badge(value string) LinkAction {
	la.set("badge", value)
	return la
}

// Block sets whether to display as a block.
func (la LinkAction) Block(value bool) LinkAction {
	la.set("block", value)
	return la
}

// Body sets the child content.
func (la LinkAction) Body(value ...any) LinkAction {
	la.set("body", value)
	return la
}

// ClassName sets the container CSS class name.
func (la LinkAction) ClassName(value string) LinkAction {
	la.set("className", value)
	return la
}

// Close sets whether to close the dialog after action.
func (la LinkAction) Close(value string) LinkAction {
	la.set("close", value)
	return la
}

// ConfirmText sets the confirmation text.
func (la LinkAction) ConfirmText(value string) LinkAction {
	la.set("confirmText", value)
	return la
}

// ConfirmTitle sets the confirmation dialog title.
func (la LinkAction) ConfirmTitle(value any) LinkAction {
	la.set("confirmTitle", value)
	return la
}

// CountDown sets the countdown time in seconds.
func (la LinkAction) CountDown(value string) LinkAction {
	la.set("countDown", value)
	return la
}

// CountDownTpl sets the countdown template text.
func (la LinkAction) CountDownTpl(value string) LinkAction {
	la.set("countDownTpl", value)
	return la
}

// Disabled sets whether the action is disabled.
func (la LinkAction) Disabled(value bool) LinkAction {
	la.set("disabled", value)
	return la
}

// DisabledOn sets the expression to disable the action.
func (la LinkAction) DisabledOn(value string) LinkAction {
	la.set("disabledOn", value)
	return la
}

// DisabledTip sets the tooltip text when disabled.
func (la LinkAction) DisabledTip(value string) LinkAction {
	la.set("disabledTip", value)
	return la
}

// EditorSetting sets the editor configuration.
func (la LinkAction) EditorSetting(value string) LinkAction {
	la.set("editorSetting", value)
	return la
}

// Hidden sets whether the action is hidden.
func (la LinkAction) Hidden(value bool) LinkAction {
	la.set("hidden", value)
	return la
}

// HiddenOn sets the expression to hide the action.
func (la LinkAction) HiddenOn(value string) LinkAction {
	la.set("hiddenOn", value)
	return la
}

// HotKey sets the keyboard shortcut.
func (la LinkAction) HotKey(value string) LinkAction {
	la.set("hotKey", value)
	return la
}

// Icon sets the button icon.
func (la LinkAction) Icon(value string) LinkAction {
	la.set("icon", value)
	return la
}

// IconClassName sets the CSS class name for the icon.
func (la LinkAction) IconClassName(value string) LinkAction {
	la.set("iconClassName", value)
	return la
}

// ID sets the ID for user behavior tracking.
func (la LinkAction) ID(value string) LinkAction {
	la.set("id", value)
	return la
}

// Label sets the button text.
func (la LinkAction) Label(value string) LinkAction {
	la.set("label", value)
	return la
}

// Level sets the button style.
func (la LinkAction) Level(value string) LinkAction {
	la.set("level", value)
	return la
}

// Link sets the link URL.
func (la LinkAction) Link(value string) LinkAction {
	la.set("link", value)
	return la
}

// LoadingClassName sets the CSS class name for loading state.
func (la LinkAction) LoadingClassName(value string) LinkAction {
	la.set("loadingClassName", value)
	return la
}

// LoadingOn sets the expression to show loading state.
func (la LinkAction) LoadingOn(value string) LinkAction {
	la.set("loadingOn", value)
	return la
}

// MergeData sets whether to merge data into the parent scope.
func (la LinkAction) MergeData(value bool) LinkAction {
	la.set("mergeData", value)
	return la
}

// OnClick sets the custom event handler.
func (la LinkAction) OnClick(value string) LinkAction {
	la.set("onClick", value)
	return la
}

// OnEvent sets the event action configuration.
func (la LinkAction) OnEvent(value any) LinkAction {
	la.set("onEvent", value)
	return la
}

// Primary sets whether the button is primary.
func (la LinkAction) Primary(value bool) LinkAction {
	la.set("primary", value)
	return la
}

// RequireSelected sets whether selection is required for batch actions.
func (la LinkAction) RequireSelected(value bool) LinkAction {
	la.set("requireSelected", value)
	return la
}

// Required sets the required fields for form validation.
func (la LinkAction) Required(value string) LinkAction {
	la.set("required", value)
	return la
}

// RightIcon sets the right-side button icon.
func (la LinkAction) RightIcon(value string) LinkAction {
	la.set("rightIcon", value)
	return la
}

// RightIconClassName sets the CSS class name for the right icon.
func (la LinkAction) RightIconClassName(value string) LinkAction {
	la.set("rightIconClassName", value)
	return la
}

// Size sets the button size.
func (la LinkAction) Size(value string) LinkAction {
	la.set("size", value)
	return la
}

// Static sets whether to display statically.
func (la LinkAction) Static(value bool) LinkAction {
	la.set("static", value)
	return la
}

// StaticClassName sets the CSS class name for static display.
func (la LinkAction) StaticClassName(value string) LinkAction {
	la.set("staticClassName", value)
	return la
}

// StaticInputClassName sets the CSS class name for static input.
func (la LinkAction) StaticInputClassName(value string) LinkAction {
	la.set("staticInputClassName", value)
	return la
}

// StaticLabelClassName sets the CSS class name for static label.
func (la LinkAction) StaticLabelClassName(value string) LinkAction {
	la.set("staticLabelClassName", value)
	return la
}

// StaticOn sets the expression for static display.
func (la LinkAction) StaticOn(value string) LinkAction {
	la.set("staticOn", value)
	return la
}

// StaticPlaceholder sets the placeholder for static display.
func (la LinkAction) StaticPlaceholder(value string) LinkAction {
	la.set("staticPlaceholder", value)
	return la
}

// StaticSchema sets the schema.Schema for static display.
func (la LinkAction) StaticSchema(value string) LinkAction {
	la.set("staticSchema", value)
	return la
}

// Style sets the component style.
func (la LinkAction) Style(value any) LinkAction {
	la.set("style", value)
	return la
}

// Target sets the target to trigger the action.
func (la LinkAction) Target(value string) LinkAction {
	la.set("target", value)
	return la
}

// TestIdBuilder sets the test ID builder.
func (la LinkAction) TestIdBuilder(value string) LinkAction {
	la.set("testIdBuilder", value)
	return la
}

// Testid sets the test ID.
func (la LinkAction) Testid(value string) LinkAction {
	la.set("testid", value)
	return la
}

// Tooltip sets the tooltip text.
func (la LinkAction) Tooltip(value string) LinkAction {
	la.set("tooltip", value)
	return la
}

// TooltipPlacement sets the tooltip placement.
func (la LinkAction) TooltipPlacement(value string) LinkAction {
	la.set("tooltipPlacement", value)
	return la
}

// UseMobileUI sets whether to use mobile UI.
func (la LinkAction) UseMobileUI(value bool) LinkAction {
	la.set("useMobileUI", value)
	return la
}

// Visible sets whether the action is visible.
func (la LinkAction) Visible(value bool) LinkAction {
	la.set("visible", value)
	return la
}

// VisibleOn sets the expression to show the action.
func (la LinkAction) VisibleOn(value string) LinkAction {
	la.set("visibleOn", value)
	return la
}

// set sets a key-value pair.
func (la LinkAction) set(key string, value any) LinkAction {
	la[key] = value
	return la
}
