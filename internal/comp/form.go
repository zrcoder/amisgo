package comp

import (
	"net/http"

	"github.com/zrcoder/amisgo/internal/servemux"
	"github.com/zrcoder/amisgo/schema"
)

// Form represents a form renderer.
type Form schema.Schema

func NewForm(mux *http.ServeMux) Form {
	return Form{"type": "form", servemux.Key: mux}
}

func (f Form) set(key string, value any) Form {
	f[key] = value
	return f
}

// ClassName sets the class name.
func (f Form) ClassName(value string) Form {
	return f.set("classname", value)
}

// PanelClassName sets the outer panel class name.
func (f Form) PanelClassName(value string) Form {
	return f.set("panelClassName", value)
}

// Reload configures form reload.
func (f Form) Reload(value string) Form {
	return f.set("reload", value)
}

// Mode sets the display mode: normal, inline, or horizontal.
func (f Form) Mode(value string) Form {
	return f.set("mode", value)
}

// AutoFocus sets whether to auto-focus.
func (f Form) AutoFocus(value bool) Form {
	return f.set("autoFocus", value)
}

// Horizontal sets the label display ratio when mode is horizontal.
func (f Form) Horizontal(value any) Form {
	return f.set("horizontal", value)
}

// LabelAlign sets the label alignment: left or right.
func (f Form) LabelAlign(value any) Form {
	return f.set("labelAlign", value)
}

// LabelWidth sets the custom label width.
func (f Form) LabelWidth(value any) Form {
	return f.set("labelWidth", value)
}

// Body sets the form content.
func (f Form) Body(value ...any) Form {
	return f.set("body", value)
}

// ResetAfterSubmit sets whether to reset the form after submission.
func (f Form) ResetAfterSubmit(value bool) Form {
	return f.set("resetAfterSubmit", value)
}

// Rules sets the validation rules.
func (f Form) Rules(value ...any) Form {
	return f.set("rules", value)
}

// SilentPolling sets whether to poll silently.
func (f Form) SilentPolling(value bool) Form {
	return f.set("silentPolling", value)
}

// Static sets the static display class name.
func (f Form) Static(value bool) Form {
	return f.set("static", value)
}

// StaticClassName sets the static class name.
func (f Form) StaticClassName(value string) Form {
	return f.set("staticClassName", value)
}

// StaticInputClassName sets the static input class name.
func (f Form) StaticInputClassName(value string) Form {
	return f.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static label class name.
func (f Form) StaticLabelClassName(value string) Form {
	return f.set("staticLabelClassName", value)
}

// StaticOn sets the static display condition.
func (f Form) StaticOn(value string) Form {
	return f.set("staticOn", value)
}

// StaticPlaceholder sets the static placeholder.
func (f Form) StaticPlaceholder(value string) Form {
	return f.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema.Schema.
func (f Form) StaticSchema(value string) Form {
	return f.set("staticSchema", value)
}

// StopAutoRefreshWhen sets the condition to stop auto-refresh.
func (f Form) StopAutoRefreshWhen(value string) Form {
	return f.set("stopAutoRefreshWhen", value)
}

// Style sets the component style.
func (f Form) Style(value any) Form {
	return f.set("style", value)
}

// SubmitOnChange sets whether to submit the form on change.
func (f Form) SubmitOnChange(value bool) Form {
	return f.set("submitOnChange", value)
}

// SubmitOnInit sets whether to submit the form on initialization.
func (f Form) SubmitOnInit(value bool) Form {
	return f.set("submitOnInit", value)
}

// SubmitText sets the default submit button text.
func (f Form) SubmitText(value string) Form {
	return f.set("submitText", value)
}

// Tabs sets the form tabs.
func (f Form) Tabs(value string) Form {
	return f.set("tabs", value)
}

// Target sets the target form or CRUD model name.
func (f Form) Target(value string) Form {
	return f.set("target", value)
}

// TestIdBuilder sets the test ID builder.
func (f Form) TestIdBuilder(value string) Form {
	return f.set("testIdBuilder", value)
}

// Testid sets the test ID.
func (f Form) Testid(value string) Form {
	return f.set("testid", value)
}

// Title sets the form title.
func (f Form) Title(value string) Form {
	return f.set("title", value)
}

// ColumnCount sets the number of columns.
func (f Form) ColumnCount(value int) Form {
	return f.set("columnCount", value)
}

// UseMobileUI sets whether to use mobile UI.
func (f Form) UseMobileUI(value bool) Form {
	return f.set("useMobileUI", value)
}

// Visible sets whether the form is visible.
func (f Form) Visible(value bool) Form {
	return f.set("visible", value)
}

// VisibleOn sets the visibility condition.
func (f Form) VisibleOn(value string) Form {
	return f.set("visibleOn", value)
}

// WrapWithPanel sets whether to wrap the form with a panel.
func (f Form) WrapWithPanel(value bool) Form {
	return f.set("wrapWithPanel", value)
}

// Actions sets the form actions.
func (f Form) Actions(value ...Action) Form {
	return f.set("actions", value)
}

// InitApi sets the API to fetch initial data.
func (f Form) InitApi(value any) Form {
	return f.set("initApi", value)
}

// Api sets the API to save data.
func (f Form) Api(value any) Form {
	return f.set("api", value)
}

// PreventEnterSubmit prevents form submission when the Enter key is pressed
func (f Form) PreventEnterSubmit(value bool) Form {
	return f.set("preventEnterSubmit", value)
}

// Redirect specifies the URL path to redirect to after form submission
func (f Form) Redirect(value string) Form {
	return f.set("redirect", value)
}

// GetData fetches data using the provided getter function.
func (f Form) GetData(getter func() (any, error)) Form {
	return f.Api(servemux.ServeData(f.mux(), getter))
}

// Submit sets the callback logic after form submission, using the generic schema.Schema type to handle form data
// Suitable for scenarios where flexible handling of form submission is required
func (f Form) Submit(callback func(schema.Schema) error) Form {
	return f.Api(servemux.BindDataRoute(f.mux(), callback))
}

// SubmitTo submits the form data to the specified struct or map and executes a custom callback
// receiver can be a struct pointer, map, or other types that can be deserialized from JSON
// callback allows further processing of the received data
func (f Form) SubmitTo(receiver any, callback func(any) error) Form {
	return f.Api(servemux.BindRouteTo(f.mux(), receiver, callback))
}

func (f Form) mux() *http.ServeMux {
	return f[servemux.Key].(*http.ServeMux)
}
