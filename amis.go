package amisgo

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var routes = map[string]any{}

func Serve(path string, component any) {
	routes[path] = component
}

func ListenAndServe(addr string, cfg ...*Config) error {
	config := getConfig(cfg)

	if config.StaticDir != "" {
		dir := strings.Trim(config.StaticDir, "/")
		config.StaticDir = "/" + dir + "/"
		if config.StaticFS == nil {
			http.Handle(config.StaticDir, http.StripPrefix(config.StaticDir, http.FileServer(http.Dir(dir))))
		} else {
			http.Handle(config.StaticDir, http.FileServer(http.FS(config.StaticFS)))
		}
	}

	for path, com := range routes {
		http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			writeHtml(config, com, w)
		})
	}

	return http.ListenAndServe(addr, nil)
}

func GenerateStaticWebsite(outputDir string, component any, cfg ...*Config) error {
	config := getConfig(cfg)
	if outputDir == "" {
		outputDir = "."
	}
	writer := bytes.NewBuffer(nil)
	writeHtml(config, component, writer)
	return os.WriteFile(filepath.Join(outputDir, "index.html"), writer.Bytes(), 0o640)
}

func writeHtml(config *Config, component any, writer io.Writer) {
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
