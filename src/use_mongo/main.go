package main

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	MongoHost = "10.98.16.210"
	MongoPort = 37017
	MongoDB = "test_cygx1_comment_filter"
	MongoCollection = "sensitive_words"
)

func main()  {
	id := "5948dc16df9ea6474ef50146"
	addrs := fmt.Sprintf("%s:%d", MongoHost, MongoPort)
	dialInfo := mgo.DialInfo{
		Addrs: []string{addrs},
		Database: MongoDB,
	}
	session ,err  := mgo.DialWithInfo(&dialInfo)
	defer session.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	db := session.DB(MongoDB)
	collection := db.C(MongoCollection)

	s := bson.M{
		"_id":bson.ObjectIdHex(id),
	}
	//u := bson.M{
	//	"status":"100",
	//}
	var obj SensitiveWord
	collection.FindId(bson.ObjectIdHex(id)).One(&obj)
	fmt.Println(obj.CreateAt)
	obj.Status = "123"
	changeInfo, err := collection.Upsert(s, obj)
	fmt.Println(changeInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("---END---")
}

type SensitiveWord struct {
	Keyword  string    `json:"keyword" bson:"keyword"`
	Status   string    `json:"status" bson:"status"`
	Level    int       `json:"level" bson:"level"`
	Type     string    `json:"type" bson:"type"`
	CreateAt time.Time `json:"create_at" bson:"create_at"`
	UpdateAt time.Time `json:"update_at" bson:"update_at"`
}
