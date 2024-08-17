package amisgo

const htmlTemplate = `<!DOCTYPE html>
<html lang={{ .Lang }}>
  <head>
    <meta charset="UTF-8" />
    <title>{{ .Title }}</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <meta
      name="viewport"
      content="width=device-width, initial-scale=1, maximum-scale=1"
    />
    <meta http-equiv="X-UA-Compatible" content="IE=Edge" />
    {{if eq .Theme "antd"}}
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/amis@6.7.0/sdk/antd.min.css" />
    {{else if eq .Theme "ang"}}
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/amis@6.7.0/sdk/ang.min.css" />
    {{else}}
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/amis@6.7.0/sdk/sdk.min.css" />
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/amis@6.7.0/sdk/helper.min.css" />
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/amis@6.7.0/sdk/iconfont.min.css" />
    {{end}}

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
    </style>
  </head>
  <body>
    <div id="root" class="app-wrapper"></div>
    <script src="https://cdn.jsdelivr.net/npm/amis@6.7.0/sdk/sdk.min.js"></script>
    <script type="text/javascript">
      (function () {
        let amis = amisRequire('amis/embed');
        let amisJSON = {{ .AmisJson }};
        let amisScoped = amis.embed(
        '#root', 
        amisJSON, 
        {},
        {{if eq .Theme "antd"}}
        {theme: 'antd'}
        {{else if eq .Theme "ang"}}
        {theme: 'ang'}
        {{end}}
        );
      })();
    </script>
  </body>
</html>`