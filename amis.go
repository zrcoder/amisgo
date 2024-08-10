package amisgo

import (
	"encoding/json"
	"net/http"
	"text/template"

	"github.com/zrcoder/amisgo/comp"
)

type Amis struct {
	Config    Config
	Router    string // default is "/"
	component comp.Component
}

type Config struct {
	Theme string
}

func New() *Amis {
	res := &Amis{}
	res.init()
	return res
}

func (a *Amis) init() {
	a.Router = "/"
}

func (a *Amis) Theme(theme string) {
	a.Config.Theme = theme
}

func (a *Amis) Route(r string) {
	a.Router = r
}

func (a *Amis) HandleRouter() {
	http.HandleFunc(a.Router, func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.New("").Parse(htmlTemplate))
		jsonData, _ := json.Marshal(a.component)
		data := struct {
			Theme    string
			AmisJson string
		}{
			Theme:    a.Config.Theme,
			AmisJson: string(jsonData),
		}
		tmpl.Execute(w, data)
	})
}

func (a *Amis) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, nil)
}
