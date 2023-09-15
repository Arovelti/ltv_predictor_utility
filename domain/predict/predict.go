package predict

import (
	"log"

	"github.com/Arovelti/ltv_predictor_utility/adapter/algo"
	"github.com/Arovelti/ltv_predictor_utility/domain/aggregator"
	"github.com/Arovelti/ltv_predictor_utility/domain/models"
)

type Predicter interface {
	PredictLTV(aggregatedData models.AggregatedData, model string) float64
}

type Loader interface {
	LoadData(source string) (*models.UserData, error)
}

func PredictLTV(loader Loader, source, model, level string) (*[]models.AggregatedData, error) {
	userData, err := loader.LoadData(source)
	if err != nil {
		log.Fatalf("unable to load data: %v", err)
	}

	aggregateDataList := aggregator.AggregateData(userData, level)

	for i, aggregatedData := range aggregateDataList {

		ltv, err := process(aggregatedData, model)
		if err != nil {
			return nil, err
		}

		aggregateDataList[i].LTV60thDay = ltv
	}

	return &aggregateDataList, nil
}

func process(data models.AggregatedData, model string) (float64, error) {
	switch model {
	case "linear":
		return algo.LinearPredict(data), nil
	case "quadratic":
		return algo.QuadraticPredictor(data), nil
	default:
		log.Print("Unsupported model: Linear extrapolation will be used")
		return data.AverageLtv * 60.0, nil
	}
}
