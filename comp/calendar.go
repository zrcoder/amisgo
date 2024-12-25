package comp

// calendar
type calendar Schema

// Calendar 创建一个新的 Calendar 实例
func Calendar() calendar {
	return make(calendar).set("type", "calendar")
}

func (c calendar) set(key string, value any) calendar {
	c[key] = value
	return c
}

// ClassName 设置 className 属性
func (c calendar) ClassName(value string) calendar {
	return c.set("className", value)
}

// Disabled 设置 disabled 属性
func (c calendar) Disabled(value bool) calendar {
	return c.set("disabled", value)
}

// DisabledOn 设置 disabledOn 属性
func (c calendar) DisabledOn(value string) calendar {
	return c.set("disabledOn", value)
}

// EditorSetting 设置 editorSetting 属性
func (c calendar) EditorSetting(value string) calendar {
	return c.set("editorSetting", value)
}

// Hidden 设置 hidden 属性
func (c calendar) Hidden(value bool) calendar {
	return c.set("hidden", value)
}

// HiddenOn 设置 hiddenOn 属性
func (c calendar) HiddenOn(value string) calendar {
	return c.set("hiddenOn", value)
}

// ID 设置 id 属性
func (c calendar) ID(value string) calendar {
	return c.set("id", value)
}

// LargeMode 设置 largeMode 属性
func (c calendar) LargeMode(value bool) calendar {
	return c.set("largeMode", value)
}

// OnEvent 设置 onEvent 属性
func (c calendar) OnEvent(value any) calendar {
	return c.set("onEvent", value)
}

// ScheduleAction 设置 scheduleAction 属性
func (c calendar) ScheduleAction(value string) calendar {
	return c.set("scheduleAction", value)
}

// ScheduleClassNames 设置 scheduleClassNames 属性
func (c calendar) ScheduleClassNames(value string) calendar {
	return c.set("scheduleClassNames", value)
}

// Schedules 设置 schedules 属性
func (c calendar) Schedules(value string) calendar {
	return c.set("schedules", value)
}

// Static 设置 static 属性
func (c calendar) Static(value bool) calendar {
	return c.set("static", value)
}

// StaticClassName 设置 staticClassName 属性
func (c calendar) StaticClassName(value string) calendar {
	return c.set("staticClassName", value)
}

// StaticInputClassName 设置 staticInputClassName 属性
func (c calendar) StaticInputClassName(value string) calendar {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName 设置 staticLabelClassName 属性
func (c calendar) StaticLabelClassName(value string) calendar {
	return c.set("staticLabelClassName", value)
}

// StaticOn 设置 staticOn 属性
func (c calendar) StaticOn(value string) calendar {
	return c.set("staticOn", value)
}

// StaticPlaceholder 设置 staticPlaceholder 属性
func (c calendar) StaticPlaceholder(value string) calendar {
	return c.set("staticPlaceholder", value)
}

// StaticSchema 设置 staticSchema 属性
func (c calendar) StaticSchema(value string) calendar {
	return c.set("staticSchema", value)
}

// Style 设置 style 属性
func (c calendar) Style(value any) calendar {
	return c.set("style", value)
}

// TestIdBuilder 设置 testIdBuilder 属性
func (c calendar) TestIdBuilder(value string) calendar {
	return c.set("testIdBuilder", value)
}

// TestID 设置 testid 属性
func (c calendar) TestID(value string) calendar {
	return c.set("testid", value)
}

// TodayActiveStyle 设置 todayActiveStyle 属性
func (c calendar) TodayActiveStyle(value any) calendar {
	return c.set("todayActiveStyle", value)
}

// UseMobileUI 设置 useMobileUI 属性
func (c calendar) UseMobileUI(value bool) calendar {
	return c.set("useMobileUI", value)
}

// Visible 设置 visible 属性
func (c calendar) Visible(value bool) calendar {
	return c.set("visible", value)
}

// VisibleOn 设置 visibleOn 属性
func (c calendar) VisibleOn(value string) calendar {
	return c.set("visibleOn", value)
}
