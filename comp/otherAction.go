package comp

import "github.com/zrcoder/amisgo/model"

type otherAction model.Schema

func OtherAction() otherAction {
	return otherAction{}.set("type", "button").set("actionType", "prev")
}

// set sets a field value
func (oa otherAction) set(key string, value any) otherAction {
	oa[key] = value
	return oa
}

// ActionType sets the action type
func (oa otherAction) ActionType(value string) otherAction {
	return oa.set("actionType", value)
}

// ActiveClassName sets the active class name
func (oa otherAction) ActiveClassName(value string) otherAction {
	return oa.set("activeClassName", value)
}

// ActiveLevel sets the active level style
func (oa otherAction) ActiveLevel(value string) otherAction {
	return oa.set("activeLevel", value)
}

// Badge sets the badge
func (oa otherAction) Badge(value string) otherAction {
	return oa.set("badge", value)
}

// Block sets block display
func (oa otherAction) Block(value bool) otherAction {
	return oa.set("block", value)
}

// Body sets the body content
func (oa otherAction) Body(value ...any) otherAction {
	return oa.set("body", value)
}

// ClassName sets the container CSS class name
func (oa otherAction) ClassName(value string) otherAction {
	return oa.set("className", value)
}

// Close sets the close action
func (oa otherAction) Close(value string) otherAction {
	return oa.set("close", value)
}

// ConfirmText sets the confirmation text
func (oa otherAction) ConfirmText(value string) otherAction {
	return oa.set("confirmText", value)
}

// ConfirmTitle sets the confirmation title
func (oa otherAction) ConfirmTitle(value any) otherAction {
	return oa.set("confirmTitle", value)
}

// CountDown sets the countdown timer
func (oa otherAction) CountDown(value string) otherAction {
	return oa.set("countDown", value)
}

// CountDownTpl sets the countdown template
func (oa otherAction) CountDownTpl(value string) otherAction {
	return oa.set("countDownTpl", value)
}

// Disabled sets the disabled state
func (oa otherAction) Disabled(value bool) otherAction {
	return oa.set("disabled", value)
}

// DisabledOn sets the disabled expression
func (oa otherAction) DisabledOn(value string) otherAction {
	return oa.set("disabledOn", value)
}

// DisabledTip sets the disabled tooltip
func (oa otherAction) DisabledTip(value string) otherAction {
	return oa.set("disabledTip", value)
}

// EditorSetting sets the editor setting
func (oa otherAction) EditorSetting(value string) otherAction {
	return oa.set("editorSetting", value)
}

// Hidden sets the hidden state
func (oa otherAction) Hidden(value bool) otherAction {
	return oa.set("hidden", value)
}

// HiddenOn sets the hidden expression
func (oa otherAction) HiddenOn(value string) otherAction {
	return oa.set("hiddenOn", value)
}

// HotKey sets the keyboard shortcut
func (oa otherAction) HotKey(value string) otherAction {
	return oa.set("hotKey", value)
}

// Icon sets the button icon
func (oa otherAction) Icon(value string) otherAction {
	return oa.set("icon", value)
}

// IconClassName sets the icon CSS class name
func (oa otherAction) IconClassName(value string) otherAction {
	return oa.set("iconClassName", value)
}

// ID sets the button ID
func (oa otherAction) ID(value string) otherAction {
	return oa.set("id", value)
}

// Label sets the button label
func (oa otherAction) Label(value string) otherAction {
	return oa.set("label", value)
}

// Level sets the button style level
func (oa otherAction) Level(value string) otherAction {
	return oa.set("level", value)
}

// LoadingClassName sets the loading CSS class name
func (oa otherAction) LoadingClassName(value string) otherAction {
	return oa.set("loadingClassName", value)
}

// LoadingOn sets the loading expression
func (oa otherAction) LoadingOn(value string) otherAction {
	return oa.set("loadingOn", value)
}

// MergeData sets the merge data flag
func (oa otherAction) MergeData(value bool) otherAction {
	return oa.set("mergeData", value)
}

// OnClick sets the custom click event handler
func (oa otherAction) OnClick(value string) otherAction {
	return oa.set("onClick", value)
}

// OnEvent sets the event configuration
func (oa otherAction) OnEvent(value any) otherAction {
	return oa.set("onEvent", value)
}

// Primary sets the primary button flag
func (oa otherAction) Primary(value bool) otherAction {
	return oa.set("primary", value)
}

// RequireSelected sets the require selected flag
func (oa otherAction) RequireSelected(value bool) otherAction {
	return oa.set("requireSelected", value)
}

// Required sets the required field
func (oa otherAction) Required(value string) otherAction {
	return oa.set("required", value)
}

// RightIcon sets the right icon
func (oa otherAction) RightIcon(value string) otherAction {
	return oa.set("rightIcon", value)
}

// RightIconClassName sets the right icon CSS class name
func (oa otherAction) RightIconClassName(value string) otherAction {
	return oa.set("rightIconClassName", value)
}

// Size sets the button size
func (oa otherAction) Size(value string) otherAction {
	return oa.set("size", value)
}

// Static sets the static display flag
func (oa otherAction) Static(value bool) otherAction {
	return oa.set("static", value)
}

// StaticClassName sets the static display CSS class name
func (oa otherAction) StaticClassName(value string) otherAction {
	return oa.set("staticClassName", value)
}

// StaticInputClassName sets the static input CSS class name
func (oa otherAction) StaticInputClassName(value string) otherAction {
	return oa.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static label CSS class name
func (oa otherAction) StaticLabelClassName(value string) otherAction {
	return oa.set("staticLabelClassName", value)
}

// StaticOn sets the static display expression
func (oa otherAction) StaticOn(value string) otherAction {
	return oa.set("staticOn", value)
}

// StaticPlaceholder sets the static placeholder
func (oa otherAction) StaticPlaceholder(value string) otherAction {
	return oa.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema
func (oa otherAction) StaticSchema(value string) otherAction {
	return oa.set("staticSchema", value)
}

// Style sets the component style
func (oa otherAction) Style(value any) otherAction {
	return oa.set("style", value)
}

// Target sets the target for the action
func (oa otherAction) Target(value string) otherAction {
	return oa.set("target", value)
}

// TestIdBuilder sets the custom test ID builder
func (oa otherAction) TestIdBuilder(value string) otherAction {
	return oa.set("testIdBuilder", value)
}

// TestID sets the test ID
func (oa otherAction) TestID(value string) otherAction {
	return oa.set("testid", value)
}

// Tooltip sets the tooltip text
func (oa otherAction) Tooltip(value string) otherAction {
	return oa.set("tooltip", value)
}

// TooltipPlacement sets the tooltip placement
func (oa otherAction) TooltipPlacement(value string) otherAction {
	return oa.set("tooltipPlacement", value)
}

// Type sets the button type
func (oa otherAction) Type(value string) otherAction {
	return oa.set("type", value)
}

// UseMobileUI sets the mobile UI flag
func (oa otherAction) UseMobileUI(value bool) otherAction {
	return oa.set("useMobileUI", value)
}

// Visible sets the visibility
func (oa otherAction) Visible(value bool) otherAction {
	return oa.set("visible", value)
}

// VisibleOn sets the visibility expression
func (oa otherAction) VisibleOn(value string) otherAction {
	return oa.set("visibleOn", value)
}
