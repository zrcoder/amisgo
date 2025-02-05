package comp

import (
	"net/http"

	"github.com/zrcoder/amisgo/internal/servemux"
	"github.com/zrcoder/amisgo/model"
)

// Page represents an AMIS Page renderer
type Page model.Schema

func NewPage(mux *http.ServeMux) Page {
	return Page{"type": "page", servemux.Key: mux}
}

// set sets a field value
func (p Page) set(key string, value any) Page {
	p[key] = value
	return p
}

// Aside sets the aside area
func (p Page) Aside(value ...any) Page {
	return p.set("aside", value)
}

// AsideClassName sets the CSS class name for the aside area
func (p Page) AsideClassName(value string) Page {
	return p.set("asideClassName", value)
}

// AsideMaxWidth sets the max width for the aside area
func (p Page) AsideMaxWidth(value int) Page {
	return p.set("asideMaxWidth", value)
}

// AsideMinWidth sets the min width for the aside area
func (p Page) AsideMinWidth(value int) Page {
	return p.set("asideMinWidth", value)
}

// AsideResizor sets whether the aside area is resizable
func (p Page) AsideResizor(value bool) Page {
	return p.set("asideResizor", value)
}

// AsideSticky sets whether the aside content is sticky
func (p Page) AsideSticky(value bool) Page {
	return p.set("asideSticky", value)
}

// Body sets the body content
func (p Page) Body(value ...any) Page {
	return p.set("body", value)
}

// BodyClassName sets the CSS class name for the body area
func (p Page) BodyClassName(value string) Page {
	return p.set("bodyClassName", value)
}

// ClassName sets the container class name
func (p Page) ClassName(value string) Page {
	return p.set("className", value)
}

// CSS sets custom page-level styles
func (p Page) CSS(value any) Page {
	return p.set("css", value)
}

// CSSVars sets CSS variables
func (p Page) CSSVars(value any) Page {
	return p.set("cssVars", value)
}

// Data sets initial page-level data
func (p Page) Data(value any) Page {
	return p.set("data", value)
}

// Definitions sets configuration definitions
func (p Page) Definitions(value string) Page {
	return p.set("definitions", value)
}

// Disabled sets whether the page is disabled
func (p Page) Disabled(value bool) Page {
	return p.set("disabled", value)
}

// DisabledOn sets the expression for disabling the page
func (p Page) DisabledOn(value string) Page {
	return p.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (p Page) EditorSetting(value string) Page {
	return p.set("editorSetting", value)
}

// HeaderClassName sets the CSS class name for the header container
func (p Page) HeaderClassName(value string) Page {
	return p.set("headerClassName", value)
}

// Hidden sets whether the page is hidden
func (p Page) Hidden(value bool) Page {
	return p.set("hidden", value)
}

// HiddenOn sets the expression for hiding the page
func (p Page) HiddenOn(value string) Page {
	return p.set("hiddenOn", value)
}

// ID sets the unique component ID
func (p Page) ID(value string) Page {
	return p.set("id", value)
}

// InitApi sets the initialization API
func (p Page) InitApi(value string) Page {
	return p.set("initApi", value)
}

// InitData sets the initialization data
func (p Page) InitData(getter func() (any, error)) Page {
	return p.InitApi(servemux.ServeData(p.mux(), getter))
}

// InitFetch sets whether to fetch data by default
func (p Page) InitFetch(value bool) Page {
	return p.set("initFetch", value)
}

// InitFetchOn sets the expression for fetching data by default
func (p Page) InitFetchOn(value string) Page {
	return p.set("initFetchOn", value)
}

// Interval sets the refresh interval (ms, minimum 1000)
func (p Page) Interval(value int) Page {
	return p.set("interval", value)
}

// LoadingConfig sets the loading configuration
func (p Page) LoadingConfig(value string) Page {
	return p.set("loadingConfig", value)
}

// Messages sets the message configuration
func (p Page) Messages(value string) Page {
	return p.set("messages", value)
}

// MobileCSS sets the styles for mobile devices
func (p Page) MobileCSS(value any) Page {
	return p.set("mobileCSS", value)
}

// Name sets the component name
func (p Page) Name(value string) Page {
	return p.set("name", value)
}

// OnEvent sets the event action configuration
func (p Page) OnEvent(value any) Page {
	return p.set("onEvent", value)
}

// PullRefresh sets the pull-to-refresh configuration
func (p Page) PullRefresh(value model.PullRefresh) Page {
	return p.set("pullRefresh", value)
}

// Regions sets the regions to display
func (p Page) Regions(value string) Page {
	return p.set("regions", value)
}

// Remark sets the page description (string or remark)
func (p Page) Remark(value any) Page {
	return p.set("remark", value)
}

// ShowErrorMsg sets whether to show error messages
func (p Page) ShowErrorMsg(value bool) Page {
	return p.set("showErrorMsg", value)
}

// SilentPolling sets whether to load data silently
func (p Page) SilentPolling(value bool) Page {
	return p.set("silentPolling", value)
}

// Static sets whether to display statically
func (p Page) Static(value bool) Page {
	return p.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (p Page) StaticClassName(value string) Page {
	return p.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (p Page) StaticInputClassName(value string) Page {
	return p.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (p Page) StaticLabelClassName(value string) Page {
	return p.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display
func (p Page) StaticOn(value string) Page {
	return p.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (p Page) StaticPlaceholder(value string) Page {
	return p.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (p Page) StaticSchema(value string) Page {
	return p.set("staticSchema", value)
}

// StopAutoRefreshWhen sets the condition to stop auto-refresh
func (p Page) StopAutoRefreshWhen(value string) Page {
	return p.set("stopAutoRefreshWhen", value)
}

// Style sets custom styles
func (p Page) Style(value any) Page {
	return p.set("style", value)
}

// SubTitle sets the page subtitle
func (p Page) SubTitle(value any) Page {
	return p.set("subTitle", value)
}

// TestIdBuilder sets the custom test ID builder
func (p Page) TestIdBuilder(value string) Page {
	return p.set("testIdBuilder", value)
}

// Testid sets the test ID
func (p Page) Testid(value string) Page {
	return p.set("testid", value)
}

// Title sets the page title
func (p Page) Title(value any) Page {
	return p.set("title", value)
}

// Toolbar sets the toolbar content
func (p Page) Toolbar(value ...any) Page {
	return p.set("toolbar", value)
}

// ToolbarClassName sets the CSS class name for the toolbar container
func (p Page) ToolbarClassName(value string) Page {
	return p.set("toolbarClassName", value)
}

// UseMobileUI sets whether to use mobile UI styles
func (p Page) UseMobileUI(value bool) Page {
	return p.set("useMobileUI", value)
}

// Visible sets whether the page is visible
func (p Page) Visible(value bool) Page {
	return p.set("visible", value)
}

// VisibleOn sets the expression for visibility
func (p Page) VisibleOn(value string) Page {
	return p.set("visibleOn", value)
}

func (p Page) mux() *http.ServeMux {
	return p[servemux.Key].(*http.ServeMux)
}
