package comp

import (
	"github.com/zrcoder/amisgo/internal/servermux"
	"github.com/zrcoder/amisgo/model"
)

// service represents a service configuration.
type service model.Schema

// Service creates a new Service instance.
func Service() service {
	return service{"type": "service"}
}

func (s service) set(key string, value any) service {
	s[key] = value
	return s
}

// Api sets the API for data fetching.
func (s service) Api(value string) service {
	return s.set("api", value)
}

// GetData fetches data using the internal API.
func (s service) GetData(getter func() (any, error)) service {
	return s.Api(servermux.ServeData(getter))
}

// Body sets the content area.
func (s service) Body(value ...any) service {
	return s.set("body", value)
}

// Data sets the data.
func (s service) Data(value any) service {
	return s.set("data", value)
}

// ClassName sets the container CSS class name.
func (s service) ClassName(value string) service {
	return s.set("className", value)
}

// DataProvider sets the data provider.
func (s service) DataProvider(value string) service {
	return s.set("dataProvider", value)
}

// Disabled sets whether the service is disabled.
func (s service) Disabled(value bool) service {
	return s.set("disabled", value)
}

// DisabledOn sets the expression to disable the service.
func (s service) DisabledOn(value string) service {
	return s.set("disabledOn", value)
}

// EditorSetting sets the editor configuration.
func (s service) EditorSetting(value string) service {
	return s.set("editorSetting", value)
}

// FetchOn sets the fetch condition.
func (s service) FetchOn(value string) service {
	return s.set("fetchOn", value)
}

// Hidden sets whether the service is hidden.
func (s service) Hidden(value bool) service {
	return s.set("hidden", value)
}

// HiddenOn sets the expression to hide the service.
func (s service) HiddenOn(value string) service {
	return s.set("hiddenOn", value)
}

// ID sets the unique component ID.
func (s service) ID(value string) service {
	return s.set("id", value)
}

// InitFetch sets whether to fetch initially.
func (s service) InitFetch(value bool) service {
	return s.set("initFetch", value)
}

// InitFetchOn sets the expression to determine initial fetch.
func (s service) InitFetchOn(value string) service {
	return s.set("initFetchOn", value)
}

// InitFetchSchema sets whether to fetch schema initially.
func (s service) InitFetchSchema(value bool) service {
	return s.set("initFetchSchema", value)
}

// InitFetchSchemaOn sets the expression to determine initial schema fetch.
func (s service) InitFetchSchemaOn(value string) service {
	return s.set("initFetchSchemaOn", value)
}

// Interval sets the interval for polling.
func (s service) Interval(value string) service {
	return s.set("interval", value)
}

// LoadingConfig sets the loading configuration.
func (s service) LoadingConfig(value string) service {
	return s.set("loadingConfig", value)
}

// Messages sets the message configurations.
func (s service) Messages(value string) service {
	return s.set("messages", value)
}

// Name sets the component name.
func (s service) Name(value string) service {
	return s.set("name", value)
}

// OnEvent sets the event action configuration.
func (s service) OnEvent(value any) service {
	return s.set("onEvent", value)
}

// model.SchemaApi sets the API for schema fetching.
func (s service) SchemaApi(value string) service {
	return s.set("schemaApi", value)
}

// ShowErrorMsg sets whether to show error messages.
func (s service) ShowErrorMsg(value bool) service {
	return s.set("showErrorMsg", value)
}

// SilentPolling sets whether to poll silently.
func (s service) SilentPolling(value bool) service {
	return s.set("silentPolling", value)
}

// Static sets whether to display statically.
func (s service) Static(value bool) service {
	return s.set("static", value)
}

// StaticClassName sets the static display class name.
func (s service) StaticClassName(value string) service {
	return s.set("staticClassName", value)
}

// StaticInputClassName sets the static input value class name.
func (s service) StaticInputClassName(value string) service {
	return s.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static label class name.
func (s service) StaticLabelClassName(value string) service {
	return s.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display.
func (s service) StaticOn(value string) service {
	return s.set("staticOn", value)
}

// StaticPlaceholder sets the static placeholder.
func (s service) StaticPlaceholder(value string) service {
	return s.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema.
func (s service) StaticSchema(value string) service {
	return s.set("staticSchema", value)
}

// StopAutoRefreshWhen sets the condition to stop auto refresh.
func (s service) StopAutoRefreshWhen(value string) service {
	return s.set("stopAutoRefreshWhen", value)
}

// Style sets the component style.
func (s service) Style(value any) service {
	return s.set("style", value)
}

// TestIdBuilder sets the test ID builder.
func (s service) TestIdBuilder(value string) service {
	return s.set("testIdBuilder", value)
}

// Testid sets the test ID.
func (s service) Testid(value string) service {
	return s.set("testid", value)
}

// UseMobileUI sets whether to use mobile UI.
func (s service) UseMobileUI(value bool) service {
	return s.set("useMobileUI", value)
}

// Visible sets whether the service is visible.
func (s service) Visible(value bool) service {
	return s.set("visible", value)
}

// VisibleOn sets the expression to determine visibility.
func (s service) VisibleOn(value string) service {
	return s.set("visibleOn", value)
}

// Ws sets the WebSocket URL.
func (s service) Ws(value string) service {
	return s.set("ws", value)
}
