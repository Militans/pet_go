package main

import (
	"context"
	"fmt"
	"start_chi/application"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())

	if err != nil {
		fmt.Println("no", err)
	}
}
