package amisgo

import (
	"encoding/json"
	"net/http"
	"text/template"

	"github.com/zrcoder/amisgo/comp"
)

type Amis struct {
	Theme string
	Lang  string
	Title string
}

func (a *Amis) Route(patten string, component comp.AmisComp) {
	http.HandleFunc(patten, func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.New("").Parse(htmlTemplate))
		amisJson, _ := json.Marshal(component)
		data := struct {
			*Amis
			AmisJson string
		}{
			Amis:     a,
			AmisJson: string(amisJson),
		}
		err := tmpl.Execute(w, data)
		if err != nil {
			panic(err)
		}
	})
}

func (a *Amis) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, nil)
}
