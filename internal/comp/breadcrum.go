package comp

import "github.com/zrcoder/amisgo/schema"

// Breadcrumb represents a Breadcrumb navigation component
type Breadcrumb schema.Schema

func NewBreadcrumb() Breadcrumb {
	return Breadcrumb{"type": "breadcrumb"}
}

func (b Breadcrumb) set(key string, value any) Breadcrumb {
	b[key] = value
	return b
}

// ClassName sets the CSS class name for the outer container
func (b Breadcrumb) ClassName(value string) Breadcrumb {
	return b.set("className", value)
}

// DropdownClassName sets the CSS class name for the dropdown menu
func (b Breadcrumb) DropdownClassName(value string) Breadcrumb {
	return b.set("dropdownClassName", value)
}

// DropdownItemClassName sets the CSS class name for dropdown menu items
func (b Breadcrumb) DropdownItemClassName(value string) Breadcrumb {
	return b.set("dropdownItemClassName", value)
}

// ItemClassName sets the CSS class name for navigation items
func (b Breadcrumb) ItemClassName(value string) Breadcrumb {
	return b.set("itemClassName", value)
}

// Items configures the breadcrumb items
func (b Breadcrumb) Items(value ...BreadcrumbItem) Breadcrumb {
	return b.set("items", value)
}

// LabelMaxLength sets the maximum display length for labels
func (b Breadcrumb) LabelMaxLength(value string) Breadcrumb {
	return b.set("labelMaxLength", value)
}

// Separator sets the separator character between breadcrumb items
func (b Breadcrumb) Separator(value string) Breadcrumb {
	return b.set("separator", value)
}

// SeparatorClassName sets the CSS class name for the separator
func (b Breadcrumb) SeparatorClassName(value string) Breadcrumb {
	return b.set("separatorClassName", value)
}

// Source configures dynamic data source for breadcrumb items
func (b Breadcrumb) Source(value string) Breadcrumb {
	return b.set("source", value)
}

// TooltipPosition sets the position of tooltip popups
func (b Breadcrumb) TooltipPosition(value string) Breadcrumb {
	return b.set("tooltipPosition", value)
}
