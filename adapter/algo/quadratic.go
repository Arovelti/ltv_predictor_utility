package algo

import "github.com/Arovelti/ltv_predictor_utility/domain/models"

func QuadraticPredictor(aggregatedData models.AggregatedData) float64 {
	return aggregatedData.AverageLtv * 60.0 * 60.0
}
