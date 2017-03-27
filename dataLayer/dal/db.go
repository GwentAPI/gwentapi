package dal

import (
	"gopkg.in/mgo.v2"
)

var dataStore *DataStore

const database string = "gwentapi"

type DataStore struct {
	session *mgo.Session
}

func (ds *DataStore) Collection(colName string) *mgo.Collection {
	return ds.session.DB(database).C(colName)
}

func InitDBWithInfo(info *mgo.DialInfo) {

	_, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
	}

	//Gwentapi = session.DB("gwentapi")
}

func (ds *DataStore) GetSession() *mgo.Session {
	if ds.session == nil {
		var err error
		ds.session, err = mgo.Dial("localhost")
		ds.session.SetMode(mgo.Monotonic, true)
		if err != nil {
			panic(err)
		}
	}

	return ds.session.Copy()
}

func (ds *DataStore) Close() {
	ds.session.Close()
}
