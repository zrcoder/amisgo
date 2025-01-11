package comp

import "github.com/zrcoder/amisgo/model"

// Calendar represents a calendar component renderer
type calendar model.Schema

// Calendar creates a new Calendar instance
func Calendar() calendar {
	return calendar{"type": "calendar"}
}

func (c calendar) set(key string, value any) calendar {
	c[key] = value
	return c
}

// ClassName sets the CSS class name for the calendar
func (c calendar) ClassName(value string) calendar {
	return c.set("className", value)
}

// Disabled enables or disables the calendar component
func (c calendar) Disabled(value bool) calendar {
	return c.set("disabled", value)
}

// DisabledOn sets a conditional expression for disabling the calendar
func (c calendar) DisabledOn(value string) calendar {
	return c.set("disabledOn", value)
}

// EditorSetting configures editor-specific settings
func (c calendar) EditorSetting(value string) calendar {
	return c.set("editorSetting", value)
}

// Hidden controls the visibility of the calendar
func (c calendar) Hidden(value bool) calendar {
	return c.set("hidden", value)
}

// HiddenOn sets a conditional expression for hiding the calendar
func (c calendar) HiddenOn(value string) calendar {
	return c.set("hiddenOn", value)
}

// ID sets a unique identifier for the calendar component
func (c calendar) ID(value string) calendar {
	return c.set("id", value)
}

// LargeMode enables or disables large display mode
func (c calendar) LargeMode(value bool) calendar {
	return c.set("largeMode", value)
}

// OnEvent configures event-driven actions
func (c calendar) OnEvent(value any) calendar {
	return c.set("onEvent", value)
}

// ScheduleAction sets the action for schedule-related interactions
func (c calendar) ScheduleAction(value string) calendar {
	return c.set("scheduleAction", value)
}

// ScheduleClassNames sets CSS class names for schedule elements
func (c calendar) ScheduleClassNames(value string) calendar {
	return c.set("scheduleClassNames", value)
}

// Schedules configures the calendar's schedule data
func (c calendar) Schedules(value string) calendar {
	return c.set("schedules", value)
}

// Static determines if the calendar is statically displayed
func (c calendar) Static(value bool) calendar {
	return c.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (c calendar) StaticClassName(value string) calendar {
	return c.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (c calendar) StaticInputClassName(value string) calendar {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (c calendar) StaticLabelClassName(value string) calendar {
	return c.set("staticLabelClassName", value)
}

// StaticOn sets a conditional expression for static display
func (c calendar) StaticOn(value string) calendar {
	return c.set("staticOn", value)
}

// StaticPlaceholder defines a placeholder for empty static values
func (c calendar) StaticPlaceholder(value string) calendar {
	return c.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (c calendar) StaticSchema(value string) calendar {
	return c.set("staticSchema", value)
}

// Style sets custom inline styles
func (c calendar) Style(value any) calendar {
	return c.set("style", value)
}

// TestIdBuilder configures test ID generation
func (c calendar) TestIdBuilder(value string) calendar {
	return c.set("testIdBuilder", value)
}

// TestID sets a specific test identifier
func (c calendar) TestID(value string) calendar {
	return c.set("testid", value)
}

// TodayActiveStyle sets the style for the current day's active state
func (c calendar) TodayActiveStyle(value any) calendar {
	return c.set("todayActiveStyle", value)
}

// UseMobileUI enables or disables mobile UI styling
func (c calendar) UseMobileUI(value bool) calendar {
	return c.set("useMobileUI", value)
}

// Visible controls the overall visibility of the calendar
func (c calendar) Visible(value bool) calendar {
	return c.set("visible", value)
}

// VisibleOn sets a conditional expression for calendar visibility
func (c calendar) VisibleOn(value string) calendar {
	return c.set("visibleOn", value)
}
