package comp

import "github.com/zrcoder/amisgo/model"

// drawerAction represents a button in a drawer.
type drawerAction model.Schema

// DrawerAction creates a new DrawerAction instance and sets the default type and actionType.
func DrawerAction() drawerAction {
	return make(drawerAction).set("type", "button").set("actionType", "drawer")
}

func (d drawerAction) set(key string, value any) drawerAction {
	d[key] = value
	return d
}

// ActiveClassName specifies the class name when the button is active.
func (d drawerAction) ActiveClassName(value string) drawerAction {
	return d.set("activeClassName", value)
}

// ActiveLevel specifies the style when the button is active.
func (d drawerAction) ActiveLevel(value string) drawerAction {
	return d.set("activeLevel", value)
}

// Align specifies the alignment of the button.
func (d drawerAction) Align(value string) drawerAction {
	return d.set("align", value)
}

// Badge specifies the badge of the button.
func (d drawerAction) Badge(value string) drawerAction {
	return d.set("badge", value)
}

// Block specifies whether the button is a block element.
func (d drawerAction) Block(value bool) drawerAction {
	return d.set("block", value)
}

// Body specifies the body of the button.
func (d drawerAction) Body(value ...any) drawerAction {
	return d.set("body", value)
}

// ClassName specifies the class name of the button.
func (d drawerAction) ClassName(value string) drawerAction {
	return d.set("className", value)
}

// Close specifies whether the drawer should be closed after the button is clicked.
func (d drawerAction) Close(value string) drawerAction {
	return d.set("close", value)
}

// ConfirmText specifies the confirm text of the button.
func (d drawerAction) ConfirmText(value string) drawerAction {
	return d.set("confirmText", value)
}

// ConfirmTitle specifies the confirm title of the button.
func (d drawerAction) ConfirmTitle(value any) drawerAction {
	return d.set("confirmTitle", value)
}

// CountDown specifies the count down time of the button.
func (d drawerAction) CountDown(value string) drawerAction {
	return d.set("countDown", value)
}

// CountDownTpl specifies the count down template of the button.
func (d drawerAction) CountDownTpl(value string) drawerAction {
	return d.set("countDownTpl", value)
}

// Data specifies the data mapping of the button.
func (d drawerAction) Data(value model.Data) drawerAction {
	return d.set("data", value)
}

// Disabled specifies whether the button is disabled.
func (d drawerAction) Disabled(value bool) drawerAction {
	return d.set("disabled", value)
}

// DisabledOn specifies the disabled expression of the button.
func (d drawerAction) DisabledOn(value string) drawerAction {
	return d.set("disabledOn", value)
}

// DisabledTip specifies the disabled tip of the button.
func (d drawerAction) DisabledTip(value string) drawerAction {
	return d.set("disabledTip", value)
}

// Drawer specifies the drawer details of the button.
func (d drawerAction) Drawer(value string) drawerAction {
	return d.set("drawer", value)
}

// EditorSetting specifies the editor setting of the button, which can be ignored at runtime.
func (d drawerAction) EditorSetting(value string) drawerAction {
	return d.set("editorSetting", value)
}

// Hidden specifies whether the button is hidden.
func (d drawerAction) Hidden(value bool) drawerAction {
	return d.set("hidden", value)
}

// HiddenOn specifies the hidden expression of the button.
func (d drawerAction) HiddenOn(value string) drawerAction {
	return d.set("hiddenOn", value)
}

// HotKey specifies the hot key of the button.
func (d drawerAction) HotKey(value string) drawerAction {
	return d.set("hotKey", value)
}

// Icon specifies the icon of the button.
func (d drawerAction) Icon(value string) drawerAction {
	return d.set("icon", value)
}

// IconClassName specifies the icon class name of the button.
func (d drawerAction) IconClassName(value string) drawerAction {
	return d.set("iconClassName", value)
}

// ID specifies the ID of the button.
func (d drawerAction) ID(value string) drawerAction {
	return d.set("id", value)
}

// Label specifies the label of the button.
func (d drawerAction) Label(value string) drawerAction {
	return d.set("label", value)
}

// Level specifies the level of the button.
func (d drawerAction) Level(value string) drawerAction {
	return d.set("level", value)
}

// LoadingClassName specifies the loading class name of the button.
func (d drawerAction) LoadingClassName(value string) drawerAction {
	return d.set("loadingClassName", value)
}

// LoadingOn specifies whether the loading effect is enabled.
func (d drawerAction) LoadingOn(value string) drawerAction {
	return d.set("loadingOn", value)
}

// MergeData specifies whether the drawer should merge the data to the parent scope.
func (d drawerAction) MergeData(value bool) drawerAction {
	return d.set("mergeData", value)
}

// NextCondition specifies whether the button has a next condition.
func (d drawerAction) NextCondition(value string) drawerAction {
	return d.set("nextCondition", value)
}

// OnClick specifies the custom event handler of the button.
func (d drawerAction) OnClick(value string) drawerAction {
	return d.set("onClick", value)
}

// OnEvent specifies the event action configuration of the button.
func (d drawerAction) OnEvent(value any) drawerAction {
	return d.set("onEvent", value)
}

// Primary specifies whether the button is the primary button.
func (d drawerAction) Primary(value bool) drawerAction {
	return d.set("primary", value)
}

// Redirect specifies the redirect configuration of the button.
func (d drawerAction) Redirect(value string) drawerAction {
	return d.set("redirect", value)
}

// Reload specifies the reload configuration of the button.
func (d drawerAction) Reload(value string) drawerAction {
	return d.set("reload", value)
}

// RequireSelected specifies whether the button requires selected items.
func (d drawerAction) RequireSelected(value bool) drawerAction {
	return d.set("requireSelected", value)
}

// Required specifies whether the button requires the specified fields to be validated.
func (d drawerAction) Required(value string) drawerAction {
	return d.set("required", value)
}

// RightIcon specifies the right icon of the button.
func (d drawerAction) RightIcon(value string) drawerAction {
	return d.set("rightIcon", value)
}

// RightIconClassName specifies the right icon class name of the button.
func (d drawerAction) RightIconClassName(value string) drawerAction {
	return d.set("rightIconClassName", value)
}

// Size specifies the size of the button.
func (d drawerAction) Size(value string) drawerAction {
	return d.set("size", value)
}

// Static specifies whether the button is static.
func (d drawerAction) Static(value bool) drawerAction {
	return d.set("static", value)
}

// StaticClassName specifies the static class name of the button.
func (d drawerAction) StaticClassName(value string) drawerAction {
	return d.set("staticClassName", value)
}

// StaticInputClassName specifies the static input class name of the button.
func (d drawerAction) StaticInputClassName(value string) drawerAction {
	return d.set("staticInputClassName", value)
}

// StaticLabelClassName specifies the static label class name of the button.
func (d drawerAction) StaticLabelClassName(value string) drawerAction {
	return d.set("staticLabelClassName", value)
}

// StaticOn specifies whether the button is static based on the expression.
func (d drawerAction) StaticOn(value string) drawerAction {
	return d.set("staticOn", value)
}

// StaticPlaceholder specifies the static placeholder of the button.
func (d drawerAction) StaticPlaceholder(value string) drawerAction {
	return d.set("staticPlaceholder", value)
}

// StaticSchema specifies the static schema of the button.
func (d drawerAction) StaticSchema(value string) drawerAction {
	return d.set("staticSchema", value)
}

// Style specifies the style of the button.
func (d drawerAction) Style(value any) drawerAction {
	return d.set("style", value)
}

// Target specifies who should trigger the action.
func (d drawerAction) Target(value string) drawerAction {
	return d.set("target", value)
}

// TestIdBuilder specifies the test ID builder of the button.
func (d drawerAction) TestIdBuilder(value string) drawerAction {
	return d.set("testIdBuilder", value)
}

// Testid specifies the test ID of the button.
func (d drawerAction) Testid(value string) drawerAction {
	return d.set("testid", value)
}

// Tooltip specifies the tooltip of the button.
func (d drawerAction) Tooltip(value string) drawerAction {
	return d.set("tooltip", value)
}

// TooltipPlacement specifies the tooltip placement of the button.
func (d drawerAction) TooltipPlacement(value string) drawerAction {
	return d.set("tooltipPlacement", value)
}

// Type specifies the type of the button.
func (d drawerAction) Type(value string) drawerAction {
	return d.set("type", value)
}

// UseMobileUI specifies whether the button should use the mobile UI.
func (d drawerAction) UseMobileUI(value bool) drawerAction {
	return d.set("useMobileUI", value)
}

// Visible specifies whether the button is visible.
func (d drawerAction) Visible(value bool) drawerAction {
	return d.set("visible", value)
}

// VisibleOn specifies whether the button is visible based on the expression.
func (d drawerAction) VisibleOn(value string) drawerAction {
	return d.set("visibleOn", value)
}
