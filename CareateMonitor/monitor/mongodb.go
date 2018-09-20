package monitor

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

// MongoDB : 写入MongoDB
type MongoDB struct {
	connectionString string
	dbName           string
	collectionName   string
	connection       *mgo.Session
}

//InitMongoDBConnection : 初始化MongoDB连接
func InitMongoDBConnection(connectionString string, dbName string, collectionName string) (mongoDB *MongoDB) {
	if session, err := mgo.Dial(connectionString); err != nil {
		fmt.Println("连接失败")
	} else {
		fmt.Println("连接成功")
		mongoDB = &MongoDB{
			connectionString: connectionString,
			dbName:           dbName,
			collectionName:   collectionName,
			connection:       session, //总连接
		}
	}
	return
}
func (mongoDB *MongoDB) write(data *MonitorMessage) (err error) {
	copyConnection := mongoDB.connection.Copy()
	defer copyConnection.Close()
	err = copyConnection.DB(mongoDB.dbName).C(mongoDB.collectionName).Insert(data) //.curCollection.Insert(data)
	return
}

//Close : 关闭连接
func (mongoDB *MongoDB) Close() {
	mongoDB.connection.Close()
}
