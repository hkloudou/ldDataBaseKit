package ldDataBaseKit

import (
	"errors"
	"os"
	"path"
	"strings"
	"time"

	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/config/env"
	mgo "gopkg.in/mgo.v2"
)

var configfile string
var configtype string
var (
	addrs    string
	database string
	username string
	password string
)

var mErr error

//GetDataBaseName GetDataBaseName
func GetDataBaseName() string {
	return database
}

// Err get err
func Err() error {
	return mErr
}

func init() {
	if err := initConfig(); err != nil {
		mErr = err
		return
	}
}

func initConfig() error {
	// readfrom ini
	configfile = env.Get("DB_CONFIGFILE", "conf/database.ini")
	configtype = env.Get("DB_CONFIGTYPE", "ini")
	addrs = env.Get("DB_MONGO_ADDRS", "")
	database = env.Get("DB_MONGO_DATABASE", "")
	username = env.Get("DB_MONGO_USERNAME", "")
	password = env.Get("DB_MONGO_PASSWORD", "")

	//if addrs not exist or database name not exist,then read from ini file
	if addrs == "" || database == "" {
		if err := readFromConfigFile(); err != nil {
			return err
		}
	}
	if addrs == "" {
		return errors.New("addrs not define in env and configfile")
	} else if database == "" {
		return errors.New("database not define in env and configfile")
	} else if username == "" {
		return errors.New("username not define in env and configfile")
	} else if password == "" {
		return errors.New("password not define in env and configfile")
	}
	dialInfo := &mgo.DialInfo{
		Addrs:     strings.Split(addrs, ","),
		Direct:    false,
		Timeout:   time.Second * 5,
		Database:  database,
		Username:  username,
		Password:  password,
		PoolLimit: 4096,
	}
	if err := InitMongoDBSession(dialInfo, mgo.Strong); err != nil {
		//log.Println("链接主数据库失败", err)
		return err
	}
	return nil
}

func readFromConfigFile() error {
	// makesure path exist
	os.MkdirAll(path.Dir(configfile), 0777)
	iniconf, err := config.NewConfig(configtype, configfile)
	if err != nil {
		return errors.New("config file " + configfile + " read error")
	}
	addrs = iniconf.String("mongo::addrs")
	database = iniconf.String("mongo::database")
	username = iniconf.String("mongo::username")
	password = iniconf.String("mongo::password")
	return nil
}

//Init Init
func Init() {

}
