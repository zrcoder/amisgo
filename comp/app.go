package comp

import "github.com/zrcoder/amisgo/model"

type app model.Schema

// App creates an application for implementing multi-page interfaces
// Suitable for full-screen mode. Not recommended for partial rendering.
func App() app {
	return app{}.set("type", "app")
}

func (a app) set(key string, value any) app {
	a[key] = value
	return a
}

// Api configures the page configuration interface
// Use this to remotely fetch page configurations
// Returns configuration path json>data>pages
// Refer to the pages attribute for specific format
func (a app) Api(value any) app {
	return a.set("api", value)
}

// BrandName sets the application name
func (a app) BrandName(value string) app {
	return a.set("brandName", value)
}

// Logo sets the application logo
// Supports image URLs or SVG content
func (a app) Logo(value string) app {
	return a.set("logo", value)
}

// ClassName sets the CSS class name
func (a app) ClassName(value string) app {
	return a.set("className", value)
}

// Header configures the top area of the application
func (a app) Header(value ...any) app {
	return a.set("header", value)
}

// AsideBefore sets the area before the page menu
func (a app) AsideBefore(value ...any) app {
	return a.set("asideBefore", value)
}

// AsideAfter sets the area after the page menu
func (a app) AsideAfter(value ...any) app {
	return a.set("asideAfter", value)
}

// Footer configures the page footer
func (a app) Footer(value ...any) app {
	return a.set("footer", value)
}

// Pages configures specific page settings
// Typically an array where the first layer represents groups
// Usually only requires configuring labels
// If no grouping is needed, leave unconfigured
// Actual pages should be configured from the second layer (in the children of the first layer)
func (a app) Pages(value ...MPageItem) app {
	return a.set("pages", value)
}
