package ldDataBaseKit

import (
	"errors"

	mgo "gopkg.in/mgo.v2"
)

var mMgoSession *mgo.Session

//ReInitMongoDBSession ReInitMongoDBSession
func ReInitMongoDBSession(dialInfo *mgo.DialInfo, mode mgo.Mode) error {
	var errDial error
	mMgoSession, errDial = mgo.DialWithInfo(dialInfo)
	if errDial != nil {

		return errors.New("hcLionDataBase.InitMongoDBSession:" + "Connect error " + errDial.Error())
	}
	mMgoSession.SetMode(mode, true)
	return nil
}

//GetMongoDBSession GetMongoDBSession
func GetMongoDBSession() (*mgo.Session, error) {
	if mMgoSession == nil {
		return nil, errors.New("please InitMongoDBSession before use GetMongoDBSession")
	}
	return mMgoSession.Clone(), nil
}

//CloseMongo Close a mongodb connect
func CloseMongo() {
	mMgoSession.Close()
}
