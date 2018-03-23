# ldDataBaseKit [![Build Status](https://travis-ci.org/hkloudou/ldDataBaseKit.svg?branch=master)](https://travis-ci.org/hkloudou/ldDataBaseKit) [![Build Status](https://godoc.org/hkloudou/ldDataBaseKit.svg?status.svg)](https://godoc.org/github.com/hkloudou/ldDataBaseKit)
- ldDataBaseKit is a databasekit writen by golang
- deppend on mgo (gopkg.in/mgo.v2)
### download and update
`go get -u -v github.com/hkloudou/ldDataBaseKit`

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

if err:=ldDataBaseKit.Err();err!=nil{
  panic(err)
}

dbse, errDB := ldDataBaseKit.GetMongoDBSession()
if errDB != nil {
  //log.Println("database session clone error")
  return
}

db := dbse.DB("dbname").C("collection")
if err := db.Insert(bson.M{}); err != nil {
}
```
