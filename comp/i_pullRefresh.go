package comp

type MPullRefresh Schema

func PullRefresh() MPullRefresh {
	return MPullRefresh{}
}

func (p MPullRefresh) set(key string, value interface{}) MPullRefresh {
	p[key] = value
	return p
}

func (p MPullRefresh) Disabled(value bool) MPullRefresh {
	return p.set("disabled", value)
}

func (p MPullRefresh) PullingText(value string) MPullRefresh {
	return p.set("pullingText", value)
}

func (p MPullRefresh) LoosingText(value string) MPullRefresh {
	return p.set("loosingText", value)
}
