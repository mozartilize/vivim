package vivim

import "github.com/spf13/viper"

func readConfig(name string, defaults map[string]interface{}) *viper.Viper {
	v := viper.New()
	for key, value := range defaults {
		v.SetDefault(key, value)
	}
	v.SetConfigName(name)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	return v
}
