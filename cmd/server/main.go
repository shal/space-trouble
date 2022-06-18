package main

import (
	"context"
	"flag"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/opencars/space-trouble/pkg/api/http"
	"github.com/opencars/space-trouble/pkg/logger"

	"github.com/opencars/space-trouble/pkg/config"
	"github.com/opencars/space-trouble/pkg/domain/service"
	"github.com/opencars/space-trouble/pkg/store/sqlstore"
)

func main() {
	cfg := flag.String("config", "config/config.yaml", "Path to the configuration file")
	port := flag.Int("port", 8080, "Port of the server")

	flag.Parse()

	conf, err := config.New(*cfg)
	if err != nil {
		logger.Fatalf("config: %v", err)
	}

	logger.NewLogger(logger.LogLevel(conf.Log.Level), conf.Log.Mode == "dev")

	store, err := sqlstore.New(&conf.DB)
	if err != nil {
		logger.Fatalf("store: %v", err)
	}

	svc := service.NewCustomerService(store.Booking())

	addr := ":" + strconv.Itoa(*port)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	logger.Infof("Listening on %s...", addr)
	if err := http.Start(ctx, addr, &conf.Server, svc); err != nil {
		logger.Fatalf("http server failed: %v", err)
	}
}
