package parser

import (
	"github.com/spf13/viper"
	"github.com/step/saurontypes"
)

func ParseConfig(viper viper.Viper) saurontypes.SauronConfig {
	sauronConfig  := saurontypes.SauronConfig{}
	viper.Unmarshal(&sauronConfig)
	return sauronConfig
}