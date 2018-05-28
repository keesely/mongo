package tests

import (
	"encoding/json"
	"fmt"
	"github.com/keesely/mongo"
	"github.com/keesely/mongo/models"
	"testing"
	"time"
)

func Test(t *testing.T) {

	dsn := "mongodb://localhost:27017/go-test"
	m := mongo.NewDB(dsn, "test")

	// 自增ObjectId 写入
	row := &models.Tests{
		Test: "Hello World",
	}
	err := m.Insert(row)

	if err != nil {
		t.Error("Expected: ", err)
	} else {
		fmt.Println("Saveed :", row.ID)
	}

	// 获取单条数据
	result := &models.Tests{}
	m.Find(mongo.M{"_id": row.ID}).One(result)
	fmt.Printf("%+v \n", result)

	// 更新数据
	if err := m.Update(row.ID, mongo.M{"test": "Updated Text", "updated_at": time.Now()}); err != nil {
		t.Error(err)
	} else {
		fmt.Printf("Updated : %s \n", row.ID)
	}

	// 所有
	resultAll := []models.Tests{}
	m.Find(nil).All(&resultAll)
	js, _ := json.Marshal(resultAll)
	fmt.Printf("%+v\n", string(js))

	// 条件列表
	//result3 := []interface{}{}
	//result3 := []models.Tests{}
	//m.Find(mongo.M{"shared": "locked"}).All(&result3)
	//fmt.Printf("%+v \n", result3)
	//js, _ = json.Marshal(result3)
	//fmt.Println(string(js))
	//for i := range result3 {
	//ref := result3[i]
	//fmt.Printf("%+v\n", ref.Shared)
	//}
}
