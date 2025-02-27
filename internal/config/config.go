package config

// var (
// 	JwtSecret = []byte("your_secret_key")
// 	DSN = "host=87.242.100.33 user=root password=reg2025 dbname=fullstack port=5432 sslmode=disable"
// )

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env       string        `yaml:"env" env-default:"local"`
	APP       APPConfig     `yaml:"app"`
	TokenTTL  time.Duration `yaml:"token_ttl" env-default:"1h"`
	JwtSecret string        `yaml:"JwtSecret"`
	DSN       DSNConfig     `yaml:"dsn"`
}

type APPConfig struct {
	Port    string        `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

type DSNConfig struct {
	Host    string `yaml:"host"`
	User    string `yaml:"user"`
	Pass    string `yaml:"password"`
	DBName  string `yaml:"dbname"`
	Port    string `yaml:"port"`
	SSLMode string `yaml:"sslmode"`
}

func MustLoad() *Config {
	configPath := fetchConfigPath()
	if configPath == "" {
		panic("config path is empty")
	}

	return MustLoadPath(configPath)
}

func MustLoadPath(configPath string) *Config {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
