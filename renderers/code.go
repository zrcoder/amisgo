package renderers

// Code 继承自 BaseRenderer
type Code struct {
	*BaseRenderer
}

// NewCode 创建一个新的 Code 实例
func NewCode() *Code {
	c := &Code{
		BaseRenderer: NewBaseRenderer(),
	}
	c.set("type", "code")
	return c
}

func (c *Code) set(key string, value interface{}) *Code {
	c.BaseRenderer.set(key, value)
	return c
}

// ClassName 设置 className 属性
func (c *Code) ClassName(value string) *Code {
	return c.set("className", value)
}

// EditorTheme 设置 editorTheme 属性
func (c *Code) EditorTheme(value string) *Code {
	return c.set("editorTheme", value)
}

// Language 设置 language 属性
func (c *Code) Language(value string) *Code {
	return c.set("language", value)
}

// Name 设置 name 属性
func (c *Code) Name(value string) *Code {
	return c.set("name", value)
}

// TabSize 设置 tabSize 属性
func (c *Code) TabSize(value string) *Code {
	return c.set("tabSize", value)
}

// Value 设置 value 属性
func (c *Code) Value(value string) *Code {
	return c.set("value", value)
}

// WordWrap 设置 wordWrap 属性
func (c *Code) WordWrap(value bool) *Code {
	return c.set("wordWrap", value)
}
