package comp

import "github.com/zrcoder/amisgo/schema"

// Badge represents a Badge renderer
type Badge schema.Schema

func NewBadge() Badge {
	return make(Badge)
}

func (b Badge) set(key string, value any) Badge {
	b[key] = value
	return b
}

// Animation enables or disables badge animation
func (b Badge) Animation(value bool) Badge {
	return b.set("animation", value)
}

// ClassName sets the CSS class name
func (b Badge) ClassName(value string) Badge {
	return b.set("className", value)
}

// Level sets the notification type
func (b Badge) Level(value string) Badge {
	return b.set("level", value)
}

// Mode sets the badge type
func (b Badge) Mode(value string) Badge {
	return b.set("mode", value)
}

// Offset adjusts the badge position relative to its default position
func (b Badge) Offset(value string) Badge {
	return b.set("offset", value)
}

// OverflowCount sets the maximum number before truncation
func (b Badge) OverflowCount(value string) Badge {
	return b.set("overflowCount", value)
}

// Position sets the badge position
func (b Badge) Position(value string) Badge {
	return b.set("position", value)
}

// Size sets the badge size
func (b Badge) Size(value string) Badge {
	return b.set("size", value)
}

// Style sets custom inline styles for the badge
func (b Badge) Style(value any) Badge {
	return b.set("style", value)
}

// Text sets the text content
func (b Badge) Text(value string) Badge {
	return b.set("text", value)
}

// VisibleOn sets a conditional expression for badge visibility
func (b Badge) VisibleOn(value string) Badge {
	return b.set("visibleOn", value)
}
