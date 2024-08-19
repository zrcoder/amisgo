package model

// Schema 通用 model，键未指定的 map
type Schema map[string]any

func (s Schema) Set(key string, value any) Schema {
	s[key] = value
	return s
}
