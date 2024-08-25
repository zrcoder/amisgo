package comp

// schema 通用 model，键未指定的 map
type (
	schema map[string]any
	Schema schema
	Data   schema
)

func (s Schema) Set(key string, value any) Schema {
	s[key] = value
	return s
}
