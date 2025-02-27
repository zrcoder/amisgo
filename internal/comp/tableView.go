package comp

import "github.com/zrcoder/amisgo/schema"

type (
	TableView schema.Schema
	Tr        schema.Schema
	Td        schema.Schema
	Tcol      schema.Schema
)

func NewTableView() TableView {
	return TableView{"type": "table-view"}
}

func NewTr() Tr {
	return Tr{}
}

func NewTd() Td {
	return Td{}
}

func NewTCol() Tcol {
	return Tcol{}
}

func (tv TableView) Width(value any) TableView {
	return tv.set("width", value)
}

func (tv TableView) Padding(value any) TableView {
	return tv.set("padding", value)
}

func (tv TableView) Border(value bool) TableView {
	return tv.set("border", value)
}

func (tv TableView) BorderColor(value string) TableView {
	return tv.set("borderColor", value)
}

func (tv TableView) Cols(value ...Tcol) TableView {
	return tv.set("cols", value)
}

func (tv TableView) Trs(value ...Tr) TableView {
	return tv.set("trs", value)
}

func (tv TableView) Caption(value string) TableView {
	return tv.set("caption", value)
}

// CaptionSide sets the caption position, top or bottom, default is top
func (tv TableView) CaptionSide(value string) TableView {
	return tv.set("captionSide", value)
}

func (t Tcol) Span(value int) Tcol {
	return t.set("span", value)
}

func (t Tcol) Style(value any) Tcol {
	return t.set("style", value)
}

func (t Tr) Height(value any) Tr {
	return t.set("height", value)
}

func (t Tr) Background(value string) Tr {
	return t.set("background", value)
}

func (t Tr) Tds(value ...Td) Tr {
	return t.set("tds", value)
}

func (t Tr) VisibleOn(value string) Tr {
	return t.set("visibleOn", value)
}

func (t Td) VisibleOn(value string) Td {
	return t.set("visibleOn", value)
}

func (t Td) Background(value string) Td {
	return t.set("background", value)
}

func (t Td) Color(value string) Td {
	return t.set("color", value)
}

func (t Td) Bold(value bool) Td {
	return t.set("bold", value)
}

func (t Td) Width(value any) Td {
	return t.set("width", value)
}

func (t Td) Padding(value any) Td {
	return t.set("padding", value)
}

func (t Td) Align(value string) Td {
	return t.set("align", value)
}

func (t Td) Valign(value string) Td {
	return t.set("valign", value)
}

func (t Td) Colspan(value int) Td {
	return t.set("colspan", value)
}

func (t Td) Rowspan(value int) Td {
	return t.set("rowspan", value)
}

func (t Td) Body(value any) Td {
	return t.set("body", value)
}

func (t Td) Style(value any) Td {
	return t.set("style", value)
}

func (tv TableView) set(key string, value any) TableView {
	tv[key] = value
	return tv
}

func (t Tcol) set(key string, value any) Tcol {
	t[key] = value
	return t
}

func (t Tr) set(key string, value any) Tr {
	t[key] = value
	return t
}

func (t Td) set(key string, value any) Td {
	t[key] = value
	return t
}
