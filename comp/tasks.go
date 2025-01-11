package comp

import "github.com/zrcoder/amisgo/model"

// tasks renderer, documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/tasks

type tasks model.Schema

// Tasks creates a new Tasks instance
func Tasks() tasks {
	return tasks{"type": "tasks"}
}

// set sets a property value and returns the updated Tasks instance
func (t tasks) set(key string, value any) tasks {
	t[key] = value
	return t
}

// BtnClassName sets the CSS class name for the button
func (t tasks) BtnClassName(value string) tasks {
	return t.set("btnClassName", value)
}

// BtnText sets the button text
func (t tasks) BtnText(value string) tasks {
	return t.set("btnText", value)
}

// CanRetryStatusCode sets the status code for retryable tasks
func (t tasks) CanRetryStatusCode(value string) tasks {
	return t.set("canRetryStatusCode", value)
}

// CheckApi sets the API for checking task status
func (t tasks) CheckApi(value string) tasks {
	return t.set("checkApi", value)
}

// ClassName sets the CSS class name for the container
func (t tasks) ClassName(value string) tasks {
	return t.set("className", value)
}

// Disabled sets whether the component is disabled
func (t tasks) Disabled(value bool) tasks {
	return t.set("disabled", value)
}

// DisabledOn sets the expression to determine if the component is disabled
func (t tasks) DisabledOn(value string) tasks {
	return t.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (t tasks) EditorSetting(value string) tasks {
	return t.set("editorSetting", value)
}

// ErrorStatusCode sets the error status code
func (t tasks) ErrorStatusCode(value string) tasks {
	return t.set("errorStatusCode", value)
}

// FinishStatusCode sets the finish status code
func (t tasks) FinishStatusCode(value string) tasks {
	return t.set("finishStatusCode", value)
}

// Hidden sets whether the component is hidden
func (t tasks) Hidden(value bool) tasks {
	return t.set("hidden", value)
}

// HiddenOn sets the expression to determine if the component is hidden
func (t tasks) HiddenOn(value string) tasks {
	return t.set("hiddenOn", value)
}

// ID sets the unique ID for the component
func (t tasks) ID(value string) tasks {
	return t.set("id", value)
}

// InitialStatusCode sets the initial status code
func (t tasks) InitialStatusCode(value string) tasks {
	return t.set("initialStatusCode", value)
}

// Interval sets the interval for checking task status
func (t tasks) Interval(value string) tasks {
	return t.set("interval", value)
}

// Items sets the items for the tasks
func (t tasks) Items(value ...any) tasks {
	return t.set("items", value)
}

// LoadingConfig sets the loading configuration
func (t tasks) LoadingConfig(value string) tasks {
	return t.set("loadingConfig", value)
}

// LoadingStatusCode sets the loading status code
func (t tasks) LoadingStatusCode(value string) tasks {
	return t.set("loadingStatusCode", value)
}

// Name sets the name of the component
func (t tasks) Name(value string) tasks {
	return t.set("name", value)
}

// OnEvent sets the event action configuration
func (t tasks) OnEvent(value any) tasks {
	return t.set("onEvent", value)
}

// OperationLabel sets the operation column label
func (t tasks) OperationLabel(value string) tasks {
	return t.set("operationLabel", value)
}

// ReSubmitApi sets the API for resubmitting failed tasks
func (t tasks) ReSubmitApi(value string) tasks {
	return t.set("reSubmitApi", value)
}

// ReadyStatusCode sets the ready status code
func (t tasks) ReadyStatusCode(value string) tasks {
	return t.set("readyStatusCode", value)
}

// RemarkLabel sets the remark column label
func (t tasks) RemarkLabel(value string) tasks {
	return t.set("remarkLabel", value)
}

// RetryBtnClassName sets the CSS class name for the retry button
func (t tasks) RetryBtnClassName(value string) tasks {
	return t.set("retryBtnClassName", value)
}

// RetryBtnText sets the retry button text
func (t tasks) RetryBtnText(value string) tasks {
	return t.set("retryBtnText", value)
}

// Static sets whether the component is displayed statically
func (t tasks) Static(value bool) tasks {
	return t.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (t tasks) StaticClassName(value string) tasks {
	return t.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (t tasks) StaticInputClassName(value string) tasks {
	return t.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (t tasks) StaticLabelClassName(value string) tasks {
	return t.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the component is displayed statically
func (t tasks) StaticOn(value string) tasks {
	return t.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (t tasks) StaticPlaceholder(value string) tasks {
	return t.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (t tasks) StaticSchema(value string) tasks {
	return t.set("staticSchema", value)
}

// StatusLabel sets the status column label
func (t tasks) StatusLabel(value string) tasks {
	return t.set("statusLabel", value)
}

// StatusLabelMap sets the status label map
func (t tasks) StatusLabelMap(value string) tasks {
	return t.set("statusLabelMap", value)
}

// StatusTextMap sets the status text map
func (t tasks) StatusTextMap(value string) tasks {
	return t.set("statusTextMap", value)
}

// Style sets the component style
func (t tasks) Style(value any) tasks {
	return t.set("style", value)
}

// SubmitApi sets the API for submitting tasks
func (t tasks) SubmitApi(value string) tasks {
	return t.set("submitApi", value)
}

// TableClassName sets the CSS class name for the table
func (t tasks) TableClassName(value string) tasks {
	return t.set("tableClassName", value)
}

// TaskNameLabel sets the task name column label
func (t tasks) TaskNameLabel(value string) tasks {
	return t.set("taskNameLabel", value)
}

// TestIdBuilder sets the test ID builder
func (t tasks) TestIdBuilder(value string) tasks {
	return t.set("testIdBuilder", value)
}

// Testid sets the test ID
func (t tasks) Testid(value string) tasks {
	return t.set("testid", value)
}

// UseMobileUI sets whether to use mobile UI
func (t tasks) UseMobileUI(value bool) tasks {
	return t.set("useMobileUI", value)
}

// Visible sets whether the component is visible
func (t tasks) Visible(value bool) tasks {
	return t.set("visible", value)
}

// VisibleOn sets the expression to determine if the component is visible
func (t tasks) VisibleOn(value string) tasks {
	return t.set("visibleOn", value)
}
