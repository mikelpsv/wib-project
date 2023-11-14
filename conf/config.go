package conf

import (
	"github.com/joho/godotenv"
	app "github.com/mlplabs/app-utils"
	"os"
	"strconv"
)

type AppCfg struct {
	EnvProduction bool
	EnvLocal      bool
	AppAddr       string
	AppPort       string
	DbHost        string
	DbPort        string
	DbName        string
	DbUser        string
	DbPassword    string
}

var Cfg AppCfg

func ReadEnv() {
	err := godotenv.Load()
	if err != nil {
		app.Log.Error.Println("error reading .env file")
	}

	Cfg.EnvProduction = os.Getenv("ENV") == "production"
	Cfg.EnvLocal = os.Getenv("ENV") == "local"

	Cfg.AppAddr = os.Getenv("APP_ADDR")
	Cfg.AppPort = os.Getenv("APP_PORT")

	Cfg.DbHost = os.Getenv("DB_HOST")
	Cfg.DbPort = os.Getenv("DB_PORT")
	Cfg.DbName = os.Getenv("DB_NAME")
	Cfg.DbUser = os.Getenv("DB_USER")
	Cfg.DbPassword = os.Getenv("DB_PASS")
}

func ReadIntValue(envString string, defVal int) (int, error) {
	val, err := strconv.Atoi(envString)
	if err != nil {
		return defVal, err
	}
	return val, nil
}
