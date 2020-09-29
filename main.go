package main

import (
	"fmt"
	"time"

	"github.com/Noah-Huppert/goconf"
)

// Config determines the details of prometheus canary metrics.
type Config struct {
	// Metrics
	Metrics []struct {
		// BaseTime is the time of day when an alert will fire.
		BaseTime time.Time `validate:"required" default:"T00:00:00Z"`
	} `validate:"required"`
}

func main() {
	// Load configuration
	loader := goconf.NewLoader()

	loader.AddConfigPath("/etc/foo/foo.*")
	loader.AddConfigPath("/etc/foo.d/*")

	config := Config{}
	err := loader.Load(&config)
	panic(err)
}
