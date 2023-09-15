package aggregator

import (
	"github.com/Arovelti/ltv_predictor_utility/entity"
)

func AggregateData(userData *entity.UserData, aggregateLevel string) []entity.AggregatedData {
	aggregatedDataMap := make(map[string]entity.AggregatedData)

	for _, ud := range userData.Data {
		var key string

		switch aggregateLevel {
		case "country":
			key = ud.Country
		case "campaign":
			key = ud.CampaignID
		}

		aggregatedData, found := aggregatedDataMap[key]
		if !found {
			aggregatedData = entity.AggregatedData{
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

	var aggregatedDataList []entity.AggregatedData
	for _, aggregatedData := range aggregatedDataMap {
		aggregatedDataList = append(aggregatedDataList, aggregatedData)
	}

	return aggregatedDataList
}
