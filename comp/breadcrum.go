package comp

import "github.com/zrcoder/amisgo/model"

// Breadcrumb represents a breadcrumb navigation component
type breadcrumb model.Schema

// Breadcrumb creates a new Breadcrumb instance
func Breadcrumb() breadcrumb {
	return breadcrumb{"type": "breadcrumb"}
}

func (b breadcrumb) set(key string, value any) breadcrumb {
	b[key] = value
	return b
}

// ClassName sets the CSS class name for the outer container
func (b breadcrumb) ClassName(value string) breadcrumb {
	return b.set("className", value)
}

// DropdownClassName sets the CSS class name for the dropdown menu
func (b breadcrumb) DropdownClassName(value string) breadcrumb {
	return b.set("dropdownClassName", value)
}

// DropdownItemClassName sets the CSS class name for dropdown menu items
func (b breadcrumb) DropdownItemClassName(value string) breadcrumb {
	return b.set("dropdownItemClassName", value)
}

// ItemClassName sets the CSS class name for navigation items
func (b breadcrumb) ItemClassName(value string) breadcrumb {
	return b.set("itemClassName", value)
}

// Items configures the breadcrumb items
func (b breadcrumb) Items(value ...any) breadcrumb {
	return b.set("items", value)
}

// LabelMaxLength sets the maximum display length for labels
func (b breadcrumb) LabelMaxLength(value string) breadcrumb {
	return b.set("labelMaxLength", value)
}

// Separator sets the separator character between breadcrumb items
func (b breadcrumb) Separator(value string) breadcrumb {
	return b.set("separator", value)
}

// SeparatorClassName sets the CSS class name for the separator
func (b breadcrumb) SeparatorClassName(value string) breadcrumb {
	return b.set("separatorClassName", value)
}

// Source configures dynamic data source for breadcrumb items
func (b breadcrumb) Source(value string) breadcrumb {
	return b.set("source", value)
}

// TooltipPosition sets the position of tooltip popups
func (b breadcrumb) TooltipPosition(value string) breadcrumb {
	return b.set("tooltipPosition", value)
}
