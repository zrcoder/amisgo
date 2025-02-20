package comp

import "github.com/zrcoder/amisgo/schema"

// Calendar represents a Calendar component renderer
type Calendar schema.Schema

func NewCalendar() Calendar {
	return Calendar{"type": "calendar"}
}

func (c Calendar) set(key string, value any) Calendar {
	c[key] = value
	return c
}

// ClassName sets the CSS class name for the calendar
func (c Calendar) ClassName(value string) Calendar {
	return c.set("className", value)
}

// Disabled enables or disables the calendar component
func (c Calendar) Disabled(value bool) Calendar {
	return c.set("disabled", value)
}

// DisabledOn sets a conditional expression for disabling the calendar
func (c Calendar) DisabledOn(value string) Calendar {
	return c.set("disabledOn", value)
}

// EditorSetting configures editor-specific settings
func (c Calendar) EditorSetting(value string) Calendar {
	return c.set("editorSetting", value)
}

// Hidden controls the visibility of the calendar
func (c Calendar) Hidden(value bool) Calendar {
	return c.set("hidden", value)
}

// HiddenOn sets a conditional expression for hiding the calendar
func (c Calendar) HiddenOn(value string) Calendar {
	return c.set("hiddenOn", value)
}

// ID sets a unique identifier for the calendar component
func (c Calendar) ID(value string) Calendar {
	return c.set("id", value)
}

// LargeMode enables or disables large display mode
func (c Calendar) LargeMode(value bool) Calendar {
	return c.set("largeMode", value)
}

// OnEvent configures event-driven actions
func (c Calendar) OnEvent(value Event) Calendar {
	return c.set("onEvent", value)
}

// ScheduleAction sets the action for schedule-related interactions
func (c Calendar) ScheduleAction(value string) Calendar {
	return c.set("scheduleAction", value)
}

// ScheduleClassNames sets CSS class names for schedule elements
func (c Calendar) ScheduleClassNames(value string) Calendar {
	return c.set("scheduleClassNames", value)
}

// Schedules configures the calendar's schedule data
func (c Calendar) Schedules(value string) Calendar {
	return c.set("schedules", value)
}

// Static determines if the calendar is statically displayed
func (c Calendar) Static(value bool) Calendar {
	return c.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (c Calendar) StaticClassName(value string) Calendar {
	return c.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (c Calendar) StaticInputClassName(value string) Calendar {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (c Calendar) StaticLabelClassName(value string) Calendar {
	return c.set("staticLabelClassName", value)
}

// StaticOn sets a conditional expression for static display
func (c Calendar) StaticOn(value string) Calendar {
	return c.set("staticOn", value)
}

// StaticPlaceholder defines a placeholder for empty static values
func (c Calendar) StaticPlaceholder(value string) Calendar {
	return c.set("staticPlaceholder", value)
}

// StaticSchema sets the schema.Schema for static display
func (c Calendar) StaticSchema(value string) Calendar {
	return c.set("staticSchema", value)
}

// Style sets custom inline styles
func (c Calendar) Style(value any) Calendar {
	return c.set("style", value)
}

// TestIdBuilder configures test ID generation
func (c Calendar) TestIdBuilder(value string) Calendar {
	return c.set("testIdBuilder", value)
}

// TestID sets a specific test identifier
func (c Calendar) TestID(value string) Calendar {
	return c.set("testid", value)
}

// TodayActiveStyle sets the style for the current day's active state
func (c Calendar) TodayActiveStyle(value any) Calendar {
	return c.set("todayActiveStyle", value)
}

// UseMobileUI enables or disables mobile UI styling
func (c Calendar) UseMobileUI(value bool) Calendar {
	return c.set("useMobileUI", value)
}

// Visible controls the overall visibility of the calendar
func (c Calendar) Visible(value bool) Calendar {
	return c.set("visible", value)
}

// VisibleOn sets a conditional expression for calendar visibility
func (c Calendar) VisibleOn(value string) Calendar {
	return c.set("visibleOn", value)
}
