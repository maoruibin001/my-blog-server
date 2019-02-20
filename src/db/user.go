package db

import (
	"fmt"
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
	Token string `json:"token"`
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
	return UserMultiFindByKV("name", name)
}

func UserSingleFindByKV(key string, v interface{}) UserSchema {

	c, session := GetCollect("my-blog-2", "user")
	defer session.Close()

	results := []UserSchema{}

	utils.HandleError("find error: ", c.Find(bson.M{key: v}).All(&results))

	result := UserSchema{}
	if len(results) > 0 {
		result = results[0]
	}
	return result
}

func UserMultiFindByKV( key string, v interface{}) []UserSchema {

	c, session := GetCollect("my-blog-2", "user")
	defer session.Close()

	results := []UserSchema{}

	utils.HandleError("find error: ", c.Find(bson.M{key: v}).All(&results))


	return results
}
func FindByPhone(phoneNumber string) UserSchema {

	return UserSingleFindByKV("phone", phoneNumber)
}

func FindById(id string) UserSchema {

	return UserSingleFindByKV("id", id)
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

func ChangeUser(id, name, age, phone, salt string) error {
	c, session := GetCollect("my-blog-2", "user")
	defer session.Close()

	selector := bson.M{"id": id}
	data := bson.M{"name": name, "age": age, "phone": phone, "salt": salt}

	err := c.Update(selector, bson.M{"$set": data})

	return err
}

func RemoveUser(k string, v interface{}) error {
	c, session := GetCollect("my-blog-2", "user")
	defer session.Close()

	err := c.Remove(bson.M{k: v})

	return err
}
//func InitUserData()  {
//	c, session := GetCollect("my-blog-2", "user")
//	defer session.Close()
//	count, err := c.Count()
//	utils.HandleError("查找错误：", err)
//	if count == 0 {
//		fmt.Println("数据库为空，初始化数据...")
//		//Name string `json:"name"`
//		//Age  string    `json:"age"`
//		//Salt string
//		//Password string `json:"password"`
//		//Phone string `json:"phone"`
//		InsertUser(UserInfo{"mao", "20", "", "123", "123"})
//	}
//}
