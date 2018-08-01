package println

import (
	"log"
	"println/orm/config"
)

type DbConfig struct {
	Driver   string
	User     string
	PassWord string
	ShowSQL  bool
	count    string
	DbName   string
	Path     string
	Dsn      string
	Config   *config.Config
}

var ConfigMap = map[string]*DbConfig{}

func ReadConfig(filePath string) map[string]*DbConfig {
	configTmp, err := config.ReadDefault(filePath)
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range configTmp.CacheData {
		if _, ok := ConfigMap[k]; !ok {
			ConfigMap[k] = &DbConfig{}
		}
		ConfigMap[k].Driver = v["driver"].V
		ConfigMap[k].User = v["user"].V
		ConfigMap[k].PassWord = v["pass"].V
		ConfigMap[k].Dsn = v["dsn"].V
		ConfigMap[k].DbName = v["db_name"].V
		ConfigMap[k].count = v["count"].V
		ConfigMap[k].Path = v["path"].V
	}

	return ConfigMap
}
