// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"context"
	"flag"

	"storj.io/storj/pkg/process"
)

func main() {
	flag.Set("metrics.interval", "1s")
	process.Must(process.Main(process.ServiceFunc(run)))
}

func run(ctx context.Context) error {
	// just go to sleep and let the background telemetry start sending
	select {}
}
