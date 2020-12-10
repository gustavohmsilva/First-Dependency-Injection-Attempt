package main

import "github.com/gustavohmsilva/test-dependency-injection/app"

func main() {
	err := app.Start("localhost", "8080")
	if err != nil {
		panic(err)
	}
}
