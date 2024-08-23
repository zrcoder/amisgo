package amisgo

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	"github.com/zrcoder/amisgo/comp"
)

func ListenAndServe(addr string, component comp.AmisComp, cfg ...*Config) error {
	config := getConfig(cfg)
	http.HandleFunc(config.Route, func(w http.ResponseWriter, r *http.Request) {
		writeHtml(config, component, w)
	})
	return http.ListenAndServe(addr, nil)
}

func GenerateStaticWebsite(outputDir string, component comp.AmisComp, cfg ...*Config) error {
	config := getConfig(cfg)
	if outputDir == "" {
		outputDir = "."
	}
	name := "index.html"
	writer := bytes.NewBuffer(nil)
	writeHtml(config, component, writer)
	err := os.WriteFile(filepath.Join(outputDir, name), writer.Bytes(), 0o640)
	if err != nil {
		return err
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
