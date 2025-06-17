package conversion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloat64ToInt(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected int
	}{
		{
			name:     "Valid float64 - no decimal",
			input:    float64(42),
			expected: 42,
		},
		{
			name:     "Valid float64 - with decimal",
			input:    float64(42.9),
			expected: 42, // truncated
		},
		{
			name:     "Invalid type - string",
			input:    "42.0",
			expected: 0,
		},
		{
			name:     "Invalid type - int",
			input:    42,
			expected: 0,
		},
		{
			name:     "Nil input",
			input:    nil,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Float64ToInt(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFloat64ToInt32(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected int32
	}{
		{
			name:     "Valid float64 - integer value",
			input:    float64(100),
			expected: int32(100),
		},
		{
			name:     "Valid float64 - with decimals",
			input:    float64(123.456),
			expected: int32(123), // decimals truncated
		},
		{
			name:     "Invalid type - string",
			input:    "123.456",
			expected: 0,
		},
		{
			name:     "Invalid type - int",
			input:    123,
			expected: 0,
		},
		{
			name:     "Nil input",
			input:    nil,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Float64ToInt32(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFloat64ToFloat32(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected float32
	}{
		{
			name:     "Valid float64 - integer value",
			input:    float64(100),
			expected: float32(100),
		},
		{
			name:     "Valid float64 - with decimals",
			input:    float64(123.456),
			expected: float32(123.456),
		},
		{
			name:     "Invalid type - string",
			input:    "123.456",
			expected: 0,
		},
		{
			name:     "Invalid type - int",
			input:    123,
			expected: 0,
		},
		{
			name:     "Nil input",
			input:    nil,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Float64ToFloat32(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestStringToInt(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Valid number string",
			input:    "123",
			expected: 123,
		},
		{
			name:     "Negative number string",
			input:    "-45",
			expected: -45,
		},
		{
			name:     "Zero",
			input:    "0",
			expected: 0,
		},
		{
			name:     "Invalid string - letters",
			input:    "abc",
			expected: 0,
		},
		{
			name:     "Invalid string - alphanumeric",
			input:    "123abc",
			expected: 0,
		},
		{
			name:     "Empty string",
			input:    "",
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StringToInt(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFormatNumberUs(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{
			name:     "Positive float",
			input:    float64(1234.56),
			expected: "$1,234",
		},
		{
			name:     "Zero",
			input:    float64(0),
			expected: "$0",
		},
		{
			name:     "Negative float",
			input:    float64(-987.65),
			expected: "$-987",
		},
		{
			name:     "Non-float input",
			input:    "abc",
			expected: "$0",
		},
		{
			name:     "Nil input",
			input:    nil,
			expected: "$0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatNumberUs(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFormatHourMinute(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected string
	}{
		{
			name:     "Zero minutes",
			input:    0,
			expected: "0h 0m",
		},
		{
			name:     "Less than an hour",
			input:    45,
			expected: "0h 45m",
		},
		{
			name:     "Exactly one hour",
			input:    60,
			expected: "1h 0m",
		},
		{
			name:     "More than an hour",
			input:    135,
			expected: "2h 15m",
		},
		{
			name:     "Multiple hours",
			input:    480,
			expected: "8h 0m",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatHourMinute(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
