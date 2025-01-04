package configurator

import (
	"strconv"
)

// GetEnv is a helper function that retrieves the value of an environment variable.
// If the environment variable is not set, it returns the default value.
func GetEnv(key, defaultValue string) string {
	value := defaultIfEmpty(getenv(key), defaultValue)
	return value
}

// GetEnvAsInt is a helper function that retrieves the value of an environment variable as an integer.
// If the environment variable is not set or cannot be converted to an integer, it returns the default value.
func GetEnvAsInt(key string, defaultValue int) int {
	value := GetEnv(key, strconv.Itoa(defaultValue))
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return intValue
}

// GetEnvAsBool is a helper function that retrieves the value of an environment variable as a boolean.
// If the environment variable is not set or cannot be converted to a boolean, it returns the default value.
func GetEnvAsBool(key string, defaultValue bool) bool {
	value := GetEnv(key, strconv.FormatBool(defaultValue))
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}
	return boolValue
}
