package comp

import (
	"net/http"

	"github.com/zrcoder/amisgo/internal/servemux"
	"github.com/zrcoder/amisgo/model"
)

// Action represents an Action button. Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/Action
type Action model.Schema

func NewAction(mux *http.ServeMux) Action {
	return Action{"type": "action", servemux.Key: mux}
}

func NewButton(mux *http.ServeMux) Action {
	return NewAction(mux).set("type", "button")
}

func NewSubmit(mux *http.ServeMux) Action {
	return NewAction(mux).set("type", "submit")
}

// ActionType sets the core configuration for the action type.
// Possible values: ajax, link, url, drawer, dialog, confirm, cancel, prev, next, copy, close, button, reset, submit, clear.
func (a Action) ActionType(value string) Action {
	return a.set("actionType", value)
}

// Dialog configures the dialog that appears when the button is clicked.
func (a Action) Dialog(value Dialog) Action {
	return a.set("dialog", value)
}

// Drawer configures the drawer that appears when the button is clicked.
func (a Action) Drawer(value any) Action {
	return a.set("drawer", value)
}

// Toast configures the toast that appears when the button is clicked.
func (a Action) Toast(value model.Toast) Action {
	return a.set("toast", value)
}

// Transform transforms the src component value with the provided function and renders the result to the dst component.
func (a Action) Transform(transfor func(input any) (any, error), src, dst string) Action {
	return a.TransformMultiple(func(d model.Schema) (model.Schema, error) {
		output, err := transfor(d[src])
		if err != nil {
			return nil, err
		}
		return model.Schema{dst: output}, nil
	}, src)
}

// TransformMultiple transforms the src components' values with the provided function and renders the result to multiple destinations.
func (a Action) TransformMultiple(transfor func(model.Schema) (model.Schema, error), src ...string) Action {
	route, data := servemux.TransformMultiple(a.mux(), src, transfor)
	return a.ActionType("ajax").Api(
		model.Schema{
			"url":  route,
			"data": data,
			"__amisgo_resp": model.Schema{
				"200": model.Schema{
					"then": model.NewEventAction().ActionType("setValue").Args(model.Schema{"value": "${__amisgo__resp}"}),
				},
			},
		},
	)
}

// ActiveClassName sets the class name for the active state of the button.
func (a Action) ActiveClassName(value string) Action {
	return a.set("activeClassName", value)
}

// ActiveLevel sets the style for the active state of the button.
func (a Action) ActiveLevel(value string) Action {
	return a.set("activeLevel", value)
}

// Api sets the API address/description.
func (a Action) Api(value any) Action {
	return a.set("api", value)
}

// ClassName adds a class name to the button.
func (a Action) ClassName(value string) Action {
	return a.set("className", value)
}

// Close specifies whether to close the current dialog or drawer after the action is performed.
func (a Action) Close(value string) Action {
	return a.set("close", value)
}

// ConfirmText sets the text to be displayed in the confirmation dialog before the action is performed.
func (a Action) ConfirmText(value string) Action {
	return a.set("confirmText", value)
}

// ConfirmTitle sets the title of the confirmation dialog.
func (a Action) ConfirmTitle(value any) Action {
	return a.set("confirmTitle", value)
}

// DisabledTip sets the text to be displayed when the button is disabled.
func (a Action) DisabledTip(value any) Action {
	return a.set("disabledTip", value)
}

// Icon sets the icon for the button, e.g., "fa fa-plus".
func (a Action) Icon(value string) Action {
	return a.set("icon", value)
}

// IconClassName adds a class name to the icon.
func (a Action) IconClassName(value string) Action {
	return a.set("iconClassName", value)
}

// Redirect sets the relative path for single-page navigation.
func (a Action) Redirect(value string) Action {
	return a.set("redirect", value)
}

// Label sets the button text.
func (a Action) Label(value string) Action {
	return a.set("label", value)
}

// Level sets the button style.
// Possible values: link, primary, enhance, secondary, info, success, warning, danger, light, dark, default.
func (a Action) Level(value string) Action {
	return a.set("level", value)
}

// Link sets the link for the button.
func (a Action) Link(value string) Action {
	return a.set("link", value)
}

// Reload specifies the target component(s) to be refreshed after the action is performed.
func (a Action) Reload(value string) Action {
	return a.set("reload", value)
}

// ReloadWindow refreshes the current page.
func (a Action) ReloadWindow() Action {
	return a.Reload("window")
}

// Required sets the required fields for the form before the action is performed.
func (a Action) Required(value ...string) Action {
	return a.set("required", value)
}

// RightIcon sets the icon on the right side of the button text.
func (a Action) RightIcon(value string) Action {
	return a.set("rightIcon", value)
}

// RightIconClassName adds a class name to the right icon.
func (a Action) RightIconClassName(value string) Action {
	return a.set("rightIconClassName", value)
}

// Size sets the button size. Possible values: xs, sm, md, lg.
func (a Action) Size(value string) Action {
	return a.set("size", value)
}

// Tooltip sets the text to be displayed when the mouse hovers over the button.
func (a Action) Tooltip(value string) Action {
	return a.set("tooltip", value)
}

// TooltipPlacement sets the position of the tooltip. Possible values: top, bottom, left, right.
func (a Action) TooltipPlacement(value string) Action {
	return a.set("tooltipPlacement", value)
}

// TooltipTrigger sets the trigger for the tooltip. Possible values: hover, focus.
func (a Action) TooltipTrigger(value string) Action {
	return a.set("tooltipTrigger", value)
}

// Badge sets the badge for the button. Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/badge
func (v Action) Badge(value string) Action {
	return v.set("badge", value)
}

// Block sets whether the button is displayed as a block element.
func (v Action) Block(value bool) Action {
	return v.set("block", value)
}

// Body sets the child content of the button.
func (v Action) Body(value ...any) Action {
	return v.set("body", value)
}

// CountDown sets the countdown time (in seconds) after the button is clicked.
func (v Action) CountDown(value string) Action {
	return v.set("countDown", value)
}

// CountDownTpl sets the custom countdown text.
func (v Action) CountDownTpl(value string) Action {
	return v.set("countDownTpl", value)
}

// Disabled sets whether the button is disabled.
func (v Action) Disabled(value bool) Action {
	return v.set("disabled", value)
}

// DisabledOn sets the expression to determine whether the button is disabled.
func (v Action) DisabledOn(value string) Action {
	return v.set("disabledOn", value)
}

// DownloadFileName sets the file name for the download.
func (v Action) DownloadFileName(value string) Action {
	return v.set("downloadFileName", value)
}

// EditorSetting sets the editor configuration, which can be ignored at runtime.
func (v Action) EditorSetting(value string) Action {
	return v.set("editorSetting", value)
}

// Hidden sets whether the button is hidden.
func (v Action) Hidden(value bool) Action {
	return v.set("hidden", value)
}

// HiddenOn sets the expression to determine whether the button is hidden.
func (v Action) HiddenOn(value string) Action {
	return v.set("hiddenOn", value)
}

// HotKey sets the keyboard shortcut for the button.
func (v Action) HotKey(value string) Action {
	return v.set("hotKey", value)
}

// ID sets the ID for the button, mainly used for user behavior tracking.
func (v Action) ID(value string) Action {
	return v.set("id", value)
}

// Url sets the URL for the button.
func (v Action) Url(value string) Action {
	return v.set("url", value)
}

// Loading sets whether to show the loading effect on the button.
func (a Action) Loading(value bool) Action {
	return a.set("loading", value)
}

// LoadingClassName sets the CSS class name for the loading effect.
func (v Action) LoadingClassName(value string) Action {
	return v.set("loadingClassName", value)
}

// LoadingOn sets the expression to determine whether to show the loading effect.
func (v Action) LoadingOn(value string) Action {
	return v.set("loadingOn", value)
}

// MergeData sets whether to merge the data from the dialog or drawer into the parent scope.
func (v Action) MergeData(value bool) Action {
	return v.set("mergeData", value)
}

// OnClick sets the custom event handler function for the button.
func (v Action) OnClick(value string) Action {
	return v.set("onClick", value)
}

// OnEvent sets the event action configuration for the button.
func (v Action) OnEvent(value any) Action {
	return v.set("onEvent", value)
}

// Primary sets whether the button is a primary button.
func (v Action) Primary(value bool) Action {
	return v.set("primary", value)
}

// RequireSelected sets whether the button requires selected elements to be clickable when used as a batch operation button.
func (v Action) RequireSelected(value bool) Action {
	return v.set("requireSelected", value)
}

// Static sets whether the button is displayed statically.
func (v Action) Static(value bool) Action {
	return v.set("static", value)
}

// StaticClassName sets the CSS class name for the static display of the form item.
func (v Action) StaticClassName(value string) Action {
	return v.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for the static display of the form item value.
func (v Action) StaticInputClassName(value string) Action {
	return v.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for the static display of the form item label.
func (v Action) StaticLabelClassName(value string) Action {
	return v.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine whether the button is displayed statically.
func (v Action) StaticOn(value string) Action {
	return v.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder text for the static display.
func (v Action) StaticPlaceholder(value string) Action {
	return v.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for the static display mode.
func (v Action) StaticSchema(value string) Action {
	return v.set("staticSchema", value)
}

// Style sets the style for the button.
func (v Action) Style(value any) Action {
	return v.set("style", value)
}

// Target sets the target component(s) that can trigger this action.
func (v Action) Target(value string) Action {
	return v.set("target", value)
}

// TestIdBuilder sets the function to generate test IDs.
func (v Action) TestIdBuilder(value string) Action {
	return v.set("testIdBuilder", value)
}

// Testid sets the test ID for the button.
func (v Action) Testid(value string) Action {
	return v.set("testid", value)
}

// UseMobileUI sets whether to disable mobile UI styles.
func (v Action) UseMobileUI(value bool) Action {
	return v.set("useMobileUI", value)
}

// Visible sets whether the button is visible.
func (v Action) Visible(value bool) Action {
	return v.set("visible", value)
}

// VisibleOn sets the expression to determine whether the button is visible.
func (v Action) VisibleOn(value string) Action {
	return v.set("visibleOn", value)
}

func (a Action) set(key string, value any) Action {
	a[key] = value
	return a
}

func (a Action) mux() *http.ServeMux {
	return a[servemux.Key].(*http.ServeMux)
}
