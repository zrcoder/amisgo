package comp

import "github.com/zrcoder/amisgo/schema"

// Tasks renderer, documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/Tasks
type Tasks schema.Schema

func NewTasks() Tasks {
	return Tasks{"type": "tasks"}
}

// set sets a property value and returns the updated Tasks instance
func (t Tasks) set(key string, value any) Tasks {
	t[key] = value
	return t
}

// BtnClassName sets the CSS class name for the button
func (t Tasks) BtnClassName(value string) Tasks {
	return t.set("btnClassName", value)
}

// BtnText sets the button text
func (t Tasks) BtnText(value string) Tasks {
	return t.set("btnText", value)
}

// CanRetryStatusCode sets the status code for retryable tasks
func (t Tasks) CanRetryStatusCode(value string) Tasks {
	return t.set("canRetryStatusCode", value)
}

// CheckApi sets the API for checking task status
func (t Tasks) CheckApi(value string) Tasks {
	return t.set("checkApi", value)
}

// ClassName sets the CSS class name for the container
func (t Tasks) ClassName(value string) Tasks {
	return t.set("className", value)
}

// Disabled sets whether the component is disabled
func (t Tasks) Disabled(value bool) Tasks {
	return t.set("disabled", value)
}

// DisabledOn sets the expression to determine if the component is disabled
func (t Tasks) DisabledOn(value string) Tasks {
	return t.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (t Tasks) EditorSetting(value string) Tasks {
	return t.set("editorSetting", value)
}

// ErrorStatusCode sets the error status code
func (t Tasks) ErrorStatusCode(value string) Tasks {
	return t.set("errorStatusCode", value)
}

// FinishStatusCode sets the finish status code
func (t Tasks) FinishStatusCode(value string) Tasks {
	return t.set("finishStatusCode", value)
}

// Hidden sets whether the component is hidden
func (t Tasks) Hidden(value bool) Tasks {
	return t.set("hidden", value)
}

// HiddenOn sets the expression to determine if the component is hidden
func (t Tasks) HiddenOn(value string) Tasks {
	return t.set("hiddenOn", value)
}

// ID sets the unique ID for the component
func (t Tasks) ID(value string) Tasks {
	return t.set("id", value)
}

// InitialStatusCode sets the initial status code
func (t Tasks) InitialStatusCode(value string) Tasks {
	return t.set("initialStatusCode", value)
}

// Interval sets the interval for checking task status
func (t Tasks) Interval(value string) Tasks {
	return t.set("interval", value)
}

// Items sets the items for the tasks
func (t Tasks) Items(value ...any) Tasks {
	return t.set("items", value)
}

// LoadingConfig sets the loading configuration
func (t Tasks) LoadingConfig(value string) Tasks {
	return t.set("loadingConfig", value)
}

// LoadingStatusCode sets the loading status code
func (t Tasks) LoadingStatusCode(value string) Tasks {
	return t.set("loadingStatusCode", value)
}

// Name sets the name of the component
func (t Tasks) Name(value string) Tasks {
	return t.set("name", value)
}

// OnEvent sets the event action configuration
func (t Tasks) OnEvent(value Event) Tasks {
	return t.set("onEvent", value)
}

// OperationLabel sets the operation column label
func (t Tasks) OperationLabel(value string) Tasks {
	return t.set("operationLabel", value)
}

// ReSubmitApi sets the API for resubmitting failed tasks
func (t Tasks) ReSubmitApi(value string) Tasks {
	return t.set("reSubmitApi", value)
}

// ReadyStatusCode sets the ready status code
func (t Tasks) ReadyStatusCode(value string) Tasks {
	return t.set("readyStatusCode", value)
}

// RemarkLabel sets the remark column label
func (t Tasks) RemarkLabel(value string) Tasks {
	return t.set("remarkLabel", value)
}

// RetryBtnClassName sets the CSS class name for the retry button
func (t Tasks) RetryBtnClassName(value string) Tasks {
	return t.set("retryBtnClassName", value)
}

// RetryBtnText sets the retry button text
func (t Tasks) RetryBtnText(value string) Tasks {
	return t.set("retryBtnText", value)
}

// Static sets whether the component is displayed statically
func (t Tasks) Static(value bool) Tasks {
	return t.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (t Tasks) StaticClassName(value string) Tasks {
	return t.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (t Tasks) StaticInputClassName(value string) Tasks {
	return t.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (t Tasks) StaticLabelClassName(value string) Tasks {
	return t.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the component is displayed statically
func (t Tasks) StaticOn(value string) Tasks {
	return t.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (t Tasks) StaticPlaceholder(value string) Tasks {
	return t.set("staticPlaceholder", value)
}

// StaticSchema sets the schema.Schema for static display
func (t Tasks) StaticSchema(value string) Tasks {
	return t.set("staticSchema", value)
}

// StatusLabel sets the status column label
func (t Tasks) StatusLabel(value string) Tasks {
	return t.set("statusLabel", value)
}

// StatusLabelMap sets the status label map
func (t Tasks) StatusLabelMap(value string) Tasks {
	return t.set("statusLabelMap", value)
}

// StatusTextMap sets the status text map
func (t Tasks) StatusTextMap(value string) Tasks {
	return t.set("statusTextMap", value)
}

// Style sets the component style
func (t Tasks) Style(value any) Tasks {
	return t.set("style", value)
}

// SubmitApi sets the API for submitting tasks
func (t Tasks) SubmitApi(value string) Tasks {
	return t.set("submitApi", value)
}

// TableClassName sets the CSS class name for the table
func (t Tasks) TableClassName(value string) Tasks {
	return t.set("tableClassName", value)
}

// TaskNameLabel sets the task name column label
func (t Tasks) TaskNameLabel(value string) Tasks {
	return t.set("taskNameLabel", value)
}

// TestIdBuilder sets the test ID builder
func (t Tasks) TestIdBuilder(value string) Tasks {
	return t.set("testIdBuilder", value)
}

// Testid sets the test ID
func (t Tasks) Testid(value string) Tasks {
	return t.set("testid", value)
}

// UseMobileUI sets whether to use mobile UI
func (t Tasks) UseMobileUI(value bool) Tasks {
	return t.set("useMobileUI", value)
}

// Visible sets whether the component is visible
func (t Tasks) Visible(value bool) Tasks {
	return t.set("visible", value)
}

// VisibleOn sets the expression to determine if the component is visible
func (t Tasks) VisibleOn(value string) Tasks {
	return t.set("visibleOn", value)
}
