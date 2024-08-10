package comp

type Ipage Option

func Page() Ipage {
	return Ipage{"type": "page"}
}

func (p Ipage) Body(input ...any) Component {
	p["body"] = input
	return p
}

func (p Ipage) Title(input any) Ipage {
	p["title"] = input
	return p
}

func (p Ipage) SubTitle(input any) Ipage {
	p["subTitle"] = input
	return p
}

func (p Ipage) Remark(input any) Ipage {
	p["remark"] = input
	return p
}

func (p Ipage) Aside(input any) Ipage {
	p["aside"] = input
	return p
}

func (p Ipage) AsideResizor(input bool) Ipage {
	p["asideResizor"] = input
	return p
}

func (p Ipage) AsideMinWidth(input any) Ipage {
	p["asideMinWidth"] = input
	return p
}

func (p Ipage) AsideMaxWidth(input any) Ipage {
	p["asideMaxWidth"] = input
	return p
}

func (p Ipage) AsideSticky(input bool) Ipage {
	p["asideSticky"] = input
	return p
}

func (p Ipage) Toolbar(input any) Ipage {
	p["toolbar"] = input
	return p
}

func (p Ipage) ClassName(input string) Ipage {
	p["className"] = input
	return p
}

func (p Ipage) CssVars(input any) Ipage {
	p["cssVars"] = input
	return p
}

func (p Ipage) ToolbarClassName(input string) Ipage {
	p["toolbarClassName"] = input
	return p
}

func (p Ipage) BodyClassName(input string) Ipage {
	p["bodyClassName"] = input
	return p
}

func (p Ipage) AsideClassName(input string) Ipage {
	p["asideClassName"] = input
	return p
}

func (p Ipage) HeaderClassName(input string) Ipage {
	p["headerClassName"] = input
	return p
}

func (p Ipage) InitApi(input any) Ipage {
	p["initApi"] = input
	return p
}

func (p Ipage) InitFetch(input bool) Ipage {
	p["initFetch"] = input
	return p
}

func (p Ipage) InitFetchOn(input any) Ipage {
	p["initFetchOn"] = input
	return p
}

func (p Ipage) Interval(input any) Ipage {
	p["interval"] = input
	return p
}

func (p Ipage) SilentPolling(input bool) Ipage {
	p["silentPolling"] = input
	return p
}

func (p Ipage) StopAutoRefreshWhen(input any) Ipage {
	p["stopAutoRefreshWhen"] = input
	return p
}

func (p Ipage) PullRefresh(input any) Ipage {
	p["pullRefresh"] = input
	return p
}
