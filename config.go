package vivim

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func readConfig(path string, defaults map[string]interface{}) *viper.Viper {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	v := viper.New()
	for key, value := range defaults {
		v.SetDefault(key, value)
	}
	v.SetEnvPrefix("vivim")
	v.SetConfigFile(path)
	v.AllowEmptyEnv(true)
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			logger.Error("Error while reading config file ", zap.Error(err))
		} else {
			// Config file was found but another error was produced
			logger.Fatal("Error while reading config file ", zap.Error(err))
		}
	}
	return v
}
