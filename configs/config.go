package configs

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/JamshedJ/REST-api/pkg/glog"
)

type DBConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DBName   string `json:"dbname"`
}

func InitConfig() (dsn string, err error) {
	logger := glog.NewLogger()
	var config struct {
		DB DBConfig `json:"db"`
	}

	file, err := os.Open("configs/config.json")
	if err != nil {
		logger.Fatal().Err(err).Msg("error opening 'config.json' file")
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		logger.Fatal().Err(err).Msg("error reading file contents")
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		logger.Fatal().Err(err).Msg("error unmarshaling data to config")
	}

	dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		config.DB.User, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.DBName)

	return dsn, nil
}
