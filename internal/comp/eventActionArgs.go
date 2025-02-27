package comp

import "github.com/zrcoder/amisgo/schema"

type EventActionArgs schema.Schema

func NewEventActionArgs() EventActionArgs {
	return EventActionArgs{}
}

func (a EventActionArgs) Set(key string, value any) EventActionArgs {
	a[key] = value
	return a
}

func (a EventActionArgs) Msg(value string) EventActionArgs {
	return a.Set("msg", value)
}

func (a EventActionArgs) MsgType(value string) EventActionArgs {
	return a.Set("msgType", value)
}

func (a EventActionArgs) Position(value string) EventActionArgs {
	return a.Set("position", value)
}

func (a EventActionArgs) CloseButton(value bool) EventActionArgs {
	return a.Set("closeButton", value)
}

func (a EventActionArgs) ShowIcon(value bool) EventActionArgs {
	return a.Set("showIcon", value)
}

func (a EventActionArgs) Timeout(value float64) EventActionArgs {
	return a.Set("timeout", value)
}

func (a EventActionArgs) Url(value string) EventActionArgs {
	return a.Set("url", value)
}

func (a EventActionArgs) Link(value string) EventActionArgs {
	return a.Set("link", value)
}

func (a EventActionArgs) Blank(value bool) EventActionArgs {
	return a.Set("timeout", value)
}

func (a EventActionArgs) Params(value any) EventActionArgs {
	return a.Set("params", value)
}

func (a EventActionArgs) TargetType(value string) EventActionArgs {
	return a.Set("targetType", value)
}

func (a EventActionArgs) Delta(value int) EventActionArgs {
	return a.Set("delta", value)
}

func (a EventActionArgs) CopyFormat(value string) EventActionArgs {
	return a.Set("copyFormat", value)
}

func (a EventActionArgs) Content(value string) EventActionArgs {
	return a.Set("content", value)
}

func (a EventActionArgs) ID(value string) EventActionArgs {
	return a.Set("id", value)
}

func (a EventActionArgs) Key(value string) EventActionArgs {
	return a.Set("key", value)
}

func (a EventActionArgs) Value(value any) EventActionArgs {
	return a.Set("value", value)
}

func (a EventActionArgs) Time(value int) EventActionArgs {
	return a.Set("time", value)
}

func (a EventActionArgs) EventName(value string) EventActionArgs {
	return a.Set("eventName", value)
}

func (a EventActionArgs) LoopName(value string) EventActionArgs {
	return a.Set("loopName", value)
}
