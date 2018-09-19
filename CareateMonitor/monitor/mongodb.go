package monitor

import (
	"gopkg.in/mgo.v2"
)

// MongoDB : 写入MongoDB
type MongoDB struct {
	connectionString string
	dbName           string
	collectionName   string
	connection       *mgo.Session
	currCollection   *mgo.Collection
}

//InitMongoDBConnection : 初始化MongoDB连接
func InitMongoDBConnection(connectionString string, dbName string, collectionName string) (mongoDB *MongoDB) {
	if session, err := mgo.Dial(connectionString); err != nil {
		panic("连接失败")
	} else {
		mongoDB = &MongoDB{
			connectionString: connectionString,
			connection:       session,
			currCollection:   session.DB(mongoDB.dbName).C(mongoDB.collectionName),
		}
	}
	return
}
func (mongoDB *MongoDB) write(data *MonitorMessage) {
	mongoDB.currCollection.Insert(data)
}
