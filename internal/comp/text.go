package comp

import "github.com/zrcoder/amisgo/model"

type Text model.Schema

func NewText() Text {
	return Text{"type": "text"}
}

func (t Text) set(key string, value any) Text {
	t[key] = value
	return t
}

func (t Text) Label(value string) Text {
	return t.set("label", value)
}
