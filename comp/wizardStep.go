package comp

import "github.com/zrcoder/amisgo/model"

// wizardStep represents a step in a wizard form

type wizardStep model.Schema

// WizardStep creates a new WizardStep instance
func WizardStep() wizardStep {
	return wizardStep{"type": "wizard-step"}
}

func (ws wizardStep) set(key string, value any) wizardStep {
	ws[key] = value
	return ws
}

// Actions sets the actions buttons displayed at the bottom
func (ws wizardStep) Actions(value string) wizardStep {
	return ws.set("actions", value)
}

// AffixFooter fixes the footer buttons at the bottom
func (ws wizardStep) AffixFooter(value bool) wizardStep {
	return ws.set("affixFooter", value)
}

// Api sets the API for saving form data
func (ws wizardStep) Api(value string) wizardStep {
	return ws.set("api", value)
}

// AsyncApi sets the API for polling until finished
func (ws wizardStep) AsyncApi(value string) wizardStep {
	return ws.set("asyncApi", value)
}

// AutoFocus sets whether to autofocus the first form element
func (ws wizardStep) AutoFocus(value bool) wizardStep {
	return ws.set("autoFocus", value)
}

// Body sets the form items
func (ws wizardStep) Body(value ...any) wizardStep {
	return ws.set("body", value)
}

// CheckInterval sets the polling interval, effective when asyncApi is set
func (ws wizardStep) CheckInterval(value string) wizardStep {
	return ws.set("checkInterval", value)
}

// ClassName sets the CSS class name for the container
func (ws wizardStep) ClassName(value string) wizardStep {
	return ws.set("className", value)
}

// ClearAfterSubmit clears the form after submission
func (ws wizardStep) ClearAfterSubmit(value bool) wizardStep {
	return ws.set("clearAfterSubmit", value)
}

// ClearPersistDataAfterSubmit clears local cache after successful submission
func (ws wizardStep) ClearPersistDataAfterSubmit(value bool) wizardStep {
	return ws.set("clearPersistDataAfterSubmit", value)
}

// ColumnCount sets the number of columns for form items
func (ws wizardStep) ColumnCount(value string) wizardStep {
	return ws.set("columnCount", value)
}

// Data sets the form data
func (ws wizardStep) Data(value any) wizardStep {
	return ws.set("data", value)
}

// Debug enables debug mode to show form data at the top
func (ws wizardStep) Debug(value bool) wizardStep {
	return ws.set("debug", value)
}

// DebugConfig sets the debug panel configuration
func (ws wizardStep) DebugConfig(value string) wizardStep {
	return ws.set("debugConfig", value)
}

// Description sets the description
func (ws wizardStep) Description(value string) wizardStep {
	return ws.set("description", value)
}

// Disabled disables the step
func (ws wizardStep) Disabled(value bool) wizardStep {
	return ws.set("disabled", value)
}

// DisabledOn sets the expression to disable the step
func (ws wizardStep) DisabledOn(value string) wizardStep {
	return ws.set("disabledOn", value)
}

// EditorSetting sets the editor configuration, ignored at runtime
func (ws wizardStep) EditorSetting(value string) wizardStep {
	return ws.set("editorSetting", value)
}

// Feedback sets the form feedback
func (ws wizardStep) Feedback(value string) wizardStep {
	return ws.set("feedback", value)
}

// FieldSet sets the field set
func (ws wizardStep) FieldSet(value string) wizardStep {
	return ws.set("fieldSet", value)
}

// FinishedField sets the field name to determine completion
func (ws wizardStep) FinishedField(value string) wizardStep {
	return ws.set("finishedField", value)
}

// Hidden hides the step
func (ws wizardStep) Hidden(value bool) wizardStep {
	return ws.set("hidden", value)
}

// HiddenOn sets the expression to hide the step
func (ws wizardStep) HiddenOn(value string) wizardStep {
	return ws.set("hiddenOn", value)
}

// Horizontal sets the horizontal layout ratio
func (ws wizardStep) Horizontal(value string) wizardStep {
	return ws.set("horizontal", value)
}

// Icon sets the icon
func (ws wizardStep) Icon(value string) wizardStep {
	return ws.set("icon", value)
}

// ID sets the unique component ID for logging
func (ws wizardStep) ID(value string) wizardStep {
	return ws.set("id", value)
}

// InitApi sets the API for initializing form data
func (ws wizardStep) InitApi(value string) wizardStep {
	return ws.set("initApi", value)
}

// InitAsyncApi sets the API for fetching initial data
func (ws wizardStep) InitAsyncApi(value string) wizardStep {
	return ws.set("initAsyncApi", value)
}

// InitCheckInterval sets the polling interval for initAsyncApi
func (ws wizardStep) InitCheckInterval(value string) wizardStep {
	return ws.set("initCheckInterval", value)
}

// InitFetch sets whether to fetch data initially
func (ws wizardStep) InitFetch(value bool) wizardStep {
	return ws.set("initFetch", value)
}

// InitFetchOn sets the expression to fetch data initially
func (ws wizardStep) InitFetchOn(value string) wizardStep {
	return ws.set("initFetchOn", value)
}

// InitFinishedField sets the field name to determine initial completion
func (ws wizardStep) InitFinishedField(value string) wizardStep {
	return ws.set("initFinishedField", value)
}

// Interval sets the interval for polling initApi
func (ws wizardStep) Interval(value string) wizardStep {
	return ws.set("interval", value)
}

// Jumpable sets whether the step is jumpable
func (ws wizardStep) Jumpable(value bool) wizardStep {
	return ws.set("jumpable", value)
}

// JumpableOn sets the expression to determine if the step is jumpable
func (ws wizardStep) JumpableOn(value string) wizardStep {
	return ws.set("jumpableOn", value)
}

// Label sets the label
func (ws wizardStep) Label(value string) wizardStep {
	return ws.set("label", value)
}

// LabelAlign sets the alignment of the form label
func (ws wizardStep) LabelAlign(value string) wizardStep {
	return ws.set("labelAlign", value)
}

// LabelWidth sets the custom width of the label, default unit is px
func (ws wizardStep) LabelWidth(value string) wizardStep {
	return ws.set("labelWidth", value)
}

// Messages sets the message configuration
func (ws wizardStep) Messages(value string) wizardStep {
	return ws.set("messages", value)
}

// Mode sets the default display mode for form items
func (ws wizardStep) Mode(value string) wizardStep {
	return ws.set("mode", value)
}

// Name sets the component name
func (ws wizardStep) Name(value string) wizardStep {
	return ws.set("name", value)
}

// OnEvent sets the event action configuration
func (ws wizardStep) OnEvent(value any) wizardStep {
	return ws.set("onEvent", value)
}

// PanelClassName sets the panel class name
func (ws wizardStep) PanelClassName(value string) wizardStep {
	return ws.set("panelClassName", value)
}

// PersistData enables local data caching
func (ws wizardStep) PersistData(value string) wizardStep {
	return ws.set("persistData", value)
}

// PersistDataKeys sets the keys to be cached locally
func (ws wizardStep) PersistDataKeys(value string) wizardStep {
	return ws.set("persistDataKeys", value)
}

// PreventEnterSubmit disables form submission on Enter key press
func (ws wizardStep) PreventEnterSubmit(value bool) wizardStep {
	return ws.set("preventEnterSubmit", value)
}

// PrimaryField sets the primary key field
func (ws wizardStep) PrimaryField(value string) wizardStep {
	return ws.set("primaryField", value)
}

// PromptPageLeave enables page leave prompt
func (ws wizardStep) PromptPageLeave(value bool) wizardStep {
	return ws.set("promptPageLeave", value)
}

// PromptPageLeaveMessage sets the page leave prompt message
func (ws wizardStep) PromptPageLeaveMessage(value string) wizardStep {
	return ws.set("promptPageLeaveMessage", value)
}

// Redirect sets the redirect URL
func (ws wizardStep) Redirect(value string) wizardStep {
	return ws.set("redirect", value)
}

// Reload sets the reload URL
func (ws wizardStep) Reload(value string) wizardStep {
	return ws.set("reload", value)
}

// ResetAfterSubmit resets the form after submission
func (ws wizardStep) ResetAfterSubmit(value bool) wizardStep {
	return ws.set("resetAfterSubmit", value)
}

// Rules sets the validation rules
func (ws wizardStep) Rules(value string) wizardStep {
	return ws.set("rules", value)
}

// SilentPolling enables silent polling
func (ws wizardStep) SilentPolling(value bool) wizardStep {
	return ws.set("silentPolling", value)
}

// Static sets the static display class name
func (ws wizardStep) Static(value bool) wizardStep {
	return ws.set("static", value)
}

// StaticClassName sets the static display CSS class name
func (ws wizardStep) StaticClassName(value string) wizardStep {
	return ws.set("staticClassName", value)
}

// StaticInputClassName sets the static display input class name
func (ws wizardStep) StaticInputClassName(value string) wizardStep {
	return ws.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static display label class name
func (ws wizardStep) StaticLabelClassName(value string) wizardStep {
	return ws.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display
func (ws wizardStep) StaticOn(value string) wizardStep {
	return ws.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (ws wizardStep) StaticPlaceholder(value string) wizardStep {
	return ws.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema
func (ws wizardStep) StaticSchema(value string) wizardStep {
	return ws.set("staticSchema", value)
}

// StopAutoRefreshWhen sets the condition to stop auto-refresh
func (ws wizardStep) StopAutoRefreshWhen(value string) wizardStep {
	return ws.set("stopAutoRefreshWhen", value)
}

// Style sets the component style
func (ws wizardStep) Style(value any) wizardStep {
	return ws.set("style", value)
}

// SubTitle sets the subtitle
func (ws wizardStep) SubTitle(value any) wizardStep {
	return ws.set("subTitle", value)
}

// SubmitOnChange submits the form on change
func (ws wizardStep) SubmitOnChange(value bool) wizardStep {
	return ws.set("submitOnChange", value)
}

// SubmitOnInit submits the form on initialization
func (ws wizardStep) SubmitOnInit(value bool) wizardStep {
	return ws.set("submitOnInit", value)
}

// SubmitText sets the default submit button text
func (ws wizardStep) SubmitText(value string) wizardStep {
	return ws.set("submitText", value)
}

// Tabs sets the tabs
func (ws wizardStep) Tabs(value string) wizardStep {
	return ws.set("tabs", value)
}

// Target sets the target for form submission
func (ws wizardStep) Target(value string) wizardStep {
	return ws.set("target", value)
}

// TestIdBuilder sets the test ID builder
func (ws wizardStep) TestIdBuilder(value string) wizardStep {
	return ws.set("testIdBuilder", value)
}

// Testid sets the test ID
func (ws wizardStep) Testid(value string) wizardStep {
	return ws.set("testid", value)
}

// Title sets the form title
func (ws wizardStep) Title(value any) wizardStep {
	return ws.set("title", value)
}

// UseMobileUI disables mobile UI styles
func (ws wizardStep) UseMobileUI(value bool) wizardStep {
	return ws.set("useMobileUI", value)
}

// Value sets the value
func (ws wizardStep) Value(value string) wizardStep {
	return ws.set("value", value)
}

// Visible sets the visibility
func (ws wizardStep) Visible(value bool) wizardStep {
	return ws.set("visible", value)
}

// VisibleOn sets the expression for visibility
func (ws wizardStep) VisibleOn(value string) wizardStep {
	return ws.set("visibleOn", value)
}

// WrapWithPanel wraps the step with a panel
func (ws wizardStep) WrapWithPanel(value bool) wizardStep {
	return ws.set("wrapWithPanel", value)
}
