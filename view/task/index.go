package task

import (
    "html/template"
)


const templateString = `
<html>
<head>
    <title>{{.Title}}</title>
</head>
<body>
    <h1>{{.Title}}</h1>
    <p>{{.Message}}</p>
</body>
</html>
`

func Index() *template.Template {
	// テンプレートをパースして取得する
	template, err := template.New("test").Parse(templateString)
	if err != nil {
    	// エラー処理
	}
	return template
}