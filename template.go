package amisgo

const htmlTemplate = `<!DOCTYPE html>
<html lang="{{ .Lang }}">
  <head>
    <meta charset="UTF-8" />
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <meta
      name="viewport"
      content="width=device-width, initial-scale=1, maximum-scale=1"
    />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
     <title>{{ .Title }}</title>
    {{if .Icon}}
    <link rel="icon" href="{{.Icon}}" />
    {{end}}
    {{if .Theme}}
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/amis@6.8.0/sdk/{{.Theme}}.min.css" />
    {{else}}
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/amis@6.8.0/sdk/sdk.min.css" />
    {{end}}
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/amis@6.8.0/sdk/helper.min.css" />
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/amis@6.8.0/sdk/iconfont.min.css" />
    <style>
      html,
      body,
      .app-wrapper {
        position: relative;
        width: 100%;
        height: 100%;
        margin: 0;
        padding: 0;
      }
      {{.CustomCSS}}
    </style>
  </head>
  <body>
    <div id="root" class="app-wrapper"></div>
    <script src="https://cdn.jsdelivr.net/npm/amis@6.8.0/sdk/sdk.min.js"></script>
    {{if .CustomJS}}
    <script type="text/javascript">
    {{.CustomJS}}
    </script> 
    {{end}}
    <script type="text/javascript">
      (function () {
        let amis = amisRequire('amis/embed');
        let amisScoped = amis.embed(
        '#root', 
        {{ .AmisJson }}, 
        {
          {{if eq .Lang "en"}}
          locale: 'en-US'
          {{end}}
        },
        {
          {{if .Theme}}
          theme: '{{ .Theme }}',
          {{end}}
          replaceText: {
            HOST: "{{ .Host }}"
          }
        }
        );
      })();
    </script>
  </body>
</html>`
