# ldDataBaseKit
- ldDataBaseKit is a databasekit writen by golang
- deppend on mgo
### download and update
`go get -u -v github.com/hkloudou/ldDataBaseKit`

### example
``` go
dbse, errDB := ldDataBaseKit.GetMongoDBSession()
if errDB != nil {
  //log.Println("database session clone error")
  return
}

db := dbse.DB("dbname").C("collection")
if err := db.Insert(bson.M{}); err != nil {
}
```
