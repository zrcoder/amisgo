package comp

// code
type code schema

// Code 创建一个新的 Code 实例
func Code() code {
	return make(code).set("type", "code")
}

func (c code) set(key string, value any) code {
	c[key] = value
	return c
}

// ClassName 设置 className 属性
func (c code) ClassName(value string) code {
	return c.set("className", value)
}

// EditorTheme 设置 editorTheme 属性
func (c code) EditorTheme(value string) code {
	return c.set("editorTheme", value)
}

// Language 设置 language 属性
func (c code) Language(value string) code {
	return c.set("language", value)
}

// Name 设置 name 属性
func (c code) Name(value string) code {
	return c.set("name", value)
}

// TabSize 设置 tabSize 属性
func (c code) TabSize(value string) code {
	return c.set("tabSize", value)
}

// Value 设置 value 属性
func (c code) Value(value string) code {
	return c.set("value", value)
}

// WordWrap 设置 wordWrap 属性
func (c code) WordWrap(value bool) code {
	return c.set("wordWrap", value)
}
