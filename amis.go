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

type Config struct {
	Theme string
	Lang  string
	Title string
}

type templateData struct {
	*Config
	AmisJson string
}

var routers = map[string]comp.AmisComp{}

func Route(patten string, component comp.AmisComp) {
	routers[patten] = component
}

func ListenAndServe(addr string, cfg ...*Config) error {
	config := getConfig(cfg)
	for patten, component := range routers {
		http.HandleFunc(patten, func(w http.ResponseWriter, r *http.Request) {
			err := writeTemplateData(config, component, w)
			if err != nil {
				panic(err)
			}
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
		err := writeTemplateData(config, component, writer)
		if err != nil {
			return err
		}
		err = os.WriteFile(name, writer.Bytes(), 0o640)
		if err != nil {
			return err
		}
	}
	return nil
}

func getConfig(cfg []*Config) *Config {
	config := &Config{
		Title: "Amisgo",
		Lang:  "zh",
	}
	if len(cfg) > 0 {
		config = cfg[0]
	}
	return config
}

func writeTemplateData(config *Config, component comp.AmisComp, writer io.Writer) error {
	tmpl := template.Must(template.New("").Parse(htmlTemplate))
	amisJson, _ := json.Marshal(component)
	data := &templateData{
		Config:   config,
		AmisJson: string(amisJson),
	}
	return tmpl.Execute(writer, data)
}
