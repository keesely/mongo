/*************************************************************************
   > File Name: writer.go
   > Author: Kee
   > Mail: chinboy2012@gmail.com
   > Created Time: 2018.05.28
************************************************************************/
package mongo

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"reflect"
	"strings"
	"time"
)

// 插入数据
func (this *DB) Insert(d interface{}) error {
	PrepareValue(d)
	return this.C(this.CName).Insert(d)
}

// 更新数据
func (this *DB) Update(_where interface{}, data interface{}) error {
	//selector := M{"_id": ObjectId(_id)}
	selector := parSelector(_where)
	_data := M{"$set": data}
	return this.C(this.CName).Update(selector, _data)
}

func (this *DB) Remove(_where interface{}) error {
	selector := parSelector(_where)
	return this.C(this.CName).Remove(selector)
}

func parSelector(_where interface{}) M {
	switch _where.(type) {
	case string:
		return M{"_id": ObjectId(_where.(string))}
	case bson.ObjectId:
		return M{"_id": _where.(bson.ObjectId)}
	default:
		return _where.(M)
	}
}

func PrepareValue(d interface{}) {
	s := reflect.ValueOf(d).Elem()
	tf := s.Type()

	for i := 0; i < s.NumField(); i++ {
		def := tf.Field(i).Tag.Get("def")
		if def != "" {
			switch def {
			case "ObjectId":
				s.Field(i).Set(reflect.ValueOf(ObjectId()))
			case "Now":
				s.Field(i).Set(reflect.ValueOf(time.Now()))
			default:
				//s.Field(i).Set(reflect.ValueOf(def))
			}
		}

		svt := fmt.Sprintf("%s", s.Field(i).Type())
		ss := fmt.Sprintf("%s", s.Field(i))

		if (svt == "string" && ss == "") || !s.Field(i).IsValid() {
			defVal := tf.Field(i).Tag.Get("defval")
			if defVal != "" {
				s.Field(i).Set(reflect.ValueOf(defVal))
			}
		}
	}
}

func ObjectId(id ...string) bson.ObjectId {
	if id == nil {
		return bson.NewObjectId()
	}
	idx := strings.Join(id, "")
	return bson.ObjectIdHex(idx)
}

func StrId(id bson.ObjectId) string {
	return fmt.Sprintf("%x", string(id))
}
