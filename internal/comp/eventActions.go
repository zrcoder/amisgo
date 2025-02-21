package comp

import "github.com/zrcoder/amisgo/schema"

type EventActions schema.Schema

func NewEventActions(value ...EventAction) EventActions {
	res := EventActions{}
	if len(value) > 0 {
		res.Actions(value...)
	}
	return res
}

func (e EventActions) set(key string, value any) EventActions {
	e[key] = value
	return e
}

func (e EventActions) Actions(value ...EventAction) EventActions {
	return e.set("actions", value)
}
