package main

/*
To run the code above, follow these steps:

1) Save the code in a file called main.go.
2) Install Go, if you haven't already, by following the instructions on the official Go website.
3) Open a terminal and navigate to the directory where the main.go file is located.
4) Run the command go run main.go to start the server.
5) Open a web browser and navigate to http://localhost:8080. You should see the HTML page with the title "Hello" and a body containing "OST".
*/

import (
	"fmt"
	"net/http"
)

const htmlTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Hello</title>
</head>
<body>
	<h1>OST</h1>
</body>
</html>
`

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, htmlTemplate)
	})

	port := "8080"
	fmt.Printf("Server listening on port %s...\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
