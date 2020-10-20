package main

var listTemplate = `
<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <title>{{ .Title }}</title>
</head>
<body>
	<ul>
		<li><a href="..">..</a></li>
		{{ range .Files }}
		<li><a href="{{ . }}">{{ . }}</a></li>
		{{ end }}
  </ul>
</body>
</html>
`
