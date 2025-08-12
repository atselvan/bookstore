package main

import (
	"log/slog"
	"os"

	"github.com/atselvan/bkst/app"
)

func main() {
	err := app.Start()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
