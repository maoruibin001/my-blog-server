package db

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"album-server/src/utils"
)

type UserInfo struct {
	Name string `json:"name"`
	Salt string
	Password string `json:"password"`
	Phone string `json:"phone"`
	IsKeeper int `json:"isKeeper"`
}
type UserSchema struct {
	Name string `json:"name"`
	Salt string
	Password string `json:"password"`
	Phone string `json:"phone"`
	IsKeeper int  `json:"isKeeper"`
	_id bson.ObjectId
	Id int `json:"id"`
	Token string `json:"token"`
}

func InsertUser(data UserInfo) UserSchema {
	c, session := GetCollect(utils.GetDbName(), "user")
	defer session.Close()
	m := UserSchema{}
	m.Phone = data.Phone
	m.Name = data.Name
	m.Password = data.Password
	m.Salt = data.Salt
	m.IsKeeper = data.IsKeeper
	// m.IsKeeper = data.IsKeeper

	m._id = bson.NewObjectId()
	m.Id = generationNameId("user")

	utils.HandleError("insert error: ", c.Insert(&m))
	fmt.Println("插入一条数据", m)

	return m

}

func FindByName(name string) []UserSchema {
	return UserMultiFindByKV("name", name)
}

func UserSingleFindByKV(key string, v interface{}) UserSchema {

	c, session := GetCollect(utils.GetDbName(), "user")
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

	c, session := GetCollect(utils.GetDbName(), "user")
	defer session.Close()

	results := []UserSchema{}

	utils.HandleError("find error: ", c.Find(bson.M{key: v}).All(&results))


	return results
}
func FindByPhone(phoneNumber string) UserSchema {

	return UserSingleFindByKV("phone", phoneNumber)
}

func FindById(id int) UserSchema {

	return UserSingleFindByKV("id", id)
}

func CreateUser(name, phone, password, salt string, isKeeper int) UserSchema {
	var userInfo = UserInfo{}
	userInfo.Name = name
	userInfo.Phone = phone
	userInfo.Password = password
	userInfo.Salt = salt
	userInfo.IsKeeper = isKeeper
	return InsertUser(userInfo)
}

func ChangeUser(id int, name, phone, salt string, isKeeper int) error {
	c, session := GetCollect(utils.GetDbName(), "user")
	defer session.Close()

	selector := bson.M{"id": id}
	data := bson.M{"name": name,"phone": phone, "salt": salt, "iskeeper": isKeeper}

	err := c.Update(selector, bson.M{"$set": data})

	return err
}

func RemoveUser(k string, v interface{}) error {
	c, session := GetCollect(utils.GetDbName(), "user")
	defer session.Close()

	err := c.Remove(bson.M{k: v})

	return err
}
