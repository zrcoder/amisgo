package comp

import "github.com/zrcoder/amisgo/model"

// dialogAction represents a dialog action button configuration.
type dialogAction model.Schema

// DialogAction creates a new dialog action button instance and sets default settings.
func DialogAction() dialogAction {
	return make(dialogAction).set("type", "button").set("actionType", "dialog")
}

func (da dialogAction) set(key string, value any) dialogAction {
	da[key] = value
	return da
}

// ActiveClassName sets the active class name.
func (da dialogAction) ActiveClassName(value string) dialogAction {
	return da.set("activeClassName", value)
}

// ActiveLevel sets the active level.
func (da dialogAction) ActiveLevel(value string) dialogAction {
	return da.set("activeLevel", value)
}

// Badge sets the badge.
func (da dialogAction) Badge(value string) dialogAction {
	return da.set("badge", value)
}

// Block sets whether the button is displayed as a block.
func (da dialogAction) Block(value bool) dialogAction {
	return da.set("block", value)
}

// Body sets the button body.
func (da dialogAction) Body(value ...any) dialogAction {
	return da.set("body", value)
}

// ClassName sets the button class name.
func (da dialogAction) ClassName(value string) dialogAction {
	return da.set("className", value)
}

// Close sets whether the dialog is closed after the action is completed.
func (da dialogAction) Close(value string) dialogAction {
	return da.set("close", value)
}

// ConfirmText sets the confirmation text.
func (da dialogAction) ConfirmText(value string) dialogAction {
	return da.set("confirmText", value)
}

// ConfirmTitle sets the confirmation title.
func (da dialogAction) ConfirmTitle(value any) dialogAction {
	return da.set("confirmTitle", value)
}

// CountDown sets the count down.
func (da dialogAction) CountDown(value string) dialogAction {
	return da.set("countDown", value)
}

// CountDownTpl sets the count down template.
func (da dialogAction) CountDownTpl(value string) dialogAction {
	return da.set("countDownTpl", value)
}

// Data sets the data mapping.
func (da dialogAction) Data(value any) dialogAction {
	return da.set("data", value)
}

// Dialog sets the dialog configuration.
func (da dialogAction) Dialog(value string) dialogAction {
	return da.set("dialog", value)
}

// Disabled sets whether the button is disabled.
func (da dialogAction) Disabled(value bool) dialogAction {
	return da.set("disabled", value)
}

// DisabledOn sets the disabled expression.
func (da dialogAction) DisabledOn(value string) dialogAction {
	return da.set("disabledOn", value)
}

// DisabledTip sets the disabled tip.
func (da dialogAction) DisabledTip(value string) dialogAction {
	return da.set("disabledTip", value)
}

// EditorSetting sets the editor setting.
func (da dialogAction) EditorSetting(value string) dialogAction {
	return da.set("editorSetting", value)
}

// Hidden sets whether the button is hidden.
func (da dialogAction) Hidden(value bool) dialogAction {
	return da.set("hidden", value)
}

// HiddenOn sets the hidden expression.
func (da dialogAction) HiddenOn(value string) dialogAction {
	return da.set("hiddenOn", value)
}

// HotKey sets the hot key.
func (da dialogAction) HotKey(value string) dialogAction {
	return da.set("hotKey", value)
}

// Icon sets the button icon.
func (da dialogAction) Icon(value string) dialogAction {
	return da.set("icon", value)
}

// IconClassName sets the icon class name.
func (da dialogAction) IconClassName(value string) dialogAction {
	return da.set("iconClassName", value)
}

// ID sets the button ID.
func (da dialogAction) ID(value string) dialogAction {
	return da.set("id", value)
}

// Label sets the button label.
func (da dialogAction) Label(value string) dialogAction {
	return da.set("label", value)
}

// Level sets the button level.
func (da dialogAction) Level(value string) dialogAction {
	return da.set("level", value)
}

// LoadingClassName sets the loading class name.
func (da dialogAction) LoadingClassName(value string) dialogAction {
	return da.set("loadingClassName", value)
}

// LoadingOn sets the loading expression.
func (da dialogAction) LoadingOn(value string) dialogAction {
	return da.set("loadingOn", value)
}

// MergeData sets whether the data is merged into the parent scope.
func (da dialogAction) MergeData(value bool) dialogAction {
	return da.set("mergeData", value)
}

// NextCondition sets the next condition expression.
func (da dialogAction) NextCondition(value string) dialogAction {
	return da.set("nextCondition", value)
}

// OnClick sets the click event handler.
func (da dialogAction) OnClick(value string) dialogAction {
	return da.set("onClick", value)
}

// OnEvent sets the event action configuration.
func (da dialogAction) OnEvent(value any) dialogAction {
	return da.set("onEvent", value)
}

// Primary sets whether the button is primary.
func (da dialogAction) Primary(value bool) dialogAction {
	return da.set("primary", value)
}

// Redirect sets the redirect URL.
func (da dialogAction) Redirect(value string) dialogAction {
	return da.set("redirect", value)
}

// Reload sets the reload action.
func (da dialogAction) Reload(value string) dialogAction {
	return da.set("reload", value)
}

// RequireSelected sets whether the button requires selection.
func (da dialogAction) RequireSelected(value bool) dialogAction {
	return da.set("requireSelected", value)
}

// Required sets the required fields.
func (da dialogAction) Required(value string) dialogAction {
	return da.set("required", value)
}

// RightIcon sets the right icon.
func (da dialogAction) RightIcon(value string) dialogAction {
	return da.set("rightIcon", value)
}

// RightIconClassName sets the right icon class name.
func (da dialogAction) RightIconClassName(value string) dialogAction {
	return da.set("rightIconClassName", value)
}

// Size sets the button size.
func (da dialogAction) Size(value string) dialogAction {
	return da.set("size", value)
}

// Static sets whether the button is static.
func (da dialogAction) Static(value bool) dialogAction {
	return da.set("static", value)
}

// StaticClassName sets the static class name.
func (da dialogAction) StaticClassName(value string) dialogAction {
	return da.set("staticClassName", value)
}

// StaticInputClassName sets the static input class name.
func (da dialogAction) StaticInputClassName(value string) dialogAction {
	return da.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static label class name.
func (da dialogAction) StaticLabelClassName(value string) dialogAction {
	return da.set("staticLabelClassName", value)
}

// StaticOn sets the static expression.
func (da dialogAction) StaticOn(value string) dialogAction {
	return da.set("staticOn", value)
}

// StaticPlaceholder sets the static placeholder.
func (da dialogAction) StaticPlaceholder(value string) dialogAction {
	return da.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema.
func (da dialogAction) StaticSchema(value string) dialogAction {
	return da.set("staticSchema", value)
}

// Style sets the button style.
func (da dialogAction) Style(value any) dialogAction {
	return da.set("style", value)
}

// Target sets the target.
func (da dialogAction) Target(value string) dialogAction {
	return da.set("target", value)
}

// TestIdBuilder sets the test ID builder.
func (da dialogAction) TestIdBuilder(value string) dialogAction {
	return da.set("testIdBuilder", value)
}

// Testid sets the test ID.
func (da dialogAction) Testid(value string) dialogAction {
	return da.set("testid", value)
}

// Tooltip sets the tooltip.
func (da dialogAction) Tooltip(value string) dialogAction {
	return da.set("tooltip", value)
}

// TooltipPlacement sets the tooltip placement.
func (da dialogAction) TooltipPlacement(value string) dialogAction {
	return da.set("tooltipPlacement", value)
}

// Type sets the button type.
func (da dialogAction) Type(value string) dialogAction {
	return da.set("type", value)
}

// UseMobileUI sets whether the button uses mobile UI.
func (da dialogAction) UseMobileUI(value bool) dialogAction {
	return da.set("useMobileUI", value)
}

// Visible sets whether the button is visible.
func (da dialogAction) Visible(value bool) dialogAction {
	return da.set("visible", value)
}

// VisibleOn sets the visible expression.
func (da dialogAction) VisibleOn(value string) dialogAction {
	return da.set("visibleOn", value)
}
