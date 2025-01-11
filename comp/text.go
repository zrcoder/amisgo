package comp

import "github.com/zrcoder/amisgo/model"

type text model.Schema

func Text() text {
	return text{"type": "text"}
}

func (t text) set(key string, value any) text {
	t[key] = value
	return t
}

func (t text) Label(value string) text {
	return t.set("label", value)
}
