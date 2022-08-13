package configmanager

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
)

type AppConfig struct {
	Port           string `json:"port"`
	LogFilePath    string `json:"logFilePath"`
	Mysql          Mysql  `json:"mysql"`
	APIEp          string `json:"apiep"`
	PaymentLinkURL string `json:"paymentLinkURL"`
	XApiVersion    string `json:"x-Api-Version"`
	XClientId      string `json:"x-Client-Id"`
	XClientSecret  string `json:"x-Client-Secret"`
}

type Mysql struct {
	DriverName     string `json:"driverName"`
	DataSourceName string `json:"dataSourceName"`
	DB             *sql.DB
}

var (
	AppConf *AppConfig
)

func InitAppConfig(file string) error {
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(raw, &AppConf); err != nil {
		return err
	}
	return nil
}

func GetAppConfig() *AppConfig {
	return AppConf
}
