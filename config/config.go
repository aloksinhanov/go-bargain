package config

import (
	"encoding/json"
	"log"
	"os"
	"reflect"
)

type Config struct {
	Server ServerConfig
}

type ServerConfig struct {
	ListenURL  string `json:"listenURL"`
	URLPrefix  string `json:"prefixURL"`
	APIVersion string `json:"APIVersion"`
}

func Load() *Config {
	cfg := &Config{}
	// serverCfg := os.Getenv("BARGAIN_SERVER_CONFIG")
	// if serverCfg == "" {
	// 	log.Fatalf("Invalid BARGAIN_SERVER_CONFIG: %v", serverCfg)
	// }
	// err := json.Unmarshal([]byte(serverCfg), &cfg.Server)
	// if err != nil {
	// 	log.Fatalf("Failed to unmarshal BARGAIN_SERVER_CONFIG: %v", serverCfg)
	// }

	getEnv("BARGAIN_SERVER_CONFIG", &cfg.Server)
	return cfg
}

//obj must be address of the variable where the config needs to be unmarshalled into
func getEnv(envKey string, obj interface{}) {
	if reflect.ValueOf(obj).Kind() != reflect.Ptr {
		log.Fatalf("obj t%v for envKey not a pointer", envKey)
	}

	serverCfg := os.Getenv(envKey)
	if serverCfg == "" {
		log.Fatalf("Invalid %v: %v", envKey, serverCfg)
	}

	err := json.Unmarshal([]byte(serverCfg), obj)
	if err != nil {
		log.Fatalf("Failed to unmarshal BARGAIN_SERVER_CONFIG: %v", serverCfg)
	}
}
