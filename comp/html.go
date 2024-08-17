package comp

type Html = Tpl

func NewHtml() Html {
	return NewTpl().set("type", "html")
}
