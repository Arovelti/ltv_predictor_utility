package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Arovelti/ltv_predictor_utility/aggregator"
	"github.com/Arovelti/ltv_predictor_utility/loader"
	"github.com/Arovelti/ltv_predictor_utility/predictor"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-ctx.Done()
		cancel()
	}()

	modelFlag := flag.String("model", "linear", "Model selection (linear or quadratic)")
	sourceFlag := flag.String("source", "test_data.csv", "Source file with data (JSON or CSV)")
	aggregateFlag := flag.String("aggregate", "country", "Data aggregation level (country or campaign)")
	flag.Parse()

	if *sourceFlag == "" {
		log.Fatal("the source flag is missed")
		return
	}

	var l loader.User
	ud, err := l.LoadData(*sourceFlag)
	if err != nil {
		log.Fatalf("unable to load data: %v", err)
	}

	aggregateDataList := aggregator.AggregateData(ud, *aggregateFlag)

	for _, aggregatedData := range aggregateDataList {
		ltv := predictor.PredictLTV(aggregatedData, *modelFlag)
		fmt.Printf("%s: %.2f\n", aggregatedData.Key, ltv)
	}

}
