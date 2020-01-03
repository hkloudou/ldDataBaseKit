package lddatabasekit

/*
 * @Author: hkloudou
 * @Github: https://github.com/hkloudou/
 * @LastEditors: 卢教(aven) hkloudou@gmail.com
 * @Date: 2018-07-07 02:45:59
 * @LastEditTime: 2019-03-28 22:14:58
 */

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
