package comp

import "github.com/zrcoder/amisgo/schema"

type Amis schema.Schema

func NewAmis() Amis {
	return Amis{"type": "amis"}
}

func (a Amis) set(key string, value any) Amis {
	a[key] = value
	return a
}

func (a Amis) Name(value string) Amis {
	return a.set("name", value)
}

func (a Amis) Props(value any) Amis {
	return a.set("props", value)
}
