package cmd

import "os"

func commandExit(cfg *Config, params ...string) error {
	defer os.Exit(0)
	return nil
}
