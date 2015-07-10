{{ define "Index" }}
<!DOCTYPE html>
<head>
	<title>gletter Demo</title>
</head>

<body>
	<p>This is a simply <b>gletter</b> demo form</p>
	<form method="POST" action="/send">
		<label>Name</label><br />
		<input type="text" name="name" /><br /><br />
		<label>E-Mail</label><br />
		<input type="text" name="email" /><br /><br />
		<label>Phone</label><br />
		<input type="text" name="phone" /><br /><br />
		<label>Message</label><br />
		<textarea name="message"></textarea><br /><br />
		<input type="submit" value="Send Mail" />
	</form>
</body>
</html>
{{ end }}
