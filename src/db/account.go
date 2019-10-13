package db

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"album-server/src/utils"
	"album-server/src/config"
)

type AccountSchema struct {
	OpenId    string   `json:"openid"`
	NickName       string      `json:"nickName"`
	Name string `json:"name"`
	Sex   int   `json:"sex"`
	Province 		   string `json:"province"`
	City          string   `json:"city"`
	Country    string   `json:"country"`
	HeadImgUrl       string      `json:"headImgUrl"`
	Language 		   string `json:"language"`
	UId int `json:"uid"`
	Score int `json:"score"`
	QrImg string `json:"qrImg"`
	VipFee int `json:"vipFee"`
	SingFee int `json:"singFee"`

}
type WxAccountSchema struct {
	OpenId    string   `json:"openid"`
	NickName       string      `json:"nickname"`
	Sex   int   `json:"sex"`
	Province 		   string `json:"province"`
	City          string   `json:"city"`
	Country    string   `json:"country"`
	HeadImgUrl       string      `json:"headimgurl"`
	Language 		   string `json:"language"`
	UId int `json:"uid"`
	Name string `json:"name"`
}


func getQrImg(Id int) string {
	return ""
}
func InsertAccount(data WxAccountSchema) AccountSchema {
	c, session := GetCollect("userdb", "user")
	defer session.Close()
	m := AccountSchema{}


	m.OpenId = data.OpenId
	m.NickName = data.NickName
	if (data.Name == "") {
		m.Name = data.NickName
	} else {
		m.Name = data.Name
	}
	m.Sex = data.Sex
	m.Province = data.Province
	m.City = data.City
	m.Country = data.Country
	m.HeadImgUrl = data.HeadImgUrl
	m.Language = data.Language
	m.UId = data.UId
	m.Score = 0
	m.QrImg = getQrImg(m.UId)
	m.VipFee = config.VIPFEE
	m.SingFee = config.SINGFEE

	utils.HandleError("insert error: ", c.Insert(&m))
	fmt.Println("插入一条数据", data)

	return m

}

func AccountSingleFindByKV(key string, v interface{}) AccountSchema {

	c, session := GetCollect("userdb", "user")
	defer session.Close()

	results := []AccountSchema{}

	utils.HandleError("find error: ", c.Find(bson.M{key: v}).All(&results))

	result := AccountSchema{}
	if len(results) > 0 {
		result = results[0]
	}
	return result
}

func FindByOpenId(id string) AccountSchema {

	return AccountSingleFindByKV("openId", id)
}

func CreateAccount(accountInfo WxAccountSchema) AccountSchema {
	accountInfo.UId = generationNameId("uid")
	return InsertAccount(accountInfo)
}
