package comp

import "github.com/zrcoder/amisgo/schema"

// NavOverflow represents the navigation overflow configuration.
type NavOverflow schema.Schema

func NewNavOverflow() NavOverflow {
	return NavOverflow{}
}

// set sets a property and returns the instance.
func (n NavOverflow) set(key string, value any) NavOverflow {
	n[key] = value
	return n
}

// Enable sets whether responsive overflow is enabled.
func (n NavOverflow) Enable(value bool) NavOverflow {
	return n.set("enable", value)
}

// ItemWidth sets the width of navigation items.
func (n NavOverflow) ItemWidth(value string) NavOverflow {
	return n.set("itemWidth", value)
}

// MaxVisibleCount sets the maximum number of visible items.
func (n NavOverflow) MaxVisibleCount(value string) NavOverflow {
	return n.set("maxVisibleCount", value)
}

// OverflowClassName sets the CSS class name for the overflow button.
func (n NavOverflow) OverflowClassName(value string) NavOverflow {
	return n.set("overflowClassName", value)
}

// OverflowIndicator sets the icon for the overflow button.
func (n NavOverflow) OverflowIndicator(value string) NavOverflow {
	return n.set("overflowIndicator", value)
}

// OverflowLabel sets the text for the overflow button.
func (n NavOverflow) OverflowLabel(value string) NavOverflow {
	return n.set("overflowLabel", value)
}

// OverflowListClassName sets the CSS class name for the overflow list.
func (n NavOverflow) OverflowListClassName(value string) NavOverflow {
	return n.set("overflowListClassName", value)
}

// OverflowPopoverClassName sets the CSS class name for the popover.
func (n NavOverflow) OverflowPopoverClassName(value string) NavOverflow {
	return n.set("overflowPopoverClassName", value)
}

// OverflowSuffix sets the suffix node for the navigation list.
func (n NavOverflow) OverflowSuffix(value string) NavOverflow {
	return n.set("overflowSuffix", value)
}

// Style sets custom styles.
func (n NavOverflow) Style(value any) NavOverflow {
	return n.set("style", value)
}

// WrapperComponent sets the outer tag name for wrapping navigation.
func (n NavOverflow) WrapperComponent(value string) NavOverflow {
	return n.set("wrapperComponent", value)
}
