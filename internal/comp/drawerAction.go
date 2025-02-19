package comp

import "github.com/zrcoder/amisgo/schema"

// DrawerAction represents a button in a drawer.
type DrawerAction schema.Schema

// NewDrawerAction creates a new NewDrawerAction instance and sets the default type and actionType.
func NewDrawerAction() DrawerAction {
	return DrawerAction{"type": "button", "actionType": "drawer"}
}

func (d DrawerAction) set(key string, value any) DrawerAction {
	d[key] = value
	return d
}

// ActiveClassName specifies the class name when the button is active.
func (d DrawerAction) ActiveClassName(value string) DrawerAction {
	return d.set("activeClassName", value)
}

// ActiveLevel specifies the style when the button is active.
func (d DrawerAction) ActiveLevel(value string) DrawerAction {
	return d.set("activeLevel", value)
}

// Align specifies the alignment of the button.
func (d DrawerAction) Align(value string) DrawerAction {
	return d.set("align", value)
}

// Badge specifies the badge of the button.
func (d DrawerAction) Badge(value string) DrawerAction {
	return d.set("badge", value)
}

// Block specifies whether the button is a block element.
func (d DrawerAction) Block(value bool) DrawerAction {
	return d.set("block", value)
}

// Body specifies the body of the button.
func (d DrawerAction) Body(value ...any) DrawerAction {
	return d.set("body", value)
}

// ClassName specifies the class name of the button.
func (d DrawerAction) ClassName(value string) DrawerAction {
	return d.set("className", value)
}

// Close specifies whether the drawer should be closed after the button is clicked.
func (d DrawerAction) Close(value string) DrawerAction {
	return d.set("close", value)
}

// ConfirmText specifies the confirm text of the button.
func (d DrawerAction) ConfirmText(value string) DrawerAction {
	return d.set("confirmText", value)
}

// ConfirmTitle specifies the confirm title of the button.
func (d DrawerAction) ConfirmTitle(value any) DrawerAction {
	return d.set("confirmTitle", value)
}

// CountDown specifies the count down time of the button.
func (d DrawerAction) CountDown(value string) DrawerAction {
	return d.set("countDown", value)
}

// CountDownTpl specifies the count down template of the button.
func (d DrawerAction) CountDownTpl(value string) DrawerAction {
	return d.set("countDownTpl", value)
}

// Data specifies the data mapping of the button.
func (d DrawerAction) Data(value any) DrawerAction {
	return d.set("data", value)
}

// Disabled specifies whether the button is disabled.
func (d DrawerAction) Disabled(value bool) DrawerAction {
	return d.set("disabled", value)
}

// DisabledOn specifies the disabled expression of the button.
func (d DrawerAction) DisabledOn(value string) DrawerAction {
	return d.set("disabledOn", value)
}

// DisabledTip specifies the disabled tip of the button.
func (d DrawerAction) DisabledTip(value string) DrawerAction {
	return d.set("disabledTip", value)
}

// Drawer specifies the drawer details of the button.
func (d DrawerAction) Drawer(value any) DrawerAction {
	return d.set("drawer", value)
}

// EditorSetting specifies the editor setting of the button, which can be ignored at runtime.
func (d DrawerAction) EditorSetting(value string) DrawerAction {
	return d.set("editorSetting", value)
}

// Hidden specifies whether the button is hidden.
func (d DrawerAction) Hidden(value bool) DrawerAction {
	return d.set("hidden", value)
}

// HiddenOn specifies the hidden expression of the button.
func (d DrawerAction) HiddenOn(value string) DrawerAction {
	return d.set("hiddenOn", value)
}

// HotKey specifies the hot key of the button.
func (d DrawerAction) HotKey(value string) DrawerAction {
	return d.set("hotKey", value)
}

// Icon specifies the icon of the button.
func (d DrawerAction) Icon(value string) DrawerAction {
	return d.set("icon", value)
}

// IconClassName specifies the icon class name of the button.
func (d DrawerAction) IconClassName(value string) DrawerAction {
	return d.set("iconClassName", value)
}

// ID specifies the ID of the button.
func (d DrawerAction) ID(value string) DrawerAction {
	return d.set("id", value)
}

// Label specifies the label of the button.
func (d DrawerAction) Label(value string) DrawerAction {
	return d.set("label", value)
}

// Level specifies the level of the button.
func (d DrawerAction) Level(value string) DrawerAction {
	return d.set("level", value)
}

// LoadingClassName specifies the loading class name of the button.
func (d DrawerAction) LoadingClassName(value string) DrawerAction {
	return d.set("loadingClassName", value)
}

// LoadingOn specifies whether the loading effect is enabled.
func (d DrawerAction) LoadingOn(value string) DrawerAction {
	return d.set("loadingOn", value)
}

// MergeData specifies whether the drawer should merge the data to the parent scope.
func (d DrawerAction) MergeData(value bool) DrawerAction {
	return d.set("mergeData", value)
}

// NextCondition specifies whether the button has a next condition.
func (d DrawerAction) NextCondition(value string) DrawerAction {
	return d.set("nextCondition", value)
}

// OnClick specifies the custom event handler of the button.
func (d DrawerAction) OnClick(value string) DrawerAction {
	return d.set("onClick", value)
}

// OnEvent specifies the event action configuration of the button.
func (d DrawerAction) OnEvent(value any) DrawerAction {
	return d.set("onEvent", value)
}

// Primary specifies whether the button is the primary button.
func (d DrawerAction) Primary(value bool) DrawerAction {
	return d.set("primary", value)
}

// Redirect specifies the redirect configuration of the button.
func (d DrawerAction) Redirect(value string) DrawerAction {
	return d.set("redirect", value)
}

// Reload specifies the reload configuration of the button.
func (d DrawerAction) Reload(value string) DrawerAction {
	return d.set("reload", value)
}

// RequireSelected specifies whether the button requires selected items.
func (d DrawerAction) RequireSelected(value bool) DrawerAction {
	return d.set("requireSelected", value)
}

// Required specifies whether the button requires the specified fields to be validated.
func (d DrawerAction) Required(value string) DrawerAction {
	return d.set("required", value)
}

// RightIcon specifies the right icon of the button.
func (d DrawerAction) RightIcon(value string) DrawerAction {
	return d.set("rightIcon", value)
}

// RightIconClassName specifies the right icon class name of the button.
func (d DrawerAction) RightIconClassName(value string) DrawerAction {
	return d.set("rightIconClassName", value)
}

// Size specifies the size of the button.
func (d DrawerAction) Size(value string) DrawerAction {
	return d.set("size", value)
}

// Static specifies whether the button is static.
func (d DrawerAction) Static(value bool) DrawerAction {
	return d.set("static", value)
}

// StaticClassName specifies the static class name of the button.
func (d DrawerAction) StaticClassName(value string) DrawerAction {
	return d.set("staticClassName", value)
}

// StaticInputClassName specifies the static input class name of the button.
func (d DrawerAction) StaticInputClassName(value string) DrawerAction {
	return d.set("staticInputClassName", value)
}

// StaticLabelClassName specifies the static label class name of the button.
func (d DrawerAction) StaticLabelClassName(value string) DrawerAction {
	return d.set("staticLabelClassName", value)
}

// StaticOn specifies whether the button is static based on the expression.
func (d DrawerAction) StaticOn(value string) DrawerAction {
	return d.set("staticOn", value)
}

// StaticPlaceholder specifies the static placeholder of the button.
func (d DrawerAction) StaticPlaceholder(value string) DrawerAction {
	return d.set("staticPlaceholder", value)
}

// StaticSchema specifies the static schema.Schema of the button.
func (d DrawerAction) StaticSchema(value string) DrawerAction {
	return d.set("staticSchema", value)
}

// Style specifies the style of the button.
func (d DrawerAction) Style(value any) DrawerAction {
	return d.set("style", value)
}

// Target specifies who should trigger the action.
func (d DrawerAction) Target(value string) DrawerAction {
	return d.set("target", value)
}

// TestIdBuilder specifies the test ID builder of the button.
func (d DrawerAction) TestIdBuilder(value string) DrawerAction {
	return d.set("testIdBuilder", value)
}

// Testid specifies the test ID of the button.
func (d DrawerAction) Testid(value string) DrawerAction {
	return d.set("testid", value)
}

// Tooltip specifies the tooltip of the button.
func (d DrawerAction) Tooltip(value string) DrawerAction {
	return d.set("tooltip", value)
}

// TooltipPlacement specifies the tooltip placement of the button.
func (d DrawerAction) TooltipPlacement(value string) DrawerAction {
	return d.set("tooltipPlacement", value)
}

// Type specifies the type of the button.
func (d DrawerAction) Type(value string) DrawerAction {
	return d.set("type", value)
}

// UseMobileUI specifies whether the button should use the mobile UI.
func (d DrawerAction) UseMobileUI(value bool) DrawerAction {
	return d.set("useMobileUI", value)
}

// Visible specifies whether the button is visible.
func (d DrawerAction) Visible(value bool) DrawerAction {
	return d.set("visible", value)
}

// VisibleOn specifies whether the button is visible based on the expression.
func (d DrawerAction) VisibleOn(value string) DrawerAction {
	return d.set("visibleOn", value)
}
