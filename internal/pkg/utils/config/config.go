package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

func LoadConfigs(patterns []string, conf interface{}) error {
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	for i := range patterns {
		matches, err := filepath.Glob(patterns[i])
		if err != nil {
			return err
		}

		for i2 := range matches {
			if _, err := os.Stat(matches[i2]); err == nil {
				viper.SetConfigFile(matches[i2])
				if err := viper.MergeInConfig(); err != nil {
					return err
				}
				viper.WatchConfig()
			}
		}
	}

	return viper.Unmarshal(conf)
}
