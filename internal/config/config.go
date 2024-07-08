package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		Env string `yaml:"env"`

		HTTPServer `yaml:"http"`
		Postgres   `yaml:"postgres"`

		Clients   ClientsConfig `yaml:"clients"`
		AppSecret string        `yaml:"app_secret" env-requred:"true" env:"APP_SECRET"`
	}

	HTTPServer struct {
		Address     string        `yaml:"address" env-default:"localhost:8080"`
		Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
		IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
	}

	Postgres struct {
		Url string `yaml:"url"`
	}

	ClientsConfig struct {
		SSO Client `yaml: "sso"`
	}

	Client struct {
		Address      string        `yaml:"address"`
		Timeout      time.Duration `yaml:"timeout"`
		RetriesCount int           `yaml:"retriesCount"`
		Insecure     bool          `yaml:"insecure"`
	}
)

func MustLoad() *Config {
	path := fetchConfigPath()

	if path == "" {
		panic("config path is empty")
	}

	return MustLoadPath(path)
}

func MustLoadPath(configPath string) *Config {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		panic("failed to read config" + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var res string

	// --config="path/to/config.yaml"
	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
