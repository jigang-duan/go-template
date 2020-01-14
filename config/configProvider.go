package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Provider interface {
	GetString(key string) string
	GetInt(key string) int
	GetBool(key string) bool
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringSlice(key string) []string
	Get(key string) interface{}
	Set(key string, value interface{})
	IsSet(key string) bool
}

func FromConfigString(config, configType string) (Provider, error) {
	v := viper.New()
	v.SetConfigType(configType)
	if err := v.ReadConfig(strings.NewReader(config)); err != nil {
		return nil, err
	}
	return v, nil
}
