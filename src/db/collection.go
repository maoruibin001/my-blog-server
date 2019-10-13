package db

import (
	"album-server/src/utils"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type CollectionSchema struct {
	Uid int `json:"uid"`
	Pid int `json:"pid"`
	Id int `json:"id"`
	CreateDate int64 `json:"createDate"`
	CreateDateStr string `json:"createDateStr"`
}

type CollectionRetData struct {
	Collections []CollectionSchema `json:"collections"`
	Count int `json:"count"`
	IsEnd int `json:"isEnd"`
}
func InsertCollection(data CollectionSchema) CollectionSchema {
	c, session := GetCollect(utils.GetDbName(), "collection")
	defer session.Close()
	utils.HandleError("insert error: ", c.Insert(&data))
	fmt.Println("插入一条数据", data)

	return data

}


func CollectionSingleFindByKV(condition bson.M) CollectionSchema {

	c, session := GetCollect(utils.GetDbName(), "collection")
	defer session.Close()

	results := []CollectionSchema{}

	fmt.Println("condition: ", condition)
	utils.HandleError("find error: ", c.Find(condition).All(&results))

	fmt.Println("result www: ", results)
	result := CollectionSchema{}
	if len(results) > 0 {
		result = results[0]
	}
	return result
}


func GetCollections(condition bson.M, pageSize, pageNo int) (CollectionRetData, error)  {
	c, session := GetCollect(utils.GetDbName(), "collection")
	defer session.Close()

	results := []CollectionSchema{}
	var err error = nil
	var count = 0
	var ret = CollectionRetData{}

	fmt.Println("condition:", condition)
	err = c.Find(condition).Limit(pageSize).Skip((pageNo - 1) * pageSize).Sort("-createdate").All(&results)
	count, err = c.Find(condition).Count()
	fmt.Println("results:", results)

	ret.Collections = results
	ret.Count = count
	if pageNo * pageSize >= count {
		ret.IsEnd = 1
	} else {
		ret.IsEnd = 0
	}
	return ret, err
}

func CollectionMultiFindByKV(condition bson.M) []CollectionSchema {

	c, session := GetCollect(utils.GetDbName(), "collection")
	defer session.Close()

	results := []CollectionSchema{}

	utils.HandleError("find error: ", c.Find(condition).All(&results))


	return results
}

func CreateCollection(pid, uid int) CollectionSchema {
	m := CollectionSchema{}
	m.Id = generationNameId("collection")
	m.Pid = pid
	m.Uid = uid
	m.CreateDate = time.Now().UnixNano() / 1e6
	m.CreateDateStr = time.Now().Format("2006年01月02日 15时04分05秒")
	return InsertCollection(m)
}
func RemoveCollection(condition bson.M) error {
	c, session := GetCollect(utils.GetDbName(), "collection")
	defer session.Close()

	_, err := c.RemoveAll(condition)


	fmt.Println("取消收藏", err)

	return err
}
