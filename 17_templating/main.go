package main

import (
	"fmt"
)

func main() {
	name := "Todd Mcloed"

	tpl := `
  <!DOCTYPE html>
  <html lang="en">
  <head>
  <meta charset="UTF-8">
  <title>hello world</title>
  </head>
  <body>
  <h1>` + name + `
  </h1>
  </body>
  </html>
  `
	fmt.Println(tpl)
}
