package db

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"album-server/src/utils"
	"time"
)

type ImgInfoSchema struct {
	ThumbUrl string `json:"thumbUrl"`
	Url string `json:"url"`
}
type ProductMoveSchema struct {
	Start int `json:"start"`
	End int `json:"end"`
	Seq int `json:"seq"`
}

type ProductSchema struct {
	Id int `json:"id"`
	LId int `json:"lId"`
	Seq int `json:"seq"`
	DescImg  string    `json:"descImg"`
	DescImgThumb string	 `json:"descImgThumb"`
	//GifImgThumb  string    `json:"gifImgThumb"`
	GifImg string	 `json:"gifImg"`
	Name  string    `json:"name"`
	OriginFile string	 `json:"originFile"`
	Prize int `json:"prize"`
	MainImgList []ImgInfoSchema `json:"mainImgList"`
	CreateDate int64 `json:"createDate"`
	CreateDateStr string `json:"createDateStr"`
	ModifyDate int64 `json:"modifyDate"`
	ModifyDateStr string `json:"modifyDateStr"`
	//Children []ProductSchema `json:"children"`
}
type RetData struct {
	Products []ProductSchema `json:"products"`
	Count int `json:"count"`
	IsEnd int `json:"isEnd"`
}
func InsertProduct(data ProductSchema) ProductSchema {
	c, session := GetCollect("album", "product")
	defer session.Close()
	m := ProductSchema{}
	m.Id = data.Id
	m.Seq = data.Seq
	m.LId = data.LId
	m.Name = data.Name
	m.DescImg = data.DescImg
	m.DescImgThumb = data.DescImgThumb
	m.GifImg = data.GifImg
	m.OriginFile = data.OriginFile
	m.Prize = data.Prize
	m.MainImgList = data.MainImgList
	m.CreateDate = data.CreateDate
	m.CreateDateStr = data.CreateDateStr
	m.ModifyDate = data.ModifyDate
	m.ModifyDateStr = data.ModifyDateStr


	utils.HandleError("insert error: ", c.Insert(&m))
	fmt.Println("插入一条数据", m)

	return m

}

//func FindByArtcleId(id int) []ArticleSchema {
//	return ArticleMultiFindByKV("aid", id)
//}

func ProductSingleFindByKV(condition bson.M) ProductSchema {

	c, session := GetCollect("album", "product")
	defer session.Close()

	results := []ProductSchema{}

	fmt.Println("condition: ", condition)
	utils.HandleError("find error: ", c.Find(condition).All(&results))

	fmt.Println("result www: ", results)
	result := ProductSchema{}
	if len(results) > 0 {
		result = results[0]
	}
	return result
}


func GetSomeProducts(conditions bson.M, pageSize, pageNo int)  (RetData, error) {
	c, session := GetCollect("album", "product")
	defer session.Close()


	var ret = RetData{}
	results := []ProductSchema{}
	var err error = nil

	count, err := c.Find(conditions).Count()

	err = c.Find(conditions).Limit(pageSize).Skip((pageNo - 1) * pageSize).Sort("-seq").All(&results)


	fmt.Println("results:", results)
	ret.Products = results
	ret.Count = count
	if pageNo * pageSize >= count {
		ret.IsEnd = 1
	} else {
		ret.IsEnd = 0
	}
	return ret, err
}


func GetProducts(condition bson.M, pageSize, pageNo int) (RetData, error)  {
	c, session := GetCollect("album", "product")
	defer session.Close()

	results := []ProductSchema{}
	var err error = nil
	var count = 0
	var ret = RetData{}

	fmt.Println("condition:", condition)
	err = c.Find(condition).Limit(pageSize).Skip((pageNo - 1) * pageSize).Sort("seq").All(&results)
	count, err = c.Find(condition).Count()
	fmt.Println("results:", results)

	ret.Products = results
	ret.Count = count
	if pageNo * pageSize >= count {
		ret.IsEnd = 1
	} else {
		ret.IsEnd = 0
	}
	return ret, err
}

func ProductMultiFindByKV(condition bson.M) []ProductSchema {

	c, session := GetCollect("album", "product")
	defer session.Close()

	results := []ProductSchema{}

	utils.HandleError("find error: ", c.Find(condition).All(&results))


	return results
}

func generationId() int {
	counter, session := GetCollect("album", "acounter")
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
func CreateProduct(name,descImg,descImgThumb,gifImg,originFile string, prize, lId int, mainImgList []ImgInfoSchema) ProductSchema {
	m := ProductSchema{}
	m.Id = generationId()
	m.Seq = generationNameId("pId")
	m.LId = lId
	m.Name = name
	m.DescImg = descImg
	m.DescImgThumb = descImgThumb
	m.GifImg = gifImg
	m.OriginFile = originFile
	m.Prize = prize
	m.MainImgList = mainImgList
	m.CreateDate = time.Now().UnixNano() / 1e6
	m.CreateDateStr = time.Now().Format("2006年01月02日 15时04分05秒")
	m.ModifyDate = time.Now().UnixNano() / 1e6
	m.ModifyDateStr = time.Now().Format("2006年01月02日 15时04分05秒")
	return InsertProduct(m)
}

func ChangeProduct(name,descImg,descImgThumb,gifImg,originFile string, prize, lId,id int, mainImgList []ImgInfoSchema, seq int) error {
	c, session := GetCollect("album", "product")
	defer session.Close()

	selector := bson.M{"id": id}
	modifyDate := time.Now().UnixNano() / 1e6
	modifyDateStr := time.Now().Format("2006年01月02日 15时04分05秒")

	//data := bson.M{"name": name,"descImg": descImg,"descImgThumb": descImgThumb,"gifImg": gifImg, "originFile": originFile, "prize": prize, "lId": lId, "mainImgList": mainImgList}
	data := bson.M{"name": name,"descimg": descImg,"descimgthumb": descImgThumb,"gifimg": gifImg, "originfile": originFile, "prize": prize, "lid": lId, "mainimglist": mainImgList, "modifydate":modifyDate, "modifydatestr":modifyDateStr, "seq": seq}

	err := c.Update(selector, bson.M{"$set": data})

	return err
}

func RemoveProduct(k string, v interface{}) error {
	c, session := GetCollect("album", "product")
	defer session.Close()

	err := c.Remove(bson.M{k: v})

	return err
}
func IniProductData()  {
	c, session := GetCollect("album", "product")
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
