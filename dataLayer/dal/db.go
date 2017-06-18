package dal

import (
	"github.com/GwentAPI/gwentapi/configuration"
	"gopkg.in/mgo.v2"
)

var mainDataStore *DataStore

var database string

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
		ds.session = mainDataStore.session.Copy()
	} /* else {
		ds.session.Copy()
	}*/
	// Useless
	return ds.session.Copy()
}

func (ds *DataStore) Close() {
	ds.session.Close()
}

func ShutDown() {
	if mainDataStore != nil && mainDataStore.session != nil {
		mainDataStore.session.Close()
	}
}

func init() {
	config := configuration.GetConfig()
	database = config.App.DbName

	mainDataStore = &DataStore{}
	var err error
	mainDataStore.session, err = mgo.Dial(config.Database.URL)
	if err != nil {
		panic(err)
	}
	mainDataStore.session.SetMode(mgo.Monotonic, true)
}
