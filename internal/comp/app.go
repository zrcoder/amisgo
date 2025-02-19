package comp

import (
	"github.com/zrcoder/amisgo/schema"
)

type App schema.Schema

// NewApp creates an application for implementing multi-page interfaces
// Suitable for full-screen mode. Not recommended for partial rendering.
func NewApp() App {
	return App{"type": "app"}
}

func (a App) set(key string, value any) App {
	a[key] = value
	return a
}

// Api configures the page configuration interface
// Use this to remotely fetch page configurations
// Returns configuration path json>data>pages
// Refer to the pages attribute for specific format
func (a App) Api(value any) App {
	return a.set("api", value)
}

// BrandName sets the application name
func (a App) BrandName(value string) App {
	return a.set("brandName", value)
}

// Logo sets the application logo
// Supports image URLs or SVG content
func (a App) Logo(value string) App {
	return a.set("logo", value)
}

// ClassName sets the CSS class name
func (a App) ClassName(value string) App {
	return a.set("className", value)
}

// Header configures the top area of the application
func (a App) Header(value ...any) App {
	return a.set("header", value)
}

// AsideBefore sets the area before the page menu
func (a App) AsideBefore(value ...any) App {
	return a.set("asideBefore", value)
}

// AsideAfter sets the area after the page menu
func (a App) AsideAfter(value ...any) App {
	return a.set("asideAfter", value)
}

// Footer configures the page footer
func (a App) Footer(value ...any) App {
	return a.set("footer", value)
}

// Pages configures specific page settings
// Typically an array where the first layer represents groups
// Usually only requires configuring labels
// If no grouping is needed, leave unconfigured
// Actual pages should be configured from the second layer (in the children of the first layer)
func (a App) Pages(value ...PageItem) App {
	return a.set("pages", value)
}
