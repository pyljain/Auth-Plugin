package config

import "gopkg.in/yaml.v3"

type Config struct {
	Auth ConfigAuth `yaml:"auth"`
}

type ConfigAuth struct {
	AuthURL      string `yaml:"auth_url"`
	TokenURL     string `yaml:"token_url"`
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
	RedirectURI  string `yaml:"redirect_uri"`
}

func LoadConfig(config []byte) (*Config, error) {
	c := Config{}

	err := yaml.Unmarshal(config, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
