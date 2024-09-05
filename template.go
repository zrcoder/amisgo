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
    {{if .Icon}}
    <link rel="shortcut icon" href="{{.Icon}}" type="image/x-icon">
    {{end}}
    {{if .Theme}}
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/amis@6.7.0/sdk/{{.Theme}}.min.css" />
    {{else}}
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/amis@6.7.0/sdk/sdk.min.css" />
    {{end}}
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/amis@6.7.0/sdk/helper.min.css" />
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/amis@6.7.0/sdk/iconfont.min.css" />
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

/* 定义SVG变量 */
:root {
    --expand-icon: url('data:image/svg+xml;utf8,<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><rect width="100%" height="100%" fill="black" /><path fill="white" d="M256 8c137 0 248 111 248 248S393 504 256 504 8 393 8 256 119 8 256 8zm113.9 231L234.4 103.5c-9.4-9.4-24.6-9.4-33.9 0l-17 17c-9.4 9.4-9.4 24.6 0 33.9L285.1 256 183.5 357.6c-9.4 9.4-9.4 24.6 0 33.9l17 17c9.4 9.4 24.6 9.4 33.9 0L369.9 273c9.4-9.4 9.4-24.6 0-34z" /></svg>');
}

/* 隐藏原始图标 */
.dark-Table-expandBtn::before,
.dark-Table-expandBtn.is-active::before {
    visibility: hidden !important;
}

/* 添加自定义的展开图标 */
.dark-Table-expandBtn::after {
    content: '' !important;
    display: inline-block !important;
    text-align: center !important;
    cursor: pointer !important;
    transition: transform ease-in-out var(--animation-duration), top ease-in-out var(--animation-duration) !important;
    position: absolute !important;
    transform-origin: 50% 50% !important;
    width: 1rem !important; /* 调整图标宽度 */
    height: 1rem !important; /* 调整图标高度 */
    top: 50% !important;
    left: 0 !important;
    transform: translateY(-50%) !important;
    background: var(--expand-icon) no-repeat center center !important;
    background-size: contain !important;
    z-index: 1 !important; /* 确保新图标在原始图标之上 */
}

/* 添加自定义的收缩图标 */
.dark-Table-expandBtn.is-active::after {
    content: '' !important;
    display: inline-block !important;
    text-align: center !important;
    cursor: pointer !important;
    transition: transform ease-in-out var(--animation-duration), top ease-in-out var(--animation-duration) !important;
    position: absolute !important;
    transform-origin: 50% 50% !important;
    width: 1rem !important; /* 调整图标宽度 */
    height: 1rem !important; /* 调整图标高度 */
    top: 50% !important;
    left: 0 !important;
    transform: translateY(-50%) rotate(90deg) !important;
    background: var(--expand-icon) no-repeat center center !important;
    background-size: contain !important;
    z-index: 1 !important; /* 确保新图标在原始图标之上 */
}
    </style>
  </head>
  <body>
    <div id="root" class="app-wrapper"></div>
    <script src="https://cdn.jsdelivr.net/npm/amis@6.7.0/sdk/sdk.min.js"></script>
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
        {{if .Theme}}
        {theme: '{{.Theme}}'}
        {{end}}
        );
      })();
    </script>
  </body>
</html>`
