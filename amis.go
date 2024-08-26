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

func ListenAndServe[T comp.AmisComp](component T, cfg ...*Config) error {
	config := getConfig(cfg)
	if config.AssertsPath != "" {
		var handler http.Handler
		if config.AssertsFS != nil {
			handler = http.FileServer(http.FS(config.AssertsFS))
		} else {
			handler = http.FileServer(http.Dir(config.AssertsPath))
		}
		http.Handle(config.AssertsPath, handler)
	}
	http.HandleFunc(config.Path, func(w http.ResponseWriter, r *http.Request) {
		writeHtml(config, component, w)
	})
	return http.ListenAndServe(config.Addr, nil)
}

func GenerateStaticWebsite[T comp.AmisComp](outputDir string, component T, cfg ...*Config) error {
	config := getConfig(cfg)
	if outputDir == "" {
		outputDir = "."
	}
	writer := bytes.NewBuffer(nil)
	writeHtml(config, component, writer)
	return os.WriteFile(filepath.Join(outputDir, "index.html"), writer.Bytes(), 0o640)
}

func writeHtml[T comp.AmisComp](config *Config, component T, writer io.Writer) {
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
