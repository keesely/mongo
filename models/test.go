/*************************************************************************
   > File Name: test.go
   > Author: Kee
   > Mail: chinboy2012@gmail.com
   > Created Time: 2018.05.27
************************************************************************/
package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Tests struct {
	ID    bson.ObjectId `bson:"_id"        json:"_id"             def:"ObjectId"`
	Test  string        `bson:"test"       json:"test,omitempty"  defval:"testing"`
	CDate time.Time     `bson:"created_at" json:"created_at"      def:"Now"`
	UDate time.Time     `bson:"updated_at" json:"updated_at"      def:"Now"`
	DDate time.Time     `bson:"deleted_at" json:"deleted_at,omitempty"`
}
