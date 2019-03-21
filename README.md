<!--
 * @Author: hkloudou
 * @Github: https://github.com/hkloudou/
 * @LastEditors: 卢教(aven) hkloudou@gmail.com
 * @Date: 2018-07-07 02:45:59
 * @LastEditTime: 2019-03-22 03:46:58
 -->
# lddatabasekit [![Build Status](https://travis-ci.org/hkloudou/lddatabasekit.svg?branch=master)](https://travis-ci.org/hkloudou/lddatabasekit) [![Build Status](https://godoc.org/hkloudou/lddatabasekit.svg?status.svg)](https://godoc.org/github.com/hkloudou/lddatabasekit)
- lddatabasekit is a databasekit writen by golang
- deppend on mgo (gopkg.in/mgo.v2)
### download and update
`go get -u -v github.com/hkloudou/lddatabasekit`

### enviroment
``` sh
ENV DB_CONFIGFILE #default conf/database.ini(the config file path)
ENV DB_CONFIGTYPE #default ini (format of config file ex.https://github.com/astaxie/beego/tree/master/config)

ENV DB_MONGO_ADDRS  #default "" example localhost,8.1.213.1
ENV DB_MONGO_DATABASE   #deault "" example testdatabase
ENV DB_MONGO_USERNAME   #default ""
ENV DB_MONGO_PASSWORD   #default ""
```

``` go
if DB_MONGO_ADDRS == "" || DB_MONGO_DATABASE == "" {
  // read from DB_CONFIGFILE
} else {
  // use parame from env define
}
```
### example
``` go

if err:=lddatabasekit.Err();err!=nil{
  panic(err)
}

dbse, errDB := lddatabasekit.GetMongoDBSession()
if errDB != nil {
  //log.Println("database session clone error")
  return
}

db := dbse.DB("dbname").C("collection")
if err := db.Insert(bson.M{}); err != nil {
}
```
