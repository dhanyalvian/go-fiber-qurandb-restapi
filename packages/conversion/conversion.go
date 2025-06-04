package conversion

import (
	"fmt"
	"strconv"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func Float64ToInt(value interface{}) int {
	var result int

	if v, ok := value.(float64); ok {
		result = int(v)
	}

	return result
}

func Float64ToInt32(value interface{}) int32 {
	var result int32

	if v, ok := value.(float64); ok {
		result = int32(v)
	}

	return result
}

func Float64ToFloat32(value interface{}) float32 {
	var result float32

	if v, ok := value.(float64); ok {
		result = float32(v)

	}

	return result
}

func StringToInt(val string) int {
	result, err := strconv.Atoi(val)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
	}

	return result
}

func FormatNumberUs(v interface{}) string {
	val := Float64ToInt32(v)
	p := message.NewPrinter(language.English)
	fn := p.Sprintf("%d", val)

	return fmt.Sprintf("$%s", fn)
}

func FormatHourMinute(minutes int) string {
	hours := minutes / 60
	remainderMinutes := minutes % 60

	return fmt.Sprintf("%dh %dm", hours, remainderMinutes)
}
