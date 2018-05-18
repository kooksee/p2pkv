package config

import (
	"time"

	"github.com/patrickmn/go-cache"
)

func GetCfg() *Config {
	once.Do(func() {
		instance = &Config{
			UdpPort:  8080,
			UdpHost:  "0.0.0.0",
			HttpHost: "0.0.0.0",
			HttpPort: 8081,
			Debug:    true,
			LogLevel: "debug",
			Cache:    cache.New(time.Minute, 5*time.Minute),
			Home:     "kdata",
		}
	})
	return instance
}
