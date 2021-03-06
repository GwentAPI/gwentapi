package dal

import (
	"crypto/tls"
	"github.com/GwentAPI/gwentapi/configuration"
	"gopkg.in/mgo.v2"
	"log"
	"net"
	"time"
)

var mainDataStore *DataStore

type DataStore struct {
	session *mgo.Session
}

func (ds *DataStore) Collection(colName string) *mgo.Collection {
	// Use database name specified in DialWithInfo when parameter is empty
	// Fallback to "test" if also empty.
	return ds.session.DB("").C(colName)
}

func InitDBWithInfo(info *mgo.DialInfo) {

	_, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
	}

	//Gwentapi = session.DB("gwentapi")
}

func (ds *DataStore) GetSession() {
	if ds.session == nil {
		ds.session = mainDataStore.session.Copy()
	} /* else {
		ds.session.Copy()
	}*/
	// Useless
	//return ds.session.Copy()
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
	tlsConfig := &tls.Config{}

	addrs := config.Database.Addrs

	dialInfo := &mgo.DialInfo{
		Addrs:    addrs,
		Database: config.Database.Database,
		Source:   config.Database.Authentication.AuthenticationDatabase,
		Username: config.Database.Authentication.Username,
		Password: config.Database.Authentication.Password,
		Timeout:  10 * time.Second,
	}

	mainDataStore = &DataStore{}
	var err error

	if config.Database.UseSSL {
		dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
			conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
			return conn, err
		}
	}

	log.Println("Attempting to establish a database connection...")

	mainDataStore.session, err = mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Fatal("Failed to establish mongoDB connection: ", err)
	}
	mainDataStore.session.SetMode(mgo.Monotonic, true)
	mainDataStore.session.SetSocketTimeout(10 * time.Second)
	log.Println("Database connection established.")
}
