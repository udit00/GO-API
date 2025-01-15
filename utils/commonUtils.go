package utils

import (
	"fmt"
	"strconv"
	"time"
)

func GetCurrentDateTimeForSql() time.Time {
	utcTime := time.Now().UTC()
	fmt.Print(utcTime)
	return utcTime
}

func GetCurrentDateTimeForSqlString() string {
	dateTimeStr := GetCurrentDateTimeForSql().Format(DATE_TIME_DEFAULT_FORMAT)
	return dateTimeStr
}

func ConvertDateTimeToGoLangTime(timeString string) (time.Time, error) {
	// Input datetime string
	// dateTimeStr := "2025-01-20 13:01:11"
	dateTimeStr := timeString

	// Correct format string
	layout := "2006-01-02 15:04:05"

	// Parse the datetime string
	parsedTime, err := time.Parse(layout, dateTimeStr)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return time.Now(), err
	}

	fmt.Println("Parsed time:", parsedTime)
	return parsedTime, nil
}

func IsNullOrEmpty(str string) bool {
	if str == "" {
		return true
		// } else if str == nil {
		// return true
	}
	return false
}

func ConvertIntToString(valueToChange int) string {
	// if(valueToChange == nil) {
	// 	return ""
	// }
	return strconv.Itoa(valueToChange)
}
