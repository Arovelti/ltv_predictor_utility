package predictor

import (
	"log"

	"github.com/Arovelti/ltv_predictor_utility/entity"
)

func PredictLTV(aggregatedData entity.AggregatedData, model string) float64 {
	switch model {
	case "linear":
		return aggregatedData.AverageLtv * 60.0
	case "quadratic":
		return aggregatedData.AverageLtv * 60.0 * 60.0
	default:
		log.Print("Unsupported model: Linear extrapolation will be used")
		return aggregatedData.AverageLtv * 60.0
	}
}
