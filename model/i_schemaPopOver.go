package model

type SchemaPopOver Schema

func NewSchemaPopOver() SchemaPopOver {
	return SchemaPopOver{}
}

func (s SchemaPopOver) set(key string, value any) SchemaPopOver {
	s[key] = value
	return s
}

// Body sets the body content
func (s SchemaPopOver) Body(value ...any) SchemaPopOver {
	return s.set("body", value)
}

// ClassName sets the class name
func (s SchemaPopOver) ClassName(value string) SchemaPopOver {
	return s.set("className", value)
}

// Mode sets the pop-over mode, optional values: dialog, drawer, popOver
func (s SchemaPopOver) Mode(value string) SchemaPopOver {
	return s.set("mode", value)
}

// Offset sets the offset value
func (s SchemaPopOver) Offset(value string) SchemaPopOver {
	return s.set("offset", value)
}

// PopOverClassName sets the class name of the pop-over container
func (s SchemaPopOver) PopOverClassName(value string) SchemaPopOver {
	return s.set("popOverClassName", value)
}

// PopOverEnableOn configures whether the current row is enabled, using an expression (syntax: `data.xxx > 5`)
func (s SchemaPopOver) PopOverEnableOn(value string) SchemaPopOver {
	return s.set("popOverEnableOn", value)
}

// Position sets the pop-over position, optional values: center, left-top, left-top-left-top, ..., fixed-right-bottom
func (s SchemaPopOver) Position(value string) SchemaPopOver {
	return s.set("position", value)
}

// ShowIcon sets whether to show the view more icon, usually a magnifying glass icon
func (s SchemaPopOver) ShowIcon(value bool) SchemaPopOver {
	return s.set("showIcon", value)
}

// Size sets the size of the pop-over window, optional values: sm, md, lg, xl
func (s SchemaPopOver) Size(value string) SchemaPopOver {
	return s.set("size", value)
}

// Title sets the title content
func (s SchemaPopOver) Title(value any) SchemaPopOver {
	return s.set("title", value)
}

// Trigger sets the trigger condition, default is click, optional values: click, hover
func (s SchemaPopOver) Trigger(value string) SchemaPopOver {
	return s.set("trigger", value)
}
