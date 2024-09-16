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

var tmpl = template.Must(template.New("").Parse(htmlTemplate))

func Serve(path string, component any) {
	routes[path] = component
}

func Redirect(src, dst string) {
	http.HandleFunc(src, func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, dst, http.StatusTemporaryRedirect)
	})
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
		com := com
		http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			writeHtml(config, com, w)
		})
	}

	return http.ListenAndServe(addr, nil)
}

func GenerateStaticWebsite(outputDir string, cfg ...*Config) error {
	config := getConfig(cfg)
	if outputDir == "" {
		outputDir = "."
	}
	for path, com := range routes {
		path = strings.TrimLeft(path, "/")
		if path == "" {
			path = "index"
		}
		path += ".html"
		writer := bytes.NewBuffer(nil)
		writeHtml(config, com, writer)
		err := os.WriteFile(filepath.Join(outputDir, path), writer.Bytes(), 0o640)
		if err != nil {
			return err
		}
	}
	return nil
}

func writeHtml(config *Config, component any, writer io.Writer) {
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
