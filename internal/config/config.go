package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `yaml:"env" env-default:"local" env-required:"true"`
	HTTPServer  `yaml:"HTTPServer"`
	DBServer    `yaml:"DBServer"`
	S3Config    `yaml:"S3Config"`
	RedisConfig `yaml:"RedisConfig"`
}

type HTTPServer struct {
	Host        string        `yaml:"Host" env-default:"localhost"`
	Port        string        `yaml:"Port" env-default:"8080"`
	Timeout     time.Duration `yaml:"Timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"IdleTimeout" env-default:"60s"`
}

type DBServer struct {
	Host     string `yaml:"Host" env-default:"localhost"`
	Port     int    `yaml:"Port" env-default:"5432"`
	Username string `yaml:"Username" env-required:"true"`
	DBName   string `yaml:"DBName" env-required:"true"`
	SSLMode  string `yaml:"SSLMode" env-default:"disable"`
	Password string `yaml:"Password" env-required:"true"`
}
type S3Config struct {
	Region          string `yaml:"Region" env-required:"true"`
	AccessKeyID     string `yaml:"AccessKeyID" env-required:"true"`
	SecretAccessKey string `yaml:"SecretAccessKey" env-required:"true"`
	Endpoint        string `yaml:"Endpoint"`
	DisableSSL      bool   `yaml:"DisableSSL"`
	ForcePathStyle  bool   `yaml:"ForcePathStyle"`
}

type RedisConfig struct {
	Addr        string        `yaml:"Addr" env-required:"true"`
	Password    string        `yaml:"Password" env-required:"true"`
	User        string        `yaml:"User" env-required:"true"`
	DB          int           `yaml:"DB" env-default:"0"`
	MaxRetries  int           `yaml:"MaxRetries" env-default:"5"`
	DialTimeout time.Duration `yaml:"DialTimeout" env-default:"10s"`
	Timeout     time.Duration `yaml:"Timeout" env-default:"5s"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
