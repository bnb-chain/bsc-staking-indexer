package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/node-real/go-pkg/log"

	"github.com/bnb-chain/bsc-staking-indexer/cc"
	"github.com/bnb-chain/bsc-staking-indexer/indexer"
	"github.com/bnb-chain/bsc-staking-indexer/store"
)

var configPath = flag.String("config", "./configs/config-example.toml", "Give a config file path")

const svc = "bsc-staking-indexer"

func main() {
	defer log.Stop()

	flag.Parse()

	cfg := LoadConfig(*configPath)
	initLogger(&cfg.Log)

	log.Infow("start", "svc", svc, "configPath", *configPath, "config", cfg)

	st, err := store.NewStore(cfg.Store)
	if err != nil {
		log.Panicw("failed to init store", "err", err.Error())
	}

	ctx, cancelCtx := context.WithCancel(context.Background())

	i, err := indexer.New(ctx, cfg.Indexer, st)
	if err != nil {
		log.Panicw("failed to init indexer", "err", err.Error())
	}

	go i.Start(ctx)

	reporter := cc.New(cfg.CC, st)
	reporter.Start()

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigterm
	log.Infof("received signal:%v, gracefully shutdown", sig.String())
	cancelCtx()
}

func initLogger(cfg *LogConfig) {
	lvl, _ := log.ParseLevel(cfg.Level)
	log.Init(lvl, log.StandardizePath(cfg.RootDir, svc))
}
