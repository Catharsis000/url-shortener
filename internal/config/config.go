package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

// парсинг конфига использует библиотеку клиненв
type Config struct { // создаём объект конфига, он соответствует ямл файл
	Env         string `yaml:"env" env-default:"local"`          // default путь при запуске  в "" стракт теги
	StoragePath string `yaml:"storage_path" env-required:"true"` // required кастом путь
	HTTPServer  `yaml:"http_server"`
}
type HTTPServer struct { // отдельный путь с вложенными параметрами
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	TdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func MustLoad() *Config { // фунуция, которая прочитает файл с конфига и создаст и заполнит объект конфиг, который написали
	// берем конфиг паз

	defaultConfigPath := "/shortcut_project6/url-shortener/config/local.yaml" // путь по которому будем брать локал ямл в конфиг паз
	if err := os.Setenv("CONFIG_PATH", defaultConfigPath); err != nil {
		log.Fatal("Error setting environment variable:", err)
		return nil
	}

	configPath := os.Getenv("CONFIG_PATH") // указываем где будем считывать файл с конфигом
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}
	// check if file exists проверяем существует ли файл
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exists: %s", configPath)
	}

	var cfg Config // объявляем объект конфига

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil { // считываем файл по пути, который указан
		log.Fatalf("cannot read config: %s", err)
	}
	return &cfg

}
