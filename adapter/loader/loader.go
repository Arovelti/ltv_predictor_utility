package loader

import (
	"embed"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/Arovelti/ltv_predictor_utility/domain/models"
	"github.com/Arovelti/ltv_predictor_utility/pkg/helper"
)

const (
	jsonType string = "json"
	csvType  string = "csv"
)

//go:embed test_data/*
var content embed.FS

type Loader func(source string) (*models.UserData, error)

func (l Loader) LoadData(source string) (*models.UserData, error) {
	return l(source)
}

func NewLoader(source string) Loader {
	return func(source string) (*models.UserData, error) {
		f := fmt.Sprintf("test_data/%s", source)

		reader, err := content.ReadFile(f)
		if err != nil {
			fmt.Println(err)
		}

		userData := &models.UserData{}

		switch source {
		case "test_data.json":
			jsonData, err := loadJson(reader)
			if err != nil {
				return nil, err
			}

			userData.Data = jsonData
			userData.DataType = jsonType

		case "test_data.csv":
			s := helper.SliceByteToString(reader)
			reader := csv.NewReader(strings.NewReader(s))

			csvData, err := loadCSV(reader)
			if err != nil {
				return nil, err
			}

			userData.Data = csvData
			userData.DataType = csvType

		default:
			return nil, errors.New("unsupported source data format")
		}

		return userData, nil
	}
}

func loadJson(reader []byte) ([]models.LoadUserData, error) {
	var data []models.LoadUserData

	if err := json.Unmarshal(reader, &data); err != nil {
		return nil, fmt.Errorf("unable to unmarshal sourse data: %v", err)
	}

	return data, nil
}

func loadCSV(reader *csv.Reader) ([]models.LoadUserData, error) {
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("unable to read records from csv file: %v", err)
	}

	var data []models.LoadUserData

	for _, record := range records[1:] {

		userID := record[0]
		campaignID := record[1]
		country := record[2]
		ltv1, _ := strconv.ParseFloat(record[3], 64)
		ltv2, _ := strconv.ParseFloat(record[4], 64)
		ltv3, _ := strconv.ParseFloat(record[5], 64)
		ltv4, _ := strconv.ParseFloat(record[6], 64)
		ltv5, _ := strconv.ParseFloat(record[7], 64)
		ltv6, _ := strconv.ParseFloat(record[8], 64)
		ltv7, _ := strconv.ParseFloat(record[9], 64)

		user := models.LoadUserData{
			UserID:     userID,
			CampaignID: campaignID,
			Country:    country,
			LTV1:       ltv1,
			LTV2:       ltv2,
			LTV3:       ltv3,
			LTV4:       ltv4,
			LTV5:       ltv5,
			LTV6:       ltv6,
			LTV7:       ltv7,
		}

		data = append(data, user)
	}

	return data, nil
}
