package main

import (
	"app-cli/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Start app-cli")

	app := app.Gerar()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
