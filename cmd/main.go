package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Arovelti/ltv_predictor_utility/adapter/loader"
	"github.com/Arovelti/ltv_predictor_utility/config/cli"
	"github.com/Arovelti/ltv_predictor_utility/domain/models"
	"github.com/Arovelti/ltv_predictor_utility/domain/predict"
)

type Predictor interface {
	PredictLTV(aggregatedData models.AggregatedData, model string) float64
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-ctx.Done()
		cancel()
	}()

	flags, err := cli.GetFlags()
	if err != nil {
		log.Fatalf("unable to parse flgs^ %v", err)
	}

	l := loader.NewLoader(flags.Source)

	predict.PredictLTV(l, flags.Source, flags.Model, flags.AggregatorType)
}
