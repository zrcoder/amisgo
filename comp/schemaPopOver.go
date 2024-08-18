package comp

// SchemaPopOver
//
// @author  slowlyo
// @version 6.7.0
type SchemaPopOver Schema

// NewSchemaPopOver 创建一个新的 SchemaPopOver 实例
func NewSchemaPopOver() SchemaPopOver {
	return SchemaPopOver{}
}

func (s SchemaPopOver) set(key string, value interface{}) SchemaPopOver {
	s[key] = value
	return s
}

// Body
func (s SchemaPopOver) Body(value ...interface{}) SchemaPopOver {
	return s.set("body", value)
}

// ClassName 类名
func (s SchemaPopOver) ClassName(value string) SchemaPopOver {
	return s.set("className", value)
}

// Mode 弹出模式 可选值: dialog | drawer | popOver
func (s SchemaPopOver) Mode(value string) SchemaPopOver {
	return s.set("mode", value)
}

// Offset 偏移量
func (s SchemaPopOver) Offset(value string) SchemaPopOver {
	return s.set("offset", value)
}

// PopOverClassName 弹框外层类名
func (s SchemaPopOver) PopOverClassName(value string) SchemaPopOver {
	return s.set("popOverClassName", value)
}

// PopOverEnableOn 配置当前行是否启动，要用表达式 (表达式，语法 `data.xxx > 5`。)
func (s SchemaPopOver) PopOverEnableOn(value string) SchemaPopOver {
	return s.set("popOverEnableOn", value)
}

// Position 弹出位置 可选值: center | left-top | left-top-left-top | left-top-left-center | left-top-left-bottom | left-top-center-top | left-top-center-center | left-top-center-bottom | left-top-right-top | left-top-right-center | left-top-right-bottom | right-top | right-top-left-top | right-top-left-center | right-top-left-bottom | right-top-center-top | right-top-center-center | right-top-center-bottom | right-top-right-top | right-top-right-center | right-top-right-bottom | left-bottom | left-bottom-left-top | left-bottom-left-center | left-bottom-left-bottom | left-bottom-center-top | left-bottom-center-center | left-bottom-center-bottom | left-bottom-right-top | left-bottom-right-center | left-bottom-right-bottom | right-bottom | right-bottom-left-top | right-bottom-left-center | right-bottom-left-bottom | right-bottom-center-top | right-bottom-center-center | right-bottom-center-bottom | right-bottom-right-top | right-bottom-right-center | right-bottom-right-bottom | fixed-center | fixed-left-top | fixed-right-top | fixed-left-bottom | fixed-right-bottom
func (s SchemaPopOver) Position(value string) SchemaPopOver {
	return s.set("position", value)
}

// ShowIcon 是否显示查看更多的 icon，通常是放大图标。
func (s SchemaPopOver) ShowIcon(value bool) SchemaPopOver {
	return s.set("showIcon", value)
}

// Size 是弹窗形式的时候有用。 可选值: sm | md | lg | xl
func (s SchemaPopOver) Size(value string) SchemaPopOver {
	return s.set("size", value)
}

// Title 标题
func (s SchemaPopOver) Title(value string) SchemaPopOver {
	return s.set("title", value)
}

// Trigger 触发条件，默认是 click 可选值: click | hover
func (s SchemaPopOver) Trigger(value string) SchemaPopOver {
	return s.set("trigger", value)
}
