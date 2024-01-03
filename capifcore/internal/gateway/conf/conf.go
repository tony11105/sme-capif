// conf.go
package conf

import (
	"os"

	yaml "gopkg.in/yaml.v2"
)

type RouteConfig struct {
	ID         string   `yaml:"id"`
	URI        string   `yaml:"uri"`
	Predicates []string `yaml:"predicates"`
}

type GatewayConfig struct {
	Routes []RouteConfig `yaml:"routes"`
}

type Config struct {
	Gateway GatewayConfig `yaml:"gateway"`
}

func LoadConfig(filename string) (*Config, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
