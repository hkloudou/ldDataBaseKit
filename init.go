package lddatabasekit

/*
 * @Author: hkloudou
 * @Github: https://github.com/hkloudou/
 * @LastEditors: 卢教(aven) hkloudou@gmail.com
 * @Date: 2018-07-07 02:45:59
 * @LastEditTime: 2019-03-28 22:14:38
 */

import (
	"errors"
	"log"
	"os"
	"strings"
	"time"

	mgo "gopkg.in/mgo.v2"
)

//var configfile string
//var configtype string
var (
	addrs          string
	database       string
	username       string
	password       string
	replicaSetName string
	timeout        = time.Second * 5
	ConnectError   error
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
	addrs = getEnv("DB_MONGO_ADDRS", "127.0.0.1")
	database = getEnv("DB_MONGO_DATABASE", "admin")
	username = getEnv("DB_MONGO_USERNAME", "admin")
	password = getEnv("DB_MONGO_PASSWORD", "")

	//link mongo
	if len(os.Getenv("MONGO_PORT_27017_TCP_ADDR")) > 0 {
		addrs = "mongo"
		addrs = getEnv("MONGO_PORT_27017_TCP_ADDR", addrs)
		addrs = addrs + ":" + getEnv("MONGO_PORT_27017_TCP_PORT", "27017")
		username = getEnv("MONGO_ENV_MONGO_INITDB_ROOT_USERNAME", username)
		password = getEnv("MONGO_ENV_MONGO_INITDB_ROOT_PASSWORD", password)
		database = getEnv("MONGO_ENV_MONGO_INITDB_ROOT_USERNAME", database)
	}

	replicaSetName = getEnv("DB_MONGO_REPLICASETNAME", "")

	if t, err := time.ParseDuration(getEnv("DB_MONGO_TIMEOUT", "5s")); err == nil {
		timeout = t
	}

	//if addrs not exist or database name not exist,then read from ini file

	/*
		if addrs == "" || database == "" {
			if err := readFromConfigFile(); err != nil {
				return err
			}
		}

	*/

	//log.Println(addrs, database, username, password)
	if addrs == "" {
		return errors.New("addrs not define in env and configfile")
	} else if database == "" {
		return errors.New("database not define in env and configfile")
	} else if username == "" {
		return errors.New("username not define in env and configfile")
	} else if password == "" {
		return errors.New("password not define in env and configfile")
	}
	log.Println("mongo:", "addr:", strings.Split(addrs, ","))
	dialInfo := &mgo.DialInfo{
		Addrs:     strings.Split(addrs, ","),
		Direct:    false,
		Timeout:   timeout,
		Database:  database,
		Username:  username,
		Password:  password,
		PoolLimit: 4096,
	}
	if replicaSetName != "" {
		dialInfo.ReplicaSetName = replicaSetName
	}
	if ConnectError = ReInitMongoDBSession(dialInfo, mgo.Strong); ConnectError != nil {
		//log.Println("链接主数据库失败", err)
		return ConnectError
	}
	return nil
}

//Init Init
func Init() {

}
