package chart

import (
	"fmt"
	"strconv"
	"time"
)

// ValueFormatter is a function that takes a value and produces a string.
type ValueFormatter func(v interface{}) string

// DateValueFormatter is a ValueFormatter for timestamps.
func DateValueFormatter(v interface{}) string {
	return FormatTime(v, DefaultDateFormat)
}

// DateHourValueFormatter is a ValueFormatter for timestamps.
func DateHourValueFormatter(v interface{}) string {
	return FormatTime(v, DefaultDateHourFormat)
}

// DateMinuteValueFormatter is a ValueFormatter for timestamps.
func DateMinuteValueFormatter(v interface{}) string {
	return FormatTime(v, DefaultDateMinuteFormat)
}

// DateTimeValueFormatter is a ValueFormatter for timestamps.
func DateTimeValueFormatter(v interface{}) string {
	return FormatTime(v, DefaultDateTimeFormat)
}

// HourValueFormatter is a ValueFormatter for timestamps.
func HourValueFormatter(v interface{}) string {
	return FormatTime(v, DefaultHourFormat)
}

// MinuteValueFormatter is a ValueFormatter for timestamps.
func MinuteValueFormatter(v interface{}) string {
	return FormatTime(v, DefaultMinuteFormat)
}

// TimeValueFormatter is a ValueFormatter for timestamps.
func TimeValueFormatter(v interface{}) string {
	return FormatTime(v, DefaultTimeFormat)
}

// TimeValueFormatterWithFormat returns a time formatter with a given format.
func TimeValueFormatterWithFormat(format string) ValueFormatter {
	return func(v interface{}) string {
		return FormatTime(v, format)
	}
}

// TimeValueFormatterWithFormat is a ValueFormatter for timestamps with a given format.
func FormatTime(v interface{}, dateFormat string) string {
	if typed, isTyped := v.(time.Time); isTyped {
		return typed.Format(dateFormat)
	}
	if typed, isTyped := v.(int64); isTyped {
		return time.Unix(0, typed).Format(dateFormat)
	}
	if typed, isTyped := v.(float64); isTyped {
		return time.Unix(0, int64(typed)).Format(dateFormat)
	}
	return ""
}

// IntValueFormatter is a ValueFormatter for float64.
func IntValueFormatter(v interface{}) string {
	switch v.(type) {
	case int:
		return strconv.Itoa(v.(int))
	case int64:
		return strconv.FormatInt(v.(int64), 10)
	case float32:
		return strconv.FormatInt(int64(v.(float32)), 10)
	case float64:
		return strconv.FormatInt(int64(v.(float64)), 10)
	default:
		return ""
	}
}

// FloatValueFormatter is a ValueFormatter for float64.
func FloatValueFormatter(v interface{}) string {
	return FloatValueFormatterWithFormat(v, DefaultFloatFormat)
}

// PercentValueFormatter is a formatter for percent values.
// NOTE: it normalizes the values, i.e. multiplies by 100.0.
func PercentValueFormatter(v interface{}) string {
	if typed, isTyped := v.(float64); isTyped {
		return FloatValueFormatterWithFormat(typed*100.0, DefaultPercentValueFormat)
	}
	return ""
}

// FloatValueFormatterWithFormat is a ValueFormatter for float64 with a given format.
func FloatValueFormatterWithFormat(v interface{}, floatFormat string) string {
	if typed, isTyped := v.(int); isTyped {
		return fmt.Sprintf(floatFormat, float64(typed))
	}
	if typed, isTyped := v.(int64); isTyped {
		return fmt.Sprintf(floatFormat, float64(typed))
	}
	if typed, isTyped := v.(float32); isTyped {
		return fmt.Sprintf(floatFormat, typed)
	}
	if typed, isTyped := v.(float64); isTyped {
		return fmt.Sprintf(floatFormat, typed)
	}
	return ""
}

// KValueFormatter is a formatter for K values.
func KValueFormatter(k float64, vf ValueFormatter) ValueFormatter {
	return func(v interface{}) string {
		return fmt.Sprintf("%0.0fσ %s", k, vf(v))
	}
}
