package comp

// Markdown rendering

type markdown Schema

// Markdown creates a new Markdown instance
func Markdown() markdown {
	return markdown{}.set("type", "markdown")
}

func (m markdown) set(key string, value any) markdown {
	m[key] = value
	return m
}

// ClassName sets the class name
func (m markdown) ClassName(value string) markdown {
	m.set("className", value)
	return m
}

// Name sets the name
func (m markdown) Name(value string) markdown {
	m.set("name", value)
	return m
}

// Options sets the Markdown options
func (m markdown) Options(value Schema) markdown {
	m.set("options", value)
	return m
}

// Src sets the external source URL
func (m markdown) Src(value string) markdown {
	m.set("src", value)
	return m
}

// Value sets the static value
func (m markdown) Value(value string) markdown {
	m.set("value", value)
	return m
}
