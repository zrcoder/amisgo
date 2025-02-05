package model

type PullRefresh Schema

func NewPullRefresh() PullRefresh {
	return PullRefresh{}
}

func (p PullRefresh) set(key string, value interface{}) PullRefresh {
	p[key] = value
	return p
}

func (p PullRefresh) Disabled(value bool) PullRefresh {
	return p.set("disabled", value)
}

func (p PullRefresh) PullingText(value string) PullRefresh {
	return p.set("pullingText", value)
}

func (p PullRefresh) LoosingText(value string) PullRefresh {
	return p.set("loosingText", value)
}
