package model

type PullRefreshSchema Schema

func PullRefresh() PullRefreshSchema {
	return PullRefreshSchema{}
}

func (p PullRefreshSchema) set(key string, value interface{}) PullRefreshSchema {
	p[key] = value
	return p
}

func (p PullRefreshSchema) Disabled(value bool) PullRefreshSchema {
	return p.set("disabled", value)
}

func (p PullRefreshSchema) PullingText(value string) PullRefreshSchema {
	return p.set("pullingText", value)
}

func (p PullRefreshSchema) LoosingText(value string) PullRefreshSchema {
	return p.set("loosingText", value)
}
