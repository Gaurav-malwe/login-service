package utils

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Gaurav-malwe/login-service/config"
	"github.com/Gaurav-malwe/login-service/internal/constants"
	log "github.com/Gaurav-malwe/login-service/utils/logging"
	"go.mongodb.org/mongo-driver/bson"
)

func GetGOProfile() string {
	var env = os.Getenv("GO_PROFILE")
	if env == "" {
		env = constants.GoProfile_Local
	}
	return env
}

func ToBSON(obj interface{}) (bson.M, error) {
	res := bson.M{}

	if obj == nil {
		return res, nil
	}

	b, err := bson.Marshal(obj)
	if err != nil {
		return nil, err
	}

	err = bson.Unmarshal(b, &res)
	if err != nil {
		return nil, err
	}

	return res, err
}

// Add utility functions here
func StringToInt(intString string) int {
	res, err := strconv.Atoi(intString)
	if err != nil {
		log.Error("Utils::StringToInt::Error: ", err)
		return 0
	}
	return res
}

func StringToFloat(floatString string) float64 {
	res, err := strconv.ParseFloat(floatString, 64)
	if err != nil {
		log.Error("Utils::StringToFloat::Error: ", err)
		return 0
	}
	return res
}

func RateFromDateToTime(rateFromDate string) time.Time {
	layout := "1/2/2006"
	parsedTime, err := time.Parse(layout, rateFromDate)
	if err != nil {
		log.Error("Utils::DateToTime::Error", err)
		return time.Time{}
	}
	return parsedTime
}

func ParseDate(date string) (time.Time, bool) {

	parsedTime, err := time.Parse(time.RFC3339, date)
	if err == nil {

		return parsedTime, true
	}
	parsedTime, err = time.Parse("20060102T150405Z", date)
	if err == nil {
		return parsedTime, true
	}
	log.Error("Utils::ParseDate::Error", err)
	return parsedTime, false
}

func NavDateToTime(navDate string) time.Time {
	layout := "02-01-2006"
	parsedTime, err := time.Parse(layout, navDate)
	if err != nil {
		log.Error("Utils::DateToTime::Error", err)
		return time.Time{}
	}
	return parsedTime
}

func TrimSpace(str string) string {
	return strings.TrimSpace(str)
}

func IsLeap(date time.Time) bool {
	year := date.Year()

	// Leap year logic
	if year%4 == 0 {
		// If divisible by 4
		if year%100 != 0 || (year%100 == 0 && year%400 == 0) {
			// Not divisible by 100 unless divisible by 400
			return true
		}
	}

	return false
}

func GetCurrentTimeForDB() time.Time {
	return time.Now().UTC()
}

func GetEnvOrDefaultInt(key string, defaultValue int64) int64 {
	cfg := config.GetConfig()
	value := cfg.GetInt64(key)
	if value == 0 {
		value = defaultValue
	}
	return value
}

func StringToInt64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

func CastValue(value string, valueType string) (interface{}, error) {
	switch valueType {
	case "number":
		res, err := strconv.ParseFloat(value, 64)
		if err != nil {
			log.Error("Utils::CastValue::Error:ParseFloat ", err)
			return nil, err
		}
		return res, nil
	case "datetime":
		return ParseDateToTime(value)
	case "bool":
		return value == "true", nil
	default:
		return value, nil
	}
}

func ParseDateToTime(date string) (time.Time, error) {
	layout := "2006-01-02T15:04:05Z"
	parsedTime, err := time.Parse(layout, date)
	if err != nil {
		log.Error("Utils::DateToTime::Error", err)
		return time.Time{}, err
	}
	return parsedTime, nil
}
