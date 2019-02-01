package db

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"my-blog-server/src/utils"
)

type UserInfo struct {
	Name string `json:"name"`
	Age  string    `json:"age"`
	Salt string
	Password string `json:"password"`
	Phone string `json:"phone"`

}
type UserSchema struct {
	Name string `json:"name"`
	Age  string    `json:"age"`
	Salt string
	Password string `json:"password"`
	Phone string `json:"phone"`
	_id bson.ObjectId
	Id string `json:"id"`
}

func InsertUser(data UserInfo) UserSchema {
	c, session := GetCollect("my-blog-2", "user")
	defer session.Close()
	m := UserSchema{}
	m.Phone = data.Phone
	m.Age = data.Age
	m.Name = data.Name
	m.Password = data.Password
	m.Salt = data.Salt

	m._id = bson.NewObjectId()
	m.Id = bson.NewObjectId().Hex()

	utils.HandleError("insert error: ", c.Insert(&m))
	fmt.Println("插入一条数据", m)

	return m

}

func FindByName(name string) []UserSchema {
	c, session := GetCollect("my-blog-2", "user")
	defer session.Close()
	results := UserFindByKV(c, "name", name)
	return results                                                  }

func UserFindByKV(c *mgo.Collection, key string, v interface{}) []UserSchema {

	results := []UserSchema{}

	utils.HandleError("find error: ", c.Find(bson.M{key: v}).All(&results))

	return results
}
func FindByPhone(phoneNumber string) UserSchema {

	c, session := GetCollect("my-blog-2", "user")
	defer session.Close()
	results := UserFindByKV(c, "phone", phoneNumber)

	result := UserSchema{}
	if len(results) > 0 {
		result = results[0]
	}
	return result
}

func CreateUser(name, age, phone, password, salt string) UserSchema {
	var userInfo = UserInfo{}
	userInfo.Name = name
	userInfo.Phone = phone
	userInfo.Password = password
	userInfo.Salt = salt
	userInfo.Age = age
	return InsertUser(userInfo)
}

func InitUserData()  {
	c, session := GetCollect("my-blog-2", "user")
	defer session.Close()
	count, err := c.Count()
	utils.HandleError("查找错误：", err)
	if count == 0 {
		fmt.Println("数据库为空，初始化数据...")
		//var i int
		//for i = 0; i < 10; i ++ {
		//	//InsertUser(UserInfo{"mao", i, strconv.Itoa(i)})
		//}

	}
}
