package comp

import (
	"github.com/zrcoder/amisgo/internal/servermux"
	"github.com/zrcoder/amisgo/model"
)

// form represents a form renderer.

type form model.Schema

// Form creates a new Form instance.
func Form() form {
	return make(form).set("type", "form")
}

func (f form) set(key string, value any) form {
	f[key] = value
	return f
}

// ClassName sets the class name.
func (f form) ClassName(value string) form {
	return f.set("classname", value)
}

// PanelClassName sets the outer panel class name.
func (f form) PanelClassName(value string) form {
	return f.set("panelClassName", value)
}

// Reload configures form reload.
func (f form) Reload(value string) form {
	return f.set("reload", value)
}

// Mode sets the display mode: normal, inline, or horizontal.
func (f form) Mode(value string) form {
	return f.set("mode", value)
}

// AutoFocus sets whether to auto-focus.
func (f form) AutoFocus(value bool) form {
	return f.set("autoFocus", value)
}

// Horizontal sets the label display ratio when mode is horizontal.
func (f form) Horizontal(value any) form {
	return f.set("horizontal", value)
}

// LabelAlign sets the label alignment: left or right.
func (f form) LabelAlign(value any) form {
	return f.set("labelAlign", value)
}

// LabelWidth sets the custom label width.
func (f form) LabelWidth(value any) form {
	return f.set("labelWidth", value)
}

// Body sets the form content.
func (f form) Body(value ...any) form {
	return f.set("body", value)
}

// ResetAfterSubmit sets whether to reset the form after submission.
func (f form) ResetAfterSubmit(value bool) form {
	return f.set("resetAfterSubmit", value)
}

// Rules sets the validation rules.
func (f form) Rules(value ...any) form {
	return f.set("rules", value)
}

// SilentPolling sets whether to poll silently.
func (f form) SilentPolling(value bool) form {
	return f.set("silentPolling", value)
}

// Static sets the static display class name.
func (f form) Static(value bool) form {
	return f.set("static", value)
}

// StaticClassName sets the static class name.
func (f form) StaticClassName(value string) form {
	return f.set("staticClassName", value)
}

// StaticInputClassName sets the static input class name.
func (f form) StaticInputClassName(value string) form {
	return f.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static label class name.
func (f form) StaticLabelClassName(value string) form {
	return f.set("staticLabelClassName", value)
}

// StaticOn sets the static display condition.
func (f form) StaticOn(value string) form {
	return f.set("staticOn", value)
}

// StaticPlaceholder sets the static placeholder.
func (f form) StaticPlaceholder(value string) form {
	return f.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema.
func (f form) StaticSchema(value string) form {
	return f.set("staticSchema", value)
}

// StopAutoRefreshWhen sets the condition to stop auto-refresh.
func (f form) StopAutoRefreshWhen(value string) form {
	return f.set("stopAutoRefreshWhen", value)
}

// Style sets the component style.
func (f form) Style(value any) form {
	return f.set("style", value)
}

// SubmitOnChange sets whether to submit the form on change.
func (f form) SubmitOnChange(value bool) form {
	return f.set("submitOnChange", value)
}

// SubmitOnInit sets whether to submit the form on initialization.
func (f form) SubmitOnInit(value bool) form {
	return f.set("submitOnInit", value)
}

// SubmitText sets the default submit button text.
func (f form) SubmitText(value string) form {
	return f.set("submitText", value)
}

// Tabs sets the form tabs.
func (f form) Tabs(value string) form {
	return f.set("tabs", value)
}

// Target sets the target form or CRUD model name.
func (f form) Target(value string) form {
	return f.set("target", value)
}

// TestIdBuilder sets the test ID builder.
func (f form) TestIdBuilder(value string) form {
	return f.set("testIdBuilder", value)
}

// Testid sets the test ID.
func (f form) Testid(value string) form {
	return f.set("testid", value)
}

// Title sets the form title.
func (f form) Title(value string) form {
	return f.set("title", value)
}

// ColumnCount sets the number of columns.
func (f form) ColumnCount(value int) form {
	return f.set("columnCount", value)
}

// UseMobileUI sets whether to use mobile UI.
func (f form) UseMobileUI(value bool) form {
	return f.set("useMobileUI", value)
}

// Visible sets whether the form is visible.
func (f form) Visible(value bool) form {
	return f.set("visible", value)
}

// VisibleOn sets the visibility condition.
func (f form) VisibleOn(value string) form {
	return f.set("visibleOn", value)
}

// WrapWithPanel sets whether to wrap the form with a panel.
func (f form) WrapWithPanel(value bool) form {
	return f.set("wrapWithPanel", value)
}

// Actions sets the form actions.
func (f form) Actions(value ...action) form {
	return f.set("actions", value)
}

// InitApi sets the API to fetch initial data.
func (f form) InitApi(value any) form {
	return f.set("initApi", value)
}

// Api sets the API to save data.
func (f form) Api(value any) form {
	return f.set("api", value)
}

// PreventEnterSubmit prevents form submission when the Enter key is pressed
func (f form) PreventEnterSubmit(value bool) form {
	return f.set("preventEnterSubmit", value)
}

// Redirect specifies the URL path to redirect to after form submission
func (f form) Redirect(value string) form {
	return f.set("redirect", value)
}

// GetData fetches data using the provided getter function.
func (s form) GetData(getter func() (any, error)) form {
	return s.Api(servermux.ServeData(getter))
}

// Submit sets the callback logic after form submission, using the generic Data type to handle form data
// Suitable for scenarios where flexible handling of form submission is required
func (f form) Submit(callback func(model.Schema) error) form {
	return f.Api(servermux.BindDataRoute(callback))
}

// SubmitTo submits the form data to the specified struct or map and executes a custom callback
// receiver can be a struct pointer, map, or other types that can be deserialized from JSON
// callback allows further processing of the received data
func (f form) SubmitTo(receiver any, callback func(any) error) form {
	return f.Api(servermux.BindRouteTo(receiver, callback))
}
