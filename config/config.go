package config

import (
	"os"
	"strconv"
)

type Config struct {
	NumWorkers int
	NumJobs    int
}

func Load() Config {
	return Config{
		NumWorkers: getEnvAsInt("NUM_WORKERS", 3),
		NumJobs:    getEnvAsInt("NUM_JOBS", 10),
	}
}

func getEnvAsInt(name string, defaultVal int) int {
	if env, exists := os.LookupEnv(name); exists {
		if val, err := strconv.Atoi(env); err == nil {
			return val
		}
	}

	return defaultVal
}
