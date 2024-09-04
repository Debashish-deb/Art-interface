package main

// For Homepage
const homePageTemplate = `
<!DOCTYPE html>
<html>
<head>
	<title>Encoder/Decoder</title>
	<style>
		body {
			text-align: center;
			font-family: Arial, sans-serif;
			background-color: #f0f0f0;
			margin: 0;
			padding: 20px;
		}
		textarea {
			width: 50%;
			height: 200px;
			margin-bottom: 10px;
			padding: 10px;
			font-family: Arial, Arial;
		}
		input[type="submit"] {
			padding: 10px 20px;
			background-color: #7b0000;
			color: white;
			border: none;
			cursor: pointer;
			margin-top: 10px;
		}
		h1, h2 {
			margin-bottom: 10px;
		}
		
	</style>
</head>
<body>
	<h1 style="color: blue;"">Welcome to</h1>
	<h1 style="color: blue;"">Ultimate Encoder/Decoder Hub</h1>
	<h2>Encode Text</h2>
	<form action="/encode" method="post">
		<textarea name="text" placeholder="Please, enter text to encode..."></textarea><br>
		<input type="submit" value="Encode">
		<br>
		<br>
	</form>
	<h2>Decode Text</h2>
	<form action="/decode" method="post">
		<textarea name="text" placeholder="Please, enter encoded text to decode..."></textarea><br>
		<input type="submit" value="Decode">
	</form>
</body>
</html>`

// For Resultpage
const resultPageTemplate = `
<!DOCTYPE html>
<html>
<head>
	<title>Result</title>
	<style>
		body {
			text-align: center;
			font-family: Arial, Arial;
			margin: 0;
			padding: 20px;
		}
		pre {
			text-align: left;
			background-color: white;
			padding: 20px;
			white-space: pre-wrap;
			font-family: "Courier New", Courier, monospace;
			word-wrap: break-word;
		}
		a {
			display: inline-block;
			margin-top: 20px;
			padding: 10px 20px;
			background-color: #7b0000;
			color: white;
			text-decoration: none;
		}
	</style>
</head>
<body>
	<h1>{{.Title}}</h1>
	<pre>{{.Result}}</pre>
	<a href="/">Return to Homepage</a>
</body>
</html>`
