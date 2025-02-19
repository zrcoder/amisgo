package comp

import "github.com/zrcoder/amisgo/schema"

type Event schema.Schema

func NewEvent() Event {
	return Event{}
}

func (e Event) Set(key string, value any) Event {
	e[key] = value
	return e
}

func (e Event) Click(value any) Event {
	return e.Set("click", value)
}

func (e Event) MouseEnter(value any) Event {
	return e.Set("mouseenter", value)
}

func (e Event) MouseLeave(value any) Event {
	return e.Set("mouseleave", value)
}

func (e Event) RowClick(value any) Event {
	return e.Set("rowClick", value)
}

func (e Event) Change(value any) Event {
	return e.Set("change", value)
}

func (e Event) SubmitSucc(value any) Event {
	return e.Set("submitSucc", value)
}

func (e Event) Confirm(value any) Event {
	return e.Set("confirm", value)
}

func (e Event) Cancel(value any) Event {
	return e.Set("cancel", value)
}
