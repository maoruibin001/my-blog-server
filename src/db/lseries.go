package db

import (
	"album-server/src/utils"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"time"
)

//type ImgInfoSchema struct {
//	ThumbUrl string `json:"thumbUrl"`
//	Url string `json:"url"`
//}
type LseriesMoveSchema struct {
	Start int `json:"start"`
	End int `json:"end"`
	Name  string    `json:"name"`
	Seq int `json:"seq"`
	Type string `json:"type"`
}
type LseriesSchema struct {
	BId int `json:"bId"`
	LId int`json:"lId"`
	Name  string    `json:"name"`
	Seq int `json:"seq"`
	MainImg string `json:"mainImg"`
	MainImgThumb string `json:"mainImgThumb"`
	Children []LseriesSchema `json:"children"`
	CreateDate int64 `json:"createDate"`
	CreateDateStr string `json:"createDateStr"`
	ModifyDate int64 `json:"modifyDate"`
	ModifyDateStr string `json:"modifyDateStr"`
}
type LRetData struct {
	Lseries []LseriesSchema `json:"lseries"`
	Count int `json:"count"`
	IsEnd int `json:"isEnd"`
}
func InsertLseries(data LseriesSchema) LseriesSchema {
	c, session := GetCollect("album", "lseries")
	defer session.Close()

	utils.HandleError("insert error: ", c.Insert(&data))
	fmt.Println("插入一条数据", data)

	return data

}

func LseriesSingleFindByKV(condition bson.M) LseriesSchema {

	c, session := GetCollect("album", "lseries")
	defer session.Close()

	results := []LseriesSchema{}

	fmt.Println("condition: ", condition)
	utils.HandleError("find error: ", c.Find(condition).All(&results))

	fmt.Println("result www: ", results)
	result := LseriesSchema{}
	if len(results) > 0 {
		result = results[0]
	}
	return result
}


func GetLserieses(condition bson.M, pageSize, pageNo int) (LRetData, error)  {
	c, session := GetCollect("album", "lseries")
	defer session.Close()

	results := []LseriesSchema{}
	var err error = nil
	var count = 0
	var ret = LRetData{}

	fmt.Println("condition:", condition)
	err = c.Find(condition).Limit(pageSize).Skip((pageNo - 1) * pageSize).Sort("seq").All(&results)
	count, err = c.Find(condition).Count()
	fmt.Println("results:", results)

	ret.Lseries = results
	ret.Count = count
	if pageNo * pageSize >= count {
		ret.IsEnd = 1
	} else {
		ret.IsEnd = 0
	}
	return ret, err
}

//func generationNameId(name string) int {
//	counter, session := GetCollect("album", name)
//	defer session.Close()
//
//	change := mgo.Change{
//		Update:    bson.M{"$inc": bson.M{"id": 1}},
//		Upsert:    true,
//		ReturnNew: true,
//	}
//	doc := struct{ Id int }{}
//	if _, err := counter.Find(bson.M{}).Apply(change, &doc); err != nil {
//		utils.HandleError("查找出错", err)
//	}
//	log.Println("doc:", doc)
//	return doc.Id
//}
func CreateLseries(bId int,name, mainImg, mainImgThumb string) LseriesSchema {
	m := LseriesSchema{}
	m.BId = bId
	m.LId = generationNameId("bId")
	m.Seq = generationNameId("lseq")
	m.Name = name
	m.MainImg = mainImg
	m.MainImgThumb = mainImgThumb
	m.CreateDate = time.Now().UnixNano() / 1e6
	m.CreateDateStr = time.Now().Format("2006年01月02日 15时04分05秒")
	m.ModifyDate = time.Now().UnixNano() / 1e6
	m.ModifyDateStr = time.Now().Format("2006年01月02日 15时04分05秒")
	return InsertLseries(m)
}

func ChangeLseries(id int, name, mainImg, mainImgThumb string, seq int) error {
	c, session := GetCollect("album", "lseries")
	defer session.Close()

	selector := bson.M{"lid": id}
	modifyDate := time.Now().UnixNano() / 1e6
	modifyDateStr := time.Now().Format("2006年01月02日 15时04分05秒")


	data := bson.M{"name": name, "seq": seq, "mainimg": mainImg, "mainimgthumb": mainImgThumb, "modifydate":modifyDate, "modifydatestr":modifyDateStr}

	err := c.Update(selector, bson.M{"$set": data})

	return err
}

func RemoveLseries(k string, v interface{}) error {
	c, session := GetCollect("album", "lseries")
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
