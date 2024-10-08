{{define "base"}}
<!doctype html>
<html lang='en'>
	<head>
		<meta charset='utf-8'>
		<title>{{template "title" .}} - Shanraq</title>
	</head>
	<body>
		<header>
			<!-- Invoke the navigation template -->
			{{template "nav" .}}
		</header>
		<main>
			{{template "main" .}}
		</main>
		<!-- Invoke the footer template -->
		{{template "footer" .}}
	</body>
</html>
{{end}}