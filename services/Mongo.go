package services

import(
	mgo "gopkg.in/mgo.v2"
)

func ConnectDB() (*mgo.Database, error){
	session, err := mgo.Dial("localhost")
	if err != nil {
		return nil, err
	}
	db := session.DB("golang-test")

	return db, nil
}