# keesely/mongo

基于gopkg/mgo.v2 封装的GO MongoDB驱动

## 安装与使用:

### 获取代码:

```
go get github.com/keesely/mongo
```

### 使用：

#### 连接

```

import "github.com/keesely/mongo"

// mongodb://{host}:{port}/{database}
dsn := "mongodb://locahost:27017/go-test"

// dsn And collection name
m := mongo.NewDB(dsn, "collection")

```

### 插入

```
row := &models.Models{ Test: "Hello world" }

if err := m.Insert(row); err != nil {
  t.Error(err)
}
```

### 查询

```
result := &models.Models{}
_id := "5b0b60899ae16b40cb0ac6c4"

// 根据ID查询
// m.Find(_id).One(result)
a
// 条件查询
m.Find(mongo.M{"_id" : mongo.ObjectId(_id)}).ONe(result)

```

### 更新

```
_id := "5b0b60899ae16b40cb0ac6c4"
data := mongo.M{ "test" : "Updated Text", "updated_at" : time.Now() }

// 根据ID更新
if err := m.Updated(_id, data); err != nil {
  t.Error(err)
}

// 根据条件更新
where := mongo.M{ "_id" : mongo.ObjectId(_id) }

if err := m.Updated(where, data); err != nil {
  t.Error(err)
}
```

### 删除

```
// ID删除
_id := "5b0b60899ae16b40cb0ac6c4"

if err := m.Remove(_id); err != nil {
  t.Error(err)
}

// 条件删除
where := mongo.M{ "_id" : mongo.ObjectId(_id) }

if err := m.Remove(where); err != nil {
  t.Error(err)
}

```
