package main

import (
	"os"
)

const ENVIRONMENT string = "production"

var config_env = map[string]map[string]string{
	"development": {
		"API_URL": "localhost:8080",
		"DB_HOST": "127.0.0.1",
		"DB_PORT": "3306",
		"DB_USER": "root",
		"DB_PASS": "",
		"DB_NAME": "db_user_product",
	},
	"production": {
		"API_URL": "167.172.73.163:9090",
		"DB_HOST": "167.172.73.163",
		"DB_PORT": "3306",
		"DB_USER": "mathias2",
		"DB_PASS": "DmPUpSN8E394zRTi",
		"DB_NAME": "mathias2",
	},
}

var CONFIG = config_env[ENVIRONMENT]

func InitConfig() {
	for key := range CONFIG {
		CONFIG[key] = GetConfig(key, CONFIG[key])
		os.Setenv(key, CONFIG[key])
	}
}

func GetConfig(key string, config string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return config
}
