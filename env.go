package env

import (
	"os"
	"strconv"
)

func GetString(name, defaultValue string) string {
	v := os.Getenv(name)
	if v == "" {
		return defaultValue
	}
	return v
}

func GetInt(name string, defaultValue int) int {
	v := os.Getenv(name)
	if v == "" {
		return defaultValue
	}
	if i, err := strconv.Atoi(v); err != nil {
		return defaultValue
	} else {
		return i
	}
}

func GetInt64(name string, defaultValue int64) int64 {
	v := os.Getenv(name)
	if v == "" {
		return defaultValue
	}
	if i, err := strconv.ParseInt(v, 10, 64); err != nil {
		return defaultValue
	} else {
		return i
	}
}

func GetUint64(name string, defaultValue uint64) uint64 {
	v := os.Getenv(name)
	if v == "" {
		return defaultValue
	}
	if i, err := strconv.ParseUint(v, 10, 64); err != nil {
		return defaultValue
	} else {
		return i
	}
}

func GetFloat(name string, defaultValue float64) float64 {
	v := os.Getenv(name)
	if v == "" {
		return defaultValue
	}
	if f, err := strconv.ParseFloat(v, 64); err != nil {
		return defaultValue
	} else {
		return f
	}
}

func GetBool(name string, defaultValue bool) bool {
	v := os.Getenv(name)
	if v == "" {
		return defaultValue
	}
	if b, err := strconv.ParseBool(v); err != nil {
		return defaultValue
	} else {
		return b
	}
}
