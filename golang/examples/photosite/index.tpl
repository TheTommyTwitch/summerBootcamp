<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Photos</title>
</head>
<body>
  {{ range .Photos}}
  <img src="{{.}}">
  {{ end }}

</body>
</html>
