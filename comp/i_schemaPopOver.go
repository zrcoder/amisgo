package comp

import "github.com/zrcoder/amisgo/model"

// schemaPopOver

type schemaPopOver model.Schema

// model.SchemaPopOver creates a new model.SchemaPopOver instance
func SchemaPopOver() schemaPopOver {
	return schemaPopOver{}
}

func (s schemaPopOver) set(key string, value any) schemaPopOver {
	s[key] = value
	return s
}

// Body sets the body content
func (s schemaPopOver) Body(value ...any) schemaPopOver {
	return s.set("body", value)
}

// ClassName sets the class name
func (s schemaPopOver) ClassName(value string) schemaPopOver {
	return s.set("className", value)
}

// Mode sets the pop-over mode, optional values: dialog, drawer, popOver
func (s schemaPopOver) Mode(value string) schemaPopOver {
	return s.set("mode", value)
}

// Offset sets the offset value
func (s schemaPopOver) Offset(value string) schemaPopOver {
	return s.set("offset", value)
}

// PopOverClassName sets the class name of the pop-over container
func (s schemaPopOver) PopOverClassName(value string) schemaPopOver {
	return s.set("popOverClassName", value)
}

// PopOverEnableOn configures whether the current row is enabled, using an expression (syntax: `data.xxx > 5`)
func (s schemaPopOver) PopOverEnableOn(value string) schemaPopOver {
	return s.set("popOverEnableOn", value)
}

// Position sets the pop-over position, optional values: center, left-top, left-top-left-top, ..., fixed-right-bottom
func (s schemaPopOver) Position(value string) schemaPopOver {
	return s.set("position", value)
}

// ShowIcon sets whether to show the view more icon, usually a magnifying glass icon
func (s schemaPopOver) ShowIcon(value bool) schemaPopOver {
	return s.set("showIcon", value)
}

// Size sets the size of the pop-over window, optional values: sm, md, lg, xl
func (s schemaPopOver) Size(value string) schemaPopOver {
	return s.set("size", value)
}

// Title sets the title content
func (s schemaPopOver) Title(value any) schemaPopOver {
	return s.set("title", value)
}

// Trigger sets the trigger condition, default is click, optional values: click, hover
func (s schemaPopOver) Trigger(value string) schemaPopOver {
	return s.set("trigger", value)
}
