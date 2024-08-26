package configuration

import (
	"booking/internal/infrastrucuture/logging"

	"github.com/spf13/viper"
)

func LoadConfig(logger *logging.ZapLogger) (*Configs, error) {
	var err error
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./configs")

		viper.SetDefault("server.port", 8080)

		if err = viper.ReadInConfig(); err != nil {
			logger.Error("Read config error", map[string]interface{}{})
			return
		}

		if err = viper.Unmarshal(&config); err != nil {
			logger.Error("Can't unmarshal data to struct", map[string]interface{}{})
			return
		}

		logger.Info("Succesfully loaded configs", map[string]interface{}{"config_file": "config.yaml"})
	})
	return config, nil
}
