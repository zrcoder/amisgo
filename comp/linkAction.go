package comp

import "github.com/zrcoder/amisgo/model"

// linkAction represents a link action component.
type linkAction model.Schema

// LinkAction creates a new LinkAction instance with default type and actionType.
func LinkAction() linkAction {
	return make(linkAction).set("type", "button").set("actionType", "link")
}

// ActiveClassName sets the active class name.
func (la linkAction) ActiveClassName(value string) linkAction {
	la.set("activeClassName", value)
	return la
}

// ActiveLevel sets the active level style.
func (la linkAction) ActiveLevel(value string) linkAction {
	la.set("activeLevel", value)
	return la
}

// Badge sets the badge.
func (la linkAction) Badge(value string) linkAction {
	la.set("badge", value)
	return la
}

// Block sets whether to display as a block.
func (la linkAction) Block(value bool) linkAction {
	la.set("block", value)
	return la
}

// Body sets the child content.
func (la linkAction) Body(value ...any) linkAction {
	la.set("body", value)
	return la
}

// ClassName sets the container CSS class name.
func (la linkAction) ClassName(value string) linkAction {
	la.set("className", value)
	return la
}

// Close sets whether to close the dialog after action.
func (la linkAction) Close(value string) linkAction {
	la.set("close", value)
	return la
}

// ConfirmText sets the confirmation text.
func (la linkAction) ConfirmText(value string) linkAction {
	la.set("confirmText", value)
	return la
}

// ConfirmTitle sets the confirmation dialog title.
func (la linkAction) ConfirmTitle(value any) linkAction {
	la.set("confirmTitle", value)
	return la
}

// CountDown sets the countdown time in seconds.
func (la linkAction) CountDown(value string) linkAction {
	la.set("countDown", value)
	return la
}

// CountDownTpl sets the countdown template text.
func (la linkAction) CountDownTpl(value string) linkAction {
	la.set("countDownTpl", value)
	return la
}

// Disabled sets whether the action is disabled.
func (la linkAction) Disabled(value bool) linkAction {
	la.set("disabled", value)
	return la
}

// DisabledOn sets the expression to disable the action.
func (la linkAction) DisabledOn(value string) linkAction {
	la.set("disabledOn", value)
	return la
}

// DisabledTip sets the tooltip text when disabled.
func (la linkAction) DisabledTip(value string) linkAction {
	la.set("disabledTip", value)
	return la
}

// EditorSetting sets the editor configuration.
func (la linkAction) EditorSetting(value string) linkAction {
	la.set("editorSetting", value)
	return la
}

// Hidden sets whether the action is hidden.
func (la linkAction) Hidden(value bool) linkAction {
	la.set("hidden", value)
	return la
}

// HiddenOn sets the expression to hide the action.
func (la linkAction) HiddenOn(value string) linkAction {
	la.set("hiddenOn", value)
	return la
}

// HotKey sets the keyboard shortcut.
func (la linkAction) HotKey(value string) linkAction {
	la.set("hotKey", value)
	return la
}

// Icon sets the button icon.
func (la linkAction) Icon(value string) linkAction {
	la.set("icon", value)
	return la
}

// IconClassName sets the CSS class name for the icon.
func (la linkAction) IconClassName(value string) linkAction {
	la.set("iconClassName", value)
	return la
}

// ID sets the ID for user behavior tracking.
func (la linkAction) ID(value string) linkAction {
	la.set("id", value)
	return la
}

// Label sets the button text.
func (la linkAction) Label(value string) linkAction {
	la.set("label", value)
	return la
}

// Level sets the button style.
func (la linkAction) Level(value string) linkAction {
	la.set("level", value)
	return la
}

// Link sets the link URL.
func (la linkAction) Link(value string) linkAction {
	la.set("link", value)
	return la
}

// LoadingClassName sets the CSS class name for loading state.
func (la linkAction) LoadingClassName(value string) linkAction {
	la.set("loadingClassName", value)
	return la
}

// LoadingOn sets the expression to show loading state.
func (la linkAction) LoadingOn(value string) linkAction {
	la.set("loadingOn", value)
	return la
}

// MergeData sets whether to merge data into the parent scope.
func (la linkAction) MergeData(value bool) linkAction {
	la.set("mergeData", value)
	return la
}

// OnClick sets the custom event handler.
func (la linkAction) OnClick(value string) linkAction {
	la.set("onClick", value)
	return la
}

// OnEvent sets the event action configuration.
func (la linkAction) OnEvent(value any) linkAction {
	la.set("onEvent", value)
	return la
}

// Primary sets whether the button is primary.
func (la linkAction) Primary(value bool) linkAction {
	la.set("primary", value)
	return la
}

// RequireSelected sets whether selection is required for batch actions.
func (la linkAction) RequireSelected(value bool) linkAction {
	la.set("requireSelected", value)
	return la
}

// Required sets the required fields for form validation.
func (la linkAction) Required(value string) linkAction {
	la.set("required", value)
	return la
}

// RightIcon sets the right-side button icon.
func (la linkAction) RightIcon(value string) linkAction {
	la.set("rightIcon", value)
	return la
}

// RightIconClassName sets the CSS class name for the right icon.
func (la linkAction) RightIconClassName(value string) linkAction {
	la.set("rightIconClassName", value)
	return la
}

// Size sets the button size.
func (la linkAction) Size(value string) linkAction {
	la.set("size", value)
	return la
}

// Static sets whether to display statically.
func (la linkAction) Static(value bool) linkAction {
	la.set("static", value)
	return la
}

// StaticClassName sets the CSS class name for static display.
func (la linkAction) StaticClassName(value string) linkAction {
	la.set("staticClassName", value)
	return la
}

// StaticInputClassName sets the CSS class name for static input.
func (la linkAction) StaticInputClassName(value string) linkAction {
	la.set("staticInputClassName", value)
	return la
}

// StaticLabelClassName sets the CSS class name for static label.
func (la linkAction) StaticLabelClassName(value string) linkAction {
	la.set("staticLabelClassName", value)
	return la
}

// StaticOn sets the expression for static display.
func (la linkAction) StaticOn(value string) linkAction {
	la.set("staticOn", value)
	return la
}

// StaticPlaceholder sets the placeholder for static display.
func (la linkAction) StaticPlaceholder(value string) linkAction {
	la.set("staticPlaceholder", value)
	return la
}

// StaticSchema sets the schema for static display.
func (la linkAction) StaticSchema(value string) linkAction {
	la.set("staticSchema", value)
	return la
}

// Style sets the component style.
func (la linkAction) Style(value any) linkAction {
	la.set("style", value)
	return la
}

// Target sets the target to trigger the action.
func (la linkAction) Target(value string) linkAction {
	la.set("target", value)
	return la
}

// TestIdBuilder sets the test ID builder.
func (la linkAction) TestIdBuilder(value string) linkAction {
	la.set("testIdBuilder", value)
	return la
}

// Testid sets the test ID.
func (la linkAction) Testid(value string) linkAction {
	la.set("testid", value)
	return la
}

// Tooltip sets the tooltip text.
func (la linkAction) Tooltip(value string) linkAction {
	la.set("tooltip", value)
	return la
}

// TooltipPlacement sets the tooltip placement.
func (la linkAction) TooltipPlacement(value string) linkAction {
	la.set("tooltipPlacement", value)
	return la
}

// UseMobileUI sets whether to use mobile UI.
func (la linkAction) UseMobileUI(value bool) linkAction {
	la.set("useMobileUI", value)
	return la
}

// Visible sets whether the action is visible.
func (la linkAction) Visible(value bool) linkAction {
	la.set("visible", value)
	return la
}

// VisibleOn sets the expression to show the action.
func (la linkAction) VisibleOn(value string) linkAction {
	la.set("visibleOn", value)
	return la
}

// set sets a key-value pair.
func (la linkAction) set(key string, value any) linkAction {
	la[key] = value
	return la
}
