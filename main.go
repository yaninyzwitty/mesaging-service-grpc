package main

import (
	"log/slog"
	"os"

	"github.com/yaninyzwitty/go-grpc-messaging/pkg"
)

func main() {
	var cfg pkg.Config
	file, err := os.Open("config.yaml")
	if err != nil {
		slog.Error("failed to open config file", "error", err)
		os.Exit(1)
	}
	defer file.Close()

	if err := cfg.LoadFile(file); err != nil {
		slog.Error("failed to load file", "error", err)
		os.Exit(1)
	}
}
