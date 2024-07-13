/*
Copyright © 2024 Obakeng Mosadi mosadiobakeng7@gmail.com
*/
package main

import (
	"log/slog"

	"github.com/joho/godotenv"
	"github.com/obakeng-develops/houdini/cmd"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file", "err", err)
	}

	cmd.Execute()
}
