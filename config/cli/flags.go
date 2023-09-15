package cli

import (
	"errors"
	"flag"
	"log"
)

type Flags struct {
	Model          string
	Source         string
	AggregatorType string
}

func GetFlags() (*Flags, error) {
	var f Flags

	modelFlag := flag.String("model", "linear", "Model selection (linear or quadratic)")
	sourceFlag := flag.String("source", "test_data.csv", "Source file with data (JSON or CSV)")
	aggregateFlag := flag.String("aggregate", "country", "Data aggregation level (country or campaign)")
	flag.Parse()

	if *sourceFlag == "" {
		log.Fatal("the source flag is missed")
		return nil, errors.New("the source flag is not found")
	}

	f = Flags{
		Model:          *modelFlag,
		Source:         *sourceFlag,
		AggregatorType: *aggregateFlag,
	}

	return &f, nil
}
