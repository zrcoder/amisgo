package comp

type pullRefresh schema

func PullRefresh() pullRefresh {
	return pullRefresh{}
}

func (p pullRefresh) set(key string, value interface{}) pullRefresh {
	p[key] = value
	return p
}

func (p pullRefresh) Disabled(value bool) pullRefresh {
	return p.set("disabled", value)
}

func (p pullRefresh) PullingText(value string) pullRefresh {
	return p.set("pullingText", value)
}

func (p pullRefresh) LoosingText(value string) pullRefresh {
	return p.set("loosingText", value)
}
