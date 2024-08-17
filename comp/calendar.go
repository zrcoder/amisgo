package comp

// Calendar 继承自 BaseRenderer
type Calendar BaseRenderer

// NewCalendar 创建一个新的 Calendar 实例
func NewCalendar() Calendar {
	c := Calendar(make(BaseRenderer))
	c.set("type", "calendar")
	return c
}

func (c Calendar) set(key string, value interface{}) Calendar {
	c[key] = value
	return c
}

// ClassName 设置 className 属性
func (c Calendar) ClassName(value string) Calendar {
	return c.set("className", value)
}

// Disabled 设置 disabled 属性
func (c Calendar) Disabled(value bool) Calendar {
	return c.set("disabled", value)
}

// DisabledOn 设置 disabledOn 属性
func (c Calendar) DisabledOn(value string) Calendar {
	return c.set("disabledOn", value)
}

// EditorSetting 设置 editorSetting 属性
func (c Calendar) EditorSetting(value string) Calendar {
	return c.set("editorSetting", value)
}

// Hidden 设置 hidden 属性
func (c Calendar) Hidden(value bool) Calendar {
	return c.set("hidden", value)
}

// HiddenOn 设置 hiddenOn 属性
func (c Calendar) HiddenOn(value string) Calendar {
	return c.set("hiddenOn", value)
}

// ID 设置 id 属性
func (c Calendar) ID(value string) Calendar {
	return c.set("id", value)
}

// LargeMode 设置 largeMode 属性
func (c Calendar) LargeMode(value bool) Calendar {
	return c.set("largeMode", value)
}

// OnEvent 设置 onEvent 属性
func (c Calendar) OnEvent(value string) Calendar {
	return c.set("onEvent", value)
}

// ScheduleAction 设置 scheduleAction 属性
func (c Calendar) ScheduleAction(value string) Calendar {
	return c.set("scheduleAction", value)
}

// ScheduleClassNames 设置 scheduleClassNames 属性
func (c Calendar) ScheduleClassNames(value string) Calendar {
	return c.set("scheduleClassNames", value)
}

// Schedules 设置 schedules 属性
func (c Calendar) Schedules(value string) Calendar {
	return c.set("schedules", value)
}

// Static 设置 static 属性
func (c Calendar) Static(value bool) Calendar {
	return c.set("static", value)
}

// StaticClassName 设置 staticClassName 属性
func (c Calendar) StaticClassName(value string) Calendar {
	return c.set("staticClassName", value)
}

// StaticInputClassName 设置 staticInputClassName 属性
func (c Calendar) StaticInputClassName(value string) Calendar {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName 设置 staticLabelClassName 属性
func (c Calendar) StaticLabelClassName(value string) Calendar {
	return c.set("staticLabelClassName", value)
}

// StaticOn 设置 staticOn 属性
func (c Calendar) StaticOn(value string) Calendar {
	return c.set("staticOn", value)
}

// StaticPlaceholder 设置 staticPlaceholder 属性
func (c Calendar) StaticPlaceholder(value string) Calendar {
	return c.set("staticPlaceholder", value)
}

// StaticSchema 设置 staticSchema 属性
func (c Calendar) StaticSchema(value string) Calendar {
	return c.set("staticSchema", value)
}

// Style 设置 style 属性
func (c Calendar) Style(value string) Calendar {
	return c.set("style", value)
}

// TestIdBuilder 设置 testIdBuilder 属性
func (c Calendar) TestIdBuilder(value string) Calendar {
	return c.set("testIdBuilder", value)
}

// TestID 设置 testid 属性
func (c Calendar) TestID(value string) Calendar {
	return c.set("testid", value)
}

// TodayActiveStyle 设置 todayActiveStyle 属性
func (c Calendar) TodayActiveStyle(value string) Calendar {
	return c.set("todayActiveStyle", value)
}

// UseMobileUI 设置 useMobileUI 属性
func (c Calendar) UseMobileUI(value bool) Calendar {
	return c.set("useMobileUI", value)
}

// Visible 设置 visible 属性
func (c Calendar) Visible(value bool) Calendar {
	return c.set("visible", value)
}

// VisibleOn 设置 visibleOn 属性
func (c Calendar) VisibleOn(value string) Calendar {
	return c.set("visibleOn", value)
}
