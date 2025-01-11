package comp

import (
	"github.com/zrcoder/amisgo/internal/servermux"
	"github.com/zrcoder/amisgo/model"
)

// page represents an AMIS page renderer

type page model.Schema

// Page creates a new Page instance
func Page() page {
	return page{"type": "page"}
}

// set sets a field value
func (p page) set(key string, value any) page {
	p[key] = value
	return p
}

// Aside sets the aside area
func (p page) Aside(value ...any) page {
	return p.set("aside", value)
}

// AsideClassName sets the CSS class name for the aside area
func (p page) AsideClassName(value string) page {
	return p.set("asideClassName", value)
}

// AsideMaxWidth sets the max width for the aside area
func (p page) AsideMaxWidth(value int) page {
	return p.set("asideMaxWidth", value)
}

// AsideMinWidth sets the min width for the aside area
func (p page) AsideMinWidth(value int) page {
	return p.set("asideMinWidth", value)
}

// AsideResizor sets whether the aside area is resizable
func (p page) AsideResizor(value bool) page {
	return p.set("asideResizor", value)
}

// AsideSticky sets whether the aside content is sticky
func (p page) AsideSticky(value bool) page {
	return p.set("asideSticky", value)
}

// Body sets the body content
func (p page) Body(value ...any) page {
	return p.set("body", value)
}

// BodyClassName sets the CSS class name for the body area
func (p page) BodyClassName(value string) page {
	return p.set("bodyClassName", value)
}

// ClassName sets the container class name
func (p page) ClassName(value string) page {
	return p.set("className", value)
}

// CSS sets custom page-level styles
func (p page) CSS(value any) page {
	return p.set("css", value)
}

// CSSVars sets CSS variables
func (p page) CSSVars(value any) page {
	return p.set("cssVars", value)
}

// Data sets initial page-level data
func (p page) Data(value any) page {
	return p.set("data", value)
}

// Definitions sets configuration definitions
func (p page) Definitions(value string) page {
	return p.set("definitions", value)
}

// Disabled sets whether the page is disabled
func (p page) Disabled(value bool) page {
	return p.set("disabled", value)
}

// DisabledOn sets the expression for disabling the page
func (p page) DisabledOn(value string) page {
	return p.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (p page) EditorSetting(value string) page {
	return p.set("editorSetting", value)
}

// HeaderClassName sets the CSS class name for the header container
func (p page) HeaderClassName(value string) page {
	return p.set("headerClassName", value)
}

// Hidden sets whether the page is hidden
func (p page) Hidden(value bool) page {
	return p.set("hidden", value)
}

// HiddenOn sets the expression for hiding the page
func (p page) HiddenOn(value string) page {
	return p.set("hiddenOn", value)
}

// ID sets the unique component ID
func (p page) ID(value string) page {
	return p.set("id", value)
}

// InitApi sets the initialization API
func (p page) InitApi(value string) page {
	return p.set("initApi", value)
}

// InitData sets the initialization data
func (p page) InitData(getter func() (any, error)) page {
	return p.InitApi(servermux.ServeData(getter))
}

// InitFetch sets whether to fetch data by default
func (p page) InitFetch(value bool) page {
	return p.set("initFetch", value)
}

// InitFetchOn sets the expression for fetching data by default
func (p page) InitFetchOn(value string) page {
	return p.set("initFetchOn", value)
}

// Interval sets the refresh interval (ms, minimum 1000)
func (p page) Interval(value int) page {
	return p.set("interval", value)
}

// LoadingConfig sets the loading configuration
func (p page) LoadingConfig(value string) page {
	return p.set("loadingConfig", value)
}

// Messages sets the message configuration
func (p page) Messages(value string) page {
	return p.set("messages", value)
}

// MobileCSS sets the styles for mobile devices
func (p page) MobileCSS(value any) page {
	return p.set("mobileCSS", value)
}

// Name sets the component name
func (p page) Name(value string) page {
	return p.set("name", value)
}

// OnEvent sets the event action configuration
func (p page) OnEvent(value any) page {
	return p.set("onEvent", value)
}

// PullRefresh sets the pull-to-refresh configuration
func (p page) PullRefresh(value MPullRefresh) page {
	return p.set("pullRefresh", value)
}

// Regions sets the regions to display
func (p page) Regions(value string) page {
	return p.set("regions", value)
}

// Remark sets the page description (string or remark)
func (p page) Remark(value any) page {
	return p.set("remark", value)
}

// ShowErrorMsg sets whether to show error messages
func (p page) ShowErrorMsg(value bool) page {
	return p.set("showErrorMsg", value)
}

// SilentPolling sets whether to load data silently
func (p page) SilentPolling(value bool) page {
	return p.set("silentPolling", value)
}

// Static sets whether to display statically
func (p page) Static(value bool) page {
	return p.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (p page) StaticClassName(value string) page {
	return p.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (p page) StaticInputClassName(value string) page {
	return p.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (p page) StaticLabelClassName(value string) page {
	return p.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display
func (p page) StaticOn(value string) page {
	return p.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (p page) StaticPlaceholder(value string) page {
	return p.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (p page) StaticSchema(value string) page {
	return p.set("staticSchema", value)
}

// StopAutoRefreshWhen sets the condition to stop auto-refresh
func (p page) StopAutoRefreshWhen(value string) page {
	return p.set("stopAutoRefreshWhen", value)
}

// Style sets custom styles
func (p page) Style(value any) page {
	return p.set("style", value)
}

// SubTitle sets the page subtitle
func (p page) SubTitle(value any) page {
	return p.set("subTitle", value)
}

// TestIdBuilder sets the custom test ID builder
func (p page) TestIdBuilder(value string) page {
	return p.set("testIdBuilder", value)
}

// Testid sets the test ID
func (p page) Testid(value string) page {
	return p.set("testid", value)
}

// Title sets the page title
func (p page) Title(value any) page {
	return p.set("title", value)
}

// Toolbar sets the toolbar content
func (p page) Toolbar(value ...any) page {
	return p.set("toolbar", value)
}

// ToolbarClassName sets the CSS class name for the toolbar container
func (p page) ToolbarClassName(value string) page {
	return p.set("toolbarClassName", value)
}

// UseMobileUI sets whether to use mobile UI styles
func (p page) UseMobileUI(value bool) page {
	return p.set("useMobileUI", value)
}

// Visible sets whether the page is visible
func (p page) Visible(value bool) page {
	return p.set("visible", value)
}

// VisibleOn sets the expression for visibility
func (p page) VisibleOn(value string) page {
	return p.set("visibleOn", value)
}
