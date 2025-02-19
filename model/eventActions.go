package model

type EventActions Schema

func NewEventActions() EventActions {
	return EventActions{}
}

func (e EventActions) set(key string, value any) EventActions {
	e[key] = value
	return e
}

func (e EventActions) Actions(value ...any) EventActions {
	return e.set("actions", value)
}
