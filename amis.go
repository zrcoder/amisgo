package amisgo

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"text/template"

	"github.com/zrcoder/amisgo/comp"
)

var routers = map[string]comp.AmisComp{}

func Route(patten string, component comp.AmisComp) {
	routers[patten] = component
}

func ListenAndServe(addr string, cfg ...*Config) error {
	config := getConfig(cfg)
	for patten, component := range routers {
		http.HandleFunc(patten, func(w http.ResponseWriter, r *http.Request) {
			writeHtml(config, component, w)
		})
	}
	return http.ListenAndServe(addr, nil)
}

func GenerateStaticWebsite(cfg ...*Config) error {
	config := getConfig(cfg)
	for patten, component := range routers {
		name := patten
		if patten == "/" {
			name = "index"
		}
		name += ".html"
		writer := bytes.NewBuffer(nil)
		writeHtml(config, component, writer)
		err := os.WriteFile(name, writer.Bytes(), 0o640)
		if err != nil {
			return err
		}
	}
	return nil
}

func writeHtml(config *Config, component comp.AmisComp, writer io.Writer) {
	tmpl := template.Must(template.New("").Parse(htmlTemplate))
	amisJson, _ := json.Marshal(component)

	type templateData struct {
		*Config
		AmisJson string
	}

	data := &templateData{
		Config:   config,
		AmisJson: string(amisJson),
	}

	err := tmpl.Execute(writer, data)
	if err != nil {
		panic(err)
	}
}
