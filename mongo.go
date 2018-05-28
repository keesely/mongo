/*************************************************************************
   > File Name: mongo.go
   > Author: Kee
   > Mail: chinboy2012@gmail.com
   > Created Time: 2018.05.18
************************************************************************/
package mongo

import "gopkg.in/mgo.v2"

const (
	FormatTime     = "15:04:05"
	FormatDate     = "2006-01-02"
	FormatDateTime = "2006-01-02 15:04:05"
)

type M map[string]interface{}

type DB struct {
	DSN        string
	CName      string
	session    *mgo.Session
	database   *mgo.Database
	collection *mgo.Collection
}

func NewDB(DSN string, c string) *DB {
	sess, err := mgo.Dial(DSN)

	if err != nil {
		panic(err)
	}

	sess.SetMode(mgo.Monotonic, true)
	coll := sess.DB("").C(c)

	//this.C(c)
	return &DB{DSN, c, sess, nil, coll}
}

func (this *DB) Conn() *mgo.Session {
	return this.session.Copy()
}

func (this *DB) DB(db string) *mgo.Database {
	this.database = this.Conn().DB(db)
	return this.database
}

func (this *DB) C(c string) *mgo.Collection {
	this.collection = this.Conn().DB("").C(c)
	return this.collection
}

func (this *DB) Find(q interface{}) *mgo.Query {
	return this.C(this.CName).Find(q)
}
