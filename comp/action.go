package comp

import (
	"github.com/zrcoder/amisgo/internal/servermux"
	"github.com/zrcoder/amisgo/model"
)

// action represents an action button. Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/action
type action model.Schema

// Action creates a new Action instance.
func Action() action {
	return action{}.set("type", "action")
}

// Button is an alias for Action with type "button".
func Button() action {
	return Action().set("type", "button")
}

// Submit creates an Action with type "submit".
func Submit() action {
	return Action().set("type", "submit")
}

// ActionType sets the core configuration for the action type.
// Possible values: ajax, link, url, drawer, dialog, confirm, cancel, prev, next, copy, close, button, reset, submit, clear.
func (a action) ActionType(value string) action {
	return a.set("actionType", value)
}

// Dialog configures the dialog that appears when the button is clicked.
func (a action) Dialog(value dialog) action {
	return a.set("dialog", value)
}

// Drawer configures the drawer that appears when the button is clicked.
func (a action) Drawer(value any) action {
	return a.set("drawer", value)
}

// Toast configures the toast that appears when the button is clicked.
func (a action) Toast(value toast) action {
	return a.set("toast", value)
}

// Transform transforms the src value with the provided function and renders the result to the dst component.
func (a action) Transform(src, dst, successMsg string, transfor func(input any) (any, error)) action {
	return a.TransformMultiple(successMsg, func(d model.Schema) (model.Schema, error) {
		output, err := transfor(d[src])
		if err != nil {
			return nil, err
		}
		return model.Schema{dst: output}, nil
	}, src)
}

// TransformMultiple transforms the src values with the provided function and renders the result to multiple destinations.
func (a action) TransformMultiple(successMsg string, transfor func(model.Schema) (model.Schema, error), src ...string) action {
	route, data := servermux.TransformMultiple(successMsg, transfor, src...)
	return a.ActionType("ajax").Api(
		model.Schema{
			"url":  route,
			"data": data,
			"__amisgo_resp": model.Schema{
				"200": model.Schema{
					"then": EventAction().ActionType("setValue").Args(model.Schema{"value": "${__amisgo__resp}"}),
				},
			},
		},
	)
}

// ActiveClassName sets the class name for the active state of the button.
func (a action) ActiveClassName(value string) action {
	return a.set("activeClassName", value)
}

// ActiveLevel sets the style for the active state of the button.
func (a action) ActiveLevel(value string) action {
	return a.set("activeLevel", value)
}

// Api sets the API address/description.
func (a action) Api(value any) action {
	return a.set("api", value)
}

// ClassName adds a class name to the button.
func (a action) ClassName(value string) action {
	return a.set("className", value)
}

// Close specifies whether to close the current dialog or drawer after the action is performed.
func (a action) Close(value string) action {
	return a.set("close", value)
}

// ConfirmText sets the text to be displayed in the confirmation dialog before the action is performed.
func (a action) ConfirmText(value string) action {
	return a.set("confirmText", value)
}

// ConfirmTitle sets the title of the confirmation dialog.
func (a action) ConfirmTitle(value any) action {
	return a.set("confirmTitle", value)
}

// DisabledTip sets the text to be displayed when the button is disabled.
func (a action) DisabledTip(value any) action {
	return a.set("disabledTip", value)
}

// Icon sets the icon for the button, e.g., "fa fa-plus".
func (a action) Icon(value string) action {
	return a.set("icon", value)
}

// IconClassName adds a class name to the icon.
func (a action) IconClassName(value string) action {
	return a.set("iconClassName", value)
}

// Redirect sets the relative path for single-page navigation.
func (a action) Redirect(value string) action {
	return a.set("redirect", value)
}

// Label sets the button text.
func (a action) Label(value string) action {
	return a.set("label", value)
}

// Level sets the button style.
// Possible values: link, primary, enhance, secondary, info, success, warning, danger, light, dark, default.
func (a action) Level(value string) action {
	return a.set("level", value)
}

// Link sets the link for the button.
func (a action) Link(value string) action {
	return a.set("link", value)
}

// Reload specifies the target component(s) to be refreshed after the action is performed.
func (a action) Reload(value string) action {
	return a.set("reload", value)
}

// ReloadWindow refreshes the current page.
func (a action) ReloadWindow() action {
	return a.Reload("window")
}

// Required sets the required fields for the form before the action is performed.
func (a action) Required(value ...string) action {
	return a.set("required", value)
}

// RightIcon sets the icon on the right side of the button text.
func (a action) RightIcon(value string) action {
	return a.set("rightIcon", value)
}

// RightIconClassName adds a class name to the right icon.
func (a action) RightIconClassName(value string) action {
	return a.set("rightIconClassName", value)
}

// Size sets the button size. Possible values: xs, sm, md, lg.
func (a action) Size(value string) action {
	return a.set("size", value)
}

// Tooltip sets the text to be displayed when the mouse hovers over the button.
func (a action) Tooltip(value string) action {
	return a.set("tooltip", value)
}

// TooltipPlacement sets the position of the tooltip. Possible values: top, bottom, left, right.
func (a action) TooltipPlacement(value string) action {
	return a.set("tooltipPlacement", value)
}

// TooltipTrigger sets the trigger for the tooltip. Possible values: hover, focus.
func (a action) TooltipTrigger(value string) action {
	return a.set("tooltipTrigger", value)
}

// Badge sets the badge for the button. Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/badge
func (v action) Badge(value string) action {
	return v.set("badge", value)
}

// Block sets whether the button is displayed as a block element.
func (v action) Block(value bool) action {
	return v.set("block", value)
}

// Body sets the child content of the button.
func (v action) Body(value ...any) action {
	return v.set("body", value)
}

// CountDown sets the countdown time (in seconds) after the button is clicked.
func (v action) CountDown(value string) action {
	return v.set("countDown", value)
}

// CountDownTpl sets the custom countdown text.
func (v action) CountDownTpl(value string) action {
	return v.set("countDownTpl", value)
}

// Disabled sets whether the button is disabled.
func (v action) Disabled(value bool) action {
	return v.set("disabled", value)
}

// DisabledOn sets the expression to determine whether the button is disabled.
func (v action) DisabledOn(value string) action {
	return v.set("disabledOn", value)
}

// DownloadFileName sets the file name for the download.
func (v action) DownloadFileName(value string) action {
	return v.set("downloadFileName", value)
}

// EditorSetting sets the editor configuration, which can be ignored at runtime.
func (v action) EditorSetting(value string) action {
	return v.set("editorSetting", value)
}

// Hidden sets whether the button is hidden.
func (v action) Hidden(value bool) action {
	return v.set("hidden", value)
}

// HiddenOn sets the expression to determine whether the button is hidden.
func (v action) HiddenOn(value string) action {
	return v.set("hiddenOn", value)
}

// HotKey sets the keyboard shortcut for the button.
func (v action) HotKey(value string) action {
	return v.set("hotKey", value)
}

// ID sets the ID for the button, mainly used for user behavior tracking.
func (v action) ID(value string) action {
	return v.set("id", value)
}

// Url sets the URL for the button.
func (v action) Url(value string) action {
	return v.set("url", value)
}

// Loading sets whether to show the loading effect on the button.
func (a action) Loading(value bool) action {
	return a.set("loading", value)
}

// LoadingClassName sets the CSS class name for the loading effect.
func (v action) LoadingClassName(value string) action {
	return v.set("loadingClassName", value)
}

// LoadingOn sets the expression to determine whether to show the loading effect.
func (v action) LoadingOn(value string) action {
	return v.set("loadingOn", value)
}

// MergeData sets whether to merge the data from the dialog or drawer into the parent scope.
func (v action) MergeData(value bool) action {
	return v.set("mergeData", value)
}

// OnClick sets the custom event handler function for the button.
func (v action) OnClick(value string) action {
	return v.set("onClick", value)
}

// OnEvent sets the event action configuration for the button.
func (v action) OnEvent(value any) action {
	return v.set("onEvent", value)
}

// Primary sets whether the button is a primary button.
func (v action) Primary(value bool) action {
	return v.set("primary", value)
}

// RequireSelected sets whether the button requires selected elements to be clickable when used as a batch operation button.
func (v action) RequireSelected(value bool) action {
	return v.set("requireSelected", value)
}

// Static sets whether the button is displayed statically.
func (v action) Static(value bool) action {
	return v.set("static", value)
}

// StaticClassName sets the CSS class name for the static display of the form item.
func (v action) StaticClassName(value string) action {
	return v.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for the static display of the form item value.
func (v action) StaticInputClassName(value string) action {
	return v.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for the static display of the form item label.
func (v action) StaticLabelClassName(value string) action {
	return v.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine whether the button is displayed statically.
func (v action) StaticOn(value string) action {
	return v.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder text for the static display.
func (v action) StaticPlaceholder(value string) action {
	return v.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for the static display mode.
func (v action) StaticSchema(value string) action {
	return v.set("staticSchema", value)
}

// Style sets the style for the button.
func (v action) Style(value any) action {
	return v.set("style", value)
}

// Target sets the target component(s) that can trigger this action.
func (v action) Target(value string) action {
	return v.set("target", value)
}

// TestIdBuilder sets the function to generate test IDs.
func (v action) TestIdBuilder(value string) action {
	return v.set("testIdBuilder", value)
}

// Testid sets the test ID for the button.
func (v action) Testid(value string) action {
	return v.set("testid", value)
}

// UseMobileUI sets whether to disable mobile UI styles.
func (v action) UseMobileUI(value bool) action {
	return v.set("useMobileUI", value)
}

// Visible sets whether the button is visible.
func (v action) Visible(value bool) action {
	return v.set("visible", value)
}

// VisibleOn sets the expression to determine whether the button is visible.
func (v action) VisibleOn(value string) action {
	return v.set("visibleOn", value)
}

func (a action) set(key string, value any) action {
	a[key] = value
	return a
}
