package config

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

type EnvConfigProvider struct{}

func (p *EnvConfigProvider) Load() (*Config, error) {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	p.BindEnvs(Config{})
	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}

	if reflect.DeepEqual(c, Config{}) {
		viper.SetConfigFile(".env")

		if err := viper.ReadInConfig(); err != nil {
			return nil, err
		}
		if err := viper.Unmarshal(&c); err != nil {
			return nil, err
		}
		if reflect.DeepEqual(c, Config{}) {
			return nil, fmt.Errorf("config is empty")
		}
	}

	return &c, nil
}

// https://github.com/spf13/viper/issues/188#issuecomment-399884438
func (p *EnvConfigProvider) BindEnvs(iface interface{}, parts ...string) {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)
	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)
		tv, ok := t.Tag.Lookup("mapstructure")
		if !ok {
			continue
		}
		switch v.Kind() {
		case reflect.Struct:
			p.BindEnvs(v.Interface(), append(parts, tv)...)
		default:
			viper.BindEnv(strings.Join(append(parts, tv), "."))
		}
	}
}
