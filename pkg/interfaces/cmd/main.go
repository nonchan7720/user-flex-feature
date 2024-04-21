package cmd

import "log/slog"

func Execute() {
	cmd := rootCommand()
	if err := cmd.Execute(); err != nil {
		slog.Error(err.Error())
	}
}
