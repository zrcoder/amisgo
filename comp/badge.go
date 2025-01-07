package comp

// Badge represents a badge renderer
type badge Schema

// Badge creates a new Badge instance
func Badge() badge {
	return make(badge)
}

func (b badge) set(key string, value any) badge {
	b[key] = value
	return b
}

// Animation enables or disables badge animation
func (b badge) Animation(value bool) badge {
	return b.set("animation", value)
}

// ClassName sets the CSS class name
func (b badge) ClassName(value string) badge {
	return b.set("className", value)
}

// Level sets the notification type
func (b badge) Level(value string) badge {
	return b.set("level", value)
}

// Mode sets the badge type
func (b badge) Mode(value string) badge {
	return b.set("mode", value)
}

// Offset adjusts the badge position relative to its default position
func (b badge) Offset(value string) badge {
	return b.set("offset", value)
}

// OverflowCount sets the maximum number before truncation
func (b badge) OverflowCount(value string) badge {
	return b.set("overflowCount", value)
}

// Position sets the badge position
func (b badge) Position(value string) badge {
	return b.set("position", value)
}

// Size sets the badge size
func (b badge) Size(value string) badge {
	return b.set("size", value)
}

// Style sets custom inline styles for the badge
func (b badge) Style(value any) badge {
	return b.set("style", value)
}

// Text sets the text content
func (b badge) Text(value string) badge {
	return b.set("text", value)
}

// VisibleOn sets a conditional expression for badge visibility
func (b badge) VisibleOn(value string) badge {
	return b.set("visibleOn", value)
}
