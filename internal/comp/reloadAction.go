package comp

import "github.com/zrcoder/amisgo/schema"

// ReloadAction reload action
type ReloadAction schema.Schema

func NewReloadAction() ReloadAction {
	return ReloadAction{"type": "button", "actionType": "reload"}
}

func (ra ReloadAction) set(key string, value any) ReloadAction {
	ra[key] = value
	return ra
}

// ActiveClassName sets active class name
func (ra ReloadAction) ActiveClassName(value string) ReloadAction {
	return ra.set("activeClassName", value)
}

// ActiveLevel sets active level
func (ra ReloadAction) ActiveLevel(value string) ReloadAction {
	return ra.set("activeLevel", value)
}

// Badge sets badge
func (ra ReloadAction) Badge(value string) ReloadAction {
	return ra.set("badge", value)
}

// Block sets block
func (ra ReloadAction) Block(value bool) ReloadAction {
	return ra.set("block", value)
}

// Body sets body
func (ra ReloadAction) Body(value ...any) ReloadAction {
	return ra.set("body", value)
}

// ClassName sets class name
func (ra ReloadAction) ClassName(value string) ReloadAction {
	return ra.set("className", value)
}

// Close sets close
func (ra ReloadAction) Close(value string) ReloadAction {
	return ra.set("close", value)
}

// ConfirmText sets confirm text
func (ra ReloadAction) ConfirmText(value string) ReloadAction {
	return ra.set("confirmText", value)
}

// ConfirmTitle sets confirm title
func (ra ReloadAction) ConfirmTitle(value any) ReloadAction {
	return ra.set("confirmTitle", value)
}

// CountDown sets count down
func (ra ReloadAction) CountDown(value string) ReloadAction {
	return ra.set("countDown", value)
}

// CountDownTpl sets count down template
func (ra ReloadAction) CountDownTpl(value string) ReloadAction {
	return ra.set("countDownTpl", value)
}

// Disabled sets disabled
func (ra ReloadAction) Disabled(value bool) ReloadAction {
	return ra.set("disabled", value)
}

// DisabledOn sets disabled on
func (ra ReloadAction) DisabledOn(value string) ReloadAction {
	return ra.set("disabledOn", value)
}

// DisabledTip sets disabled tip
func (ra ReloadAction) DisabledTip(value string) ReloadAction {
	return ra.set("disabledTip", value)
}

// EditorSetting sets editor setting
func (ra ReloadAction) EditorSetting(value string) ReloadAction {
	return ra.set("editorSetting", value)
}

// Hidden sets hidden
func (ra ReloadAction) Hidden(value bool) ReloadAction {
	return ra.set("hidden", value)
}

// HiddenOn sets hidden on
func (ra ReloadAction) HiddenOn(value string) ReloadAction {
	return ra.set("hiddenOn", value)
}

// HotKey sets hot key
func (ra ReloadAction) HotKey(value string) ReloadAction {
	return ra.set("hotKey", value)
}

// Icon sets icon
func (ra ReloadAction) Icon(value string) ReloadAction {
	return ra.set("icon", value)
}

// IconClassName sets icon class name
func (ra ReloadAction) IconClassName(value string) ReloadAction {
	return ra.set("iconClassName", value)
}

// ID sets id
func (ra ReloadAction) ID(value string) ReloadAction {
	return ra.set("id", value)
}

// Label sets label
func (ra ReloadAction) Label(value string) ReloadAction {
	return ra.set("label", value)
}

// Level sets level
func (ra ReloadAction) Level(value string) ReloadAction {
	return ra.set("level", value)
}

// LoadingClassName sets loading class name
func (ra ReloadAction) LoadingClassName(value string) ReloadAction {
	return ra.set("loadingClassName", value)
}

// LoadingOn sets loading on
func (ra ReloadAction) LoadingOn(value string) ReloadAction {
	return ra.set("loadingOn", value)
}

// MergeData sets merge data
func (ra ReloadAction) MergeData(value bool) ReloadAction {
	return ra.set("mergeData", value)
}

// OnClick sets on click
func (ra ReloadAction) OnClick(value string) ReloadAction {
	return ra.set("onClick", value)
}

// OnEvent sets on event
func (ra ReloadAction) OnEvent(value Event) ReloadAction {
	return ra.set("onEvent", value)
}

// Primary sets primary
func (ra ReloadAction) Primary(value bool) ReloadAction {
	return ra.set("primary", value)
}

// RequireSelected sets require selected
func (ra ReloadAction) RequireSelected(value bool) ReloadAction {
	return ra.set("requireSelected", value)
}

// Required sets required
func (ra ReloadAction) Required(value string) ReloadAction {
	return ra.set("required", value)
}

// RightIcon sets right icon
func (ra ReloadAction) RightIcon(value string) ReloadAction {
	return ra.set("rightIcon", value)
}

// RightIconClassName sets right icon class name
func (ra ReloadAction) RightIconClassName(value string) ReloadAction {
	return ra.set("rightIconClassName", value)
}

// Size sets size
func (ra ReloadAction) Size(value string) ReloadAction {
	return ra.set("size", value)
}

// Static sets static
func (ra ReloadAction) Static(value bool) ReloadAction {
	return ra.set("static", value)
}

// StaticClassName sets static class name
func (ra ReloadAction) StaticClassName(value string) ReloadAction {
	return ra.set("staticClassName", value)
}

// StaticInputClassName sets static input class name
func (ra ReloadAction) StaticInputClassName(value string) ReloadAction {
	return ra.set("staticInputClassName", value)
}

// StaticLabelClassName sets static label class name
func (ra ReloadAction) StaticLabelClassName(value string) ReloadAction {
	return ra.set("staticLabelClassName", value)
}

// StaticOn sets static on
func (ra ReloadAction) StaticOn(value string) ReloadAction {
	return ra.set("staticOn", value)
}

// StaticPlaceholder sets static placeholder
func (ra ReloadAction) StaticPlaceholder(value string) ReloadAction {
	return ra.set("staticPlaceholder", value)
}

// StaticSchema sets static schema.Schema
func (ra ReloadAction) StaticSchema(value string) ReloadAction {
	return ra.set("staticSchema", value)
}

// Style sets style
func (ra ReloadAction) Style(value any) ReloadAction {
	return ra.set("style", value)
}

// Target sets target
func (ra ReloadAction) Target(value string) ReloadAction {
	return ra.set("target", value)
}

// TestIdBuilder sets test id builder
func (ra ReloadAction) TestIdBuilder(value string) ReloadAction {
	return ra.set("testIdBuilder", value)
}

// Testid sets test id
func (ra ReloadAction) Testid(value string) ReloadAction {
	return ra.set("testid", value)
}

// Tooltip sets tooltip
func (ra ReloadAction) Tooltip(value string) ReloadAction {
	return ra.set("tooltip", value)
}

// TooltipPlacement sets tooltip placement
func (ra ReloadAction) TooltipPlacement(value string) ReloadAction {
	return ra.set("tooltipPlacement", value)
}

// Type sets type
func (ra ReloadAction) Type(value string) ReloadAction {
	return ra.set("type", value)
}

// UseMobileUI sets use mobile ui
func (ra ReloadAction) UseMobileUI(value bool) ReloadAction {
	return ra.set("useMobileUI", value)
}

// Visible sets visible
func (ra ReloadAction) Visible(value bool) ReloadAction {
	return ra.set("visible", value)
}

// VisibleOn sets visible on
func (ra ReloadAction) VisibleOn(value string) ReloadAction {
	return ra.set("visibleOn", value)
}
