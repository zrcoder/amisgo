package amisgo

import "github.com/zrcoder/amisgo/comp"

func (a *Amis) Page() comp.Ipage {
	res := comp.Page()
	a.ensureComp(res)
	return res
}

func (a *Amis) ensureComp(c comp.Component) {
	if a.component == nil {
		a.component = c
	}
}
