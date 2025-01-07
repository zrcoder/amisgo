package comp

// navOverflow represents the navigation overflow configuration.
type navOverflow Schema

// NavOverflow creates a new NavOverflow instance.
func NavOverflow() navOverflow {
	return navOverflow{}
}

// set sets a property and returns the instance.
func (n navOverflow) set(key string, value any) navOverflow {
	n[key] = value
	return n
}

// Enable sets whether responsive overflow is enabled.
func (n navOverflow) Enable(value bool) navOverflow {
	return n.set("enable", value)
}

// ItemWidth sets the width of navigation items.
func (n navOverflow) ItemWidth(value string) navOverflow {
	return n.set("itemWidth", value)
}

// MaxVisibleCount sets the maximum number of visible items.
func (n navOverflow) MaxVisibleCount(value string) navOverflow {
	return n.set("maxVisibleCount", value)
}

// OverflowClassName sets the CSS class name for the overflow button.
func (n navOverflow) OverflowClassName(value string) navOverflow {
	return n.set("overflowClassName", value)
}

// OverflowIndicator sets the icon for the overflow button.
func (n navOverflow) OverflowIndicator(value string) navOverflow {
	return n.set("overflowIndicator", value)
}

// OverflowLabel sets the text for the overflow button.
func (n navOverflow) OverflowLabel(value string) navOverflow {
	return n.set("overflowLabel", value)
}

// OverflowListClassName sets the CSS class name for the overflow list.
func (n navOverflow) OverflowListClassName(value string) navOverflow {
	return n.set("overflowListClassName", value)
}

// OverflowPopoverClassName sets the CSS class name for the popover.
func (n navOverflow) OverflowPopoverClassName(value string) navOverflow {
	return n.set("overflowPopoverClassName", value)
}

// OverflowSuffix sets the suffix node for the navigation list.
func (n navOverflow) OverflowSuffix(value string) navOverflow {
	return n.set("overflowSuffix", value)
}

// Style sets custom styles.
func (n navOverflow) Style(value any) navOverflow {
	return n.set("style", value)
}

// WrapperComponent sets the outer tag name for wrapping navigation.
func (n navOverflow) WrapperComponent(value string) navOverflow {
	return n.set("wrapperComponent", value)
}
