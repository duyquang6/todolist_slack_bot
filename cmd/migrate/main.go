// This package serve for db migration
package main

import (
	"context"
	"flag"
	"log"

	"github.com/duyquang6/go-rbac-practice/internal/buildinfo"
	"github.com/duyquang6/go-rbac-practice/pkg/logging"

	"github.com/sethvargo/go-signalcontext"
)

var (
	migrationDir = flag.String("path", "migrations/", "path to migrations folder")
)

func main() {
	flag.Parse()
	ctx, done := signalcontext.OnInterrupt()

	logger := logging.NewLoggerFromEnv().
		With("build_id", buildinfo.RBACServer.ID()).
		With("build_time", buildinfo.RBACServer.Time())
	ctx = logging.WithLogger(ctx, logger)

	defer func() {
		done()
		if r := recover(); r != nil {
			logger.Fatalw("application panic", "panic", r)
		}
	}()

	err := realMain(ctx)
	done()

	if err != nil {
		log.Fatal(err)
	}
	logger.Info("successful shutdown")
}

func realMain(ctx context.Context) error {
	logger := logging.FromContext(ctx)
	logger.Info(migrationDir)
	return nil
}
