package comp

import "github.com/zrcoder/amisgo/schema"

type OtherAction schema.Schema

func NewOtherAction() OtherAction {
	return OtherAction{"type": "button", "actionType": "prev"}
}

// set sets a field value
func (oa OtherAction) set(key string, value any) OtherAction {
	oa[key] = value
	return oa
}

// ActionType sets the action type
func (oa OtherAction) ActionType(value string) OtherAction {
	return oa.set("actionType", value)
}

// ActiveClassName sets the active class name
func (oa OtherAction) ActiveClassName(value string) OtherAction {
	return oa.set("activeClassName", value)
}

// ActiveLevel sets the active level style
func (oa OtherAction) ActiveLevel(value string) OtherAction {
	return oa.set("activeLevel", value)
}

// Badge sets the badge
func (oa OtherAction) Badge(value string) OtherAction {
	return oa.set("badge", value)
}

// Block sets block display
func (oa OtherAction) Block(value bool) OtherAction {
	return oa.set("block", value)
}

// Body sets the body content
func (oa OtherAction) Body(value ...any) OtherAction {
	return oa.set("body", value)
}

// ClassName sets the container CSS class name
func (oa OtherAction) ClassName(value string) OtherAction {
	return oa.set("className", value)
}

// Close sets the close action
func (oa OtherAction) Close(value string) OtherAction {
	return oa.set("close", value)
}

// ConfirmText sets the confirmation text
func (oa OtherAction) ConfirmText(value string) OtherAction {
	return oa.set("confirmText", value)
}

// ConfirmTitle sets the confirmation title
func (oa OtherAction) ConfirmTitle(value any) OtherAction {
	return oa.set("confirmTitle", value)
}

// CountDown sets the countdown timer
func (oa OtherAction) CountDown(value string) OtherAction {
	return oa.set("countDown", value)
}

// CountDownTpl sets the countdown template
func (oa OtherAction) CountDownTpl(value string) OtherAction {
	return oa.set("countDownTpl", value)
}

// Disabled sets the disabled state
func (oa OtherAction) Disabled(value bool) OtherAction {
	return oa.set("disabled", value)
}

// DisabledOn sets the disabled expression
func (oa OtherAction) DisabledOn(value string) OtherAction {
	return oa.set("disabledOn", value)
}

// DisabledTip sets the disabled tooltip
func (oa OtherAction) DisabledTip(value string) OtherAction {
	return oa.set("disabledTip", value)
}

// EditorSetting sets the editor setting
func (oa OtherAction) EditorSetting(value string) OtherAction {
	return oa.set("editorSetting", value)
}

// Hidden sets the hidden state
func (oa OtherAction) Hidden(value bool) OtherAction {
	return oa.set("hidden", value)
}

// HiddenOn sets the hidden expression
func (oa OtherAction) HiddenOn(value string) OtherAction {
	return oa.set("hiddenOn", value)
}

// HotKey sets the keyboard shortcut
func (oa OtherAction) HotKey(value string) OtherAction {
	return oa.set("hotKey", value)
}

// Icon sets the button icon
func (oa OtherAction) Icon(value string) OtherAction {
	return oa.set("icon", value)
}

// IconClassName sets the icon CSS class name
func (oa OtherAction) IconClassName(value string) OtherAction {
	return oa.set("iconClassName", value)
}

// ID sets the button ID
func (oa OtherAction) ID(value string) OtherAction {
	return oa.set("id", value)
}

// Label sets the button label
func (oa OtherAction) Label(value string) OtherAction {
	return oa.set("label", value)
}

// Level sets the button style level
func (oa OtherAction) Level(value string) OtherAction {
	return oa.set("level", value)
}

// LoadingClassName sets the loading CSS class name
func (oa OtherAction) LoadingClassName(value string) OtherAction {
	return oa.set("loadingClassName", value)
}

// LoadingOn sets the loading expression
func (oa OtherAction) LoadingOn(value string) OtherAction {
	return oa.set("loadingOn", value)
}

// MergeData sets the merge data flag
func (oa OtherAction) MergeData(value bool) OtherAction {
	return oa.set("mergeData", value)
}

// OnClick sets the custom click event handler
func (oa OtherAction) OnClick(value string) OtherAction {
	return oa.set("onClick", value)
}

// OnEvent sets the event configuration
func (oa OtherAction) OnEvent(value any) OtherAction {
	return oa.set("onEvent", value)
}

// Primary sets the primary button flag
func (oa OtherAction) Primary(value bool) OtherAction {
	return oa.set("primary", value)
}

// RequireSelected sets the require selected flag
func (oa OtherAction) RequireSelected(value bool) OtherAction {
	return oa.set("requireSelected", value)
}

// Required sets the required field
func (oa OtherAction) Required(value string) OtherAction {
	return oa.set("required", value)
}

// RightIcon sets the right icon
func (oa OtherAction) RightIcon(value string) OtherAction {
	return oa.set("rightIcon", value)
}

// RightIconClassName sets the right icon CSS class name
func (oa OtherAction) RightIconClassName(value string) OtherAction {
	return oa.set("rightIconClassName", value)
}

// Size sets the button size
func (oa OtherAction) Size(value string) OtherAction {
	return oa.set("size", value)
}

// Static sets the static display flag
func (oa OtherAction) Static(value bool) OtherAction {
	return oa.set("static", value)
}

// StaticClassName sets the static display CSS class name
func (oa OtherAction) StaticClassName(value string) OtherAction {
	return oa.set("staticClassName", value)
}

// StaticInputClassName sets the static input CSS class name
func (oa OtherAction) StaticInputClassName(value string) OtherAction {
	return oa.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static label CSS class name
func (oa OtherAction) StaticLabelClassName(value string) OtherAction {
	return oa.set("staticLabelClassName", value)
}

// StaticOn sets the static display expression
func (oa OtherAction) StaticOn(value string) OtherAction {
	return oa.set("staticOn", value)
}

// StaticPlaceholder sets the static placeholder
func (oa OtherAction) StaticPlaceholder(value string) OtherAction {
	return oa.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema.Schema
func (oa OtherAction) StaticSchema(value string) OtherAction {
	return oa.set("staticSchema", value)
}

// Style sets the component style
func (oa OtherAction) Style(value any) OtherAction {
	return oa.set("style", value)
}

// Target sets the target for the action
func (oa OtherAction) Target(value string) OtherAction {
	return oa.set("target", value)
}

// TestIdBuilder sets the custom test ID builder
func (oa OtherAction) TestIdBuilder(value string) OtherAction {
	return oa.set("testIdBuilder", value)
}

// TestID sets the test ID
func (oa OtherAction) TestID(value string) OtherAction {
	return oa.set("testid", value)
}

// Tooltip sets the tooltip text
func (oa OtherAction) Tooltip(value string) OtherAction {
	return oa.set("tooltip", value)
}

// TooltipPlacement sets the tooltip placement
func (oa OtherAction) TooltipPlacement(value string) OtherAction {
	return oa.set("tooltipPlacement", value)
}

// Type sets the button type
func (oa OtherAction) Type(value string) OtherAction {
	return oa.set("type", value)
}

// UseMobileUI sets the mobile UI flag
func (oa OtherAction) UseMobileUI(value bool) OtherAction {
	return oa.set("useMobileUI", value)
}

// Visible sets the visibility
func (oa OtherAction) Visible(value bool) OtherAction {
	return oa.set("visible", value)
}

// VisibleOn sets the visibility expression
func (oa OtherAction) VisibleOn(value string) OtherAction {
	return oa.set("visibleOn", value)
}
