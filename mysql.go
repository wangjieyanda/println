package println

import (
	"fmt"
	"log"

	"println/orm/jmoiron/sqlx"
	_ "println/orm/mysql"
	//"github.com/jmoiron/sqlx"
)

// "println/orm/jmoiron/sqlx"

type mysqlDB struct {
	// *sqlx.DB
	Limit   string
	Offset  string
	Where   map[string]interface{}
	Order   string
	GroupBy string
}

func DbConnect(key string) (db *sqlx.DB, err error) {
	if _, ok := ConfigMap[key]; ok {
		dirver := ConfigMap[key].Driver
		user := ConfigMap[key].User
		passWord := ConfigMap[key].PassWord
		dsn := ConfigMap[key].Dsn
		dbName := ConfigMap[key].DbName
		if len(dirver) <= 0 {
			log.Fatal("db memery err")
			return nil, err
		}
		if len(user) <= 0 {
			log.Fatal("db user err")
			return nil, err
		}
		if len(passWord) <= 0 {
			log.Fatal("db password err")
			return nil, err
		}
		if len(dsn) <= 0 {
			log.Fatal("db dsn err")
			return nil, err
		}
		if len(dbName) <= 0 {
			log.Fatal("db dbname err")
			return nil, err
		}
		url := "%s:%s@tcp(%s)/%s?charset=utf8"
		url = fmt.Sprintf(url, user, passWord, dsn, dbName)
		db, err := sqlx.Open(dirver, url)
		if err != nil {
			log.Fatal(err)
		}
		return db, err
	}
	return nil, err
}
