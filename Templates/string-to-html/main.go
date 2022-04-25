package main

import(
	"fmt"
)

func main()  {
	name1 := "Gaurav Sahil"
	// name2 := "Tanya Roy"

	text := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Document</title>
	</head>
	<body>
		<h1> `+name1+` </h1>
	</body>
	</html>`

	fmt.Println(text)
	//using go run main.go -> index.html from cmd to pipeline output to index.html file
}