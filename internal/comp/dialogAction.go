package comp

import "github.com/zrcoder/amisgo/model"

// DialogAction represents a dialog action button configuration.
type DialogAction model.Schema

func NewDialogAction() DialogAction {
	return DialogAction{"type": "button", "actionType": "dialog"}
}

func (da DialogAction) set(key string, value any) DialogAction {
	da[key] = value
	return da
}

// ActiveClassName sets the active class name.
func (da DialogAction) ActiveClassName(value string) DialogAction {
	return da.set("activeClassName", value)
}

// ActiveLevel sets the active level.
func (da DialogAction) ActiveLevel(value string) DialogAction {
	return da.set("activeLevel", value)
}

// Badge sets the badge.
func (da DialogAction) Badge(value string) DialogAction {
	return da.set("badge", value)
}

// Block sets whether the button is displayed as a block.
func (da DialogAction) Block(value bool) DialogAction {
	return da.set("block", value)
}

// Body sets the button body.
func (da DialogAction) Body(value ...any) DialogAction {
	return da.set("body", value)
}

// ClassName sets the button class name.
func (da DialogAction) ClassName(value string) DialogAction {
	return da.set("className", value)
}

// Close sets whether the dialog is closed after the action is completed.
func (da DialogAction) Close(value string) DialogAction {
	return da.set("close", value)
}

// ConfirmText sets the confirmation text.
func (da DialogAction) ConfirmText(value string) DialogAction {
	return da.set("confirmText", value)
}

// ConfirmTitle sets the confirmation title.
func (da DialogAction) ConfirmTitle(value any) DialogAction {
	return da.set("confirmTitle", value)
}

// CountDown sets the count down.
func (da DialogAction) CountDown(value string) DialogAction {
	return da.set("countDown", value)
}

// CountDownTpl sets the count down template.
func (da DialogAction) CountDownTpl(value string) DialogAction {
	return da.set("countDownTpl", value)
}

// Data sets the data mapping.
func (da DialogAction) Data(value any) DialogAction {
	return da.set("data", value)
}

// Dialog sets the dialog configuration.
func (da DialogAction) Dialog(value string) DialogAction {
	return da.set("dialog", value)
}

// Disabled sets whether the button is disabled.
func (da DialogAction) Disabled(value bool) DialogAction {
	return da.set("disabled", value)
}

// DisabledOn sets the disabled expression.
func (da DialogAction) DisabledOn(value string) DialogAction {
	return da.set("disabledOn", value)
}

// DisabledTip sets the disabled tip.
func (da DialogAction) DisabledTip(value string) DialogAction {
	return da.set("disabledTip", value)
}

// EditorSetting sets the editor setting.
func (da DialogAction) EditorSetting(value string) DialogAction {
	return da.set("editorSetting", value)
}

// Hidden sets whether the button is hidden.
func (da DialogAction) Hidden(value bool) DialogAction {
	return da.set("hidden", value)
}

// HiddenOn sets the hidden expression.
func (da DialogAction) HiddenOn(value string) DialogAction {
	return da.set("hiddenOn", value)
}

// HotKey sets the hot key.
func (da DialogAction) HotKey(value string) DialogAction {
	return da.set("hotKey", value)
}

// Icon sets the button icon.
func (da DialogAction) Icon(value string) DialogAction {
	return da.set("icon", value)
}

// IconClassName sets the icon class name.
func (da DialogAction) IconClassName(value string) DialogAction {
	return da.set("iconClassName", value)
}

// ID sets the button ID.
func (da DialogAction) ID(value string) DialogAction {
	return da.set("id", value)
}

// Label sets the button label.
func (da DialogAction) Label(value string) DialogAction {
	return da.set("label", value)
}

// Level sets the button level.
func (da DialogAction) Level(value string) DialogAction {
	return da.set("level", value)
}

// LoadingClassName sets the loading class name.
func (da DialogAction) LoadingClassName(value string) DialogAction {
	return da.set("loadingClassName", value)
}

// LoadingOn sets the loading expression.
func (da DialogAction) LoadingOn(value string) DialogAction {
	return da.set("loadingOn", value)
}

// MergeData sets whether the data is merged into the parent scope.
func (da DialogAction) MergeData(value bool) DialogAction {
	return da.set("mergeData", value)
}

// NextCondition sets the next condition expression.
func (da DialogAction) NextCondition(value string) DialogAction {
	return da.set("nextCondition", value)
}

// OnClick sets the click event handler.
func (da DialogAction) OnClick(value string) DialogAction {
	return da.set("onClick", value)
}

// OnEvent sets the event action configuration.
func (da DialogAction) OnEvent(value any) DialogAction {
	return da.set("onEvent", value)
}

// Primary sets whether the button is primary.
func (da DialogAction) Primary(value bool) DialogAction {
	return da.set("primary", value)
}

// Redirect sets the redirect URL.
func (da DialogAction) Redirect(value string) DialogAction {
	return da.set("redirect", value)
}

// Reload sets the reload action.
func (da DialogAction) Reload(value string) DialogAction {
	return da.set("reload", value)
}

// RequireSelected sets whether the button requires selection.
func (da DialogAction) RequireSelected(value bool) DialogAction {
	return da.set("requireSelected", value)
}

// Required sets the required fields.
func (da DialogAction) Required(value string) DialogAction {
	return da.set("required", value)
}

// RightIcon sets the right icon.
func (da DialogAction) RightIcon(value string) DialogAction {
	return da.set("rightIcon", value)
}

// RightIconClassName sets the right icon class name.
func (da DialogAction) RightIconClassName(value string) DialogAction {
	return da.set("rightIconClassName", value)
}

// Size sets the button size.
func (da DialogAction) Size(value string) DialogAction {
	return da.set("size", value)
}

// Static sets whether the button is static.
func (da DialogAction) Static(value bool) DialogAction {
	return da.set("static", value)
}

// StaticClassName sets the static class name.
func (da DialogAction) StaticClassName(value string) DialogAction {
	return da.set("staticClassName", value)
}

// StaticInputClassName sets the static input class name.
func (da DialogAction) StaticInputClassName(value string) DialogAction {
	return da.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static label class name.
func (da DialogAction) StaticLabelClassName(value string) DialogAction {
	return da.set("staticLabelClassName", value)
}

// StaticOn sets the static expression.
func (da DialogAction) StaticOn(value string) DialogAction {
	return da.set("staticOn", value)
}

// StaticPlaceholder sets the static placeholder.
func (da DialogAction) StaticPlaceholder(value string) DialogAction {
	return da.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema.
func (da DialogAction) StaticSchema(value string) DialogAction {
	return da.set("staticSchema", value)
}

// Style sets the button style.
func (da DialogAction) Style(value any) DialogAction {
	return da.set("style", value)
}

// Target sets the target.
func (da DialogAction) Target(value string) DialogAction {
	return da.set("target", value)
}

// TestIdBuilder sets the test ID builder.
func (da DialogAction) TestIdBuilder(value string) DialogAction {
	return da.set("testIdBuilder", value)
}

// Testid sets the test ID.
func (da DialogAction) Testid(value string) DialogAction {
	return da.set("testid", value)
}

// Tooltip sets the tooltip.
func (da DialogAction) Tooltip(value string) DialogAction {
	return da.set("tooltip", value)
}

// TooltipPlacement sets the tooltip placement.
func (da DialogAction) TooltipPlacement(value string) DialogAction {
	return da.set("tooltipPlacement", value)
}

// Type sets the button type.
func (da DialogAction) Type(value string) DialogAction {
	return da.set("type", value)
}

// UseMobileUI sets whether the button uses mobile UI.
func (da DialogAction) UseMobileUI(value bool) DialogAction {
	return da.set("useMobileUI", value)
}

// Visible sets whether the button is visible.
func (da DialogAction) Visible(value bool) DialogAction {
	return da.set("visible", value)
}

// VisibleOn sets the visible expression.
func (da DialogAction) VisibleOn(value string) DialogAction {
	return da.set("visibleOn", value)
}
