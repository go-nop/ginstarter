package configurator

import "os"

// defaultIfEmpty checks if the value is empty and returns the default value if it is.
func defaultIfEmpty[T comparable](value, defaultValue T) T {
	var zeroValue T
	if value == zeroValue {
		return defaultValue
	}
	return value
}

// getenv is a helper function that retrieves the value of an environment variable.
// If the environment variable is not set, it returns an empty string.
func getenv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	if val := os.Getenv(key); val != "" {
		return val
	}

	return ""
}
