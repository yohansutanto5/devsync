package config

import (
	"encoding/json"
	"os"
)

type db struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Vendor      string `json:"vendor"`
	Host        string `json:"host"`
	Port        string `json:"port"`
	Database    string `json:"database"`
	Schema      string `json:"schema"`
	MaxOpenConn int    `json:"max_open_conn"`
	MaxIdleConn int    `json:"max_idle_conn"`
}

type redis struct {
	Address  string `json:"address"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

type dir struct {
	Migration string `json:"sql"`
}

type cron struct {
	Task1 string `json:"task_1"`
	Task2 string `json:"task_2"`
}
type dbslave struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Database string `json:"database"`
	Schema   string `json:"schema"`
}

type auth struct {
	Token string `json:"token"`
	Key   string `json:"key"`
}

type url struct {
	Jira       string `json:"jira"`
	Confluence string `json:"confluence"`
	Jenkins    string `json:"jenkins"`
}

type kafka struct {
	Hostname      string `json:"hostname"`
	User          string `json:"user"    `
	Password      string `json:"password"`
	Port          string `json:"port" `
	AuthMethod    string `json:"authMethod"`
	ConsumerGroup string `json:"consumerGroup"`
	ApprovalTopic string `json:"approvalTopic"`
}

type mode struct {
	Mysql      bool   `json:"mysql"`
	Postgresql bool   `json:"postgresql"`
	Redis      bool   `json:"redis"`
	Kafka      bool   `json:"kafka"`
	Slave      bool   `json:"slave"`
	Run        string `json:"run"`
}

type Configuration struct {
	Db      db      `json:"db"`
	Dbslave dbslave `json:"dbslave"`
	Auth    auth    `json:"auth"`
	URL     url     `json:"url"`
	Redis   redis   `json:"redis"`
	Cron    cron    `json:"cron"`
	Mode    mode    `json:"mode"`
	Dir     dir
	Kafka   kafka `json:"kafka"`
}

func Load(env string) Configuration {

	var configFile string

	switch env {
	case "dev":
		configFile = os.Getenv("config") + "/config_dev.json"
	case "prd":
		configFile = os.Getenv("config")
	case "test":
		configFile = os.Getenv("config") + "/config_test.json"
	default:
		panic("Input ENV")
	}

	file, err := os.Open(configFile)

	if err != nil {
		panic("Error opening configuration file: " + err.Error())
	}
	defer file.Close()

	// Parse the JSON configuration
	var config Configuration
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		panic("Error decoding configuration: " + err.Error())
	}

	return config
}
