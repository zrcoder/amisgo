package comp

import (
	"net/http"

	"github.com/zrcoder/amisgo/internal/servemux"
	"github.com/zrcoder/amisgo/schema"
)

// Service represents a Service configuration.
type Service schema.Schema

func NewService(mux *http.ServeMux) Service {
	return Service{"type": "service", servemux.Key: mux}
}

func (s Service) set(key string, value any) Service {
	s[key] = value
	return s
}

// Api sets the API for data fetching.
func (s Service) Api(value string) Service {
	return s.set("api", value)
}

// GetData fetches data using the internal API.
func (s Service) GetData(getter func() (any, error)) Service {
	return s.Api(servemux.ServeData(s.mux(), getter))
}

// Body sets the content area.
func (s Service) Body(value ...any) Service {
	return s.set("body", value)
}

// Data sets the data.
func (s Service) Data(value any) Service {
	return s.set("data", value)
}

// ClassName sets the container CSS class name.
func (s Service) ClassName(value string) Service {
	return s.set("className", value)
}

// DataProvider sets the data provider.
func (s Service) DataProvider(value string) Service {
	return s.set("dataProvider", value)
}

// Disabled sets whether the service is disabled.
func (s Service) Disabled(value bool) Service {
	return s.set("disabled", value)
}

// DisabledOn sets the expression to disable the service.
func (s Service) DisabledOn(value string) Service {
	return s.set("disabledOn", value)
}

// EditorSetting sets the editor configuration.
func (s Service) EditorSetting(value string) Service {
	return s.set("editorSetting", value)
}

// FetchOn sets the fetch condition.
func (s Service) FetchOn(value string) Service {
	return s.set("fetchOn", value)
}

// Hidden sets whether the service is hidden.
func (s Service) Hidden(value bool) Service {
	return s.set("hidden", value)
}

// HiddenOn sets the expression to hide the service.
func (s Service) HiddenOn(value string) Service {
	return s.set("hiddenOn", value)
}

// ID sets the unique component ID.
func (s Service) ID(value string) Service {
	return s.set("id", value)
}

// InitFetch sets whether to fetch initially.
func (s Service) InitFetch(value bool) Service {
	return s.set("initFetch", value)
}

// InitFetchOn sets the expression to determine initial fetch.
func (s Service) InitFetchOn(value string) Service {
	return s.set("initFetchOn", value)
}

// InitFetchSchema sets whether to fetch schema.Schema initially.
func (s Service) InitFetchSchema(value bool) Service {
	return s.set("initFetchSchema", value)
}

// InitFetchSchemaOn sets the expression to determine initial schema.Schema fetch.
func (s Service) InitFetchSchemaOn(value string) Service {
	return s.set("initFetchSchemaOn", value)
}

// Interval sets the interval for polling.
func (s Service) Interval(value string) Service {
	return s.set("interval", value)
}

// LoadingConfig sets the loading configuration.
func (s Service) LoadingConfig(value string) Service {
	return s.set("loadingConfig", value)
}

// Messages sets the message configurations.
func (s Service) Messages(value string) Service {
	return s.set("messages", value)
}

// Name sets the component name.
func (s Service) Name(value string) Service {
	return s.set("name", value)
}

// OnEvent sets the event action configuration.
func (s Service) OnEvent(value Event) Service {
	return s.set("onEvent", value)
}

// SchemaApi sets the API for schema.Schema fetching.
func (s Service) SchemaApi(value string) Service {
	return s.set("schemaApi", value)
}

// ShowErrorMsg sets whether to show error messages.
func (s Service) ShowErrorMsg(value bool) Service {
	return s.set("showErrorMsg", value)
}

// SilentPolling sets whether to poll silently.
func (s Service) SilentPolling(value bool) Service {
	return s.set("silentPolling", value)
}

// Static sets whether to display statically.
func (s Service) Static(value bool) Service {
	return s.set("static", value)
}

// StaticClassName sets the static display class name.
func (s Service) StaticClassName(value string) Service {
	return s.set("staticClassName", value)
}

// StaticInputClassName sets the static input value class name.
func (s Service) StaticInputClassName(value string) Service {
	return s.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static label class name.
func (s Service) StaticLabelClassName(value string) Service {
	return s.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display.
func (s Service) StaticOn(value string) Service {
	return s.set("staticOn", value)
}

// StaticPlaceholder sets the static placeholder.
func (s Service) StaticPlaceholder(value string) Service {
	return s.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema.Schema.
func (s Service) StaticSchema(value string) Service {
	return s.set("staticSchema", value)
}

// StopAutoRefreshWhen sets the condition to stop auto refresh.
func (s Service) StopAutoRefreshWhen(value string) Service {
	return s.set("stopAutoRefreshWhen", value)
}

// Style sets the component style.
func (s Service) Style(value any) Service {
	return s.set("style", value)
}

// TestIdBuilder sets the test ID builder.
func (s Service) TestIdBuilder(value string) Service {
	return s.set("testIdBuilder", value)
}

// Testid sets the test ID.
func (s Service) Testid(value string) Service {
	return s.set("testid", value)
}

// UseMobileUI sets whether to use mobile UI.
func (s Service) UseMobileUI(value bool) Service {
	return s.set("useMobileUI", value)
}

// Visible sets whether the service is visible.
func (s Service) Visible(value bool) Service {
	return s.set("visible", value)
}

// VisibleOn sets the expression to determine visibility.
func (s Service) VisibleOn(value string) Service {
	return s.set("visibleOn", value)
}

// Ws sets the WebSocket URL.
func (s Service) Ws(value string) Service {
	return s.set("ws", value)
}

func (s Service) mux() *http.ServeMux {
	return s[servemux.Key].(*http.ServeMux)
}
