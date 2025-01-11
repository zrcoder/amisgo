package comp

import "github.com/zrcoder/amisgo/model"

// reloadAction reload action
type reloadAction model.Schema

// ReloadAction creates a new ReloadAction instance
func ReloadAction() reloadAction {
	return reloadAction{"type": "button", "actionType": "reload"}
}

func (ra reloadAction) set(key string, value any) reloadAction {
	ra[key] = value
	return ra
}

// ActiveClassName sets active class name
func (ra reloadAction) ActiveClassName(value string) reloadAction {
	return ra.set("activeClassName", value)
}

// ActiveLevel sets active level
func (ra reloadAction) ActiveLevel(value string) reloadAction {
	return ra.set("activeLevel", value)
}

// Badge sets badge
func (ra reloadAction) Badge(value string) reloadAction {
	return ra.set("badge", value)
}

// Block sets block
func (ra reloadAction) Block(value bool) reloadAction {
	return ra.set("block", value)
}

// Body sets body
func (ra reloadAction) Body(value ...any) reloadAction {
	return ra.set("body", value)
}

// ClassName sets class name
func (ra reloadAction) ClassName(value string) reloadAction {
	return ra.set("className", value)
}

// Close sets close
func (ra reloadAction) Close(value string) reloadAction {
	return ra.set("close", value)
}

// ConfirmText sets confirm text
func (ra reloadAction) ConfirmText(value string) reloadAction {
	return ra.set("confirmText", value)
}

// ConfirmTitle sets confirm title
func (ra reloadAction) ConfirmTitle(value any) reloadAction {
	return ra.set("confirmTitle", value)
}

// CountDown sets count down
func (ra reloadAction) CountDown(value string) reloadAction {
	return ra.set("countDown", value)
}

// CountDownTpl sets count down template
func (ra reloadAction) CountDownTpl(value string) reloadAction {
	return ra.set("countDownTpl", value)
}

// Disabled sets disabled
func (ra reloadAction) Disabled(value bool) reloadAction {
	return ra.set("disabled", value)
}

// DisabledOn sets disabled on
func (ra reloadAction) DisabledOn(value string) reloadAction {
	return ra.set("disabledOn", value)
}

// DisabledTip sets disabled tip
func (ra reloadAction) DisabledTip(value string) reloadAction {
	return ra.set("disabledTip", value)
}

// EditorSetting sets editor setting
func (ra reloadAction) EditorSetting(value string) reloadAction {
	return ra.set("editorSetting", value)
}

// Hidden sets hidden
func (ra reloadAction) Hidden(value bool) reloadAction {
	return ra.set("hidden", value)
}

// HiddenOn sets hidden on
func (ra reloadAction) HiddenOn(value string) reloadAction {
	return ra.set("hiddenOn", value)
}

// HotKey sets hot key
func (ra reloadAction) HotKey(value string) reloadAction {
	return ra.set("hotKey", value)
}

// Icon sets icon
func (ra reloadAction) Icon(value string) reloadAction {
	return ra.set("icon", value)
}

// IconClassName sets icon class name
func (ra reloadAction) IconClassName(value string) reloadAction {
	return ra.set("iconClassName", value)
}

// ID sets id
func (ra reloadAction) ID(value string) reloadAction {
	return ra.set("id", value)
}

// Label sets label
func (ra reloadAction) Label(value string) reloadAction {
	return ra.set("label", value)
}

// Level sets level
func (ra reloadAction) Level(value string) reloadAction {
	return ra.set("level", value)
}

// LoadingClassName sets loading class name
func (ra reloadAction) LoadingClassName(value string) reloadAction {
	return ra.set("loadingClassName", value)
}

// LoadingOn sets loading on
func (ra reloadAction) LoadingOn(value string) reloadAction {
	return ra.set("loadingOn", value)
}

// MergeData sets merge data
func (ra reloadAction) MergeData(value bool) reloadAction {
	return ra.set("mergeData", value)
}

// OnClick sets on click
func (ra reloadAction) OnClick(value string) reloadAction {
	return ra.set("onClick", value)
}

// OnEvent sets on event
func (ra reloadAction) OnEvent(value any) reloadAction {
	return ra.set("onEvent", value)
}

// Primary sets primary
func (ra reloadAction) Primary(value bool) reloadAction {
	return ra.set("primary", value)
}

// RequireSelected sets require selected
func (ra reloadAction) RequireSelected(value bool) reloadAction {
	return ra.set("requireSelected", value)
}

// Required sets required
func (ra reloadAction) Required(value string) reloadAction {
	return ra.set("required", value)
}

// RightIcon sets right icon
func (ra reloadAction) RightIcon(value string) reloadAction {
	return ra.set("rightIcon", value)
}

// RightIconClassName sets right icon class name
func (ra reloadAction) RightIconClassName(value string) reloadAction {
	return ra.set("rightIconClassName", value)
}

// Size sets size
func (ra reloadAction) Size(value string) reloadAction {
	return ra.set("size", value)
}

// Static sets static
func (ra reloadAction) Static(value bool) reloadAction {
	return ra.set("static", value)
}

// StaticClassName sets static class name
func (ra reloadAction) StaticClassName(value string) reloadAction {
	return ra.set("staticClassName", value)
}

// StaticInputClassName sets static input class name
func (ra reloadAction) StaticInputClassName(value string) reloadAction {
	return ra.set("staticInputClassName", value)
}

// StaticLabelClassName sets static label class name
func (ra reloadAction) StaticLabelClassName(value string) reloadAction {
	return ra.set("staticLabelClassName", value)
}

// StaticOn sets static on
func (ra reloadAction) StaticOn(value string) reloadAction {
	return ra.set("staticOn", value)
}

// StaticPlaceholder sets static placeholder
func (ra reloadAction) StaticPlaceholder(value string) reloadAction {
	return ra.set("staticPlaceholder", value)
}

// StaticSchema sets static schema
func (ra reloadAction) StaticSchema(value string) reloadAction {
	return ra.set("staticSchema", value)
}

// Style sets style
func (ra reloadAction) Style(value any) reloadAction {
	return ra.set("style", value)
}

// Target sets target
func (ra reloadAction) Target(value string) reloadAction {
	return ra.set("target", value)
}

// TestIdBuilder sets test id builder
func (ra reloadAction) TestIdBuilder(value string) reloadAction {
	return ra.set("testIdBuilder", value)
}

// Testid sets test id
func (ra reloadAction) Testid(value string) reloadAction {
	return ra.set("testid", value)
}

// Tooltip sets tooltip
func (ra reloadAction) Tooltip(value string) reloadAction {
	return ra.set("tooltip", value)
}

// TooltipPlacement sets tooltip placement
func (ra reloadAction) TooltipPlacement(value string) reloadAction {
	return ra.set("tooltipPlacement", value)
}

// Type sets type
func (ra reloadAction) Type(value string) reloadAction {
	return ra.set("type", value)
}

// UseMobileUI sets use mobile ui
func (ra reloadAction) UseMobileUI(value bool) reloadAction {
	return ra.set("useMobileUI", value)
}

// Visible sets visible
func (ra reloadAction) Visible(value bool) reloadAction {
	return ra.set("visible", value)
}

// VisibleOn sets visible on
func (ra reloadAction) VisibleOn(value string) reloadAction {
	return ra.set("visibleOn", value)
}
