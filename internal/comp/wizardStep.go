package comp

import "github.com/zrcoder/amisgo/schema"

// WizardStep represents a step in a wizard form
type WizardStep schema.Schema

func NewWizardStep() WizardStep {
	return WizardStep{"type": "wizard-step"}
}

func (ws WizardStep) set(key string, value any) WizardStep {
	ws[key] = value
	return ws
}

// Actions sets the actions buttons displayed at the bottom
func (ws WizardStep) Actions(value ...Action) WizardStep {
	return ws.set("actions", value)
}

// AffixFooter fixes the footer buttons at the bottom
func (ws WizardStep) AffixFooter(value bool) WizardStep {
	return ws.set("affixFooter", value)
}

// Api sets the API for saving form data
func (ws WizardStep) Api(value string) WizardStep {
	return ws.set("api", value)
}

// AsyncApi sets the API for polling until finished
func (ws WizardStep) AsyncApi(value string) WizardStep {
	return ws.set("asyncApi", value)
}

// AutoFocus sets whether to autofocus the first form element
func (ws WizardStep) AutoFocus(value bool) WizardStep {
	return ws.set("autoFocus", value)
}

// Body sets the form items
func (ws WizardStep) Body(value ...any) WizardStep {
	return ws.set("body", value)
}

// CheckInterval sets the polling interval, effective when asyncApi is set
func (ws WizardStep) CheckInterval(value string) WizardStep {
	return ws.set("checkInterval", value)
}

// ClassName sets the CSS class name for the container
func (ws WizardStep) ClassName(value string) WizardStep {
	return ws.set("className", value)
}

// ClearAfterSubmit clears the form after submission
func (ws WizardStep) ClearAfterSubmit(value bool) WizardStep {
	return ws.set("clearAfterSubmit", value)
}

// ClearPersistDataAfterSubmit clears local cache after successful submission
func (ws WizardStep) ClearPersistDataAfterSubmit(value bool) WizardStep {
	return ws.set("clearPersistDataAfterSubmit", value)
}

// ColumnCount sets the number of columns for form items
func (ws WizardStep) ColumnCount(value string) WizardStep {
	return ws.set("columnCount", value)
}

// Data sets the form data
func (ws WizardStep) Data(value any) WizardStep {
	return ws.set("data", value)
}

// Debug enables debug mode to show form data at the top
func (ws WizardStep) Debug(value bool) WizardStep {
	return ws.set("debug", value)
}

// DebugConfig sets the debug panel configuration
func (ws WizardStep) DebugConfig(value string) WizardStep {
	return ws.set("debugConfig", value)
}

// Description sets the description
func (ws WizardStep) Description(value string) WizardStep {
	return ws.set("description", value)
}

// Disabled disables the step
func (ws WizardStep) Disabled(value bool) WizardStep {
	return ws.set("disabled", value)
}

// DisabledOn sets the expression to disable the step
func (ws WizardStep) DisabledOn(value string) WizardStep {
	return ws.set("disabledOn", value)
}

// EditorSetting sets the editor configuration, ignored at runtime
func (ws WizardStep) EditorSetting(value string) WizardStep {
	return ws.set("editorSetting", value)
}

// Feedback sets the form feedback
func (ws WizardStep) Feedback(value string) WizardStep {
	return ws.set("feedback", value)
}

// FieldSet sets the field set
func (ws WizardStep) FieldSet(value string) WizardStep {
	return ws.set("fieldSet", value)
}

// FinishedField sets the field name to determine completion
func (ws WizardStep) FinishedField(value string) WizardStep {
	return ws.set("finishedField", value)
}

// Hidden hides the step
func (ws WizardStep) Hidden(value bool) WizardStep {
	return ws.set("hidden", value)
}

// HiddenOn sets the expression to hide the step
func (ws WizardStep) HiddenOn(value string) WizardStep {
	return ws.set("hiddenOn", value)
}

// Horizontal sets the horizontal layout ratio
func (ws WizardStep) Horizontal(value string) WizardStep {
	return ws.set("horizontal", value)
}

// Icon sets the icon
func (ws WizardStep) Icon(value string) WizardStep {
	return ws.set("icon", value)
}

// ID sets the unique component ID for logging
func (ws WizardStep) ID(value string) WizardStep {
	return ws.set("id", value)
}

// InitApi sets the API for initializing form data
func (ws WizardStep) InitApi(value string) WizardStep {
	return ws.set("initApi", value)
}

// InitAsyncApi sets the API for fetching initial data
func (ws WizardStep) InitAsyncApi(value string) WizardStep {
	return ws.set("initAsyncApi", value)
}

// InitCheckInterval sets the polling interval for initAsyncApi
func (ws WizardStep) InitCheckInterval(value string) WizardStep {
	return ws.set("initCheckInterval", value)
}

// InitFetch sets whether to fetch data initially
func (ws WizardStep) InitFetch(value bool) WizardStep {
	return ws.set("initFetch", value)
}

// InitFetchOn sets the expression to fetch data initially
func (ws WizardStep) InitFetchOn(value string) WizardStep {
	return ws.set("initFetchOn", value)
}

// InitFinishedField sets the field name to determine initial completion
func (ws WizardStep) InitFinishedField(value string) WizardStep {
	return ws.set("initFinishedField", value)
}

// Interval sets the interval for polling initApi
func (ws WizardStep) Interval(value string) WizardStep {
	return ws.set("interval", value)
}

// Jumpable sets whether the step is jumpable
func (ws WizardStep) Jumpable(value bool) WizardStep {
	return ws.set("jumpable", value)
}

// JumpableOn sets the expression to determine if the step is jumpable
func (ws WizardStep) JumpableOn(value string) WizardStep {
	return ws.set("jumpableOn", value)
}

// Label sets the label
func (ws WizardStep) Label(value string) WizardStep {
	return ws.set("label", value)
}

// LabelAlign sets the alignment of the form label
func (ws WizardStep) LabelAlign(value string) WizardStep {
	return ws.set("labelAlign", value)
}

// LabelWidth sets the custom width of the label, default unit is px
func (ws WizardStep) LabelWidth(value string) WizardStep {
	return ws.set("labelWidth", value)
}

// Messages sets the message configuration
func (ws WizardStep) Messages(value string) WizardStep {
	return ws.set("messages", value)
}

// Mode sets the default display mode for form items
func (ws WizardStep) Mode(value string) WizardStep {
	return ws.set("mode", value)
}

// Name sets the component name
func (ws WizardStep) Name(value string) WizardStep {
	return ws.set("name", value)
}

// OnEvent sets the event action configuration
func (ws WizardStep) OnEvent(value Event) WizardStep {
	return ws.set("onEvent", value)
}

// PanelClassName sets the panel class name
func (ws WizardStep) PanelClassName(value string) WizardStep {
	return ws.set("panelClassName", value)
}

// PersistData enables local data caching
func (ws WizardStep) PersistData(value string) WizardStep {
	return ws.set("persistData", value)
}

// PersistDataKeys sets the keys to be cached locally
func (ws WizardStep) PersistDataKeys(value string) WizardStep {
	return ws.set("persistDataKeys", value)
}

// PreventEnterSubmit disables form submission on Enter key press
func (ws WizardStep) PreventEnterSubmit(value bool) WizardStep {
	return ws.set("preventEnterSubmit", value)
}

// PrimaryField sets the primary key field
func (ws WizardStep) PrimaryField(value string) WizardStep {
	return ws.set("primaryField", value)
}

// PromptPageLeave enables page leave prompt
func (ws WizardStep) PromptPageLeave(value bool) WizardStep {
	return ws.set("promptPageLeave", value)
}

// PromptPageLeaveMessage sets the page leave prompt message
func (ws WizardStep) PromptPageLeaveMessage(value string) WizardStep {
	return ws.set("promptPageLeaveMessage", value)
}

// Redirect sets the redirect URL
func (ws WizardStep) Redirect(value string) WizardStep {
	return ws.set("redirect", value)
}

// Reload sets the reload URL
func (ws WizardStep) Reload(value string) WizardStep {
	return ws.set("reload", value)
}

// ResetAfterSubmit resets the form after submission
func (ws WizardStep) ResetAfterSubmit(value bool) WizardStep {
	return ws.set("resetAfterSubmit", value)
}

// Rules sets the validation rules
func (ws WizardStep) Rules(value string) WizardStep {
	return ws.set("rules", value)
}

// SilentPolling enables silent polling
func (ws WizardStep) SilentPolling(value bool) WizardStep {
	return ws.set("silentPolling", value)
}

// Static sets the static display class name
func (ws WizardStep) Static(value bool) WizardStep {
	return ws.set("static", value)
}

// StaticClassName sets the static display CSS class name
func (ws WizardStep) StaticClassName(value string) WizardStep {
	return ws.set("staticClassName", value)
}

// StaticInputClassName sets the static display input class name
func (ws WizardStep) StaticInputClassName(value string) WizardStep {
	return ws.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static display label class name
func (ws WizardStep) StaticLabelClassName(value string) WizardStep {
	return ws.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display
func (ws WizardStep) StaticOn(value string) WizardStep {
	return ws.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (ws WizardStep) StaticPlaceholder(value string) WizardStep {
	return ws.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema.Schema
func (ws WizardStep) StaticSchema(value string) WizardStep {
	return ws.set("staticSchema", value)
}

// StopAutoRefreshWhen sets the condition to stop auto-refresh
func (ws WizardStep) StopAutoRefreshWhen(value string) WizardStep {
	return ws.set("stopAutoRefreshWhen", value)
}

// Style sets the component style
func (ws WizardStep) Style(value any) WizardStep {
	return ws.set("style", value)
}

// SubTitle sets the subtitle
func (ws WizardStep) SubTitle(value any) WizardStep {
	return ws.set("subTitle", value)
}

// SubmitOnChange submits the form on change
func (ws WizardStep) SubmitOnChange(value bool) WizardStep {
	return ws.set("submitOnChange", value)
}

// SubmitOnInit submits the form on initialization
func (ws WizardStep) SubmitOnInit(value bool) WizardStep {
	return ws.set("submitOnInit", value)
}

// SubmitText sets the default submit button text
func (ws WizardStep) SubmitText(value string) WizardStep {
	return ws.set("submitText", value)
}

// Tabs sets the tabs
func (ws WizardStep) Tabs(value string) WizardStep {
	return ws.set("tabs", value)
}

// Target sets the target for form submission
func (ws WizardStep) Target(value string) WizardStep {
	return ws.set("target", value)
}

// TestIdBuilder sets the test ID builder
func (ws WizardStep) TestIdBuilder(value string) WizardStep {
	return ws.set("testIdBuilder", value)
}

// Testid sets the test ID
func (ws WizardStep) Testid(value string) WizardStep {
	return ws.set("testid", value)
}

// Title sets the form title
func (ws WizardStep) Title(value any) WizardStep {
	return ws.set("title", value)
}

// UseMobileUI disables mobile UI styles
func (ws WizardStep) UseMobileUI(value bool) WizardStep {
	return ws.set("useMobileUI", value)
}

// Value sets the value
func (ws WizardStep) Value(value string) WizardStep {
	return ws.set("value", value)
}

// Visible sets the visibility
func (ws WizardStep) Visible(value bool) WizardStep {
	return ws.set("visible", value)
}

// VisibleOn sets the expression for visibility
func (ws WizardStep) VisibleOn(value string) WizardStep {
	return ws.set("visibleOn", value)
}

// WrapWithPanel wraps the step with a panel
func (ws WizardStep) WrapWithPanel(value bool) WizardStep {
	return ws.set("wrapWithPanel", value)
}
