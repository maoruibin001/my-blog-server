package db

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"album-server/src/utils"
	"time"
)

//type ImgInfoSchema struct {
//	ThumbUrl string `json:"thumbUrl"`
//	Url string `json:"url"`
//}
type BseriesMoveSchema struct {
	Start int `json:"start"`
	End int `json:"end"`
	Name  string    `json:"name"`
	Seq int `json:"seq"`
	Type string `json:"type"`
}
type BseriesSchema struct {
	BId int `json:"bId"`
	Name  string    `json:"name"`
	Seq int `json:"seq"`
	Children []LseriesSchema `json:"children"`
	CreateDate int64 `json:"createDate"`
	CreateDateStr string `json:"createDateStr"`
	ModifyDate int64 `json:"modifyDate"`
	ModifyDateStr string `json:"modifyDateStr"`
}
type BRetData struct {
	Bseries []BseriesSchema `json:"bseries"`
	Count int `json:"count"`
	IsEnd int `json:"isEnd"`
}
func InsertBseries(data BseriesSchema) BseriesSchema {
	c, session := GetCollect("album", "bseries")
	defer session.Close()
	m := BseriesSchema{}
	m.BId = data.BId
	m.Seq = data.Seq
	m.Name = data.Name
	m.CreateDate = data.CreateDate
	m.CreateDateStr = data.CreateDateStr
	m.ModifyDate = data.ModifyDate
	m.ModifyDateStr = data.ModifyDateStr


	utils.HandleError("insert error: ", c.Insert(&m))
	fmt.Println("插入一条数据", m)

	return m

}

func BseriesSingleFindByKV(condition bson.M) BseriesSchema {

	c, session := GetCollect("album", "bseries")
	defer session.Close()

	results := []BseriesSchema{}

	fmt.Println("condition: ", condition)
	utils.HandleError("find error: ", c.Find(condition).All(&results))

	fmt.Println("result www: ", results)
	result := BseriesSchema{}
	if len(results) > 0 {
		result = results[0]
	}
	return result
}


func GetBserieses(condition bson.M, pageSize, pageNo int) (BRetData, error)  {
	c, session := GetCollect("album", "bseries")
	defer session.Close()

	results := []BseriesSchema{}
	var err error = nil
	var count = 0
	var ret = BRetData{}

	fmt.Println("condition:", condition)
	err = c.Find(condition).Limit(pageSize).Skip((pageNo - 1) * pageSize).Sort("seq").All(&results)
	count, err = c.Find(condition).Count()
	fmt.Println("results:", results)

	ret.Bseries = results
	ret.Count = count
	if pageNo * pageSize >= count {
		ret.IsEnd = 1
	} else {
		ret.IsEnd = 0
	}
	return ret, err
}

func generationNameId(name string) int {
	counter, session := GetCollect("album", name)
	defer session.Close()

	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"id": 1}},
		Upsert:    true,
		ReturnNew: true,
	}
	doc := struct{ Id int }{}
	if _, err := counter.Find(bson.M{}).Apply(change, &doc); err != nil {
		utils.HandleError("查找出错", err)
	}
	log.Println("doc:", doc)
	return doc.Id
}
func CreateBseries(name string) BseriesSchema {
	m := BseriesSchema{}
	m.BId = generationNameId("bId")
	m.Seq = generationNameId("bseq")
	m.Name = name
	m.CreateDate = time.Now().UnixNano() / 1e6
	m.CreateDateStr = time.Now().Format("2006年01月02日 15时04分05秒")
	m.ModifyDate = time.Now().UnixNano() / 1e6
	m.ModifyDateStr = time.Now().Format("2006年01月02日 15时04分05秒")
	return InsertBseries(m)
}

func ChangeBseries(id int, name string, seq int) error {
	c, session := GetCollect("album", "bseries")
	defer session.Close()

	selector := bson.M{"bid": id}
	modifyDate := time.Now().UnixNano() / 1e6
	modifyDateStr := time.Now().Format("2006年01月02日 15时04分05秒")

	data := bson.M{"name": name, "seq": seq, "modifydate":modifyDate, "modifydatestr":modifyDateStr}

	err := c.Update(selector, bson.M{"$set": data})

	return err
}

func RemoveBseries(k string, v interface{}) error {
	c, session := GetCollect("album", "bseries")
	defer session.Close()

	_, err := c.RemoveAll(bson.M{k: v})

	return err
}
//func IniProductData()  {
//	c, session := GetCollect("album", "product")
//	defer session.Close()
//	count, err := c.Count()
//	utils.HandleError("查找错误：", err)
//	if count == 0 {
//		fmt.Println("数据库为空，初始化数据...")
//		//var i int
//		//for i = 0; i < 10; i ++ {
//		//	//InsertUser(UserInfo{"mao", i, strconv.Itoa(i)})
//		//}
//
//	}
//}
