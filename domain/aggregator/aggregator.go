package aggregator

import "github.com/Arovelti/ltv_predictor_utility/domain/models"

func AggregateData(userData *models.UserData, level string) []models.AggregatedData {
	aggregatedDataMap := make(map[string]models.AggregatedData)

	for _, ud := range userData.Data {
		var key string

		switch level {
		case "country":
			key = ud.Country
		case "campaign":
			key = ud.CampaignID
		}

		aggregatedData, found := aggregatedDataMap[key]
		if !found {
			aggregatedData = models.AggregatedData{
				Key:        key,
				TotalLTV:   0,
				UserCount:  0,
				AverageLtv: 0,
				MaxLtv:     0,
				LTV60thDay: 0,
			}
		}

		aggregatedData.TotalLTV += ud.LTV7
		aggregatedData.UserCount++
		aggregatedData.AverageLtv = aggregatedData.TotalLTV / float64(aggregatedData.UserCount)

		if ud.LTV7 > aggregatedData.MaxLtv {
			aggregatedData.MaxLtv = ud.LTV7
		}

		aggregatedDataMap[key] = aggregatedData
	}

	var aggregatedDataList []models.AggregatedData
	for _, aggregatedData := range aggregatedDataMap {
		aggregatedDataList = append(aggregatedDataList, aggregatedData)
	}

	return aggregatedDataList
}
