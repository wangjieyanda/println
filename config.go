package println

import (
	"log"
	"println/logs"
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

func ReadConfig(filePath string) error {
	configTmp, err := config.ReadDefault(filePath)
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range configTmp.CacheData {
		if _, ok := ConfigMap[k]; !ok {
			ConfigMap[k] = &DbConfig{}
		}
		if len(v) != 0 {
			if _, ok := v["driver"]; ok {
				ConfigMap[k].Driver = v["driver"].V
			}
			if _, ok := v["user"]; ok {
				ConfigMap[k].User = v["user"].V
			}
			if _, ok := v["pass"]; ok {
				ConfigMap[k].PassWord = v["pass"].V
			}
			if _, ok := v["dsn"]; ok {
				ConfigMap[k].Dsn = v["dsn"].V
			}
			if _, ok := v["db_name"]; ok {
				ConfigMap[k].DbName = v["db_name"].V
			}
			if _, ok := v["count"]; ok {
				ConfigMap[k].count = v["count"].V
			}
			if _, ok := v["path"]; ok {
				ConfigMap[k].Path = v["path"].V
			}
		}
	}
	return nil
}

func setLogPath() {
	logConfig := new(logs.Logger)
	if _, ok := ConfigMap["log"]; ok {
		logConfig.LogPath = ConfigMap["log"].Path
	}
	
}
